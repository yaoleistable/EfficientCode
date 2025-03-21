# -*- coding: UTF-8 -*-
"""
@author: YaoLei
@license: (C) Copyright 2011-2021.
@file: 07 文件重命名 前5个字符.py
@time: 2019/10/11 18:20
"""
"""
    脚本编写目的：工作中，有时候需要把一个文件夹中的文件前几个字符删除，添加对应序号
    
"""
# alt+shift+f 代码格式化




import os
def file_rename():
    file_dir = os.getcwd()
    i = 1
    for root, dirs, files in os.walk(file_dir):
        for file in files:
            # print("旧文件名:"+file)
            file_name = os.path.splitext(file)[0]
            file_ext = os.path.splitext(file)[1]
            # print("文件名："+file_name)
            # print("拓展名："+file_ext)
            # 新文件名为：去掉前7个字符，增加数字和下划线，去掉结尾14个字符

            if file_ext == ".docx":
                new_file_name = str(i)+"_"+file[6:-13]+file_ext
                print("新文件名："+new_file_name)
                os.rename(file, new_file_name)
                i = i + 1
            elif file_ext == ".doc":
                new_file_name = str(i)+"_"+file[6:-12]+file_ext
                print("新文件名："+new_file_name)
                os.rename(file, new_file_name)
                i = i + 1


def get_file_name():
    files = os.listdir(os.getcwd())
    for file in files:
        print(file)


if __name__ == "__main__":
    file_rename()
    get_file_name()
