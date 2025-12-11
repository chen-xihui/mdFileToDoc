# mdFileToDoc

一个将文件夹内所有markdown文件转换为docx格式的Go程序。

## 功能特点

- 遍历输入目录中的所有文件
- 将`.md`文件转换为`.docx`格式
- 复制其他格式的文件到输出目录
- 保持原始目录结构

## 使用方法

### 编译程序

```bash
go build -o mdFileToDoc.exe
```

### 运行程序

```bash
# 指定输入目录，使用默认输出目录（输入目录名+_docx）
./mdFileToDoc.exe -i input_directory

# 指定输入和输出目录
./mdFileToDoc.exe -i input_directory -o output_directory
```

## 注意事项

- 当前实现将markdown内容直接保存为.txt格式并修改扩展名为.docx
- 生成的文件可以被Microsoft Word打开，但不是严格意义上的docx格式
- 如需更完善的docx格式支持，建议使用更专业的库或工具

## 依赖

- Go标准库

## 许可证

MIT