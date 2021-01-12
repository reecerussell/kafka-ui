package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"

	"github.com/gorilla/mux"

	"github.com/reecerussell/kafka-ui/config"
	"github.com/reecerussell/kafka-ui/handler"
	"github.com/reecerussell/kafka-ui/listener"
	"github.com/reecerussell/kafka-ui/logging"
	"github.com/reecerussell/kafka-ui/middleware"
	"github.com/reecerussell/kafka-ui/ws"
)

var (
	configFilename string
	hostAddress    string
)

func main() {
	flag.StringVar(&configFilename, "config-file", "config.json", "A JSON configuration file (default is config.json)")
	flag.StringVar(&hostAddress, "host", "127.0.0.1:8002", "The host address of the service.")
	flag.Parse()

	cnf := getConfig(configFilename)

	r := mux.NewRouter()
	wss := ws.New()
	mapRoutes(r, cnf, wss)

	ctx := context.Background()
	wg := sync.WaitGroup{}

	var (
		lis listener.Listener
		srv *http.Server
	)

	go func() {
		lis = startListener(ctx, cnf, &wg, wss.Send)
	}()

	go wss.Start()
	go func() {
		srv = startServer(r, &wg)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)
	<-stop

	fmt.Printf("\r") // clears "^C" on interrupt
	logging.Info("Gracfully shutting down")

	srv.Shutdown(context.Background())
	lis.Stop()

	wg.Wait()
	logging.Info("Shutdown")
}

func getConfig(filename string) *config.Config {
	logging.Info("Loading configuration from '%s'", filename)
	cnf, err := config.GetConfig(filename)
	if err != nil {
		logging.Error(err.Error())
		panic(err)
	}

	return cnf
}

func mapRoutes(r *mux.Router, cnf *config.Config, wss *ws.Server) {
	r.HandleFunc("/ws", wss.Serve)
	r.Handle("/api/topics", handler.GetTopics(cnf)).Methods("GET")
}

func startListener(ctx context.Context, cnf *config.Config, wg *sync.WaitGroup, send func(m *ws.Message)) listener.Listener {
	p := listener.NewProcessor(send)
	l := listener.New(ctx, cnf)

	go func() {
		wg.Add(1)
		defer wg.Done()

		logging.Info("Starting Kafka listener, on: %v", *cnf.Kafka.BootstrapServers)
		err := l.Start(p)
		if err != nil {
			panic(err)
		}
	}()

	return l
}

func startServer(h http.Handler, wg *sync.WaitGroup) *http.Server {
	srv := http.Server{
		Addr:    hostAddress,
		Handler: middleware.Pipe(h, middleware.NewLoggingMiddleware, middleware.NewCORSMiddleware),
	}

	go func() {
		wg.Add(1)
		defer wg.Done()

		logging.Info("Listening on: %s", hostAddress)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return &srv
}
