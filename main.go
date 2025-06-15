package main

import (
	"io"
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// Configure lumberjack for log rotation
	rotatingLog := &lumberjack.Logger{
		Filename:   "rotated-app.log", // Log file name
		MaxSize:    5,                 // Megabytes
		MaxBackups: 3,                 // Number of old log files to keep
		MaxAge:     7,                 // Days
		Compress:   true,
	}

	// Optional: also write to console
	multiOutput := os.Stdout
	writer := slog.NewJSONHandler(io.MultiWriter(rotatingLog, multiOutput), nil)

	// Create the logger
	logger := slog.New(writer)

	// Example logs
	for i := 1; i <= 1000; i++ {
		logger.Info("Processing record", "record_id", i)
	}
}
