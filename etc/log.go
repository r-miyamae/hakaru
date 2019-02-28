package etc

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type HakaruLog struct {
	At    string `json:"at"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Logger struct {
	agent *log.Logger
}

func NewLogger(path string) *Logger {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	return &Logger{agent: log.New(f, "", 0)}
}

func (l *Logger) Hakaru(name string, value string) error {
	return l.Log(HakaruLog{
		At: time.Now().Format(time.RFC3339),
		Name: name,
		Value: value,
	})
}

func (l *Logger) Log(v interface{}) error {
	if s, err := LogToString(v); err == nil {
		l.agent.Println(s)
		return nil
	} else {
		return err
	}
}

func LogToString(v interface{}) (string, error) {
	j, err := json.Marshal(v)
	return string(j), err
}
