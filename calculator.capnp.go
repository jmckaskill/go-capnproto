
type Calculator interface {
	C.Capability
	Evaluate(ctx C.Call, exp CalculatorExpression) (CalculatorValue, error)
	DefFunction(ctx C.Call, paramCount int32, body CalculatorExpression) (CalculatorFunction, error)
	GetOperator(ctx C.Call, op CalculatorOperator) (CalculatorFunction, error)
}

func Calculator_Interface(ctx CallContext, obj Capability, method uint16, params Struct) error {
	if ctx.InterfaceId() != 0x.... {
		return C.ErrWrongInterface
	}

	switch s := obj.(Calculator); method {
	case 0:
		a0 := params.GetObject(0)
		ret := ctx.CreateReturnParams(0, 1)
		r0, err := s.Evaluate(ctx, CalculatorExpression(a0))
		ret.SetObject(0, C.Object(r0))
		return err
	case 1:
		a0 := params.Get32(0)
		a1 := params.GetObject(0)
		ret := ctx.CreateReturnParams(0, 1)
		r0, err := s.DefFunction(ctx, int32(a0), CalculatorExpression(a1))
		ret.SetObject(0, C.Object(r0))
		return err
	case 2:
		a0 := params.Get16(0)
		ret := ctx.CreateReturnParams(0, 1)
		r0, err := s.GetOperator(ctx, CalculatorOperator(a0))
		ret.SetObject(0, C.Object(r0))
		return err
	default:
		return C.ErrWrongMethod
	}
}

type CalculatorExpression C.Object
type CalculatorExpression_Which uint16

func (s CalculatorExpression) Which() CalculatorExpression_Which
func (s CalculatorExpression) Literal() float64
func (s CalculatorExpression) PreviousResult() Value
func (s CalculatorExpression) Parameter() uint32
func (s CalculatorExpression) Call() CalculatorExpressionCall
func (s CalculatorExpressionCall) Function() Function
func (s CalculatorExpressionCall) Params() CalculatorExpression_List

type CalculatorValue interface {
	C.Capability
	Read(ctx C.Call) (float64, error)
}

func CalculatorValue_Interface(ctx CallContext, obj Capability, method uint16, params Struct) error {
	if ctx.InterfaceId() != 0x.... {
		return C.ErrWrongInterface
	}

	switch s := obj.(CalculatorValue); method {
	case 0:
		ret := ctx.CreateReturnParams(8, 0)
		r0, err := s.Read(ctx)
		ret.Set64(0, math.Float64bits(r0))
		return err
	default:
		return C.ErrWrongMethod
	}
}

type CalculatorFunction interface {
	C.Capability
	Call(ctx C.Call, params C.Float64List) (float64, error)
}

func CalculatorFunction_Interface(ctx CallContext, obj Capability, method uint16, params Struct) error {
	if ctx.InterfaceId() != 0x.... {
		return C.ErrWrongInterface
	}

	switch s := obj.(CalculatorValue); method {
	case 0:
		a0 := params.GetObject(0).ToFloat64List()
		ret := ctx.CreateReturnParams(8, 0)
		r0, err := s.Call(ctx, a0)
		ret.Set64(0, math.Float64bits(r0))
		return err
	default:
		return C.ErrWrongMethod
	}
}

type CalculatorOperator uint16

const (
	CALCULATOR_OPERATOR_ADD CalculatorOperator = 0
	CALCULATOR_OPERATOR_SUBTRACT = 1
	CALCULATOR_OPERATOR_MULTIPLY = 2
	CALCULATOR_OPERATOR_DIVIDE = 3
)

