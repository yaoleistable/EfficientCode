# -*- coding: UTF-8 -*-
"""
@author: YaoLei
@license: (C) Copyright 2011-2019.
@file: 获取当前文件夹内文件名.py
@time: 2019/10/11 18:20
"""
"""
    脚本编写目的：工作中，有时候需要把一个文件夹中的文件名，导出到文本文件中
    功能：把一个文件夹内文件名，写入到文本文件中，便于使用
"""

import os


# 获取当前文件夹内文件名，保存为文本文件
def get_file_name(file_dir):
    text_name = '当前文件夹内文件名.txt'
    if os.path.exists(text_name):
        os.remove(text_name)
        print('删除成功')

    for root, dirs, files in os.walk(file_dir):
        # print(root)
        # print(dirs)
        # print('文件：', files)
        for file in files:
            print(file)
            # 方法一：
            with open(text_name, "a", encoding='GBK') as f:
                file_name = str(file) + '\n'
                f.writelines(file_name)
            # 方法二：
            # with open("当前文件夹内文件名.txt", "a") as f:
            # for file in files:
            # file_name = str(j)+' '+file+'\n'
            # f.writelines(file_name)
            # 方法三：
            # s = '\n'
            # with open("当前文件夹内文件名.txt", "a") as f:  # 注意需用a模式，用w模式，每次写会清空上一次写的内容
            # f.write(s.join(file))
    print("读取当前文件夹内文件名成功！")


def get_file_name2():
    files = os.listdir(os.getcwd())
    for file in files:
        print(file[12:])


if __name__ == "__main__":
    file_dir = os.getcwd()
    # file_dir = r'/Users/yaolei/PycharmProjects/PythonBase/Python Program Design By Dong Fu Guo'
    print('当前文件夹：', file_dir)
    get_file_name(file_dir)
    # get_file_name2()
