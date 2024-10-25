package main

import (
	configuration "SpecForge_api_backend/internal/configuration"
	handler "SpecForge_api_backend/internal/handlers"
	globalUtility "SpecForge_api_backend/utilities/globalUtility"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			errTime := globalUtility.GetCurrentTimestamp()
			_, currentFile, _, _ := runtime.Caller(0)
			errorMsg := fmt.Sprintf("Some error occured in %v", currentFile)
			for i := 1; ; i++ {
				pc, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				if file == currentFile {
					funcName := runtime.FuncForPC(pc).Name()
					errorMsg = fmt.Sprintf("ERROR: [%s] [%s:%d] [%s] - %v", errTime, file, line, funcName, err)
					break
				}
			}
			fmt.Println(errorMsg)
		}
	} ()
	configuration.LoadConfig()
	startApiServer()
}

func startApiServer() {
	router := gin.New()

	//use middlewares here
	
	handler.HandleRoutes(router)

	serverPort := ""
	if len(os.Args) > 1 {
		serverPort = os.Args[1]
		if serverPort == "" {
			panic("No port passed.")
		}
	} else {
		panic("No arguments passed")
	}

	httpServer := &http.Server{
		Addr: ":" + serverPort,
		Handler: router,
	}
	go startListening(httpServer)
	closeServer(httpServer)
}

func startListening(httpServer *http.Server) {
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		errMsg := fmt.Sprintf("Some issue in listening and serve the request: %v", err)
		panic(errMsg)
	}
}

func closeServer(httpServer *http.Server) {
	quitSignChannel := make(chan os.Signal, 1)
	signal.Notify(quitSignChannel, os.Interrupt, syscall.SIGTERM)
	<-quitSignChannel
	fmt.Println("Shutting down server...")
	// Create a context with a timeout for the shutdown process
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Attempt to gracefully shut down the server
    if err := httpServer.Shutdown(ctx); err != nil {
        errMsg := fmt.Sprintf("Error during server shutdown: %v\n", err)
		panic(errMsg)
    }
    fmt.Println("Server shut down gracefully.")
}