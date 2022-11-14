package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/pprof"
	"sync"
	"time"
)

// Server describes HTTP API server
type Server struct {
	http http.Server
	mux  *http.ServeMux
}

const (
	httpAPITimeout  = time.Second * 100
	shutdownTimeout = time.Second * 10
)

func New(port int) *Server {
	s := &Server{}
	mux := http.NewServeMux()

	mux.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
	mux.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	mux.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	mux.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
	mux.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))
	s.mux = mux

	s.http = http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.TimeoutHandler(s.mux, httpAPITimeout, ""),
	}

	return s
}

// Run starts the HTTP server
func (s *Server) Run(stopCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	var err error

	go func() {
		if err2 := s.http.ListenAndServe(); err2 != http.ErrServerClosed {
			panic(fmt.Errorf("Could not start http server: %v", err2))
		}
	}()
	fmt.Printf("listening on %s \n", s.http.Addr)

	// wait for a shutdown signal on the channel
	<-stopCh

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err = s.http.Shutdown(ctx); err == nil {
		fmt.Println("Http server shut down")
		return
	}
	if err == context.DeadlineExceeded {
		fmt.Println("Shutdown timeout exceeded. closing http server")
		if err = s.http.Close(); err != nil {
			fmt.Printf("could not close http connection: %v \n", err)
		}
		return
	}
	fmt.Printf("Could not shutdown http server: %v \n", err)
}

func (s *Server) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	s.mux.HandleFunc(pattern, handler)
}

func main() {

	wg := new(sync.WaitGroup)
	wg.Add(1)

	server := New(8080)
	server.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header)
		b, _ := ioutil.ReadFile("resources/hosein.webp")
		w.Header().Add("content-type", "image/bmp")
		w.Write(b)
	})

	stop := make(chan struct{})
	go server.Run(stop, wg)

	// To shutdown the server after X seconds
	go func() {
		<-time.After(time.Second * 10)
		stop <- struct{}{}
	}()

	// Wait for the server to shutdown
	wg.Wait()
	fmt.Println("Party is over!")
}
