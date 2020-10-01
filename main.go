package logger

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/fatih/color"
)

// Logger empty struct
type Logger struct {}

func (l *Logger) Write(bytes []byte) (int, error) {
	return fmt.Fprintf(color.Output, "%s", string(bytes))
}