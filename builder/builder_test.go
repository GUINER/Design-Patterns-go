package builder

import "testing"

func TestNewTakeAwayBuilder(t *testing.T) {
	builder := NewTakeAwayBuilder()
	builder.YourStapleFood("黑椒鸡扒意面").YourSideDish("香肠").WithChili("不要辣").
		WithDrink("柠檬水").YourRemark("走青！肚子饿，请快点送到！")

	builder.CreateOrder()
}
