package main

import (
	"LearnGoService/Handlers"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	// pl := data.GetProducts()
	// fmt.Println("productlist:", pl, len(pl))
	//hh := &Handlers.Hello{}
	//gh := &Handlers.GoodBye{}
	//sm.Handle("/Goodbye", gh)
	ph := &Handlers.Products{}
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	serv := &http.Server{
		Addr:    ":9091",
		Handler: sm,
	}
	go func() {
		fmt.Println("going to Listen and Serve")
		serv.ListenAndServe()
	}()

	sgnchan := make(chan os.Signal)
	signal.Notify(sgnchan, os.Interrupt)
	signal.Notify(sgnchan, os.Kill)

	sgn := <-sgnchan
	fmt.Println("Going to shutdown:", sgn.String())
	ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	serv.Shutdown(ctx)

}
