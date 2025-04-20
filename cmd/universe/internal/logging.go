package sphere_platform_universe_internal

import (
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

// SetupLogging - setup the logging.
func SetupLogging(rootDirectory string) {

	logDirectory := filepath.Join(rootDirectory, "logs")

	logFile := time.Now().Format("2006-01-02_15-04-05") + "_log.json"

	// Ensure the log directory exists
	if err := os.MkdirAll(logDirectory, os.ModePerm); err != nil {
		panic("Failed to create log directory: " + err.Error())
	}

	// Open the log file
	filePath := filepath.Join(logDirectory, logFile)
	logFileHandle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Failed to open log file: " + err.Error())
	}

	// Create a JSON logger that writes to the file
	logger := slog.New(slog.NewJSONHandler(logFileHandle, nil))
	slog.SetDefault(logger)

	slog.Info("Logging setup completed", "logFile", filePath)
}
