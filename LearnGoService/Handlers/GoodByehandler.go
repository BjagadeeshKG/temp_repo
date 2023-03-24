package Handlers

import (
	"net/http"
)

type GoodBye struct {
}

func (h *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	//bdy, _ := ioutil.ReadAll(r.Body)
	strbdy := "ric" //string(bdy)
	rw.Write([]byte("GoodBye" + strbdy))
}
