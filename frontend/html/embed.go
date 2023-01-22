package html

import (
	"embed"
)

//go:embed index.html wasm_exec.js test.wasm wasm_exec.js
var Assets embed.FS
