package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/tealeg/xlsx"
)

func main() {
	// 定义命令行参数
	filenamePtr := flag.String("filename", "", "Excel 文件名")

	dirPtr := flag.String("dir", "", "创建文件夹")

	flag.Parse()

	// 检查参数是否为空
	if *filenamePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// 检查参数是否为空
	if *dirPtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if _, err := os.Stat(*dirPtr); os.IsNotExist(err) {
		if err := os.MkdirAll(*dirPtr, 0755); err != nil {
			fmt.Printf("创建文件夹失败: %s\n", err)
			return
		}
	}

	// 打开 Excel 文件
	xlFile, err := xlsx.OpenFile(*filenamePtr)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	// 遍历工作表

	// 遍历所有行
	for _, row := range xlFile.Sheets[0].Rows {
		// 遍历所有单元格

		for _, cell := range row.Cells {

			value, err := cell.FormattedValue()
			if err != nil {
				fmt.Println("Error reading cell value:", err)
				continue
			}

			if isValidUrl(value) {
				fmt.Print(value, "\t")

				downLoadFile(*dirPtr, value)
			}
		}
		fmt.Println()
	}

}

func downLoadFile(dir, url string) {
	// 发送HTTP请求获取图片
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("获取图片失败: %s\n", err)
		return
	}
	defer resp.Body.Close()

	// 创建本地文件并保存图片
	filename := filepath.Base(url)
	filepath := filepath.Join(dir, filename)
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Printf("创建文件失败: %s\n", err)
		return
	}
	defer file.Close()

	// 将HTTP响应体中的数据写入本地文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("保存图片失败: %s\n", err)
		return
	}

	fmt.Println("图片已保存到", filepath)
}

func isValidUrl(urlString string) bool {
	urlRegexp := regexp.MustCompile(`^(http|https):\/\/[^\s]+`)

	// 判断字符串是否符合URL正则表达式
	if urlRegexp.MatchString(urlString) {
		return true
	} else {
		return false
	}
}
