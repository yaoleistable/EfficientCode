# -*- coding: UTF-8 -*-
"""
@author: YaoLei
@license: (C) Copyright 2011-2019.
@file: 文件重命名 文件夹到文件夹.py
@time: 2019/10/12 10:07
"""
import os

"""
    脚本编写目的：工作中，有时候需要把一个文件夹中的文件名，复制出来，重命名给另一个文件夹中类似文件
    功能：以一个文件夹内文件名，重命名另一个文件夹内文件名
    src_path 原文件夹
    dst_path 目标文件夹
"""


def file_rename(src_path, dst_path):
    src_file_name = os.listdir(src_path)
    # print(len(src_file_name))
    # print(src_file_name[0])
    dst_file_name = os.listdir(dst_path)
    # print(dst_file_name[0])
    # old_name = dst_path + '\\' + src_file_name[0]
    # new_name = dst_path + '\\' + dst_file_name[0]
    # os.rename(old_name, new_name)
    i = 0
    for file in dst_file_name:
        print('文件', file)
        old_name = os.path.join(dst_path, dst_file_name[i])
        print('旧文件名', old_name)
        new_name = os.path.join(dst_path, dst_file_name[i][:10] + src_file_name[i][10:])
        print('新文件名', new_name)
        os.rename(old_name, new_name)
        i = i + 1
    print('文件重命名完成')


if __name__ == "__main__":
    src_path = r'C:\Users\yaole\Desktop\二期 建筑\5号楼2'
    dst_path = r'C:\Users\yaole\Desktop\二期 建筑\8号楼'
    file_rename(src_path, dst_path)
