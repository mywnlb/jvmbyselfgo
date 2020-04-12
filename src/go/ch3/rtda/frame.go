package rtda

type Frame struct {
	Lower *Frame
	//局部变量表指针
	LocalVars LocalVars
	//操作数栈指针
	OperandStack *OperandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		Lower:        nil,
		LocalVars:    newLocalVars(maxLocals),
		OperandStack: newOperandStack(maxStack),
	}
}
