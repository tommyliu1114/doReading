package routers

import (
	web_socket "clean-web/routers/web-socket"
	"github.com/gorilla/mux"
)

var(
	GRouter *mux.Router
)
func InitRouters() *mux.Router{
	GRouter = mux.NewRouter()
	GRouter.HandleFunc("/echo",web_socket.Chat)
	GRouter.HandleFunc("/",web_socket.FrontServe)
	return GRouter
}



