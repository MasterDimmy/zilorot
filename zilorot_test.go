package zilorot

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	logger := NewLogger("test.log", 1, 3, 1)

	fn := func() {
		for i := 0; i < 100000; i++ {
			logger.Printf("Some data")
		}
	}
	go fn()
	go fn()
	fn()
}

func NewLogger(name string, log_max_size_in_mb int, max_backups int, max_age_in_days int) *log.Logger {
	e, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("error zilorot log file: %v", err)
		os.Exit(1)
	}
	logger := log.New(e, "", log.Ldate|log.Ltime)

	logger.SetOutput(&Logger{
		Filename:   name,
		MaxSize:    log_max_size_in_mb, // megabytes
		MaxBackups: max_backups,
		MaxAge:     max_age_in_days, //days
	})

	return logger
}
