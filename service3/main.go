package main

import (
	"context"
	"fmt"
	"httpserver"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)
	//errInfo := http.ListenAndServe(":8080", mux)
	//if errInfo != nil {
	//	log.Fatalf("error %s\n", errInfo)
	//}
	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()
	log.Println("server begin")
	<-done
	log.Println("server stop")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed:%s", err)
	}
	log.Println("server exited properly\n")
}
func handleIndex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "entering service 2 handleindex")
	delay := httpserver.RandIntTime(10, 20)
	time.Sleep(time.Millisecond * time.Duration(delay))
	io.WriteString(w, "------------the http request service------------------------------------\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
	io.WriteString(w, fmt.Sprintf("res in %d ms", delay))

}
