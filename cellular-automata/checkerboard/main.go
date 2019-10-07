package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	runForever := make(chan bool)
	window := js.Global()
	canvas := window.Get("document").Call("getElementById", "canvas")
	window.Call("addEventListener", "resize", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		width, height := updateCanvasSize(window, canvas)
		fmt.Printf("resize triggered: %f x %f\n", width, height)
		return nil
	}))
	_, _ = updateCanvasSize(window, canvas)

	canvasCtx := canvas.Call("getContext", "2d")

	var renderer js.Func
	renderer = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		width, height := getWindowSize(window)
		canvasCtx.Call("clearRect", 0, 0, width, height)
		canvasCtx.Set("fillStyle", "red")
		canvasCtx.Set("strokeStyle", "red")
		canvasCtx.Set("lineWidth", 5)
		fmt.Printf("Renderer called: %s %s %f\n",
			canvasCtx.Get("fillStyle").String(), canvasCtx.Get("strokeStyle").String(), canvasCtx.Get("lineWidth").Float())
		{
			squareWidth := width / 8
			squareHeight := height / 8
			offsetWidth := squareWidth * 0.15
			offsetHeight := squareHeight * 0.15
			for i := 0; i < 8; i++ {
				for j := 0; j < 8; j++ {
					if (i+j)%2 == 0 {
						canvasCtx.Call("strokeRect", float64(i)*squareWidth+offsetWidth, float64(j)*squareHeight+offsetHeight,
							squareWidth-(2*offsetWidth), squareHeight-(2*offsetHeight))
					} else {
						canvasCtx.Call("fillRect", float64(i)*squareWidth, float64(j)*squareHeight, squareWidth, squareHeight)
					}
				}
			}
		}
		window.Call("requestAnimationFrame", renderer)
		return nil
	})
	window.Call("requestAnimationFrame", renderer)

	<-runForever
}

func getWindowSize(window js.Value) (float64, float64) {
	width := window.Get("innerWidth").Float()
	height := window.Get("innerHeight").Float()
	return width, height
}

func updateCanvasSize(window js.Value, canvas js.Value) (float64, float64) {
	width, height := getWindowSize(window)
	canvas.Set("width", width)
	canvas.Set("height", height)
	return width, height
}
