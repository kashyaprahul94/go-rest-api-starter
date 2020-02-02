package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	router "github.com/julienschmidt/httprouter"
)

// New creates the instance of new http web server \n
// It accepts a port as an argument on which http web server will listen on
func New(port string) *HTTPServer {

	router := router.New()
	router.RedirectTrailingSlash = false

	serverInstace := &HTTPServer{
		port: port,
		router: &Router{
			Router: router,
		},
	}

	return serverInstace
}

// GetSubRouter will create an empty route and return the subrouter instance
func (hs *HTTPServer) GetSubRouter(namespace string) *RouteGroup {
	return hs.router.NewGroup(namespace)
}

// GetPort will return the port used by the web server
func (hs *HTTPServer) GetPort() string {
	return hs.port
}

func (hs *HTTPServer) prepareGlobalMiddlewares() {
	hs.router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {

			header := w.Header()
			header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		w.WriteHeader(http.StatusNoContent)
	})
}

func prepareForGracefulShutdown(server *http.Server) {

	// Creat the channel for signal
	stop := make(chan os.Signal, 1)

	// Observe the `SIGINT` signal
	signal.Notify(stop, os.Interrupt)

	// Wait until we receive the message in the channel
	<-stop

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println(err.Error())
	}

	log.Println("shutting down web server")
	os.Exit(0)
}

// Listen will start listening
func (hs *HTTPServer) Listen(cb func()) {

	// Initialize the http server
	server := &http.Server{
		Addr:    strings.Join([]string{"0.0.0.0", hs.port}, ":"),
		Handler: hs.router,
	}

	// Prepare middlewares
	hs.prepareGlobalMiddlewares()

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println()
			log.Println("web server encountered error")
			fmt.Println()
			panic(err)
		}
	}()

	// Run the callbacck in a goroutine
	go cb()

	// prepare the webserver for graceful shutdown
	prepareForGracefulShutdown(server)
}
