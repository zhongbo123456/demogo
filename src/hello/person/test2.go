package main

import "fmt"

//结构体，相当java中的类
type Books struct {
	title string
	author string
	subject string
	book_id int
}


func main()  {
	fmt.Println(Books{"java","zhangsan","编程",1})
	//key-value格式
	fmt.Println(Books{title:"高数",author:"lisi",subject:"数学",book_id:2})
	//忽略字段为0或空
	fmt.Println(Books{title:"离散数学",subject:"数学"})
	var Books1 Books
	Books1.title="go language"
	Books1.author="zhaoliu"
	Books1.subject="编程"
	Books1.book_id=3
	fmt.Println(Books1.title)
	//结构体指针
	printBook(&Books1)


}
func printBook(book *Books)  {
	fmt.Println(book.title)
	fmt.Println(book.author)
	fmt.Println(book.subject)
	fmt.Println(book.book_id)
	
}
