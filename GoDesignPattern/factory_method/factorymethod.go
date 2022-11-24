package factory_method

// operator是被封在的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// OperatorFactory是工厂接口
type OperatorFactory interface {
	Create() Operator
}

// OperatorBase  是operator接口的实现的基类
type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int) {
	o.a = a
}

func (o *OperatorBase) SetB(b int) {
	o.b = b
}

type PlusOperatorFactory struct {
}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type PlusOperator struct {
	*OperatorBase
}

func (p PlusOperator) Result() int {
	return p.a + p.b
}

// 减法实现
type MinusOperatorFactory struct {
}

func (o MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type MinusOperator struct {
	*OperatorBase
}

func (o MinusOperator) Result() int {

	return o.a - o.b
}
