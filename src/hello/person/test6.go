package main

import "fmt"
//递归函数
func main()  {
	var  i int =3
	result:= Factorial(i)
	fmt.Println(result)

}
func Factorial(n int)(result int)  {
	if n>0{
		result= n*Factorial(n-1)
		return result
	}
	return 1
}