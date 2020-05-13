package main

import "fmt"
/**
结构体使用
 */
type Books struct {
	title string
	author string
	book_id int
	subject string

}

func main() {
	//创建一个新的结构体
	fmt.Println(Books{"go语言","daine",2,"数学"})
	//可以使用key-->value格式
	fmt.Println(Books{title:"python语言",author:"james",book_id:3,subject:"语文"})
	//忽略的字段为0或者空
	fmt.Println(Books{title:"java语言",author:"jordan"})

	/**
	访问结构体成员
	 */
	var book1  Books		/*声明brook1为Books类型 */
	var book2  Books		/*声明brook2为Books类型 */
	book1.title = "java教程"
	book2.title = "python教程"
	book1.author = "daine"
	book2.author = "jack"
	fmt.Printf("brook1的title是:%s\n",book1.title)
	fmt.Printf("brook2的title是:%s\n",book2.title)

	//调用函数
	printBook(book1)
	printBook(book2)


}
//你可以像其他数据类型一样将结构体类型作为参数传递给函数。并以以上实例的方式访问结构体变量：
func printBook(book Books)  {
	fmt.Printf("Book的title是%s\n",book.title)
	fmt.Printf("Book的author是%s\n",book.author)

	
}
