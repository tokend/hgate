package server

import (
	"gitlab.com/tokend/hgate/server/problem"
	"encoding/json"
	"mime"
	"strings"
)

func (action *Handler) Do(fns ...func()) {
	for _, fn := range fns {
		if action.Err != nil {
			return
		}
		fn()
	}

}

func (action *Handler) ReadBody() {
	action.rawBody = make([]byte, action.R.ContentLength)

	_, action.Err = action.R.Body.Read(action.rawBody)
	if action.Err != nil {
		action.rawBody = []byte{}
	}
}

func (action *Handler) GetJSON(dest interface{}) error {
	decoder := json.NewDecoder(action.R.Body)
	err := decoder.Decode(&dest)
	if err != nil {
		return err
	}
	defer action.R.Body.Close()
	return nil
}

func (action *Handler) ValidateRequest() {
	if !IsValidContentType(action.R.Header.Get("Content-type"), action.ContentType) {
		action.Err = &problem.UnsupportedMediaType
	}
}

func (action *Handler) SetInvalidField(name string, reason error) {
	br := problem.BadRequest

	br.Extras = map[string]interface{}{}
	br.Extras["invalid_field"] = name
	br.Extras["reason"] = reason.Error()

	action.Err = &br
}
func IsValidContentType(contentType, mimetype string) bool {
	if contentType == "" {
		return mimetype == "application/octet-stream"
	}

	for _, v := range strings.Split(contentType, ",") {
		t, _, err := mime.ParseMediaType(v)
		if err != nil {
			break
		}
		if t == mimetype {
			return true
		}
	}
	return false
}
