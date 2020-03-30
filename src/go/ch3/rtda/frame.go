package rtda

type Frame struct {
	lower *Frame
	//局部变量表指针
	localVars LocalVars
	//操作数栈指针
	operandStack *OperandStack
}
