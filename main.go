// +build js,wasm

package main

import (
	"fmt"
	"strings"
	"syscall/js"
)

const (
	signature = `
       \
        \_[_]_  
          (")  
      '--( : )--'
        (  :  )
      ""'-...-'""
`
)

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}

func registerCallbacks() {
	js.Global().Set("bdnSay", js.NewCallback(bdnSay))
}

func say(text string) string {
	if text == "" {
		text = "nothing"
	}

	sb := strings.Builder{}

	size := len(text) + 3
	sb.WriteString(printChar("_", size))
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("< %s > \n", text))
	sb.WriteString(printChar("-", size))
	sb.WriteString(signature)

	return sb.String()
}

func bdnSay(i []js.Value) {
	text := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()

	result := say(text)

	js.Global().Get("document").Call("getElementById", i[1].String()).Set("innerHTML", result)
}

func printChar(c string, time int) string {
	sb := strings.Builder{}
	for i := 0; i <= time; i++ {
		sb.WriteString(c)
	}

	return sb.String()
}
