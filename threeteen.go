package main

import "fmt"
/**
定义了一个接口Phone，接口里面有一个方法call()。
然后我们在main函数里面定义了一个Phone类型变量，
并分别为之赋值为NokiaPhone和IPhone。然后调用call()方法
 */
type Phone interface {
	call()
}

type NokiaPhone struct {

}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("i am nokia,call you")
	
}

type IPhone struct {

}

func (iphone IPhone) call() {
	fmt.Println("i am iphone,call you")

}


func main() {

	var phone  Phone
	phone = new(NokiaPhone)
	phone.call()
	phone = new(IPhone)
	phone.call()


}
