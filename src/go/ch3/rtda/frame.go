package rtda

type Frame struct {
	lower *Frame
	//局部变量表指针
	localVars LocalVars
	//操作数栈指针
	operandStack *OperandStack
}

func NewFrame(maxLocals,maxStack uint)*Frame  {
	return &Frame{
		lower:        nil,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}
