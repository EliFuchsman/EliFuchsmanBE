package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	NotFound  = NewOutput("not_found", "What you are looking for cannot be found.")
	Invalid   = NewOutput("invalid", "The value provided is invalid.")
	Internal  = NewOutput("internal_error", "An internal error occurred.")
	Forbidden = NewOutput("forbidden_error", "You are not authorized to access this resource")
)

type Output struct {
	Value string
	Desc  string
}

func NewOutput(value, desc string) *Output {
	return &Output{
		Value: value,
		Desc:  desc,
	}
}

func (o Output) String() string {
	return o.Value
}

func (o Output) ToUpper() string {
	return strings.ToUpper(o.String())
}

func (o Output) Description() string {
	return o.Desc
}

type Error struct {
	Message     string `json:"message"`
	Resource    string `json:"resource"`
	Description string `json:"description"`
}

func (e Error) Error() string {
	return e.Message
}

func write(w http.ResponseWriter, code int, data interface{}) {
	if code < 200 {
		panic(fmt.Sprintf("status code %d must be >= 200", code))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if data != nil {
		enc := json.NewEncoder(w)
		enc.SetIndent("", "")
		enc.SetEscapeHTML(false)
		if err := enc.Encode(data); err != nil {
			fields := log.Fields{"data": data, "code": code}
			log.WithFields(fields).Errorf("%+v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func Err(w http.ResponseWriter, e *Error, code int) {
	if e == nil {
		panic("error must not be empty")
	}
	write(w, code, e)
}

func New(message, resource string, output *Output) *Error {
	err := &Error{
		Message:     message,
		Resource:    resource,
		Description: output.Description(),
	}
	return err
}

func NewInternalError(resource string) *Error {
	return New(Internal.ToUpper(), resource, Internal)
}

func NewNotFoundError(resource string) *Error {
	return New(NotFound.ToUpper(), resource, NotFound)
}

func NewForbiddenError(resource string) *Error {
	return New(Forbidden.ToUpper(), resource, Forbidden)
}

func OK200(w http.ResponseWriter, data interface{}) {
	write(w, 200, data)
}

func NotFound404(w http.ResponseWriter, resource string) {
	Err(w, NewNotFoundError(resource), 404)
}

func InternalError500(w http.ResponseWriter, resource string, err error) {
	Err(w, NewInternalError(resource), 500)
}

func Forbidden403(w http.ResponseWriter, resource string) {
	Err(w, NewForbiddenError(resource), 500)
}
