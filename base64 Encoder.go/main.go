package main

import (
	"encoding/base64"
	"fmt"
)

func in(A string) {
	admin := A
	c := base64.StdEncoding.EncodeToString([]byte(admin))
	fmt.Println("编码结果:", c)
}
func put(A string) {
	admin := A
	d, ok := base64.StdEncoding.DecodeString(admin)
	if ok != nil {
		panic(fmt.Errorf("解码失败: %v", ok))
	}

	D := string(d)
	fmt.Println("解码结果:", D)
}
func Ifn(A string) {
	for i := 1; i > 0; i++ {
		if A == "exit" {
			break
		} else if A == "1" {

			var admin string
			fmt.Println("请输入你要编码的内容:")
			fmt.Scanf("%s", &admin)
			in(admin)

		} else if A == "2" {

			var admin string
			fmt.Println("请输入你要执行解码的内容:")
			fmt.Scanf("%s", &admin)
			put(admin)

		} else {
			fmt.Println("输入无效,请输入1,2或exit。")
			continue
		}
	}
}

func main() {
	for i := 1; i > 0; i++ {
		var A string
		fmt.Println("请输入你要执行的选项 1编码 2解码 exit退出:")
		fmt.Scanf("%s", &A)

	}

}
