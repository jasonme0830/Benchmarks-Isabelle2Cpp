package main

import (
	"fmt"
	"time"
	"bufio"
	"os"
	"strconv"
	. "isabelle/exported/HOL"
	. "isabelle/exported/Tree"
    . "isabelle/exported/Orderings"
    . "isabelle/exported/InsertTree"
	. "isabelle/exported/SearchTree" 
)

// 定义一个辅助函数，用于判断一个 Tree 是否是 Tip 类型
func isTip[a any](t Tree[a]) bool {
	_, ok := t.(Tip[a])
	return ok
}
// 打印函数，递归打印二叉树
func printTree[a any](t Tree[a], indent string) {
	if isTip(t) {
		fmt.Print(indent + "Tip")
		return
	}
 
	// 类型断言为 Node
	node, ok := t.(Node[a])
	if !ok {
		fmt.Println("Invalid tree structure")
		return
	}

	// 打印左子树
	fmt.Print("(")
	printTree(node.A, indent+"")
	fmt.Print(")")
	// 打印当前节点
	fmt.Print(indent + fmt.Sprintf(" Node(%v) ", node.Aa))
	// 打印右子树
	fmt.Print("(")
	printTree(node.Ab, indent+"")
	fmt.Print(")")
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
	tar, eoo := strconv.Atoi(args[1])
	if (err != nil)||(eoo != nil) {
		fmt.Printf("错误：无效的整数参数 '%s'\n", args[0])
		os.Exit(1)
	}
	fmt.Println("num: ", num, "   target: ", tar)

	// 打开文件进行读取
	file, err := os.Open("../ArandomNum/random_numbers.txt")
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer file.Close()
	var numbers []int
	scanner := bufio.NewScanner(file)
	// 从文件中读取指定数量的整数并存储到切片中
	for i := 0; i < num && scanner.Scan(); i++ {
		// 将每一行转换为整数
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("无法将行转换为整数:", err)
			continue
		}
		numbers = append(numbers, number)
	}
	// 检查扫描过程中是否有错误
	if err := scanner.Err(); err != nil {
		fmt.Println("read file error:", err)
		return
	}
	// print random numbers list
	// fmt.Print("random numbers list:  ")
	// for _, number := range numbers {
	// 	fmt.Print(number, " ")
	// }
	// fmt.Println()

	// 实例化 Ord 结构体，例如用于比较整数
	// TODO:fix the use of func "Less_eq" in the body of Inserttree
	intLess := Ord[int]{
		Less_eq: func(a, b int) bool {
			return a <= b
		},
		Less: func(a, b int) bool {
			return a < b
		},
	}
	newTree := Inserttree(intLess, numbers[0], Tree[int](Tip[int]{}))
	for i:=1; i<len(numbers); i++{
		newTree = Inserttree(intLess, numbers[i], newTree)
	}
	// printTree(newTree, "")


	//实例化 Equal 结构体，例如用于比较整数
	intEqual := Equal[int]{
		Equala: func(a, b int) bool {
			return a == b
		},
	}

	startNano := time.Now().UnixNano() // 获取纳秒时间戳
	res := Searchtree1(intEqual, tar, newTree)
	endNano := time.Now().UnixNano()
	elapsedNano := endNano - startNano
	fmt.Println("res: ", res, "  time: ", elapsedNano," ns")
}

