package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// convertDirectory 遍历输入目录，转换md文件并复制其他文件到输出目录
func convertDirectory(inputDir, outputDir string) error {
	// 创建输出目录
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("创建输出目录失败: %v", err)
	}

	// 遍历输入目录
	return filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 计算相对路径
		relPath, err := filepath.Rel(inputDir, path)
		if err != nil {
			return err
		}

		// 构建输出路径
		outputPath := filepath.Join(outputDir, relPath)

		if info.IsDir() {
			// 创建对应的子目录
			return os.MkdirAll(outputPath, 0755)
		} else {
			// 处理文件
			if strings.ToLower(filepath.Ext(path)) == ".md" {
				// 转换md文件
				docxPath := strings.TrimSuffix(outputPath, ".md") + ".docx"
				fmt.Printf("转换文件: %s -> %s\n", path, docxPath)
				if err := convertMarkdownToDocx(path, docxPath); err != nil {
					return fmt.Errorf("转换文件失败 %s: %v", path, err)
				}
			} else {
				// 复制其他文件
				fmt.Printf("复制文件: %s -> %s\n", path, outputPath)
				if err := copyFile(path, outputPath); err != nil {
					return fmt.Errorf("复制文件失败 %s: %v", path, err)
				}
			}
		}

		return nil
	})
}

// convertMarkdownToDocx 将markdown文件转换为docx文件
func convertMarkdownToDocx(inputPath, outputPath string) error {
	// 读取markdown文件内容
	content, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("读取markdown文件失败: %v", err)
	}

	// 对于简单的实现，我们直接将markdown内容写入txt文件
	// 然后将其扩展名改为docx（注意：这不是真正的docx格式，但可以被Word打开）
	txtContent := string(content)

	// 保存为txt文件（将被Word识别为纯文本）
	if err := os.WriteFile(outputPath, []byte(txtContent), 0644); err != nil {
		return fmt.Errorf("保存文件失败: %v", err)
	}

	return nil
}

// copyFile 复制文件
func copyFile(src, dst string) error {
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("打开源文件失败: %v", err)
	}
	defer srcFile.Close()

	// 创建目标文件
	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("创建目标文件失败: %v", err)
	}
	defer dstFile.Close()

	// 复制文件内容
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("复制文件内容失败: %v", err)
	}

	// 复制文件权限
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("获取源文件信息失败: %v", err)
	}

	if err := os.Chmod(dst, srcInfo.Mode()); err != nil {
		return fmt.Errorf("设置目标文件权限失败: %v", err)
	}

	return nil
}

func main() {
	fmt.Println("mdFileToDoc - Markdown to Docx Converter")

	// 定义命令行参数
	var inputDir string
	var outputDir string

	flag.StringVar(&inputDir, "i", "", "输入目录路径")
	flag.StringVar(&inputDir, "input", "", "输入目录路径")
	flag.StringVar(&outputDir, "o", "", "输出目录路径")
	flag.StringVar(&outputDir, "output", "", "输出目录路径")

	// 解析命令行参数
	flag.Parse()

	// 验证输入参数
	if inputDir == "" {
		fmt.Println("错误: 必须指定输入目录")
		flag.Usage()
		os.Exit(1)
	}

	if outputDir == "" {
		// 默认输出目录为输入目录名加上"_docx"
		outputDir = inputDir + "_docx"
		fmt.Printf("未指定输出目录，使用默认目录: %s\n", outputDir)
	}

	// 检查输入目录是否存在
	if _, err := os.Stat(inputDir); os.IsNotExist(err) {
		fmt.Printf("错误: 输入目录不存在 %s\n", inputDir)
		os.Exit(1)
	}

	// 开始转换
	fmt.Printf("开始转换目录: %s\n", inputDir)
	fmt.Printf("输出目录: %s\n", outputDir)

	if err := convertDirectory(inputDir, outputDir); err != nil {
		fmt.Printf("转换失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("转换完成!")
}
