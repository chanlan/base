package midware

import (
	"testing"
	"net/http"
	"fmt"
)

var test http.HandlerFunc = func(res http.ResponseWriter, req *http.Request) {
	fmt.Println("testing...................")
}

func Test_NewEngine(t *testing.T){
	return
	engine := new(ServerEngine)
	engine.Host = "127.0.0.1"
	engine.Port = 8080
	engine.Domain = ""
	engine.Debug = false 
	engine.WriteTimeout = DefaultWriteTimeout
	engine.ReadTimeout = DefaultReadTimeout
	routing := map[string]http.HandlerFunc{
		"/":     test,
	}
	engine.Routing(routing)
	engine.NewEngine()
}