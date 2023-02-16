package main

//抽象工厂模式的优缺点
//优点：
//1.  拥有工厂方法模式的优点
//2. 当一个产品族中的多个对象被设计成一起工作时，它能够保证客户端始终只使用同一个产品族中的对象。
//3   增加新的产品族很方便，无须修改已有系统，符合“开闭原则”。
//
//缺点：
//1. 增加新的产品等级结构麻烦，需要对原有系统进行较大的修改，甚至需要修改抽象层代码，这显然会带来较大的不便，违背了“开闭原则”。
//抽象配件
type AbstractVideoCard interface {
	ShowVideoCard()
}

type AbstractMemory interface {
	ShowMemory()
}

type AbstractCpu interface {
	ShowCpu()
}

//具体配件
//英特尔显卡
type IntelVideoCard struct {
}

func (ivc *IntelVideoCard) ShowVideoCard() {
	println("英特尔显卡")
}

//Nvidia显卡
type NvidiaVideoCard struct {
}

func (nvc *NvidiaVideoCard) ShowVideoCard() {
	println("Nvidia显卡")
}

//Kingston显卡
type KingstonVideoCard struct {
}

func (kvc *KingstonVideoCard) ShowVideoCard() {
	println("Kingston显卡")
}

//英特尔内存
type IntelMemory struct {
}

func (im *IntelMemory) ShowMemory() {
	println("英特尔内存")
}

//Nvidia内存
type NvidiaMemory struct {
}

func (nm *NvidiaMemory) ShowMemory() {
	println("Nvidia内存")
}

//Kingston内存
type KingstonMemory struct {
}

func (km *KingstonMemory) ShowMemory() {
	println("Kingston内存")
}

//英特尔cpu
type IntelCpu struct {
}

func (ic *IntelCpu) ShowCpu() {
	println("英特尔cpu")
}

//Nvidia cpu
type NvidiaCpu struct {
}

func (nc *NvidiaCpu) ShowCpu() {
	println("NvidiaCPU")
}

//Kingston cpu
type KingstonCpu struct {
}

func (kc *KingstonCpu) ShowCpu() {
	println("KingstonCPU")
}

//抽象工厂
type AbstractFactory interface {
	CreateVideoCard() AbstractVideoCard
	CreateMemory() AbstractMemory
	CreateCpu() AbstractCpu
}

//具体工厂
//英特尔工厂
type IntelFactory struct {
}

func (ifa *IntelFactory) CreateVideoCard() AbstractVideoCard {
	var videoCard AbstractVideoCard
	videoCard = new(IntelVideoCard)
	return videoCard
}

func (ifa *IntelFactory) CreateMemory() AbstractMemory {
	var memory AbstractMemory
	memory = new(IntelMemory)
	return memory
}

func (ifa *IntelFactory) CreateCpu() AbstractCpu {
	var cpu AbstractCpu
	cpu = new(IntelCpu)
	return cpu
}

//Nvidia工厂
type NvidiaFactory struct {
}

func (nfa *NvidiaFactory) CreateVideoCard() AbstractVideoCard {
	var videoCard AbstractVideoCard
	videoCard = new(NvidiaVideoCard)
	return videoCard
}

func (nfa *NvidiaFactory) CreateMemory() AbstractMemory {
	var memory AbstractMemory
	memory = new(NvidiaMemory)
	return memory
}

func (nfa *NvidiaFactory) CreateCpu() AbstractCpu {
	var cpu AbstractCpu
	cpu = new(NvidiaCpu)
	return cpu
}

//Kingston工厂
type KingstonFactory struct {
}

func (kfa *KingstonFactory) CreateVideoCard() AbstractVideoCard {
	var videoCard AbstractVideoCard
	videoCard = new(KingstonVideoCard)
	return videoCard
}

func (kfa *KingstonFactory) CreateMemory() AbstractMemory {
	var memory AbstractMemory
	memory = new(KingstonMemory)
	return memory
}

func (kfa *KingstonFactory) CreateCpu() AbstractCpu {
	var cpu AbstractCpu
	cpu = new(KingstonCpu)
	return cpu
}

func main() {
	var ifac AbstractFactory
	ifac = new(IntelFactory)
	//生产intelCPU
	var intelCpu AbstractCpu
	intelCpu = ifac.CreateCpu()
	intelCpu.ShowCpu()
	//生产intel显卡
	var intelVideoCard AbstractVideoCard
	intelVideoCard = ifac.CreateVideoCard()
	intelVideoCard.ShowVideoCard()
	//生产intel内存
	var intelMemory AbstractMemory
	intelMemory = ifac.CreateMemory()
	intelMemory.ShowMemory()

	//生产intelCpu
	var intelCpu1 AbstractCpu
	intelCpu1 = ifac.CreateCpu()
	intelCpu1.ShowCpu()
	//Nvidia显卡
	var nfac AbstractFactory
	nfac = new(NvidiaFactory)
	var nvidiaVideoCard AbstractVideoCard
	nvidiaVideoCard = nfac.CreateVideoCard()
	nvidiaVideoCard.ShowVideoCard()
	//生产Kingston内存
	var kfac AbstractFactory
	kfac = new(KingstonFactory)
	var kingstonMemory AbstractMemory
	kingstonMemory = kfac.CreateMemory()
	kingstonMemory.ShowMemory()
}
