# -*- coding: UTF-8 -*-
"""
@author: YaoLei
@license: (C) Copyright 2011-2019.
@file: 文件重命名 从文本文件.py
@time: 2019/10/12 12:45
"""
"""
    脚本编写目的：工作中，有时候需要把一个文本文件中的名字，对应的命名给文件夹中的文件，
        如CAD图纸导出的图纸命名，可截图图纸目录，用QQ识别为文字后，粘贴到文本文件中，用本脚本读取内容，命令文件夹内PDF图纸
    功能：把文本文件文件名，批量重命名给对应PDF文件
"""


# 文件重命名，从文本文件
def file_rename(txt_path, file_dir):
    import os
    with open(txt_path, 'r', encoding='UTF-8') as f:
        s = f.readlines()
        print(s)
    file_name = os.listdir(file_dir)
    print(file_name[0])
    for i in range(len(s)):
        old_name = os.path.join(file_dir, file_name[i])

        new_name = old_name[0:-4:1] + ' ' + s[i].strip() + '.pdf'
        print('旧文件:', old_name)
        print('新文件:', new_name)
        i = i + 1
        os.rename(old_name, new_name)


if __name__ == "__main__":
    txt_path = r'C:\Users\yaole\Desktop\result.txt'
    file_dir = r'C:\Users\yaole\Desktop\PDF'
    file_rename(txt_path, file_dir)
