package main

import (
	"clean-web/routers"
	"net/http"
)

func main(){
	gr := routers.InitRouters()
	http.ListenAndServe(":8080",gr)
}
