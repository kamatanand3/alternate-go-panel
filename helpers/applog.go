package helpers

import (
	"log"
	"os"
	"sync"
)

var (
	loggers   = make(map[string]*log.Logger)
	mu        sync.Mutex
	requestId string
	sourceApp string
	channel   = "app" // default channel
)

// getLogger returns a logger for a specific channel (file)
func getLogger(channel string) *log.Logger {
	mu.Lock()
	defer mu.Unlock()

	if logger, exists := loggers[channel]; exists {
		return logger
	}

	file, err := os.OpenFile("logs/"+channel+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("❌ Could not open log file:", err)
	}

	logger := log.New(file, "", log.LstdFlags)
	loggers[channel] = logger
	return logger
}

// ---------- Utility setters ----------
func SetRequestId(rid string) {
	requestId = rid
}
func GetRequestId() string {
	return requestId
}

func SetSourceApp(app string) {
	sourceApp = app
}
func GetSourceApp() string {
	return sourceApp
}

func SetChannel(ch string) {
	channel = ch
}

// ---------- Logging methods ----------
func Info(className, methodName, description string, data interface{}) {
	getLogger(channel).Printf("%s|%s|%s|%s|%v", requestId, className, methodName, description, data)
}

func Debug(className, methodName, description string, data interface{}) {
	getLogger(channel).Printf("DEBUG|%s|%s|%s|%s|%v", requestId, className, methodName, description, data)
}

func Error(className, methodName, description string, data interface{}) {
	getLogger(channel).Printf("ERROR|%s|%s|%s|%s|%v", requestId, className, methodName, description, data)
}
