package main

import (
	"fmt"
	"hotelBotGUI/internal/cons"
	"hotelBotGUI/internal/webserver"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	mymiddleware "hotelBotGUI/internal/middleware"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	zrlog "github.com/rs/zerolog/log"
)

func SimpleAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey != "secret" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func StartWebServer(host string, port int) {

	router := chi.NewRouter()

	// Chaining middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(mymiddleware.LoggingMiddleware)
	//router.Use(SimpleAuthMiddleware)

	router.Use(middleware.Timeout(60 * time.Second))

	handler := webserver.NewHandler()
	handler.Register(router)

	// if err := http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), router); err != nil {
	// 	zrlog.Fatal().Msg(err.Error())
	// }

	srv := &http.Server{
		Addr:         fmt.Sprintf("%v:%v", host, port),
		ReadTimeout:  50 * time.Second,
		WriteTimeout: 50 * time.Second,
		//TLSConfig:    tlsConfig,
		Handler: router,
	}

	zrlog.Print(srv.ListenAndServeTLS(cons.CertPaht, cons.KeyPath))

}

func main() {

	logFile, err := os.OpenFile("./temp/info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		zrlog.Error().Msg(fmt.Sprintf("the file info.log doesn't open: %v", err))
		os.Exit(1)
	}
	defer logFile.Close()

	zrlog.Logger = zerolog.New(logFile).With().Timestamp().Logger()

	err = godotenv.Load("app.env")
	if err != nil {
		zrlog.Fatal().Msg("Error loading .env file: ")
		os.Exit(1)
	}

	host := os.Getenv("WEBSERVER_HOST")
	port, err := strconv.Atoi(os.Getenv("WEBSERVER_PORT"))
	if err != nil {
		port = 8081
	}

	StartWebServer(host, port)

}
