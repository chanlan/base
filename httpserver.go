package midware

import (
	"time"
	"github.com/gorilla/mux"
	"net/http"
	"net"
	"strconv"
)

const (
	DefaultWriteTimeout = 30 * time.Second
	DefaultReadTimeout = 60 * time.Second
)

type ServerEngine struct {
	Host string
	Port int
	Domain string
	Debug bool
	WriteTimeout time.Duration
	ReadTimeout time.Duration
	Router  *mux.Router
}

//默认
func NewDefaultEngine() *ServerEngine {
	return &ServerEngine{
		Host: "127.0.0.1",
		Port: 8080,
		Domain: "",
		Debug: false, 
		WriteTimeout: DefaultWriteTimeout,
		ReadTimeout: DefaultReadTimeout,
		Router: nil,
	}
}

//注册router
func (engine *ServerEngine) Routing(routing map[string]http.HandlerFunc) {
	router := mux.NewRouter()
	for url, handler := range routing {
		router.HandleFunc(url, handler)
	}
	engine.Router = router
}

//服务器监听
func (engine *ServerEngine) NewEngine() {
	server := &http.Server {
        Addr: net.JoinHostPort(engine.Host, strconv.Itoa(engine.Port)),
        ReadTimeout: DefaultReadTimeout,
        WriteTimeout: DefaultWriteTimeout,
        Handler: engine.Router,
    }

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}