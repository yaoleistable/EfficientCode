# PDF工具集

一个简单易用的PDF工具集，支持PDF文件的合并和页面提取功能。

## 功能特点

- 📄 PDF文件合并
  - 支持批量合并指定目录下的所有PDF文件
  - 按文件名自然排序
  - 自动生成合并后的文件

- 📑 PDF页面提取
  - 支持按页码范围提取
  - 批量处理目录下的PDF文件
  - 灵活的输出目录配置

## 安装要求

- Go 1.16+
- Windows 7/10/11
- 足够的磁盘空间（建议源文件总大小的2倍）

## 快速开始

### 构建程序

```bash
# 克隆项目
git clone [your-repository-url]
cd pdftools

# 初始化模块
go mod init pdftools
go mod tidy

# 构建（标准版本）
go build -o pdftools.exe

# 构建（优化大小）
go build -ldflags="-s -w" -o pdftools.exe
```

### 使用方法

1. **合并PDF文件**

```bash
# 合并当前目录下的所有PDF
.\pdftools merge

# 合并指定目录下的PDF
.\pdftools merge -dir "D:\文档\待合并"
```

2. **提取PDF页面**

```bash
# 从指定目录提取页面
.\pdftools extract -input "D:\文档\源文件" -output "D:\文档\输出" -pages "1,3-5"
```

## 使用提示

### 文件命名建议

为确保合并顺序正确，建议使用数字前缀：

```
01-第一章.pdf
02-第二章.pdf
03-第三章.pdf
```

### 注意事项

- 使用前备份重要文件
- 确保有足够的磁盘空间
- 检查文件和目录的访问权限
- 关闭可能占用PDF文件的程序

## 常见问题

### 1. 程序无法运行
- 以管理员身份运行
- 检查杀毒软件设置
- 使用完整路径运行程序

### 2. 找不到PDF文件
- 确认文件扩展名（.pdf）
- 检查目录路径
- 验证文件访问权限

### 3. 合并/提取失败
- 检查PDF文件完整性
- 确认目录写入权限
- 关闭占用PDF的程序

## 项目结构

```
pdftools/
├── pdf/
│   ├── merger.go    # PDF合并功能
│   └── extractor.go # PDF提取功能
├── main.go          # 程序入口
├── go.mod          # 项目依赖
└── README.md       # 项目文档
```

## 贡献指南

欢迎提交问题和改进建议！

1. Fork 项目
2. 创建新分支 (`git checkout -b feature/improvement`)
3. 提交更改 (`git commit -am '添加新功能'`)
4. 推送分支 (`git push origin feature/improvement`)
5. 创建 Pull Request

## 技术支持

如遇问题，请提供：

- 完整的错误信息
- 操作系统版本
- 文件目录结构
- 执行的具体命令

## 许可证

[MIT License](LICENSE)