package order

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path"
	"sync"
	"sync/atomic"
	"time"

	"github.com/HyperloopUPV-H8/h9-backend/pkg/abstraction"
	"github.com/HyperloopUPV-H8/h9-backend/pkg/logger"
	"github.com/HyperloopUPV-H8/h9-backend/pkg/transport/packet/data"
)

const (
	Name abstraction.LoggerName = "order"
)

type Logger struct {
	// An atomic boolean is used in order to use CompareAndSwap in the Start and Stop methods
	running  *atomic.Bool
	fileLock *sync.RWMutex
	writer   io.WriteCloser
}

type Record struct {
	Packet    *data.Packet
	From      string
	To        string
	Timestamp time.Time
}

func (*Record) Name() abstraction.LoggerName {
	return Name
}
func (record *Record) GetFrom() string         { return record.From }
func (record *Record) GetTo() string           { return record.To }
func (record *Record) GetTimestamp() time.Time { return record.Timestamp }

func NewLogger() *Logger {
	return &Logger{
		running:  &atomic.Bool{},
		fileLock: &sync.RWMutex{},
		writer:   nil,
	}
}

func (sublogger *Logger) Start() error {
	if !sublogger.running.CompareAndSwap(false, true) {
		fmt.Println("Logger already running")
		return nil
	}

	fmt.Println("Logger started")
	return nil
}

func (sublogger *Logger) PushRecord(record abstraction.LoggerRecord) error {
	if !sublogger.running.Load() {
		return &logger.ErrLoggerNotRunning{
			Name:      Name,
			Timestamp: time.Now(),
		}
	}

	orderRecord, ok := record.(*Record)
	if !ok {
		return &logger.ErrWrongRecordType{
			Name:      Name,
			Timestamp: time.Now(),
			Expected:  &Record{},
			Received:  record,
		}
	}

	sublogger.fileLock.Lock()
	defer sublogger.fileLock.Unlock()

	if sublogger.writer == nil {
		filename := path.Join("logger/order", fmt.Sprintf("order_%s", logger.Timestamp.Format(time.RFC3339)), "order.csv")
		os.MkdirAll(path.Dir(filename), os.ModePerm)

		file, err := os.Create(filename)
		if err != nil {
			return &logger.ErrCreatingFile{
				Name:      Name,
				Timestamp: time.Now(),
				Inner:     err,
			}
		}
		sublogger.writer = file
	}

	csvWriter := csv.NewWriter(sublogger.writer)
	defer csvWriter.Flush()

	timestamp := orderRecord.Packet.Timestamp().Format(time.RFC3339)
	val := fmt.Sprint(orderRecord.Packet.GetValues())
	err := csvWriter.Write([]string{timestamp, val, record.GetFrom(), record.GetTo(), record.GetTimestamp().Format(time.RFC3339)})
	if err != nil {
		return err
	}

	return nil
}

func (sublogger *Logger) PullRecord(abstraction.LoggerRequest) (abstraction.LoggerRecord, error) {
	panic("TODO!")
}

func (sublogger *Logger) Stop() error {
	if !sublogger.running.CompareAndSwap(true, false) {
		fmt.Println("Logger already stopped")
		return nil
	}

	sublogger.writer.Close()

	fmt.Println("Logger stopped")
	return nil
}
