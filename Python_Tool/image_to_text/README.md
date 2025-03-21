# Image to Text Converter

这是一个基于AI的图片文字识别工具，可以批量处理指定文件夹下的图片，并将识别结果输出为Markdown文档。

## 功能特点

- 支持自定义AI服务配置（URL、模型、API key）
- 支持批量处理图片文件
- 异步处理提高效率
- 结果输出为Markdown格式

## 使用方法

1. 首先配置 `config.yaml` 文件，设置AI服务参数：
   ```yaml
   ai_service:
     url: "your_ai_service_url"
     model: "your_model_name"
     api_key: "your_api_key"
   ```

2. 配置 `userconfig.yaml` 文件，设置输入输出路径：
   ```yaml
   path_settings:
     input_dir: "path/to/images"
     output_file: "result.md"
   ```

3. 运行程序：
   ```bash
   python main.py
   ```
   或者通过命令行参数指定路径（将覆盖配置文件中的设置）：
   ```bash
   python main.py --input_dir /path/to/images --output_file result.md
   ```

## 依赖安装

```bash
pip install -r requirements.txt
```

## 配置说明

在 `config.yaml` 文件中配置以下参数：

- `url`: AI服务的API端点URL
- `model`: 使用的AI模型名称
- `api_key`: API访问密钥

在 `userconfig.yaml` 文件中配置以下参数：

- `input_dir`: 输入图片文件夹路径
- `output_file`: 输出Markdown文件路径

## 注意事项

- 支持的图片格式：jpg, jpeg, png
- 请确保有足够的磁盘空间存储输出文件
- API密钥请妥善保管，不要泄露