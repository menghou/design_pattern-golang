package strategy_pattern


type CashSuper interface {
	AcceptCash(float64) float64
}

type CashNormal struct {
}

func (n *CashNormal) AcceptCash(money float64) float64 {
	return money
}

type CashRebate struct {
	 discount float64
}

func (n *CashRebate) AcceptCash(money float64) float64 {
	return n.discount * money
}

type CashReturn struct {
	total float64
	ret   float64
}

func (n *CashReturn) AcceptCash(money float64) float64 {
	if money < n.total {
		return money
	}else{
		return money - n.ret
	}
}

type CashContext struct {
	cashStragety CashSuper
}

func (c *CashContext) AcceptCash(money float64) float64 {
	return c.cashStragety.AcceptCash(money)
}

func (c *CashContext) CashContext(str string)  {
	switch str{
	case "normal":
		c.cashStragety = CashNormal{}
	case "8 rebate":
		c.cashStragety = CashRebate{discount:0.8}
	case "300 return 100" :
		c.cashStragety = CashReturn{total:300, ret:100}
	default:
		panic("wrong input")
	}
}