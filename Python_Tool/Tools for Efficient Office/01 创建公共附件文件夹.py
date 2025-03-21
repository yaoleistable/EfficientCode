# -*- coding: UTF-8 -*-
"""
@author: YaoLei
"""

"""
    脚本编写目的：在工作中，有时会遇到需要批量新建文件夹，用于不同项目的归档
    功能：指定文件创建目录、文件名等，脚本按照文件路径，自动批量创建文件夹
"""


def mkdir(path):
    # 引入模块
    import os

    # 去除首位空格
    path = path.strip()
    # 去除尾部 \ 符号
    path = path.rstrip("\\")

    # 判断路径是否存在
    # 存在     True
    # 不存在   False
    isExists = os.path.exists(path)

    # 判断结果
    if not isExists:
        # 如果不存在则创建目录
        # 创建目录操作函数
        os.makedirs(path)

        print(path + ' 创建成功')
        return True
    else:
        # 如果目录存在则不创建，并提示目录已存在
        print(path + ' 目录已存在')
        return False


if __name__ == '__main__':
    # 定义要创建的目录
    project_name = "D:\\金融谷三期"
    path_1 = project_name + "\\1、建筑专业\\1、建筑专业图纸（总平、设计说明、构造做法表、门窗表）\\构造做法表"
    path_2 = project_name + "\\1、建筑专业\\1、建筑专业图纸（总平、设计说明、构造做法表、门窗表）\\门窗表"
    path_3 = project_name + "\\1、建筑专业\\1、建筑专业图纸（总平、设计说明、构造做法表、门窗表）\\设计说明"
    path_4 = project_name + "\\1、建筑专业\\1、建筑专业图纸（总平、设计说明、构造做法表、门窗表）\\总平"
    path_5 = project_name + "\\1、建筑专业\\2、建筑专业图纸（平立剖）"
    path_6 = project_name + "\\1、建筑专业\\3、建筑专业图纸（其他）"
    path_7 = project_name + "\\2、暖通专业\\1、暖通专业图纸（设计说明、设备表、系统图）\\设备表"
    path_8 = project_name + "\\2、暖通专业\\1、暖通专业图纸（设计说明、设备表、系统图）\\设计说明"
    path_9 = project_name + "\\2、暖通专业\\1、暖通专业图纸（设计说明、设备表、系统图）\\系统图"
    path_10 = project_name + "\\2、暖通专业\\2、暖通专业图纸（其他）"
    path_11 = project_name + "\\3、给排水专业\\1、给排水专业图纸（设计说明及系统图）"
    path_12 = project_name + "\\3、给排水专业\\2、给排水专业图纸（其他）"
    path_13 = project_name + "\\4、电气专业\\1、电气专业图纸（设计说明及系统图）"
    path_14 = project_name + "\\4、电气专业\\2、电气专业图纸（照明平面图）"
    path_15 = project_name + "\\4、电气专业\\3、电气专业图纸（其他）"
    path_16 = project_name + "\\5、结构专业\\1、结构专业图纸（设计说明）"
    path_17 = project_name + "\\5、结构专业\\2、结构专业图纸（其他）"
    path_18 = project_name + "\\6、景观专业\\1、景观专业图纸（种植施工图及苗木表）\\苗木表"
    path_19 = project_name + "\\6、景观专业\\1、景观专业图纸（种植施工图及苗木表）\\种植施工图"
    path_20 = project_name + "\\6、景观专业\\2、景观专业图纸（其他）"
    path_list = [path_1, path_2, path_3, path_4, path_4, path_5, path_6, path_7, path_8, path_9, path_10, path_11,
                 path_12, path_13, path_14, path_15, path_16, path_17, path_18, path_19, path_20]
    # 调用函数
    for path in path_list:
        mkdir(path)
