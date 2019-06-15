package abstract_factory

import "testing"

func TestNewAbstractFactory(t *testing.T) {
	mobile := &MobileA{BrandName: "HW"}
	computer := &ComputerA{BrandName: "Dell"}

	factory := NewAbstractFactory(mobile, computer)

	factory.Mobile.CreateMobile()
	factory.Computer.CreateComputer()
}
