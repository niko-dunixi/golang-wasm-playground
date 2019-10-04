package main

import (
	"syscall/js"
)

func main() {
	runForever := make(chan bool)
	window := js.Global()
	canvas := window.Get("document").Call("getElementById", "canvas")
	canvasCtx := canvas.Call("getContext", "2d")
	canvasCtx.Set("fillStyle", "red")
	canvasCtx.Set("strokeStyle", "red")
	canvasCtx.Set("lineWidth", 5)

	var renderer js.Func
	renderer = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// Make sure the canvas stretches the whole way through the page
		height := window.Get("innerHeight").Float()
		width := window.Get("innerWidth").Float()
		canvas.Set("innerHeight", height)
		canvas.Set("innerWidth", width)
		canvasCtx.Call("clearRect", 0, 0, width, height)
		{
			squareWidth := width / 8
			squareHeight := height / 8
			for i := 0; i < 8; i++ {
				for j := 0; j < 8; j++ {
					//canvasCtx.Call("fillRect", float64(i)*squareWidth, float64(j)*squareHeight, squareWidth, squareHeight)
					canvasCtx.Call("strokeRect", float64(i)*squareWidth, float64(j)*squareHeight, squareWidth, squareHeight)
				}
			}
		}
		// request a frame draw
		window.Call("requestAnimationFrame", renderer)
		return nil
	})
	window.Call("requestAnimationFrame", renderer)

	<-runForever
}
