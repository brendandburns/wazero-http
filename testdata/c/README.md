# Experimental HTTP Client and Server Example
*NB*: this example uses an experimental `wasi-http` that incorporates an
experimental HTTP client library being developed as part of the WASI specification.
Use at your own risk, things may change in the future.

## Building the client example
```sh
make main.wasm
```

### Running
```sh
make clean; make run
```

## Building the server example
```sh
make server.wasm
```

### Running
```sh
make clean; make run-server
```
