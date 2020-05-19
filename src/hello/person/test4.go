package main

import "fmt"
//range
func main()  {
	nums := []int{2, 3, 4}
	sum:=0
	for _,num:=range nums{
		sum+=num
	}
	fmt.Println(sum)

	for i,num:=range nums{
		if num==3 {
			fmt.Println(i)
		}
	}

	kvs := map[string]string{"a": "apple", "c": "banana"}
	for k,v:=range kvs{
		fmt.Println(k,v)
	}
	for i,c:=range "go"{
		fmt.Println(i,c)
	}
}
