package main

import "fmt"
//指针
func main()  {
	//声明实际变量
	var a int=10
	//声明指针
	var ip *int
	//指针变量的存储地址
	ip=&a
	fmt.Println(&a)
	//指针变量存储的地址
	fmt.Println(ip)
	//使用指针访问值
	fmt.Println(*ip)

	var ip1 *int
	fmt.Println(ip1)
	if ip1 == nil {
		fmt.Println("空指针")
	} else {
		fmt.Println("不是空指针")
	}


}
