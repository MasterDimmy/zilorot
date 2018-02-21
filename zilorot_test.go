package zilorot

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	os.Mkdir("logs", 0777)
	logger := NewLogger("./logs/3proxy_status.log", 2, 2, 2)

	//logger := NewLogger("test.log", 5, 10, 10)

	fn := func() {
		for i := 0; i < 1000000; i++ {
			logger.Printf("Some data sdfjh jhsdfk hskdfj hskdhf kshdfkshd kjh")
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
