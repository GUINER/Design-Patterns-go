package builder

import "fmt"

// 使用多个简单的对象一步一步构建成一个复杂的对象

// 外卖点餐
type TakeAway struct {
	StapleFood string `json:"staple_food"` // 主食
	SideDish   string `json:"side_dish"`   // 配菜
	Chili      string `json:"chili"`       // 辣椒
	Drink      string `json:"drink"`       // 饮料
	Remark     string `json:"remark"`      // 备注
}

func (t *TakeAway) YourStapleFood(food string) *TakeAway {
	t.StapleFood = food
	return t
}

func (t *TakeAway) YourSideDish(food string) *TakeAway {
	t.SideDish = food
	return t

}

func (t *TakeAway) WithChili(chili string) *TakeAway {
	t.Chili = chili
	return t
}

func (t *TakeAway) WithDrink(drink string) *TakeAway {
	t.Drink = drink
	return t

}

func (t *TakeAway) YourRemark(note string) *TakeAway {
	t.Remark = note
	return t

}

func (t *TakeAway) CreateOrder() {
	fmt.Printf("你的外卖：%s, %s, %s, %s \n", t.StapleFood, t.SideDish, t.Chili, t.Drink)
	fmt.Println("备注：" + t.Remark)
}

func NewTakeAwayBuilder() *TakeAway {
	return &TakeAway{}
}
