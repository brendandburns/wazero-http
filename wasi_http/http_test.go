package wasi_http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"github.com/tetratelabs/wazero/sys"
)

type handler struct {
	urls   []string
	bodies []string
}

func (h *handler) reset() {
	h.bodies = []string{}
	h.urls = []string{}
}

func (h *handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	body := ""

	if req.Body != nil {
		defer req.Body.Close()
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err.Error())
		}
		body = string(data)
	}

	res.WriteHeader(200)
	res.Write([]byte("Response"))

	h.urls = append(h.urls, req.URL.String())
	h.bodies = append(h.bodies, body)
}

func TestHttpClient(t *testing.T) {
	filePaths, _ := filepath.Glob("../testdata/c/main.wasm")
	for _, file := range filePaths {
		fmt.Printf("%v\n", file)
	}
	if len(filePaths) == 0 {
		t.Log("nothing to test")
		t.FailNow()
	}

	h := handler{}
	s := httptest.NewServer(&h)
	defer s.Close()

	expectedPaths := [][]string{
		{
			"/get?some=arg&goes=here",
			"/post",
			"/put",
		},
	}

	// TODO: Body for requests are not currently supported
	expectedBodies := [][]string{
		{
			"",
			"{\"foo\": \"bar\"}",
			"",
		},
	}

	for testIx, test := range filePaths {
		name := test
		for strings.HasPrefix(name, "../") {
			name = name[3:]
		}

		t.Run(name, func(t *testing.T) {
			bytecode, err := os.ReadFile(test)
			if err != nil {
				t.Fatal(err)
			}

			ctx := context.Background()

			runtime := wazero.NewRuntime(ctx)
			defer runtime.Close(ctx)

			wasi_snapshot_preview1.MustInstantiate(ctx, runtime)

			u, _ := url.Parse(s.URL)
			config := wazero.NewModuleConfig().
				WithStdout(os.Stdout).
				WithArgs("wasi").
				WithEnv("AUTHORITY", fmt.Sprintf("%s:%s", u.Hostname(), u.Port()))
			w := MakeWasiHTTP("2023_11_10")
			w.Instantiate(ctx, runtime)

			instance, err := runtime.InstantiateWithConfig(ctx, bytecode, config)
			if err != nil {
				switch e := err.(type) {
				case *sys.ExitError:
					if exitCode := e.ExitCode(); exitCode != 0 {
						t.Error("exit code:", exitCode)
						t.FailNow()
					}
				default:
					// t.Error(err.Error())
					// TODO: There is a problem with tearing down the module...
					fmt.Printf("Instantiating wasm module error: %s", err.Error())
				}
			}
			if instance != nil {
				if err := instance.Close(ctx); err != nil {
					t.Error("closing wasm module instance:", err)
				}
			}
			if !reflect.DeepEqual(expectedPaths[testIx], h.urls) {
				t.Errorf("Unexpected paths: %v vs %v", h.urls, expectedPaths[testIx])
			}
			if !reflect.DeepEqual(expectedBodies[testIx], h.bodies) {
				t.Errorf("Unexpected bodies: %v vs %v", h.bodies, expectedBodies[testIx])
			}

			h.reset()
		})
	}
}

func TestServer(t *testing.T) {
	filePaths := []string{"../testdata/c/server.wasm"}
	for _, file := range filePaths {
		fmt.Printf("%v\n", file)
	}
	if len(filePaths) == 0 {
		t.Log("nothing to test")
		t.FailNow()
	}

	for _, test := range filePaths {
		name := test
		for strings.HasPrefix(name, "../") {
			name = name[3:]
		}

		t.Run(name, func(t *testing.T) {
			bytecode, err := os.ReadFile(test)
			if err != nil {
				t.Fatal(err)
			}

			ctx := context.Background()

			runtime := wazero.NewRuntime(ctx)
			defer runtime.Close(ctx)

			wasi_snapshot_preview1.MustInstantiate(ctx, runtime)

			w := MakeWasiHTTP("2023_11_10")
			w.Instantiate(ctx, runtime)

			instance, err := runtime.Instantiate(ctx, bytecode)
			if err != nil {
				switch e := err.(type) {
				case *sys.ExitError:
					if exitCode := e.ExitCode(); exitCode != 0 {
						t.Error("exit code:", exitCode)
					}
				default:
					t.Error("instantiating wasm module instance:", err)
				}
			}
			if instance != nil {
				h := w.MakeHandler(ctx, instance)
				s := httptest.NewServer(h)
				defer s.Close()

				res, err := http.Get(s.URL)
				if err != nil {
					t.Error("Failed to read from server.")
				}
				defer res.Body.Close()

				data, err := ioutil.ReadAll(res.Body)
				if err != nil {
					t.Error("Failed to read body.")
				}
				u, _ := url.Parse(s.URL)
				obj := make(map[string]interface{})
				err = json.Unmarshal(data, &obj)
				if err != nil {
					t.Error("Failed to parse body.")
				}
				msg := obj["msg"].(string)
				expected := "Hello world!"
				if len(msg) != len(expected) {
					t.Errorf("lengths don't match!")
				}
				for ix := range msg {
					if msg[ix] != expected[ix] {
						t.Errorf("Character %d is wrong", ix)
					}
				}
				if strings.Compare(msg, expected) != 0 {
					t.Errorf("Unexpected message: '%s'", obj["msg"].(string))
				}
				if obj["authority"] != fmt.Sprintf("%s:%s", u.Hostname(), u.Port()) {
					t.Errorf("Unexpected authority: %s", obj["authority"])
				}

				if err := instance.Close(ctx); err != nil {
					t.Error("closing wasm module instance:", err)
				}
			}
		})
	}
}
