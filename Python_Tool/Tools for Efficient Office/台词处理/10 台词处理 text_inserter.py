# -*- coding: UTF-8 -*-
"""
@author: yaoleistable
@file: text_inserter.py
@time: 2025-03-18 09:20:17
@description: 根据标点符号在文本中插入指定文字
@version: 1.0.6

功能：
1. 从1-原始台词.txt读取原始文本
2. 从2-插入字符A.txt读取插入文字A（遇到","或"，"时插入）
3. 从3-插入字符B.txt读取插入文字B（遇到"."、"。"、"!"或"！"时插入）
4. 处理结果输出到4-修改后台词.txt并自动打开
"""

import os
import sys
import logging
import subprocess
from datetime import datetime
from typing import List, Tuple, Optional

# 配置信息
CURRENT_USER = "yaoleistable"
CURRENT_TIME = "2025-03-18 09:20:17"
INPUT_FILE_A = "1-原始台词.txt"
INPUT_FILE_B = "2-插入字符A.txt"
INPUT_FILE_C = "3-插入字符B.txt"
OUTPUT_FILE = "4-修改后台词.txt"

def setup_logging() -> None:
    """配置日志系统"""
    log_dir = "logs"
    if not os.path.exists(log_dir):
        os.makedirs(log_dir)

    log_file = os.path.join(log_dir, f"text_insert_{datetime.now().strftime('%Y%m%d')}.log")

    logging.basicConfig(
        level=logging.INFO,
        format='%(asctime)s - %(levelname)s - %(message)s',
        datefmt='%Y-%m-%d %H:%M:%S',
        handlers=[
            logging.FileHandler(log_file, encoding='utf-8'),
            logging.StreamHandler()
        ]
    )

def read_file(filename: str) -> Optional[str]:
    """
    读取文件内容

    Args:
        filename: 文件名

    Returns:
        Optional[str]: 文件内容，如果出错返回None
    """
    try:
        if not os.path.exists(filename):
            logging.error(f"文件不存在: {filename}")
            return None

        with open(filename, 'r', encoding='utf-8') as f:
            content = f.read().strip()
            return content

    except Exception as e:
        logging.error(f"读取文件 {filename} 时出错: {str(e)}")
        return None

def open_file(filepath: str) -> None:
    """
    使用系统默认程序打开文件

    Args:
        filepath: 文件路径
    """
    try:
        if sys.platform == 'win32':
            os.startfile(filepath)
        elif sys.platform == 'darwin':  # macOS
            subprocess.run(['open', filepath])
        else:  # linux
            subprocess.run(['xdg-open', filepath])
    except Exception as e:
        logging.error(f"打开文件失败: {str(e)}")
        print(f"无法自动打开文件: {str(e)}")

class TextInserter:
    def __init__(self, text_b: str, text_c: str):
        """
        初始化文本插入器

        Args:
            text_b: 遇到","或"，"时要插入的文字
            text_c: 遇到"."、"。"、"!"或"！"时要插入的文字
        """
        self.text_b = text_b
        self.text_c = text_c
        self.punctuation_b = {',', '，'}  # 英文逗号、中文逗号
        self.punctuation_c = {'.', '。', '!', '！'}  # 英文句号、中文句号、英文感叹号、中文感叹号

    def process_text(self, input_text: str) -> str:
        """
        处理文本，在指定标点符号后插入文字

        Args:
            input_text: 输入的文本

        Returns:
            str: 处理后的文本
        """
        # 将输入文本分割成字符列表
        chars = list(input_text)
        result = []
        i = 0

        while i < len(chars):
            char = chars[i]
            result.append(char)

            # 检查是否是目标标点符号
            if char in self.punctuation_b:
                result.append(self.text_b)
            elif char in self.punctuation_c:
                result.append(self.text_c)

            i += 1

        return ''.join(result)

def main():
    """主函数"""
    setup_logging()
    print("文本处理工具")
    print("=" * 50)
    print(f"当前用户: {CURRENT_USER}")
    print(f"当前时间: {CURRENT_TIME}")
    print(f"工作目录: {os.getcwd()}")
    print("=" * 50)

    try:
        # 读取输入文件
        text_a = read_file(INPUT_FILE_A)
        text_b = read_file(INPUT_FILE_B)
        text_c = read_file(INPUT_FILE_C)

        if not all([text_a, text_b, text_c]):
            print("读取文件失败，请检查输入文件是否存在且内容正确")
            return

        # 创建处理器并处理文本
        inserter = TextInserter(text_b, text_c)
        result = inserter.process_text(text_a)

        # 保存结果
        try:
            with open(OUTPUT_FILE, 'w', encoding='utf-8') as f:
                f.write(result)  # 直接写入处理后的文本

            print(f"\n处理完成！")
            print(f"- 原始文本: {text_a}")
            print(f"- 处理结果: {result}")
            print(f"- 输出文件: {OUTPUT_FILE}")

            # 自动打开输出文件
            print("\n正在打开输出文件...")
            open_file(OUTPUT_FILE)

        except Exception as e:
            logging.error(f"保存结果时出错: {str(e)}")
            print(f"保存结果失败: {str(e)}")

    except KeyboardInterrupt:
        print("\n程序被用户中断")
    except Exception as e:
        logging.error(f"程序运行出错: {str(e)}")
        print(f"程序出错: {str(e)}")
    finally:
        print("\n按任意键退出...")
        input()

if __name__ == '__main__':
    main()
