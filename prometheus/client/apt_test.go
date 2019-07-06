package prometheus_client

import (
	"context"
	"testing"

	"github.com/vadv/gopher-lua-libs/http"
	"github.com/vadv/gopher-lua-libs/strings"
	"github.com/vadv/gopher-lua-libs/time"
	lua "github.com/yuin/gopher-lua"
)

func TestApi(t *testing.T) {
	state := lua.NewState()
	state.SetContext(context.Background())
	Preload(state)
	http.Preload(state)
	strings.Preload(state)
	time.Preload(state)
	if err := state.DoFile("./test/test_api.lua"); err != nil {
		t.Fatalf("execute test: %s\n", err.Error())
	}
}