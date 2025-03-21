# -*- coding: UTF-8 -*-
"""
@author: yaoleistable
@file: pdf_merger.py
@time: 2025-03-13 10:20:55
@description: 将当前目录下的所有PDF文件合并为一个文件
@version: 1.0.2

使用说明：
1. 需要安装pypdf 5.3.1版本
2. 安装命令：pip install pypdf==5.3.1
3. 查看版本：pip show pypdf
4. 将需要合并的PDF文件放在程序同目录下
5. 运行程序，自动合并所有PDF文件
6. 合并后的文件保存为 merged_output.pdf
"""

import os
import sys
import logging
from datetime import datetime
from typing import List, Optional
from pypdf import PdfReader, PdfWriter

# 配置信息
REQUIRED_PYPDF_VERSION = "5.3.1"
OUTPUT_FILENAME = "merged_output.pdf"
CURRENT_USER = "yaoleistable"
CURRENT_TIME = "2025-03-13 10:20:55"

def setup_logging() -> None:
    """配置日志系统"""
    log_dir = "logs"
    if not os.path.exists(log_dir):
        os.makedirs(log_dir)

    log_file = os.path.join(log_dir, f"pdf_merge_{datetime.now().strftime('%Y%m%d')}.log")

    logging.basicConfig(
        level=logging.INFO,
        format='%(asctime)s - %(levelname)s - %(message)s',
        datefmt='%Y-%m-%d %H:%M:%S',
        handlers=[
            logging.FileHandler(log_file, encoding='utf-8'),
            logging.StreamHandler()
        ]
    )

def check_pypdf_version() -> bool:
    """检查pypdf版本"""
    try:
        import pypdf
        current_version = pypdf.__version__
        if current_version != REQUIRED_PYPDF_VERSION:
            print(f"警告：当前 pypdf 版本 ({current_version}) 与推荐版本 ({REQUIRED_PYPDF_VERSION}) 不符")
            print(f"建议执行: pip install pypdf=={REQUIRED_PYPDF_VERSION}")
            return False
        return True
    except ImportError:
        print("pypdf 未安装")
        return False

def get_pdf_files() -> List[str]:
    """获取当前目录下的所有PDF文件"""
    pdf_files = [f for f in os.listdir() if f.lower().endswith('.pdf')]
    # 按文件名排序，排除输出文件
    return sorted([f for f in pdf_files if f != OUTPUT_FILENAME])

def merge_pdfs(pdf_files: List[str], output_filename: str) -> bool:
    """
    合并PDF文件

    Args:
        pdf_files: PDF文件列表
        output_filename: 输出文件名

    Returns:
        bool: 合并是否成功
    """
    try:
        merger = PdfWriter()
        total_pages = 0

        print(f"\n开始合并 {len(pdf_files)} 个PDF文件...")

        for index, pdf_file in enumerate(pdf_files, 1):
            try:
                with open(pdf_file, 'rb') as file:
                    reader = PdfReader(file)
                    page_count = len(reader.pages)
                    merger.append(fileobj=file)
                    total_pages += page_count
                    print(f"\r正在处理: {pdf_file} ({index}/{len(pdf_files)}) - {page_count}页", end="")

            except Exception as e:
                logging.error(f"处理文件 {pdf_file} 时出错: {str(e)}")
                return False

        # 写入合并后的文件
        with open(output_filename, 'wb') as output:
            merger.write(output)

        print(f"\n\n合并完成！")
        print(f"- 合并文件数: {len(pdf_files)}")
        print(f"- 总页数: {total_pages}")
        print(f"- 输出文件: {output_filename}")

        return True

    except Exception as e:
        logging.error(f"合并PDF时发生错误: {str(e)}")
        return False

def main():
    """主函数"""
    setup_logging()
    print("PDF文件合并工具")
    print("=" * 50)
    print(f"当前用户: {CURRENT_USER}")
    print(f"当前时间: {CURRENT_TIME}")
    print(f"工作目录: {os.getcwd()}")

    try:
        import pypdf
        print(f"pypdf版本: {pypdf.__version__}")
    except ImportError:
        print("pypdf未安装")
    print("=" * 50)

    # 检查pypdf版本
    check_pypdf_version()

    try:
        # 获取PDF文件列表
        pdf_files = get_pdf_files()

        if not pdf_files:
            print("当前目录下没有找到PDF文件")
            return

        # 显示待处理文件
        print("\n待合并的PDF文件:")
        for i, file in enumerate(pdf_files, 1):
            print(f"{i}. {file}")

        # 记录开始时间
        start_time = datetime.now()

        # 执行合并
        if merge_pdfs(pdf_files, OUTPUT_FILENAME):
            end_time = datetime.now()
            duration = end_time - start_time
            print(f"\n处理时间: {duration.total_seconds():.1f} 秒")
        else:
            print("\n合并过程中出现错误，请查看日志文件了解详情")

    except KeyboardInterrupt:
        print("\n程序被用户中断")
    except Exception as e:
        logging.error(f"程序运行出错: {str(e)}")
    finally:
        print("\n按任意键退出...")
        input()

if __name__ == '__main__':
    main()
