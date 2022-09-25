package zilorot

import (
	"fmt"
	"log"
	"os"
	"path"
	"sync"
	"testing"
)

func TestLog(t *testing.T) {
	os.Mkdir("logs", 0777)
	logger := NewLogger("./logs/3proxy_status.log", 2, 2, 2)
	//defer logger.Writer().(*Logger).ForceCleanup()

	//logger := NewLogger("test.log", 5, 10, 10)

	var wg sync.WaitGroup

	wg.Add(3)
	fn := func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			logger.Printf("Some data sdfjh jhsdfk hskdfj hskdhf kshdfkshd kjh %10000s", "txt")
		}
	}
	go fn()
	go fn()
	fn()

	wg.Wait()
}

func NewLogger(name string, log_max_size_in_mb int, max_backups int, max_age_in_days int) *log.Logger {
	e, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("error zilorot log file: %v", err)
		if os.Stderr != nil {
			fmt.Fprintf(os.Stderr, "error zilorot log file: %v", err)
		}

		os.Exit(1)
	}
	logger := log.New(e, "", log.Ldate|log.Ltime)

	logger.SetOutput(&Logger{
		Filename:   name,
		filedir:    path.Dir(name),
		MaxSize:    log_max_size_in_mb, // megabytes
		MaxBackups: max_backups,
		MaxAge:     max_age_in_days, //days
	})

	return logger
}
