package main

import (
	"context"
	"os"
	"time"

	"github.com/bartventer/log"
)

func main() {
	logger := log.New(
		log.WithLevel(log.DebugLevel),
		log.WithWriter(os.Stdout),
		log.WithPrefix("DemoApp"),
		log.WithReportTimestamp(true),
	)

	logger.Info("Starting the application...")
	time.Sleep(2 * time.Second)

	logger.Debug("Initializing resources...")
	time.Sleep(2 * time.Second)

	logger.Info("Processing task 1...")
	time.Sleep(2 * time.Second)
	logger.Info("Task 1 completed successfully.")
	time.Sleep(2 * time.Second)

	logger.Warn("Processing task 2 with a warning...")
	time.Sleep(2 * time.Second)
	logger.Info("Task 2 completed with warnings.")
	time.Sleep(2 * time.Second)

	logger.Error("Processing task 3 encountered an error!")
	time.Sleep(2 * time.Second)
	logger.Info("Task 3 failed.")
	time.Sleep(2 * time.Second)

	// Log with context
	ctx := context.Background()
	logger.InfoContext(ctx, "Shutting down the application...")
	time.Sleep(2 * time.Second)
}
