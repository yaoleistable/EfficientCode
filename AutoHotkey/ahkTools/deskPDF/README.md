# PDF 工具箱

这是一个纯浏览器端的 PDF 工具集合,包含以下工具:

## PDFSplitterPro.html
PDF 分割工具 - 从多个 PDF 文件中灵活提取指定页面

### 功能特点
- 支持多文件批量处理
- 多种页面选择模式:单页/范围/自定义
- 自定义输出文件名格式
- 支持移除元数据
- 纯浏览器端处理,文件不会上传到服务器
- 支持拖放操作

### 如何运行
1. 双击`PDFSplitterPro.html`文件直接运行
2. 使用 Python 的简易 HTTP 服务器（推荐）
```bash
cd f:\Code\EfficientCode\AutoHotkey\ahkTools\deskPDF
python -m http.server 8080
# 然后在浏览器中访问： http://localhost:8080/PDFSplitterPro.html
```
3. 使用 Node.js 的 http-server（推荐）
```bash
npm install -g http-server # 安装 http-server
cd f:\Code\EfficientCode\AutoHotkey\ahkTools\deskPDF
http-server
# 然后在浏览器中访问： http://localhost:8080/PDFSplitterPro.html
```

## PDFMergerPro.html
PDF 合并工具 - 轻松合并多个 PDF 文件为一个文档

### 功能特点
- 支持多文件合并
- 自定义合并顺序(正常/反向/交替)
- 自定义输出文件名
- 支持移除元数据
- 纯浏览器端处理,文件不会上传到服务器
- 支持拖放操作

### 如何运行
1. 双击`PDFMergerPro.html`文件直接运行
2. 使用 Python 的简易 HTTP 服务器（推荐）
```bash
cd f:\Code\EfficientCode\AutoHotkey\ahkTools\deskPDF
python -m http.server 8080
# 然后在浏览器中访问： http://localhost:8080/PDFMergerPro.html
```
3. 使用 Node.js 的 http-server（推荐）
```bash
npm install -g http-server # 安装 http-server
cd f:\Code\EfficientCode\AutoHotkey\ahkTools\deskPDF
http-server
# 然后在浏览器中访问： http://localhost:8080/PDFMergerPro.html
```

## 为什么推荐使用 HTTP 服务器?
直接打开 HTML 文件可能会受到浏览器安全限制,影响某些功能(如文件夹选择)。使用 HTTP 服务器可以:
- 避免浏览器安全限制
- 获得更好的文件操作体验
- 支持更多高级功能

## 技术栈
- PDF-Lib.js: PDF 文件处理
- TailwindCSS: 界面样式
- FileSaver.js: 文件保存
- Font Awesome: 图标

## 注意事项
- 推荐使用最新版本的 Chrome、Firefox 或 Edge 浏览器
- 所有操作均在浏览器本地完成,不会上传文件到服务器
- 大文件处理可能需要较长时间,请耐心等待

