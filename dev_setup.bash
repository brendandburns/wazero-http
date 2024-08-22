#!/bin/bash

BINDGEN_VERSION=0.30.0

wget https://github.com/bytecodealliance/wit-bindgen/releases/download/v${BINDGEN_VERSION}/wit-bindgen-${BINDGEN_VERSION}-x86_64-linux.tar.gz
tar -xvzf wit-bindgen-${BINDGEN_VERSION}-x86_64-linux.tar.gz
sudo mv wit-bindgen-${BINDGEN_VERSION}-x86_64-linux/wit-bindgen /usr/local/bin/
rm -rf wit-bindgen-${BINDGEN_VERSION}-x86_64-linux*

WASI_SDK_VERSION=24

wget https://github.com/WebAssembly/wasi-sdk/releases/download/wasi-sdk-${WASI_SDK_VERSION}/wasi-sdk-${WASI_SDK_VERSION}.0-x86_64-linux.deb
sudo dpkg -i wasi-sdk-${WASI_SDK_VERSION}.0-x86_64-linux.deb
rm wasi-sdk-${WASI_SDK_VERSION}.0-x86_64-linux.debgit co

