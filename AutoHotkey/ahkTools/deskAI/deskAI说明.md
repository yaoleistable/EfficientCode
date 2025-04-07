deskAI说明
## 1. 构建程序
```bash
# 克隆项目
git clone [your-repository-url]
cd deskAI

# 初始化模块
go mod init deskAI
go mod tidy

# 构建（标准版本）
go build
go build -o deskAI.exe
# 构建（优化大小）
go build -ldflags="-s -w" -o deskAI.exe
```

## 2. 使用方法
```bash
# 发送笔记到 DinoAI
deskAI.exe dinoxPost "这是一条测试笔记"
# 查看帮助
deskAI.exe help
```