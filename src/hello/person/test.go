package main

import "fmt"

var g int=20
//数组
func main()  {
	var g int = 10
	fmt.Println(g)
	//数组初始化
	var balance = [3] int{1, 2, 3,}
	fmt.Println(balance)
	//不设置大小
	var b=[]int{1,2,3,4,5}
	fmt.Println(b)
	fmt.Println(b[4])
	//循环
    var c[10]int
	var i int
	for  i=0; i<10;i++{
		c[i]=100+i
	}
	fmt.Println(c)
}