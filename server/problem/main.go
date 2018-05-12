package problem

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-errors/errors"
	//"gitlab.com/distributed_lab/logan"
)

var (
	errToProblemMap = map[error]P{}
)

// RegisterError records an error -> P mapping, allowing the app to register
// specific errors that may occur in other packages to be rendered as a specific
// P instance.
//
// For example, you might want to render any sql.ErrNoRows errors as a
// problem.NotFound, and you would do so by calling:
//
// problem.RegisterError(sql.ErrNoRows, problem.NotFound) in you application
// initialization sequence
func RegisterError(err error, p P) {
	errToProblemMap[err] = p
}

// HasProblem types can be transformed into a problem.
// Implement it for custom errors.
type HasProblem interface {
	Problem() P
}

// P is a struct that represents an error response to be rendered to a connected
// client.
type P struct {
	Type     string                 `json:"type"`
	Title    string                 `json:"title"`
	Status   int                    `json:"status"`
	Detail   string                 `json:"detail,omitempty"`
	Instance string                 `json:"instance,omitempty"`
	Extras   map[string]interface{} `json:"extras,omitempty"`
}

func (p *P) Error() string {
	return fmt.Sprintf("problem: %s", p.Type)
}

// Render writes a http response to `w`, compliant with the "Problem
// Details for HTTP APIs" RFC:
//   https://tools.ietf.org/html/draft-ietf-appsawg-http-problem-00
//
// `p` is the problem, which may be either a concrete P struct, an implementor
// of the `HasProblem` interface, or an error.  Any other value for `p` will
// panic.
func Render(w http.ResponseWriter, p interface{}) {
	switch p := p.(type) {
	case P:
		render(w, p)
	case *P:
		render(w, *p)
	case HasProblem:
		render(w, p.Problem())
	case error:
		renderErr(w, p)
	default:
		panic(fmt.Sprintf("Invalid problem: %v+", p))
	}
}

func render(w http.ResponseWriter, p P) {

	w.Header().Set("Content-Type", "application/problem+json; charset=utf-8")
	js, err := json.MarshalIndent(p, "", "  ")

	if err != nil {
		//err := errors.Wrap(err, 1)
		//logan.WithStack(err).WithError(err)
		http.Error(w, "error rendering problem", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(p.Status)
	w.Write(js)
}

func renderErr(w http.ResponseWriter, err error) {
	origErr := err

	if err, ok := err.(*errors.Error); ok {
		origErr = err.Err
	}

	p, ok := errToProblemMap[origErr]

	// If this error is not a registered error
	// log it and replace it with a 500 error
	if !ok {
		//logan.WithStack(err).Error(err)
		p = ServerError
	}

	render(w, p)
}

// Well-known and reused problems below:
var (
	// NotFound is a well-known problem type.  Use it as a shortcut
	// in your actions.
	NotFound = P{
		Type:   "not_found",
		Title:  "Resource Missing",
		Status: http.StatusNotFound,
		Detail: "The resource at the url requested was not found.  This is usually " +
			"occurs for one of two reasons:  The url requested is not valid, or no " +
			"data in our database could be found with the parameters provided.",
	}

	Success = P{
		Type:   "success",
		Title:  "Success",
		Status: http.StatusOK,
	}

	// ServerError is a well-known problem type.  Use it as a shortcut
	// in your actions.
	ServerError = P{
		Type:   "server_error",
		Title:  "Internal Server Error",
		Status: http.StatusInternalServerError,
		Detail: "An error occurred while processing this request.  This is usually due " +
			"to a bug within the server software.  Trying this request again may " +
			"succeed if the bug is transient, otherwise please report this issue " +
			"to the issue tracker at: https://gitlab.com/tokend/hgate/issues." +
			" Please include this response in your issue.",
	}
	// NotImplemented is a well-known problem type.  Use it as a shortcut
	// in your actions.
	NotImplemented = P{
		Type:   "not_implemented",
		Title:  "Resource Not Yet Implemented",
		Status: http.StatusNotFound,
		Detail: "While the requested URL is expected to eventually point to a " +
			"valid resource, the work to implement the resource has not yet " +
			"been completed.",
	}

	MethodNotAllowed = P{
		Type:   "method_not_allowed",
		Title:  "Method Not Allowed",
		Status: http.StatusMethodNotAllowed,
		Detail: "The method specified in the Request-Line is not allowed " +
			"for the resource identified by the Request-URI",
	}

	// NotAcceptable is a well-known problem type.  Use it as a shortcut
	// in your actions.
	NotAcceptable = P{
		Type: "not_acceptable",
		Title: "An acceptable response content-type could not be provided for " +
			"this request",
		Status: http.StatusNotAcceptable,
	}

	// BadRequest is a well-known problem type.  Use it as a shortcut
	// in your actions.
	BadRequest = P{
		Type:   "bad_request",
		Title:  "Bad Request",
		Status: http.StatusBadRequest,
		Detail: "The request you sent was invalid in some way",
	}

	// Timeout is a well-known problem type.  Use it as a shortcut
	// in your actions.
	Timeout = P{
		Type:   "timeout",
		Title:  "Timeout",
		Status: http.StatusGatewayTimeout,
		Detail: "Your request timed out before completing.  Please try your " +
			"request again.",
	}

	// UnsupportedMediaType is a well-known problem type.  Use it as a shortcut
	// in your actions.
	UnsupportedMediaType = P{
		Type:   "unsupported_media_type",
		Title:  "Unsupported Media Type",
		Status: http.StatusUnsupportedMediaType,
		Detail: "The request has an unsupported content type. Presently, the " +
			"only supported content type is application/x-www-form-urlencoded.",
	}

	// Forbidden is a well-known problem type.  Use it as a shortcut
	// in your actions.
	Forbidden = P{
		Type:   "forbidden",
		Title:  "Forbidden",
		Status: http.StatusForbidden,
	}

	Conflict = P{
		Type:   "conflict",
		Title:  "Conflict",
		Status: http.StatusConflict,
		Detail: "Already exists",
	}
)
