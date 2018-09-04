# bdnsay
A simple example of Go Web assembly based on the go documentation :
<https://github.com/golang/go/wiki/WebAssembly>

## Run the example
```GOARCH=wasm GOOS=js go build -o bdnsay.wasm main.go```
and 
```go run html.go```
then go to http://localhost:8080/wasm_exec.html
