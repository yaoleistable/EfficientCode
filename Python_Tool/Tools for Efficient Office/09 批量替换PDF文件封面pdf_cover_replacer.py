# -*- coding: UTF-8 -*-
"""
@author: yaoleistable
@file: pdf_cover_replacer.py
@time: 2025-03-18 08:37:31
@description: 批量替换PDF文件封面，适用于扫描了一堆文件封面，要替换掉盖章封面的场景
@version: 1.0.0

使用说明：
1. 需要安装pypdf 5.3.1版本
2. 安装命令：pip install pypdf==5.3.1
3. A文件夹：存放原始PDF文件
4. B文件夹：存放对应的封面文件，封面文件顺序需要与A的一致，名称不一致没关系
5. 处理后的文件保存在 "输出文件" 目录中
"""

import os
import sys
import logging
from datetime import datetime
from typing import List, Tuple, Optional
from pypdf import PdfReader, PdfWriter

# 配置信息
REQUIRED_PYPDF_VERSION = "5.3.1"
CURRENT_USER = "yaoleistable"
CURRENT_TIME = "2025-03-18 08:37:31"

def setup_logging() -> None:
    """配置日志系统"""
    log_dir = "logs"
    if not os.path.exists(log_dir):
        os.makedirs(log_dir)
    
    log_file = os.path.join(log_dir, f"pdf_cover_replace_{datetime.now().strftime('%Y%m%d')}.log")
    
    logging.basicConfig(
        level=logging.INFO,
        format='%(asctime)s - %(levelname)s - %(message)s',
        datefmt='%Y-%m-%d %H:%M:%S',
        handlers=[
            logging.FileHandler(log_file, encoding='utf-8'),
            logging.StreamHandler()
        ]
    )

def get_pdf_files(folder: str) -> List[str]:
    """获取指定文件夹下的所有PDF文件"""
    if not os.path.exists(folder):
        return []
    return sorted([f for f in os.listdir(folder) if f.lower().endswith('.pdf')])

def replace_first_page(original_pdf: str, cover_pdf: str, output_pdf: str) -> bool:
    """
    替换PDF文件的第一页
    
    Args:
        original_pdf: 原始PDF文件路径
        cover_pdf: 封面PDF文件路径
        output_pdf: 输出PDF文件路径
    
    Returns:
        bool: 替换是否成功
    """
    try:
        # 读取原始文件
        with open(original_pdf, 'rb') as file:
            original_reader = PdfReader(file)
            if len(original_reader.pages) == 0:
                logging.error(f"文件 {original_pdf} 是空的PDF文件")
                return False
            
            # 读取封面文件
            with open(cover_pdf, 'rb') as cover_file:
                cover_reader = PdfReader(cover_file)
                if len(cover_reader.pages) == 0:
                    logging.error(f"封面文件 {cover_pdf} 是空的PDF文件")
                    return False
                
                # 创建新的PDF
                writer = PdfWriter()
                
                # 添加新封面
                writer.add_page(cover_reader.pages[0])
                
                # 添加原始文件的其余页面
                for i in range(1, len(original_reader.pages)):
                    writer.add_page(original_reader.pages[i])
                
                # 保存新文件
                with open(output_pdf, 'wb') as output_file:
                    writer.write(output_file)
                
                return True
                
    except Exception as e:
        logging.error(f"处理文件时发生错误: {str(e)}")
        return False

def process_pdf_files(folder_a: str, folder_b: str, output_folder: str) -> Tuple[int, int, List[str]]:
    """
    处理两个文件夹中的PDF文件
    
    Returns:
        Tuple[int, int, List[str]]: (成功数, 失败数, 失败文件列表)
    """
    success_count = 0
    failed_count = 0
    failed_files = []
    
    # 获取文件列表
    original_pdfs = get_pdf_files(folder_a)
    cover_pdfs = get_pdf_files(folder_b)
    
    if not original_pdfs:
        print(f"错误：在 {folder_a} 中没有找到PDF文件")
        return 0, 0, []
        
    if not cover_pdfs:
        print(f"错误：在 {folder_b} 中没有找到PDF文件")
        return 0, 0, []
        
    if len(original_pdfs) != len(cover_pdfs):
        print(f"警告：A文件夹有 {len(original_pdfs)} 个文件，B文件夹有 {len(cover_pdfs)} 个文件，数量不匹配")
    
    # 创建输出目录
    if not os.path.exists(output_folder):
        os.makedirs(output_folder)
    
    # 处理每个文件
    total_files = min(len(original_pdfs), len(cover_pdfs))
    for i in range(total_files):
        original_pdf = os.path.join(folder_a, original_pdfs[i])
        cover_pdf = os.path.join(folder_b, cover_pdfs[i])
        output_pdf = os.path.join(output_folder, original_pdfs[i])
        
        print(f"\r正在处理: {original_pdfs[i]} ({i+1}/{total_files})", end="")
        
        if replace_first_page(original_pdf, cover_pdf, output_pdf):
            success_count += 1
        else:
            failed_count += 1
            failed_files.append(original_pdfs[i])
    
    return success_count, failed_count, failed_files

def main():
    """主函数"""
    setup_logging()
    print("PDF封面替换工具")
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
    
    # 设置文件夹路径
    folder_a = r"G:\Lei\Downloads\新建文件夹\A"  # 原始PDF文件夹
    folder_b = r"G:\Lei\Downloads\新建文件夹\PDF拆页"  # 封面PDF文件夹
    output_folder = "输出文件"
    
    # 检查文件夹是否存在
    if not os.path.exists(folder_a):
        print(f"错误：找不到文件夹 {folder_a}")
        return
    
    if not os.path.exists(folder_b):
        print(f"错误：找不到文件夹 {folder_b}")
        return
    
    try:
        start_time = datetime.now()
        
        # 处理文件
        success_count, failed_count, failed_files = process_pdf_files(folder_a, folder_b, output_folder)
        
        end_time = datetime.now()
        duration = end_time - start_time
        
        print("\n" + "=" * 50)
        print(f"处理完成！")
        print(f"处理时间: {duration.total_seconds():.1f} 秒")
        print(f"成功: {success_count} 个文件")
        print(f"失败: {failed_count} 个文件")
        
        if failed_files:
            print("\n处理失败的文件:")
            for file in failed_files:
                print(f"- {file}")
        
        if success_count > 0:
            print(f"\n处理后的文件已保存到: '{os.path.join(os.getcwd(), output_folder)}' 目录")
        
    except KeyboardInterrupt:
        print("\n程序被用户中断")
    except Exception as e:
        logging.error(f"程序运行出错: {str(e)}")
    finally:
        print("\n按任意键退出...")
        input()

if __name__ == '__main__':
    main()
