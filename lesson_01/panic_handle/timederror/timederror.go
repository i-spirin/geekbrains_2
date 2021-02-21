package timederror

import (
	"fmt"
	"time"
)

// TimedError is struct for handling error with time
type TimedError struct {
	happenedAt time.Time
	err        string
}

// Error with time
func (t *TimedError) Error() string {
	return t.happenedAt.String() + " - " + t.err
}

// New error with time will be created
func New(err interface{}) error {
	return &TimedError{
		happenedAt: time.Now(),
		err:        fmt.Sprintf("%v", err),
	}
}
