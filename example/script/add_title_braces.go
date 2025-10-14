package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 输入文件和输出文件的路径
	inputFile := "D:\\go\\eth-kawa\\example\\script\\input.txt" // 替换为你的输入文件路径
	outputFile := "D:\\go\\eth-kawa\\example\\script\\out.txt"  // 处理后的输出文件

	// 打开输入文件
	inFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("无法打开输入文件: %v\n", err)
		return
	}
	defer inFile.Close()

	// 创建输出文件
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("无法创建输出文件: %v\n", err)
		return
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)
	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	// 逐行处理
	for scanner.Scan() {
		line := scanner.Text()

		// 检查是否是title行
		if strings.Contains(line, "title={") {
			// 将title={替换为title={{
			modifiedLine := strings.Replace(line, "title={", "title={{", 1)
			// 将最后的}替换为}}（但要确保只替换title字段的 closing brace）
			// 先找到最后一个}的位置
			lastBrace := strings.LastIndex(modifiedLine, "}")
			if lastBrace != -1 {
				// 在最后一个}前插入一个}
				modifiedLine = modifiedLine[:lastBrace] + "}" + modifiedLine[lastBrace:]
			}
			line = modifiedLine
		}

		// 写入处理后的行
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Printf("写入文件时出错: %v\n", err)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("读取文件时出错: %v\n", err)
		return
	}

	fmt.Printf("处理完成，结果已保存到 %s\n", outputFile)
}
