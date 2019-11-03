package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"net/http"
	"strconv"

	"github.com/google/logger"
	"github.com/pkg/errors"
)

// responseDecodeJSON ...
func responseDecodeJSON(bodyResponse io.Reader, response interface{}) error {
	var body, errBody = ioutil.ReadAll(bodyResponse)
	if errBody != nil {
		return errBody
	}

	errJSON := json.Unmarshal(body, response)
	if errJSON != nil {
		return errJSON
	}

	return nil
}

// requestEncodeJSON ...
func requestEncodeJSON(objRequest interface{}) (*bytes.Buffer, error) {
	bodyRequestJSON := new(bytes.Buffer)
	encodeJSON, erro := json.Marshal(objRequest)
	if erro != nil {
		return nil, erro
	}
	bodyRequestJSON.Write(encodeJSON)

	return bodyRequestJSON, nil
}

//DecodeBodyJSON ...
func DecodeBodyJSON(r *http.Request, v interface{}, logger *logger.Logger) error {
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		return errors.Wrap(erro, ErrorReadAllBuffer)
	}
	logger.Info(fmt.Sprintf("Request: %s", string(body)))

	if erro = json.Unmarshal(body, v); erro != nil {
		return errors.Wrap(erro, ErrorJSONUnmarshal)
	}

	return nil
}

/*
Respond: This method return one message in JSON format
*/
func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	if _, err := io.Copy(w, &buf); err != nil {
		logger.Error("Error coping buffer")
	}

}

/*
RespondErro: This method return one error message in JSON format
*/
func RespondErro(w http.ResponseWriter, r *http.Request, status int, errMsg *ErrMessage) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(errMsg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Write(buf.Bytes())
}

/*
JSONResponse: This method return one message in TEXT PLAIN format
Ex: JSONResponse(w, http.StatusBadRequest, "One ou more fields has not been informed...")
*/
func JSONResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, message)
}

//handleNotFound ...
func handleNotFound(w http.ResponseWriter, r *http.Request) {
	body := ErrMessage{Message: "URL not found",
		Code: strconv.Itoa(http.StatusNotFound)}
	Respond(w, r, http.StatusNotFound, body)
	return
}
