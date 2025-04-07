package main

import (
	"fmt"
	"time"
	"os"
	"strconv"
	. "isabelle/exported/List"
    "isabelle/exported/AddListTail" // 导入模块中的包
)

// 
func printHelper[a any](list Lista[a]) {
    switch v := list.(type) {
    case Nil[a]:
        fmt.Printf("Nil\n")
        return // 空链表直接返回
    case Cons[a]:
        fmt.Printf("%v -> ", v.A)
        printHelper(v.Aa)
    default:
        panic(fmt.Sprintf("未知链表类型: %T", list))
    }
}

func main() {
	// 获取命令行参数
	args := os.Args[1:]
 
	// 检查是否提供了参数
	if len(args) < 1 {
		fmt.Println("错误：请提供一个整数参数")
		fmt.Println("用法：程序名 <整数>")
		os.Exit(1)
	}
 
	// 尝试转换参数为整数
	num, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("错误：无效的整数参数 '%s'\n", args[0])
		os.Exit(1)
	}
	fmt.Println("num: ", num)

	xs := Nil[int]{}
	newList := AddListTail.AddListTaila[int](num, xs)
	for i:=num-1; i>0; i--{
		newList = AddListTail.AddListTaila[int](i, newList)
	}

	startNano := time.Now().UnixNano() // 获取纳秒时间戳
	newList = AddListTail.AddListTaila[int](0, newList)
	endNano := time.Now().UnixNano()
	elapsedNano := endNano - startNano

	printHelper(newList)
	fmt.Println(" time: ", elapsedNano,"ns")
}

