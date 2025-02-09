package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"the-car-wash-directory/internal/server"
	"the-car-wash-directory/internal/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func run(ctx context.Context, args []string, stderr, stdout *os.File) error {
	// Setup signal handling for graceful shutdown
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// Example: load config from env
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// ------------------------------------------------------------
	// 1. Create a COLOR-ENABLED Zap logger (development style)
	// ------------------------------------------------------------
	// We create a custom console encoder config
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder // color for INFO, WARN, ERROR, etc.
	// Optionally tweak time/other keys here if you want

	// Build a core using this console encoder
	consoleEncoder := zapcore.NewConsoleEncoder(encoderCfg)

	// Write logs to stdout (or use a file if desired)
	logLevel := zapcore.DebugLevel // show all logs
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(stdout), logLevel)

	// Construct the *zap.Logger
	logger := zap.New(core)
	defer logger.Sync()

	// Get a sugared logger
	sugar := logger.Sugar()

	// Example: log a startup message
	sugar.Infow("Starting car wash service", "time", time.Now().Format(time.RFC3339))

	// Example: create domain services
	carWashService := services.NewCarWashService(sugar)

	// Build our Gin server
	engine := server.NewServer(sugar, carWashService) // *gin.Engine

	// Start listening in a goroutine
	listenAddr := net.JoinHostPort(host, port)
	srv := &httpServer{
		engine:  engine,
		address: listenAddr,
		logger:  sugar, // *zap.SugaredLogger (colored output)
	}
	go srv.start()

	// Wait for graceful shutdown
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		srv.shutdown(10 * time.Second)
	}()
	wg.Wait()
	return nil
}

// httpServer wraps gin.Engine for graceful shutdown
type httpServer struct {
	engine  *gin.Engine
	address string
	logger  *zap.SugaredLogger
	server  *http.Server
}

func (h *httpServer) start() {
	h.server = &http.Server{
		Addr:    h.address,
		Handler: h.engine,
	}

	h.logger.Infof("Listening on %s\n", h.address)
	if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		h.logger.Infof("ListenAndServe error: %v\n", err)
	}
}

func (h *httpServer) shutdown(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	h.logger.Infof("Shutting down server gracefully...")
	if err := h.server.Shutdown(ctx); err != nil {
		h.logger.Infof("Server shutdown error: %v", err)
	}
	h.logger.Infof("Server stopped.")
}
