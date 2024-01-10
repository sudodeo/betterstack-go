package logs

import "fmt"

// BetterStackError represents an error returned by the BetterStack SDK.
type BetterStackError struct {
	StatusCode int    // The HTTP status code associated with the error.
	Message    string // The error message.
}

func (err *BetterStackError) Error() string {
	return fmt.Sprintf(`Error failed with status code: %d\nerror: %v`, err.StatusCode, err.Message)
}
