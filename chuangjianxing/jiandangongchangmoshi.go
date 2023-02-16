package main

import "fmt"

//简单工厂方法模式的优缺点
//优点：
//1. 实现了对象创建和使用的分离。
//2. 不需要记住具体类名，记住参数即可，减少使用者记忆量。
//
//缺点：
//1. 对工厂类职责过重，一旦不能工作，系统受到影响。
//2. 增加系统中类的个数，复杂度和理解度增加。
//3. 违反“开闭原则”，添加新产品需要修改工厂逻辑，工厂越来越复杂。
//
//适用场景：
//1.  工厂类负责创建的对象比较少，由于创建的对象较少，不会造成工厂方法中的业务逻辑太过复杂。
//2. 客户端只知道传入工厂类的参数，对于如何创建对象并不关心。

//适用场景：
//1.  工厂类负责创建的对象比较少，由于创建的对象较少，不会造成工厂方法中的业务逻辑太过复杂。
//2. 客户端只知道传入工厂类的参数，对于如何创建对象并不关心。

//抽象层(抽象水果)

type F interface {
	Show()
}

//具体实现层

type A struct {
}

func (a *A) Show() {
	fmt.Println("我是苹果")
}

type B struct {
}

func (b *B) Show() {
	fmt.Println("我是香蕉")
}

type P struct {
}

func (p *P) Show() {
	fmt.Println("我是梨")
}

//工厂类
type FF struct {
}

//此方法在增加水果产品类的时候会修改，违背开闭原则
func (ff *FF) CreateF(name string) F {
	var f F
	if name == "apple" {
		f = new(A)
	} else if name == "banana" {
		f = new(B)
	} else if name == "pear" {
		f = new(P)
	}
	return f
}

func main() {
	fac := new(FF)
	apple := fac.CreateF("apple")
	apple.Show()
	banana := fac.CreateF("banana")
	banana.Show()
	pear := fac.CreateF("pear")
	pear.Show()
}
