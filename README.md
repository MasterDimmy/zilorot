# zilorot
Golang library to store log file with rotation, size limitation, time-based rotating and storing data in zip-format for backups.

Based on https://github.com/natefinch/lumberjack , but added ability to store data in zip format.


Usage example:

import (
	"fmt"
	"github.com/MasterDimmy/zilorot"
)

//NewLogger return log.Logger with rotating, size limitation, time-based rotation and storing data in zip-format for backups.
//func NewLogger(name string, log_max_size_in_mb int, max_backups int, max_age_in_days int) *log.Logger

func main() {
	logger := zilorot.NewLogger("test.log", 1, 3, 1)

	fn := func() {
		for i := 0; i < 100000; i++ {
			logger.Printf("Some data")
		}
	}
	go fn()
	go fn()
	fn()
}

