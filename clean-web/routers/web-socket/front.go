package web_socket

import "net/http"

func FrontServe(w http.ResponseWriter,r * http.Request)  {
	http.ServeFile(w,r,"./frontpages/websockets.html")
}
