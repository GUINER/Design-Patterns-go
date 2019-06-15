package abstract_factory

import "fmt"

// 抽象工厂，相当于一个工厂的集合。一个工厂可以生产多个产品，如手机、电脑等，因为每个产品都是一个接口，所以一个工厂会可以有多个接口。

// --------- 产品A接口 ---------
type ProductMobile interface {
	CreateMobile()
}

// --------- 产品B接口 ---------
type ProductComputer interface {
	CreateComputer()
}

// 实现你的product
// ----------- mobile -------------
type MobileA struct {
	BrandName string `json:"brand_name"`
}

func (m *MobileA) CreateMobile() {
	fmt.Println("generate a mobile : " + m.BrandName)
}

// ----------- mobile -------------
type ComputerA struct {
	BrandName string `json:"brand_name"`
}

func (c *ComputerA) CreateComputer() {
	fmt.Println("generate a computer : " + c.BrandName)
}

// --------- 抽象工厂 ---------
type AbstractFactory struct {
	Mobile   ProductMobile
	Computer ProductComputer
}

// 创建一个工厂
func NewAbstractFactory(mobile ProductMobile, computer ProductComputer) *AbstractFactory {
	return &AbstractFactory{
		Mobile:   mobile,
		Computer: computer,
	}
}
