import os
import asyncio
import argparse
import yaml
from typing import List, Dict
from pathlib import Path
import aiohttp
import aiofiles
from google import genai
import base64

class ImageToTextConverter:
    def __init__(self, config_path: str = 'config.yaml', user_config_path: str = 'userconfig.yaml'):
        self.config = self._load_config(config_path)
        self.ai_service = self.config['ai_service']
        self.image_processing = self.config['image_processing']
        self.path_settings = self._load_user_config(user_config_path)

    def _load_config(self, config_path: str) -> Dict:
        with open(config_path, 'r', encoding='utf-8') as f:
            return yaml.safe_load(f)

    def _load_user_config(self, user_config_path: str) -> Dict:
        try:
            with open(user_config_path, 'r', encoding='utf-8') as f:
                return yaml.safe_load(f).get('path_settings', {})
        except FileNotFoundError:
            return {}

    async def process_image(self, image_path: str) -> str:
        try:
            client = genai.Client(api_key=self.ai_service['api_key'])
            with open(image_path, 'rb') as f:
                image_data = f.read()
                image_base64 = base64.b64encode(image_data).decode('utf-8')
            
            response = client.models.generate_content(
                model=self.ai_service['model'],
                contents=[{"inlineData": {
                    "mimeType": "image/jpeg",
                    "data": image_base64
                }}]
            )
            
            if response.text:
                return response.text
            else:
                return f'Error processing {image_path}: No text generated'
        except Exception as e:
            return f'Error processing {image_path}: {str(e)}'

    def is_supported_format(self, file_path: str) -> bool:
        return any(file_path.lower().endswith(fmt) for fmt in self.image_processing['supported_formats'])

    async def process_directory(self, input_dir: str = None, output_file: str = None):
        # 优先使用参数传入的路径，如果没有则使用配置文件中的路径
        input_dir = input_dir or self.path_settings.get('input_dir')
        output_file = output_file or self.path_settings.get('output_file')

        if not input_dir or not output_file:
            raise ValueError('Input directory and output file must be specified either in userconfig.yaml or as arguments')

        input_path = Path(input_dir)
        if not input_path.exists():
            raise FileNotFoundError(f'Input directory {input_dir} does not exist')

        image_files = [str(f) for f in input_path.glob('**/*') if f.is_file() and self.is_supported_format(str(f))]
        if not image_files:
            print('No supported image files found')
            return

        print(f'Found {len(image_files)} images to process')
        semaphore = asyncio.Semaphore(self.image_processing['max_concurrent'])

        async def process_with_semaphore(image_path: str):
            async with semaphore:
                return await self.process_image(image_path)

        tasks = [process_with_semaphore(image) for image in image_files]
        results = await asyncio.gather(*tasks)

        async with aiofiles.open(output_file, 'w', encoding='utf-8') as f:
            await f.write('# 图片文字识别结果\n\n')
            for image_path, text in zip(image_files, results):
                relative_path = os.path.relpath(image_path, input_dir)
                await f.write(f'## {relative_path}\n\n{text}\n\n')

def main():
    parser = argparse.ArgumentParser(description='Convert images to text using AI')
    parser.add_argument('--input_dir', help='Directory containing images (optional if specified in userconfig.yaml)')
    parser.add_argument('--output_file', help='Output markdown file path (optional if specified in userconfig.yaml)')
    parser.add_argument('--config', default='config.yaml', help='AI service configuration file path')
    parser.add_argument('--user_config', default='userconfig.yaml', help='User configuration file path')

    args = parser.parse_args()
    converter = ImageToTextConverter(args.config, args.user_config)

    try:
        asyncio.run(converter.process_directory(args.input_dir, args.output_file))
        print(f'Processing complete. Results saved to {args.output_file or converter.path_settings.get("output_file")}')
    except Exception as e:
        print(f'Error: {str(e)}')

if __name__ == '__main__':
    main()