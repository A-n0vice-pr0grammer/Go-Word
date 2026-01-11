package main

import (
	"fmt"
	"time"
)

type BookNews struct {
	Name  string
	Admin string
	Where string
} //定义书籍信息内容

var BookName = map[int]BookNews{
	1: {Name: "哈利波特", Admin: "罗琳", Where: "第五章"},
} //书籍信息

func main() {
	text()

	for {
		t()
		fmt.Printf("\n请输入你要执行的操作编号：\n")
		fmt.Println("1.存书/编辑  2.删除  3.查询  4.查询所有书  5.退出")
		var using int
		fmt.Scanf("%d", &using)
		fmt.Scanln()
		num := Using(using)

		switch num {
		case 1:
			Put()
		case 2:
			Delet()
		case 3:
			found()
		case 4:
			oprt()
		case 5:
			fmt.Println("已离开，886~")
			return
		default:
			fmt.Println("输入格式错误！请输入正确的指令!")
		}
	}

}

func Using(using int) int {
	switch using {
	case 1:
		t() //存书/修改存书信息
		fmt.Println("正在打开你的书架...")
		return 1
	case 2:
		t() //删除书籍信息
		fmt.Println("正在打开你的书架...")
		return 2
	case 3:
		t() //列出查询的书籍信息
		fmt.Println("正在打开你的书架...")
		return 3
	case 4:
		t() //列出所有书籍信息
		fmt.Println("正在打开你的书架...")
		return 4
	case 5:
		t() //退出图书馆
		fmt.Println("正在离开图书馆...")
		return 5
	default:
		return 0
	}
}
func Put() {
	t()
	//存书/修改
	fmt.Println("请输入您要存入/修改的书籍编号：")
	var key int
	fmt.Scanf("%d", &key)
	fmt.Scanln()
	fmt.Println("请输入您要存入/修改的书籍名称：")
	var Bname string
	fmt.Scanf("%s", &Bname)
	fmt.Scanln()
	fmt.Println("请输入您要存入/修改的作者姓名：")
	var Aname string
	fmt.Scanf("%s", &Aname)
	fmt.Scanln()
	fmt.Println("请输入您要存入/修改的书籍观看位置：")
	var where string
	fmt.Scanf("%s", &where)
	fmt.Scanln()

	//储存

	Name := BookNews{
		Name:  Bname,
		Admin: Aname,
		Where: where,
	}
	BookName[key] = Name

	t()
	fmt.Printf("存入/修改成功，您存入/修改书籍编号为 %d \n", key)
}
func Delet() {
	//删除书籍信息
	fmt.Println("请输入你想删除的书籍编号")
	var key int
	fmt.Scanf("%d", &key)
	fmt.Scanln()
	t()
	fmt.Println("删除中...")
	delete(BookName, key)
	t()
	fmt.Printf("**删除成功**\n")
}
func found() {
	//列出查询的书籍信息
	fmt.Println("请输入你想查询的书籍编号")
	var key int
	fmt.Scanf("%d", &key)
	fmt.Scanln()
	t()
	fmt.Println("查询中...")
	fmt.Println(BookName[key])
	t()
	fmt.Printf("查询完毕\n")

}
func oprt() {
	//列出所有书籍信息
	t()
	fmt.Println("正在为你进行所有书籍查询")
	t()
	fmt.Println("查询中...")
	for _, value := range BookName {
		t()
		fmt.Println(value)
	}
	t()
	fmt.Printf("查询完毕\n")
}
func text() {
	fmt.Println("欢迎来到我的图书馆！！")
	t()
	fmt.Println("你可以将喜欢的书存入这里，这里将成为你的专属书架")
	t()
	fmt.Println("也可以寻找看过的书籍，修改书籍的存储内容")
	t()
	fmt.Println("书架存书格式如下")
	t()
	fmt.Println("{书籍名称  作者姓名  观看章节位置}")
	t()
	fmt.Println("现在开始使用我们的图书馆吧")
	t()
	fmt.Printf("\n")
}
func t() {
	time.Sleep(1 * time.Second)
}
