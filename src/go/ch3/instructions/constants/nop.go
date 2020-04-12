package constants

import (
	"jvmbyselfgo/src/go/ch3/instructions/base"
	"jvmbyselfgo/src/go/ch3/rtda"
)

type NOP struct{}

func (NOP) FetchOperands(reader *base.BytecodeReader) {
	panic("implement me")
}

func (NOP) Execute(frame *rtda.Frame) {

}
