wasi_sdk := /usr/local/lib/wasi-sdk-21.0
cc := ${wasi_sdk}/bin/clang
ar := ${wasi_sdk}/bin/ar
lib_objs := client.o wasi_http.o

.phony: gen clean run

default: main.wasm

client.o : client.c
	${cc} -c client.c -o client.o

%.o : %.c
	@echo "Compiling c file into o file"
	${cc} -c $< -o $@

libwasihttp.a: ${lib_objs}
	${ar} cr libwasihttp.a ${lib_objs}

wasi-http: ; git clone https://github.com/WebAssembly/wasi-http; cd wasi-http; git checkout v0.2.0 ; cp ../client.wit wit/client.wit; cd ../

gen: wasi-http ; wit-bindgen c --autodrop-borrows yes ./wasi-http/wit -w client

main.wasm: gen main.o libwasihttp.a; ${cc} main.o -L. -lwasihttp -o main.wasm

main_2023_11_10.embed.wasm: main.wasm ; wasm-tools component embed ./wasi-http/wit main.wasm -o main_2023_11_10.embed.wasm -w client

wasi_snapshot_preview1.reactor.wasm: ; wget https://github.com/bytecodealliance/wasmtime/releases/download/v15.0.1/wasi_snapshot_preview1.reactor.wasm

main_2023_11_10.component.wasm: main_2023_11_10.embed.wasm wasi_snapshot_preview1.reactor.wasm; wasm-tools component new main_2023_11_10.embed.wasm -o main_2023_11_10.component.wasm --adapt wasi_snapshot_preview1.reactor.wasm

clean: ; rm -f client.c *.o client.h *.wasm *.a; rm -rf wasi-http

run: main_2023_11_10.component.wasm ; wasmtime -S http --wasm component-model main_2023_11_10.component.wasm

server.wasm: gen server.o libwasihttp.a; ${cc} server.o -L. -lwasihttp -o server.wasm

server_2023_11_10.embed.wasm: server.wasm ; wasm-tools component embed ./wasi-http/wit server.wasm -o server_2023_11_10.embed.wasm -w proxy

server_2023_11_10.component.wasm: server_2023_11_10.embed.wasm wasi_snapshot_preview1.reactor.wasm; wasm-tools component new server_2023_11_10.embed.wasm -o server_2023_11_10.component.wasm --adapt wasi_snapshot_preview1.reactor.wasm

run-server: server_2023_11_10.component.wasm ; wasmtime serve server_2023_11_10.component.wasm
