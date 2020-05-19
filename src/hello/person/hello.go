package main

import "fmt"

var x, y int
var (
	aa int
	bb bool
)
var c1,d1 int=1,2
var e1,f1  = 123,"hello"
/**
基本类型，函数
 */
func main()  {
	var a string="zhangsan"
	var b, c =1, 2
	fmt.Println(b,c)
	fmt.Println(a)
	fmt.Println("hello world")
	//默认0
	var d int
	fmt.Println(d)
	//默认false
	var e bool
	fmt.Println(e)

	var f string
	fmt.Println(f)

	var tt=true
	fmt.Println(tt)

	ee := "ni hao"
	fmt.Println(ee)

	g1,h1:=123,"world"
	fmt.Println(x,y,aa,bb,c1,d1,e1,f,g1,h1)
	const LENGTH  int=10
	const WIDTH int =5
	var area int
	area=LENGTH*WIDTH
	fmt.Printf("面积为：%d",area)
	fmt.Println()
	const (
		a3=iota
		b3
		c3
		d3
	)
	fmt.Println(a3,b3,c3,d3)
   const(
   	i=1<<iota
   	j=2<<iota
   	k
   	l
   )
	fmt.Println("i:",i)
	fmt.Println("j:",j)
	fmt.Println("k:",k)
	fmt.Println("l",l)
   //比较两个数的最大值
  var result1 =max(10,50)
   fmt.Println(result1)
  //
	s, i2 := swap("aa", "bb")
	fmt.Println(s,i2)

}
//定义一个函数返回多个值
func swap(x,y string)(string,string){
	return y,x
}


//定义一个函数，返回两个数的最大值
func  max (num1, num2  int) int{
	var result int
	if num1 > num2{
		result=num1
	}else{
		result=num2
	}
	return result
}
