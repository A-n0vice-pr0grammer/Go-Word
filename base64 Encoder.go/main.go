package main

import (
	"encoding/base64"
	"fmt"
	"time"
)

func main() {
	print()
	for {
		// 判断用户输入逻辑
		times()
		fmt.Println("请输入你要执行的操作：")
		var using string
		fmt.Scanf("%s", &using)
		fmt.Scanln()
		num := Input(using)
		times()
		switch num {
		case 1:
			close()
		case 2:
			oppen()
		case 3:
			fmt.Println("程序已正常退出，再见～")
			return
		default:
			fmt.Println("输入格式错误！请输入 1 / 2 / exit 其中一个指令")
		}
	}
}

func times() {
	time.Sleep(1 * time.Second) //停顿1s
}
func print() {
	times()
	fmt.Println("欢迎使用base64编解码程序：")
	times()
	fmt.Println("输入 1 : 进行 Base64 编码操作")
	times()
	fmt.Println("输入 2 : 进行 Base64 解码操作")
	times()
	fmt.Println("输入 exit : 退出本次操作")
	times()
}
func Input(using string) int {
	switch using {
	case "1":
		times()
		fmt.Println("接下来执行编码操作...")
		return 1
	case "2":
		times()
		fmt.Println("接下来执行解码操作...")
		return 2
	case "exit":
		times()
		fmt.Println("正在退出程序...")
		return 3
	default:
		return 0
	}
}
func close() {
	fmt.Println("请输入需要编码的字符串：")
	admin := writing()
	// 标准Base64编码
	encoded := base64.StdEncoding.EncodeToString([]byte(admin))
	fmt.Println("编码结果：", encoded)
}
func oppen() {
	fmt.Println("请输入需要解码的字符串：")
	encoded := writing()
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("格式错误:", err)
		return
	}
	fmt.Println("解码结果:", string(decoded))
}
func writing() string {
	//输入
	var str string
	fmt.Scanf("%s", &str)
	fmt.Scanln()
	return str
}

