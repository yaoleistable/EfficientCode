# -*- coding: UTF-8 -*-

"""
PDF批量转换工具
将当前目录下的所有PDF文件转换为Word文档

作者：Lei
开发时间：2025-03-14 15:50:01 UTC
版本：1.0.0
联系方式：请通过GitHub Issues反馈问题

使用说明：
1. 将本脚本放在包含PDF文件的目录中
2. 双击运行脚本，程序会自动检查并安装所需依赖
3. 转换完成的Word文件将保存在'converted_word_files'文件夹中
4. 转换日志保存在'conversion_log.txt'文件中

功能特点：
- 自动检测和安装所需依赖包
- 批量转换当前目录下所有PDF文件
- 保持原文件名，仅更改扩展名为.docx
- 详细的转换进度和状态显示
- 完整的错误处理和日志记录

注意事项：
- 请确保有足够的磁盘空间
- 转换大文件可能需要较长时间
- 建议在转换前备份重要文件
"""

import os
import sys
import subprocess
from datetime import datetime
import logging

# 脚本信息
__author__ = 'Lei'
__version__ = '1.0.0'
__date__ = '2025-03-14'

def check_and_install_dependencies():
    """检查并安装所需的依赖包"""
    required_packages = ['pdf2docx', 'tqdm']
    missing_packages = []

    def is_package_installed(package_name):
        try:
            __import__(package_name)
            return True
        except ImportError:
            return False

    print("检查依赖包...")
    for package in required_packages:
        if not is_package_installed(package):
            missing_packages.append(package)
    
    if missing_packages:
        print(f"需要安装以下依赖包: {', '.join(missing_packages)}")
        try:
            for package in missing_packages:
                print(f"正在安装 {package}...")
                subprocess.check_call([sys.executable, "-m", "pip", "install", package])
            print("所有依赖包安装完成！")
        except Exception as e:
            print(f"安装依赖包时出错: {str(e)}")
            print("请手动运行以下命令安装依赖：")
            print(f"pip install {' '.join(missing_packages)}")
            input("按回车键退出...")
            sys.exit(1)
    else:
        print("所有依赖包已安装！")

# 首先检查并安装依赖
check_and_install_dependencies()

# 现在可以安全地导入依赖包
from pdf2docx import Converter
from tqdm import tqdm

def show_welcome_message():
    """显示欢迎信息和使用说明"""
    welcome_text = f"""
PDF批量转换工具 v{__version__}
作者: {__author__}
开发时间: {__date__}
{"="*50}

使用说明：
1. 本工具将自动转换当前目录下的所有PDF文件为Word格式
2. 转换后的文件将保存在'converted_word_files'文件夹中
3. 转换过程的日志将保存在'conversion_log.txt'文件中

操作步骤：
1. 等待程序自动检查并安装必要的依赖包
2. 程序会自动扫描当前目录下的所有PDF文件
3. 转换完成后，可以：
   - 打开输出文件夹查看结果
   - 查看转换日志了解详细信息
   - 选择退出程序

注意事项：
- 转换过程中请勿关闭程序
- 确保有足够的磁盘空间
- 如遇问题，请查看转换日志
{"="*50}
"""
    print(welcome_text)

def setup_logging():
    """设置日志配置"""
    logging.basicConfig(
        level=logging.INFO,
        format='%(asctime)s - %(levelname)s - %(message)s',
        handlers=[
            logging.FileHandler('conversion_log.txt', encoding='utf-8'),
            logging.StreamHandler()
        ]
    )

def convert_pdf_to_word(current_dir='.'):
    """
    转换当前目录下的所有PDF文件为Word文档
    
    Args:
        current_dir (str): 要处理的目录路径，默认为当前目录
    """
    # 设置日志
    setup_logging()
    logger = logging.getLogger(__name__)
    
    # 创建输出目录
    output_dir = os.path.join(current_dir, 'converted_word_files')
    os.makedirs(output_dir, exist_ok=True)
    
    # 获取所有PDF文件
    pdf_files = [f for f in os.listdir(current_dir) if f.lower().endswith('.pdf')]
    
    if not pdf_files:
        logger.warning("当前目录下没有找到PDF文件！")
        return False
    
    logger.info(f"找到 {len(pdf_files)} 个PDF文件待转换")
    
    # 记录转换开始时间
    start_time = datetime.now()
    
    # 使用tqdm创建进度条
    successful_conversions = 0
    failed_conversions = 0
    failed_files = []
    
    for pdf_file in tqdm(pdf_files, desc="转换进度"):
        pdf_path = os.path.join(current_dir, pdf_file)
        # 生成输出文件名（将.pdf替换为.docx）
        word_file = os.path.splitext(pdf_file)[0] + '.docx'
        word_path = os.path.join(output_dir, word_file)
        
        try:
            # 检查文件大小
            file_size = os.path.getsize(pdf_path) / (1024 * 1024)  # 转换为MB
            logger.info(f"开始转换: {pdf_file} (大小: {file_size:.2f}MB)")
            
            # 转换文件
            cv = Converter(pdf_path)
            cv.convert(word_path)
            cv.close()
            
            successful_conversions += 1
            logger.info(f"成功转换: {pdf_file} -> {word_file}")
            
        except Exception as e:
            failed_conversions += 1
            failed_files.append(pdf_file)
            logger.error(f"转换失败 {pdf_file}: {str(e)}")
            continue
            
    # 记录转换结束时间和统计信息
    end_time = datetime.now()
    duration = (end_time - start_time).total_seconds()
    
    # 输出统计信息
    logger.info("\n转换完成！统计信息：")
    logger.info(f"总文件数: {len(pdf_files)}")
    logger.info(f"成功转换: {successful_conversions}")
    logger.info(f"转换失败: {failed_conversions}")
    logger.info(f"总耗时: {duration:.2f} 秒")
    logger.info(f"转换后的文件保存在: {output_dir}")
    
    if failed_files:
        logger.info("\n转换失败的文件:")
        for file in failed_files:
            logger.info(f"- {file}")
    
    return True

def show_menu():
    """显示交互式菜单"""
    while True:
        print("\n" + "="*50)
        print("\n请选择操作:")
        print("1. 打开输出文件夹")
        print("2. 查看转换日志")
        print("3. 退出程序")
        choice = input("\n请输入选项 (1/2/3): ")
        
        if choice == '1':
            try:
                output_dir = os.path.abspath('./converted_word_files')
                if os.path.exists(output_dir):
                    os.startfile(output_dir)  # Windows系统
                else:
                    print("输出文件夹不存在！")
            except Exception as e:
                print(f"无法打开文件夹: {str(e)}")
        elif choice == '2':
            try:
                if os.path.exists('conversion_log.txt'):
                    os.startfile('conversion_log.txt')  # Windows系统
                else:
                    print("日志文件不存在！")
            except Exception as e:
                print(f"无法打开日志文件: {str(e)}")
        elif choice == '3':
            print("\n感谢使用！程序退出...")
            break
        else:
            print("无效的选项，请重新输入！")

def main():
    try:
        show_welcome_message()
        
        print(f"当前用户: {os.getenv('USERNAME', 'Unknown')}")
        print(f"开始时间: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
        print("-" * 50)
        
        # 运行转换程序
        success = convert_pdf_to_word()
        
        if success:
            print("\n转换过程已完成！")
        else:
            print("\n当前目录没有找到PDF文件！")
        
    except KeyboardInterrupt:
        print("\n程序被用户中断")
    except Exception as e:
        print(f"程序执行出错: {str(e)}")
    finally:
        # 显示交互式菜单
        show_menu()

if __name__ == "__main__":
    main()
