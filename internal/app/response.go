package app

import (
	"encoding/json"
	"net/http"
)

// Response implements standard JSON response payload structure.
type Response struct {
	Status string          `json:"status"`
	Error  *ResponseError  `json:"error,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

// ResponseError implements the standard Error structure to return in response payloads.
type ResponseError struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

func (e ResponseError) Error() string {
	j, err := json.Marshal(e)
	if err != nil {
		return "ResponseError: " + err.Error()
	}
	return string(j)
}

// fail ends an unsuccessful JSON response with the standard
// failure format for services.
func fail(w http.ResponseWriter, status, errCode int, details ...string) {
	msg, ok := errMap[errCode]
	if !ok {
		errCode = status
		msg = http.StatusText(status)
	}
	r := &Response{
		Status: StatusFail,
		Error: &ResponseError{
			Code:    errCode,
			Message: msg,
			Details: details,
		},
	}
	j, err := json.Marshal(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(j)
}

// send sends a successful JSON response.
func send(w http.ResponseWriter, status int, result interface{}) {
	rj, err := json.Marshal(result)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	r := &Response{
		Status: StatusOK,
		Result: rj,
	}
	j, err := json.Marshal(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(j)
}

// ResponseStatus constants
const (
	StatusOK   = "ok"
	StatusFail = "nok"
)

// Err codes map body error status codes
// to body error status messages.
const (
	ErrCodeServiceUnavailable = 1
)

var errMap = map[int]string{
	ErrCodeServiceUnavailable: "Service Unavailable",
}
