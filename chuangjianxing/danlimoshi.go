package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//单例模式的优缺点
//优点：
//(1) 单例模式提供了对唯一实例的受控访问。
//(2) 节约系统资源。由于在系统内存中只存在一个对象。
//
//缺点：
//(1) 扩展略难。单例模式中没有抽象层。
//(2) 单例类的职责过重

//适用场景
//(1) 系统只需要一个实例对象，如系统要求提供一个唯一的序列号生成器或资源管理器，或者需要考虑资源消耗太大而只允许创建一个对象。
//(2) 客户调用类的单个实例只允许使用一个公共访问点，除了该公共访问点，不能通过其他途径访问该实例。

//饿汉式
var instanceInit *instance = new(instance)

type instance struct {
}

func GetInstance() *instance {
	return instanceInit
}

func main() {
	i := GetInstance()
	fmt.Println(i)
}

//懒汉式
var instanceInit1 *instance
var i uint32
var lock sync.Mutex

func GetInstance1() *instance {
	if atomic.LoadUint32(&i) == 1 {
		return instanceInit1
	}
	lock.Lock()
	defer lock.Unlock()
	if i == 0 {
		instanceInit1 = new(instance)
		atomic.StoreUint32(&i, 1)
	}
	return instanceInit1
}
