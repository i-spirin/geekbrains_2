package timederror

import (
	"fmt"
	"time"
)

// TimedError is struct for handling error with time
type TimedError struct {
	HappenedAt time.Time
	Err        string
}

// Error with time
func (t *TimedError) Error() string {
	return t.HappenedAt.String() + " - " + t.Err
}

// New error with time will be created
func New(err interface{}) error {
	return &TimedError{
		HappenedAt: time.Now(),
		Err:        fmt.Sprintf("%v", err),
	}
}
