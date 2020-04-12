package loads

import (
	"jvmbyselfgo/src/go/ch3/instructions/base"
	"jvmbyselfgo/src/go/ch3/rtda"
)

func _aload(frame rtda.Frame, index uint) {
	ref := frame.LocalVars.GetInt(index)
	frame.OperandStack.PushInt(ref)
}

type ALOAD struct {
	base.Index8Instruction
}

func (self *ALOAD) FetchOperands(reader *base.BytecodeReader) {
	panic("implement me")
}

func (ALOAD) Execute(frame *rtda.Frame) {
	panic("implement me")
}
