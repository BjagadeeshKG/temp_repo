package Handlers

import (
	"net/http"
)

type Hello struct {
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	//bdy, _ := ioutil.ReadAll(r.Body)
	//strbdy := "ric" //string(bdy)
	//rw.Write([]byte("Hello" + strbdy))
}
