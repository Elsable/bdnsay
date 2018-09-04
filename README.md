# bdnsay
A simple example of Go Web assembly based on the go documentation :
https://github.com/golang/go/wiki/WebAssembly and this article: https://tutorialedge.net/golang/go-webassembly-tutorial/

ASCII art found here: https://asciiart.website/index.php?art=holiday/christmas/snowman

bdn = Bonhomme de neige, snowman in french.

```
< Hello world > 
---------------
       \
        \_[_]_  
          (")  
      '--( : )--'
        (  :  )
      ""'-...-'""
```

## Run the example
```GOARCH=wasm GOOS=js go build -o bdnsay.wasm main.go```
and 
```go run html.go```
then go to http://localhost:8080/wasm_exec.html
