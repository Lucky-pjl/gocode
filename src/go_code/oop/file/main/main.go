package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func test() {
	file, err := os.Open("e:/go/gocode/src/go_code/oop/file/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	defer file.Close()

	// 创建一个*Reader,默认缓冲区为4096
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束一次
		fmt.Print(str)
		if err == io.EOF {
			break
		}
	}
	fmt.Println("\n---------------")
}

// ioutil.ReadFile() 一次读取文件所有内容
func ioutilt() {
	file := "e:/go/gocode/src/go_code/oop/file/test.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read file err=", err)
	}
	fmt.Printf("%v", string(content))
	fmt.Println("\n---------------")
}

// 创建文件并写入数据
func wFile() {
	// 1.打开文件
	filePath := "e:/go/gocode/src/go_code/oop/file/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err=", err)
	}
	defer file.Close()
	str := "hello,garden\n"
	// 2.写入时，使用带缓存的*Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// writer是先写入缓存，需要调用Flush()写入文件
	writer.Flush()
}

func copy() {
	filePath1 := "e:/go/gocode/src/go_code/oop/file/abc.txt"
	filePath2 := "e:/go/gocode/src/go_code/oop/file/abc2.txt"
	data, err := ioutil.ReadFile(filePath1)
	if err != nil {
		fmt.Println("read file err=", err)
		return
	}
	err = ioutil.WriteFile(filePath2, data, 0666)
	if err != nil {
		fmt.Println("write file err=", err)
	}
}

func testCopy() {
	srcFile := "d:/Study/head.jpg"
	dstFile := "d:/Study/head2.jpg"
	src, err := os.Open(srcFile)
	if err != nil {
		fmt.Printf("open file error=%v\n", err)
	}
	// 通过src获取Reader
	reader := bufio.NewReader(src)
	defer src.Close()
	// fmt.Println(reader.ReadLine())
	// 打开dstFileName
	dst, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE, 644)
	if err != nil {
		fmt.Printf("open file error=%v\n", err)
	}
	writer := bufio.NewWriter(dst)

	defer dst.Close()
	_, err2 := io.Copy(writer, reader)
	writer.Flush()
	if err2 == nil {
		fmt.Println("拷贝完成")
	} else {
		fmt.Println("error=", err2)
	}
}

func main() {
	test()
	ioutilt()
	// wFile()
	// copy()
	// testCopy()
}
