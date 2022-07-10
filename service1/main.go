package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)
	mux.HandleFunc("/healthz", healthz)
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
	delay := timeRandInt(10, 20)
	time.Sleep(time.Millisecond * time.Duration(delay))
	io.WriteString(w, "------------the http request------------------------------------\n")
	req, err := http.NewRequest("GET", "http://service2", nil)
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	lowerCaseHeader := make(http.Header)
	for key, value := range r.Header {
		lowerCaseHeader[strings.ToLower(key)] = value
	}
	log.Printf("header: %s\n", lowerCaseHeader)
	req.Header = lowerCaseHeader
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		io.WriteString(w, fmt.Sprintf("err info %s\n", err))
	} else {
		io.WriteString(w, "------------the http request------------------------------------\n")
	}
	resp.Write(w)
	io.WriteString(w, fmt.Sprintf("res in %d ms", delay))

}
func timeRandInt(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "service1\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
	io.WriteString(w, "ok\n")
}
