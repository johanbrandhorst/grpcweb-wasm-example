generate:
	go run github.com/bufbuild/buf/cmd/buf@latest generate

	rm -f ./frontend/html/index.html ./frontend/html/wasm_exec.js ./frontend/html/test.wasm
	GOOS=js GOARCH=wasm go build -o ./frontend/html/test.wasm ./frontend/frontend.go
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.html ./frontend/html/index.html
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js ./frontend/html/wasm_exec.js
	sed -i -e 's;</button>;</button>\n\t<div id=\"target\"></div>;' ./frontend/html/index.html

serve:
	go run main.go
