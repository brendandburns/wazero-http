#include "wasi_http.h"
#include <stdlib.h>

int main() {
    char buff[64 * 1024];
    wasi_http_response_t response = {
        .body = buff,
        .body_max_len = 64 * 1024
    };
    char *authority = getenv("AUTHORITY");
    wasi_http_request(GET, HTTP, authority, "/get?some=arg&goes=here", NULL, &response);
    free_response(&response);

    wasi_http_request(POST, HTTP, authority, "/post", "{\"foo\": \"bar\"}", &response);
    free_response(&response);
    
    wasi_http_request(PUT, HTTP, authority, "/put", NULL, &response);
    free_response(&response);
    return 0;
}

// The wasm component hook for the 'main'
bool exports_wasi_cli_run_run() {
    return !main();
}
