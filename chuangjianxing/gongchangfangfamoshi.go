package main

import "fmt"

//工厂方法模式的优缺点
//优点：
//1. 不需要记住具体类名，甚至连具体参数都不用记忆。
//2. 实现了对象创建和使用的分离。
//3. 系统的可扩展性也就变得非常好，无需修改接口和原类。
//4.对于新产品的创建，符合开闭原则。
//
//缺点：
//1. 增加系统中类的个数，复杂度和理解度增加。
//2. 增加了系统的抽象性和理解难度。
//
//适用场景：
//1. 客户端不知道它所需要的对象的类。
//2. 抽象工厂类通过其子类来指定创建哪个对象。

//适用场景：
//1. 客户端不知道它所需要的对象的类。
//2. 抽象工厂类通过其子类来指定创建哪个对象。

//适用场景
//(1) 系统中有多于一个的产品族。而每次只使用其中某一产品族。可以通过配置文件等方式来使得用户可以动态改变产品族，也可以很方便地增加新的产品族。
//(2) 产品等级结构稳定。设计完成之后，不会向系统中增加新的产品等级结构或者删除已有的产品等级结构。

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

//抽象工厂类
type FF interface {
	CreateF() F
}

//苹果工厂
type AF struct {
}

func (af *AF) CreateF() F {
	var f F
	f = new(A)
	return f
}

//香蕉工厂
type BF struct {
}

func (bf *BF) CreateF() F {
	var f F
	f = new(B)
	return f
}

//梨工厂
type PF struct {
}

func (pf *PF) CreateF() F {
	var f F
	f = new(P)
	return f
}

func main() {
	var afac FF
	afac = new(AF)
	apple := afac.CreateF()
	apple.Show()
	var bfac FF
	bfac = new(BF)
	banana := bfac.CreateF()
	banana.Show()
	var pfac FF
	pfac = new(PF)
	pear := pfac.CreateF()
	pear.Show()
}
