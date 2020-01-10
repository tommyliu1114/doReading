package web_socket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:1024,
		WriteBufferSize:1024,
	}
)
func Chat(w http.ResponseWriter,r *http.Request){
	conn,_ := upgrader.Upgrade(w,r,nil)
	for{
		msgType,msg,err := conn.ReadMessage()
		if err != nil{
			return
		}
		fmt.Printf("%s sent: %s\n",conn.RemoteAddr(),string(msg))
		if err = conn.WriteMessage(msgType,msg);err != nil{
			return
		}
	}
}
