# Project name
BINARY_NAME=game
WASM_OUTPUT=web/main.wasm

.PHONY: all run build-web serve clean

# Default target: build for desktop
all: run

# Run the game natively for debugging
run:
	go run .

# Build for WebAssembly
build-web:
	GOOS=js GOARCH=wasm go build -o $(WASM_OUTPUT)

# Serve the web folder with a local web server
serve:
	cd web && python3 -m http.server

# Clean generated WASM files
clean:
	rm -f $(WASM_OUTPUT)
