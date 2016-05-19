package sorter

import (
	"reflect"
	"sort"
)

type (
	FLess func(i, j int) bool
	FSwap func(i, j int)
)

type Wrapper struct {
	Length   int
	LessFunc FLess
	SwapFunc FSwap
}

func (w *Wrapper) Len() int {
	return w.Length
}

func (w *Wrapper) Less(i, j int) bool {
	return w.LessFunc(i, j)
}

func (w *Wrapper) Swap(i, j int) {
	w.SwapFunc(i, j)
}

func NewWrapper(length int, fLess FLess, fSwap FSwap) sort.Interface {
	return &Wrapper{
		Length:   length,
		LessFunc: fLess,
		SwapFunc: fSwap,
	}
}

type WrapperWith struct {
	Value    reflect.Value
	LessFunc FLess
}

func (w *WrapperWith) Len() int {
	return w.Value.Len()
}

func (w *WrapperWith) Less(i, j int) bool {
	return w.LessFunc(i, j)
}

func (w *WrapperWith) Swap(i, j int) {
	v1 := w.Value.Index(i).Interface()
	v2 := w.Value.Index(j).Interface()
	w.Value.Index(j).Set(reflect.ValueOf(v1))
	w.Value.Index(i).Set(reflect.ValueOf(v2))
}

func NewWrapperWith(arr interface{}, fLess FLess) sort.Interface {
	return &WrapperWith{
		Value:    reflect.ValueOf(arr),
		LessFunc: fLess,
	}
}
