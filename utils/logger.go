package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// LogLevel represents different log levels
type LogLevel string

const (
	DEBUG     LogLevel = "debug"
	INFO      LogLevel = "info"
	NOTICE    LogLevel = "notice"
	WARNING   LogLevel = "warning"
	ERROR     LogLevel = "error"
	CRITICAL  LogLevel = "critical"
	ALERT     LogLevel = "alert"
	EMERGENCY LogLevel = "emergency"
)

// LogChannel represents different log channels
type LogChannel string

const (
	STACK      LogChannel = "stack"
	SINGLE     LogChannel = "single"
	DAILY      LogChannel = "daily"
	APILOG     LogChannel = "apilog"
	APPLOG     LogChannel = "applog"
	QUERIESLOG LogChannel = "querieslog"
	STDERR     LogChannel = "stderr"
)


var (
	// Default log directory
	logDir   = "logs"
	logMutex sync.Mutex
)

// Initialize log directories only
func init() {
	// Create logs directory if it doesn't exist
	if err := os.MkdirAll(logDir, 0755); err != nil {
		log.Printf("Failed to create logs directory: %v", err)
	}

	// Create subdirectories for different channels
	channels := []string{"api", "app", "queries"}
	for _, channel := range channels {
		channelDir := filepath.Join(logDir, channel)
		if err := os.MkdirAll(channelDir, 0755); err != nil {
			log.Printf("Failed to create %s directory: %v", channel, err)
		}
	}
}

func GetRequestIDFromContext(c *gin.Context) string {
	if v, exists := c.Get("request_id"); exists {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return "" // or generate a fallback ID if needed
}

// addLog is the core logging function
func addLog(c *gin.Context, channel LogChannel, level LogLevel, description string, data interface{}, className, methodName string) {
	requestID := GetRequestIDFromContext(c)

	// Format the log message
	parts := []string{requestID}
	if className != "" {
		parts = append(parts, className)
	}
	if methodName != "" {
		parts = append(parts, methodName)
	}
	parts = append(parts, description, formatData(data))
	logMessage := strings.Join(parts, "|")

	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Get the appropriate log file path
	logFilePath := getLogFilePath(channel)

	// Write to log file
	writeToLogFile(logFilePath, timestamp, level, logMessage)

	//Also log to stderr for development
	// if channel == STDERR || os.Getenv("GIN_MODE") == "debug" {
	// 	log.Printf("[%s] %s: %s", level, timestamp, logMessage)
	// }
}

// AppLog logs an info message to the app log file
func AppLog(c *gin.Context, levelStr, description string, data interface{}, className, methodName string) {
	addLog(c, APPLOG, LogLevel(levelStr), description, data, className, methodName)
}

// ApiLog logs an info message to the api log file
func ApiLog(c *gin.Context, levelStr, description string, data interface{}, className, methodName string) {
	addLog(c, APILOG, LogLevel(levelStr), description, data, className, methodName)
}

func formatData(data interface{}) string {
	if data == nil {
		return "null"
	}
	b, err := json.Marshal(data)
	if err != nil {
		return fmt.Sprintf("%+v", data)
	}
	return string(b)
}

func getLogFilePath(channel LogChannel) string {
	today := time.Now().Format("2006-01-02")

	switch channel {
	case SINGLE:
		return filepath.Join(logDir, "go.log")
	case DAILY:
		return filepath.Join(logDir, fmt.Sprintf("go-%s.log", today))
	case APILOG:
		return filepath.Join(logDir, "api", fmt.Sprintf("go-%s.log", today))
	case APPLOG:
		return filepath.Join(logDir, "app", fmt.Sprintf("go-%s.log", today))
	case QUERIESLOG:
		return filepath.Join(logDir, "queries", fmt.Sprintf("go-%s.log", today))
	case STDERR:
		return "stderr"
	default:
		return filepath.Join(logDir, "go.log")
	}
}

// writeToLogFile writes the log message to the specified file
func writeToLogFile(filePath, timestamp string, level LogLevel, message string) {
	logMutex.Lock()
	defer logMutex.Unlock()
	if filePath == "stderr" {
		log.Printf("[%s] %s: %s", level, timestamp, message)
		return
	}

	// Open file in append mode, create if doesn't exist
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed to open log file %s: %v", filePath, err)
		return
	}
	defer file.Close()

	// Write log entry
	logEntry := fmt.Sprintf("[%s] %s.%s: %s\n",
		timestamp,
		os.Getenv("APP_ENV"), // or your env variable
		strings.ToUpper(string(level)),
		message,
	)
	if _, err := file.WriteString(logEntry); err != nil {
		log.Printf("Failed to write to log file %s: %v", filePath, err)
	}
	// TODO: Add log rotation in the future
}
