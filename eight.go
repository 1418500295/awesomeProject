package main

import (

	"fmt"
)

type Books struct {
	title string
	author string
	subject string

}

func printBook(book *Books)  {
	fmt.Println("Book的标题是"+book.title)
	fmt.Println("Book的作者是"+book.author)


}

func main() {
	//声明book1 book2为Books类型
	var book1  Books
	var book2  Books
	/*book1描述*/
	book1.title = "语文"
	book1.author = "小黄"
	/*book2描述*/
	book2.title = "数学"
	book2.author = "小红"

	//调用函数，传入实例
	printBook(&book1)
	printBook(&book2)

}
