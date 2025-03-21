import json
import sys
import requests
import configparser
from typing import Dict, Any, Optional
from pathlib import Path
import logging

class AITool:
    def __init__(self, config_path: str = "config.ini"):
        """初始化AI工具"""
        self.config = self._load_config(config_path)
        self.setup_logging()
        
    def setup_logging(self):
        """设置日志"""
        logging.basicConfig(
            level=logging.INFO,
            format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
            handlers=[
                logging.FileHandler('ai_tool.log'),
                logging.StreamHandler()
            ]
        )
        self.logger = logging.getLogger('AITool')

    def _load_config(self, config_path: str) -> Dict[str, Any]:
        """加载并解析配置文件"""
        if not Path(config_path).exists():
            raise FileNotFoundError(f"配置文件不存在: {config_path}")

        config = configparser.ConfigParser()
        config.read(config_path, encoding='utf-8')
        
        # 验证热键格式
        if 'Hotkey' in config:
            for key, value in config['Hotkey'].items():
                hotkey = value.split(';')[0].strip()  # 移除注释
                if not self._validate_hotkey(hotkey):
                    raise ValueError(f"无效的热键格式 '{key}': {hotkey}")
        
        # 解析Functions部分的JSON字符串
        functions = {}
        if 'Functions' in config:
            for key in config['Functions']:
                try:
                    functions[key] = json.loads(config['Functions'][key])
                except json.JSONDecodeError as e:
                    raise ValueError(f"Functions配置解析错误 '{key}': {e}")

        return {
            'api': dict(config['API']),
            'functions': functions
        }

    def _validate_hotkey(self, hotkey: str) -> bool:
        """验证热键格式是否正确"""
        valid_modifiers = ['!', '^', '+', '#']
        
        if len(hotkey) < 2:
            return False
            
        modifier = hotkey[0]
        key = hotkey[1:]
        return modifier in valid_modifiers and len(key) == 1

    def _make_api_request(self, messages: list, function_config: Dict) -> Dict:
        """发送API请求"""
        headers = {
            'Authorization': f"Bearer {self.config['api']['api_key']}",
            'Content-Type': 'application/json'
        }

        data = {
            'model': self.config['api']['model'],
            'messages': messages,
            'temperature': function_config.get('temperature', 0.7),
            'stream': function_config.get('stream', False)
        }

        try:
            response = requests.post(
                f"{self.config['api']['base_url']}/chat/completions",
                headers=headers,
                json=data,
                timeout=float(self.config['api']['timeout'])
            )
            response.raise_for_status()
            return response.json()
        except Exception as e:
            self.logger.error(f"API请求失败: {str(e)}")
            raise

    def process_text(self, function_name: str, text: str, **kwargs) -> str:
        """处理文本，根据功能名称调用相应的AI功能"""
        if function_name not in self.config['functions']:
            raise ValueError(f"未定义的功能: {function_name}")

        function_config = self.config['functions'][function_name]
        
        # 准备消息
        messages = [
            {
                'role': 'system',
                'content': function_config['system_prompt']
            },
            {
                'role': 'user',
                'content': function_config['user_prompt'].format(text=text, **kwargs)
            }
        ]

        try:
            result = self._make_api_request(messages, function_config)
            return result['choices'][0]['message']['content'].strip()
        except Exception as e:
            self.logger.error(f"处理失败: {str(e)}")
            return f"处理失败: {str(e)}"

    def detect_language(self, text: str) -> str:
        """检测文本语言"""
        return 'zh' if any('\u4e00' <= char <= '\u9fff' for char in text) else 'en'

    def translate(self, text: str) -> str:
        """翻译文本"""
        source_lang = self.detect_language(text)
        target_lang = 'English' if source_lang == 'zh' else 'Chinese'
        return self.process_text(
            'translate', 
            text,
            source_lang=source_lang,
            target_lang=target_lang
        )

    def polish(self, text: str) -> str:
        """润色文本"""
        return self.process_text('polish', text)

    def summarize(self, text: str) -> str:
        """总结文本"""
        return self.process_text('summarize', text)

def main():
    """命令行入口点"""
    if len(sys.argv) < 3:
        print("用法: python ai_tool.py <function> <text>")
        sys.exit(1)

    function_name = sys.argv[1]
    text = sys.argv[2]

    try:
        ai_tool = AITool()
        if hasattr(ai_tool, function_name):
            result = getattr(ai_tool, function_name)(text)
        else:
            result = ai_tool.process_text(function_name, text)
            
        # 将结果写入文件
        with open('result.txt', 'w', encoding='utf-8') as f:
            f.write(result)
            
    except Exception as e:
        # 错误信息写入文件
        with open('result.txt', 'w', encoding='utf-8') as f:
            f.write(f"错误: {str(e)}")

if __name__ == '__main__':
    main()