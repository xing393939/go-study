package strategy

type IStrategy interface {
	do(int, int) int
}

type add struct{}

func (*add) do(a, b int) int {
	return a + b
}

type reduce struct{}

func (*reduce) do(a, b int) int {
	return a - b
}

type Operator struct {
	strategy IStrategy
}

func (operator *Operator) setStrategy(strategy IStrategy) {
	operator.strategy = strategy
}

func (operator *Operator) calculate(a, b int) int {
	return operator.strategy.do(a, b)
}
