package main

import "syscall/js"
import "github.com/linuxsuren/audacity-label-plugin/pkg"

func main() {
	c := make(chan struct{}, 0)

	println("Go WebAssembly Initialized")

	js.Global().Set("convert", js.FuncOf(convert))

	<-c
}

func convert(this js.Value, args []js.Value) interface{} {
	var txt string
	if len(args) > 0 {
		txt = pkg.Convert(args[0].String())
	}
	return js.ValueOf(txt)
}
