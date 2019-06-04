package facotry_pattern

import (
	"guin/design_model/facotry_pattern/huawei"
	"guin/design_model/facotry_pattern/xiaomi"
)

type Phone interface {
	Call()
	Sms()
}

func GetXiaomi() Phone {
	return &xiaomi.XiaoMi{}
}

func GetHuaWei() Phone {
	return &huawei.HuaWei{}
}

func GenPhone(phoneName string) Phone {
	switch phoneName {
	case "xiaomi":
		return GetXiaomi()
	case "huawei":
		return GetHuaWei()
	default:
		panic("Invalid phone Name")
	}
}
