package main

import "fmt"

func main()  {
	var numbers=make([]int,3,5)
	printSlice(numbers)
	//空切片
	var numbers1 []int
	printSlice(numbers1)
	if numbers1==nil {
		fmt.Println("切片是空的")
	}
	//切片截取
	var numbers2 =[]int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers2)
	fmt.Println("numbers2[1:4]:",numbers2[1:4])
	fmt.Println("numbers2[:3]:",numbers2[:3])
	fmt.Println("numbers2[4:]:",numbers2[4:])
	//append和copy
	var numbers3 []int
	fmt.Println("numbers3",numbers3)
	//追加空切片
	numbers3=append(numbers3,0)
	fmt.Println("numbers3",numbers3)
	//向切片添加一个元素
	numbers3= append(numbers3, 1)
	fmt.Println("numbers3",numbers3)
	//向切片添加多个元素
	numbers3=append(numbers3,2,3)
	fmt.Println("numbers3",numbers3)
	//创建切片是numbers3的两倍
	numbers4 := make([]int, len(numbers3), (cap(numbers3))*2)
	//拷贝
	copy(numbers4,numbers3)
	printSlice(numbers4)
}
func printSlice(x []int)  {
	fmt.Println(len(x),cap(x),x)
}