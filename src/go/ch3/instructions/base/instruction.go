package base

import "jvmbyselfgo/jvmbyselfgo/src/go/ch3/rtda"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}

/**
没有操作数的指令
*/
type NoOperandsInstruction struct {
}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

func (self *NoOperandsInstruction) Execute(frame *rtda.Frame) {

}

/**
跳转指令
*/
type BranchInstruction struct {
	Offset int
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

func (self *BranchInstruction) Execute(frame *rtda.Frame) {

}

/**
存储和加载类指令存取局部变量表
*/
type Index8Instruction struct {
	Index uint
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

func (Index8Instruction) Execute(frame *rtda.Frame) {
	panic("implement me")
}

/**
运行时常量池缩影
*/
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}

func (Index16Instruction) Execute(frame *rtda.Frame) {
	panic("implement me")
}
