# -*- coding: UTF-8 -*-
"""
PDF文件批量拆分工具
功能：将PDF文件按页拆分成单页PDF文件
用途：适用于需要将多页PDF文档分割为单页文件的场景
作者：yaoleistable
日期：2025-03-13
"""

import os
import sys
import logging
from datetime import datetime
from typing import Optional

REQUIRED_PYPDF2_VERSION = "3.0.1"

def install_pypdf2() -> bool:
    """安装指定版本的 PyPDF2"""
    try:
        import pip
        print(f"正在安装 PyPDF2 {REQUIRED_PYPDF2_VERSION}...")
        pip.main(['install', f'PyPDF2=={REQUIRED_PYPDF2_VERSION}'])
        return True
    except Exception as e:
        print(f"安装失败: {e}")
        print(f"请手动执行: pip install PyPDF2=={REQUIRED_PYPDF2_VERSION}")
        return False

# 尝试导入PyPDF2，如果失败则安装指定版本
try:
    from PyPDF2 import PdfReader, PdfWriter
    import PyPDF2
    current_version = PyPDF2.__version__
    if current_version != REQUIRED_PYPDF2_VERSION:
        print(f"当前 PyPDF2 版本 ({current_version}) 与所需版本 ({REQUIRED_PYPDF2_VERSION}) 不符")
        if install_pypdf2():
            # 重新导入以确保使用新安装的版本
            import importlib
            importlib.reload(PyPDF2)
            from PyPDF2 import PdfReader, PdfWriter
        else:
            sys.exit(1)
except ImportError:
    if not install_pypdf2():
        print("PyPDF2 安装失败")
        input("\n按任意键退出...")
        sys.exit(1)
    from PyPDF2 import PdfReader, PdfWriter

def setup_logging() -> None:
    """配置日志系统"""
    log_dir = "logs"
    if not os.path.exists(log_dir):
        os.makedirs(log_dir)

    log_file = os.path.join(log_dir, f"pdf_split_{datetime.now().strftime('%Y%m%d')}.log")

    logging.basicConfig(
        level=logging.INFO,
        format='%(asctime)s - %(levelname)s - %(message)s',
        handlers=[
            logging.FileHandler(log_file, encoding='utf-8'),
            logging.StreamHandler()
        ]
    )

def split_pdf(infile: str) -> bool:
    """
    将PDF文件按页拆分为单页PDF文件

    Args:
        infile: 待拆分的PDF文件路径

    Returns:
        bool: 拆分是否成功
    """
    try:
        # 检查输入文件是否存在
        if not os.path.exists(infile):
            logging.error(f"文件不存在: {infile}")
            return False

        path = os.getcwd()
        file_name = os.path.splitext(infile)[0]
        out_path = os.path.join(path, 'PDF拆页')

        # 创建输出目录
        if not os.path.exists(out_path):
            os.makedirs(out_path)

        with open(infile, 'rb') as pdf_file:
            try:
                reader = PdfReader(pdf_file)

                if len(reader.pages) == 0:
                    logging.error(f"文件 {infile} 是空的PDF文件")
                    return False

                for i in range(len(reader.pages)):
                    writer = PdfWriter()
                    writer.add_page(reader.pages[i])
                    out_file_name = os.path.join(out_path, f"{file_name}_{i + 1}.pdf")

                    with open(out_file_name, 'wb') as outfile:
                        writer.write(outfile)

                    print(f"\r正在处理: {infile} - 第 {i + 1}/{len(reader.pages)} 页", end="")

                print(f"\n{infile} 拆分完成！共处理 {len(reader.pages)} 页")
                return True

            except Exception as e:
                logging.error(f"处理文件 {infile} 时出错: {str(e)}")
                return False

    except Exception as e:
        logging.error(f"处理文件 {infile} 时发生错误: {str(e)}")
        return False

def batch_split_files() -> None:
    """批量拆分当前目录下的所有PDF文件"""
    success_count = 0
    failed_count = 0
    failed_files = []
    start_time = datetime.now()

    try:
        pdf_files = [f for f in os.listdir() if f.lower().endswith('.pdf')]

        if not pdf_files:
            print("当前目录下没有找到PDF文件")
            return

        total_files = len(pdf_files)
        print(f"找到 {total_files} 个PDF文件")

        for index, file in enumerate(pdf_files, 1):
            print(f"\n[{index}/{total_files}] 正在处理: {file}")
            if split_pdf(file):
                success_count += 1
            else:
                failed_count += 1
                failed_files.append(file)

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
            print(f"\n拆分后的PDF文件已保存到: '{os.path.join(os.getcwd(), 'PDF拆页')}' 目录")

    except Exception as e:
        logging.error(f"批量处理时发生错误: {str(e)}")

def main():
    """主函数"""
    setup_logging()
    print("PDF文件批量拆分工具")
    print("=" * 50)
    print(f"当前用户: {os.getenv('USERNAME', 'Unknown')}")
    print(f"当前时间: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
    print(f"工作目录: {os.getcwd()}")
    print(f"PyPDF2版本: {PyPDF2.__version__}")
    print("=" * 50)

    try:
        batch_split_files()
    except KeyboardInterrupt:
        print("\n程序被用户中断")
    except Exception as e:
        logging.error(f"程序运行出错: {str(e)}")
    finally:
        print("\n按任意键退出...")
        input()

if __name__ == '__main__':
    main()
