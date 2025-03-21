# -*- coding: UTF-8 -*-
"""
@author: Lei
@file: 05 批量提取PDF文件封面V3.py
@time: 2025/03/13 18:13:41
@description: 批量提取PDF文件的第一页（封面）并另存为新的PDF文件
@version: 3.0

使用说明：
0.与05 批量提取PDF文件封面V2.py相比，v3版本使用了最新的pypdf 5.3.1，更新更好。
1. 需要安装pypdf 5.3.1版本
2. 安装命令：pip install pypdf==5.3.1
3. 查看版本：pip show pypdf
4. 将需要处理的PDF文件放在程序同目录下
5. 运行程序，自动处理所有PDF文件
6. 处理后的文件保存在"PDF拆页"目录中
"""

import os
import sys
import logging
from typing import Optional, Union, Any, Tuple
from datetime import datetime

# 指定pypdf版本
REQUIRED_PYPDF_VERSION = "5.3.1"

def setup_logging() -> None:
    """配置日志系统"""
    log_dir = "logs"
    if not os.path.exists(log_dir):
        os.makedirs(log_dir)

    log_file = os.path.join(log_dir, f"pdf_split_{datetime.now().strftime('%Y%m%d')}.log")

    logging.basicConfig(
        level=logging.INFO,
        format='%(asctime)s - %(levelname)s - %(message)s',
        datefmt='%Y-%m-%d %H:%M:%S',
        handlers=[
            logging.FileHandler(log_file, encoding='utf-8'),
            logging.StreamHandler()
        ]
    )

class PDFHandler:
    """PDF处理器类"""
    def __init__(self):
        self.Reader = None
        self.Writer = None
        self._init_pdf_components()

    def _init_pdf_components(self) -> bool:
        """初始化PDF处理组件"""
        try:
            from pypdf import PdfReader, PdfWriter
            self.Reader = PdfReader
            self.Writer = PdfWriter
            return True
        except ImportError:
            return False

    def create_reader(self, file_obj) -> Any:
        """创建PDF读取器"""
        return self.Reader(file_obj)

    def create_writer(self) -> Any:
        """创建PDF写入器"""
        return self.Writer()

    def get_page_count(self, reader) -> int:
        """获取PDF页数"""
        return len(reader.pages)

    def get_page(self, reader, page_num: int) -> Any:
        """获取指定页面"""
        return reader.pages[page_num]

    def add_page(self, writer, page) -> None:
        """添加页面到写入器"""
        writer.add_page(page)

def check_pypdf_version() -> bool:
    """检查pypdf版本"""
    try:
        import pypdf
        current_version = pypdf.__version__
        if current_version != REQUIRED_PYPDF_VERSION:
            print(f"当前 pypdf 版本 ({current_version}) 与所需版本 ({REQUIRED_PYPDF_VERSION}) 不符")
            return False
        return True
    except ImportError:
        return False

def install_dependencies() -> bool:
    """安装必要的依赖包"""
    if check_pypdf_version():
        return True

    print(f"正在安装 pypdf {REQUIRED_PYPDF_VERSION}...")
    try:
        import pip
        pip.main(['install', f'pypdf=={REQUIRED_PYPDF_VERSION}'])

        # 重新加载以确保使用新版本
        import importlib
        import pypdf
        importlib.reload(pypdf)

        if pypdf.__version__ != REQUIRED_PYPDF_VERSION:
            raise ImportError(f"无法安装指定版本 {REQUIRED_PYPDF_VERSION}")

        print(f"pypdf {pypdf.__version__} 安装成功！")
        return True
    except Exception as e:
        logging.error(f"安装 pypdf 失败: {str(e)}")
        print(f"请手动执行: pip install pypdf=={REQUIRED_PYPDF_VERSION}")
        return False

def create_output_directory(output_path: str) -> bool:
    """创建输出目录"""
    try:
        if not os.path.exists(output_path):
            os.makedirs(output_path)
        return True
    except Exception as e:
        logging.error(f"创建输出目录失败: {str(e)}")
        return False

def split_pdf(infile: str, output_dir: str = 'PDF拆页') -> bool:
    """
    拆分PDF文件的第一页（封面）

    Args:
        infile: PDF文件路径
        output_dir: 输出目录名

    Returns:
        bool: 操作是否成功
    """
    try:
        # 确保输入文件存在
        if not os.path.isfile(infile):
            logging.error(f"文件不存在: {infile}")
            return False

        path = os.getcwd()
        file_name = os.path.splitext(os.path.basename(infile))[0]
        out_path = os.path.join(path, output_dir)

        if not create_output_directory(out_path):
            return False

        handler = PDFHandler()
        if handler.Reader is None or handler.Writer is None:
            logging.error("无法初始化PDF处理组件")
            return False

        with open(infile, 'rb') as pdf_file:
            try:
                reader = handler.create_reader(pdf_file)
                writer = handler.create_writer()

                if handler.get_page_count(reader) == 0:
                    logging.error(f"文件 {infile} 是空的PDF文件")
                    return False

                # 获取并添加第一页
                first_page = handler.get_page(reader, 0)
                handler.add_page(writer, first_page)

                out_file_name = os.path.join(out_path, f"{file_name}_1.pdf")

                with open(out_file_name, 'wb') as outfile:
                    writer.write(outfile)

                logging.info(f"成功提取封面: {infile}")
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
            print(f"\r正在处理: {file} ({index}/{total_files})", end="")
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
            print(f"\n提取的PDF封面已保存到: '{os.path.join(os.getcwd(), 'PDF拆页')}' 目录")

    except Exception as e:
        logging.error(f"批量处理时发生错误: {str(e)}")

def main():
    """主函数"""
    setup_logging()
    print("PDF封面提取工具")
    print("=" * 50)
    print(f"当前用户: {os.getenv('USERNAME', 'yaoleistable')}")
    print(f"当前时间: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
    print(f"工作目录: {os.getcwd()}")

    try:
        import pypdf
        print(f"pypdf版本: {pypdf.__version__}")
    except ImportError:
        print("pypdf未安装")
    print("=" * 50)

    if not install_dependencies():
        print("程序初始化失败，请手动安装指定版本的pypdf")
        print(f"执行命令: pip install pypdf=={REQUIRED_PYPDF_VERSION}")
        print("\n按任意键退出...")
        input()
        return

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
