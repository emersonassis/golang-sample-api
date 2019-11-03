package web

import (
	"net/http"
	"strconv"

	"github.com/google/logger"
	"projects.org/sample/sample-api/core"
)

//Handler ...
type Handler struct {
	Logger *logger.Logger
}

/*
HandlerVersion: Return the version`s API
*/
func (h *Handler) HandlerVersion(w http.ResponseWriter, r *http.Request) {

	var (
		response = &VersionResponse{}
		request  = &VersionRequest{}
	)

	version := request.Version
	version.Major = "1"
	version.Minor = "0"
	version.Patch = "0"

	response.CodResponse = "201"
	response.Message = "I'm Ok"
	response.Version = &version

	core.Respond(w, r, http.StatusOK, response)
	return

}

/*
	HandleProcessObject: Example of POST...
*/
func (h *Handler) HandleProcessObject(w http.ResponseWriter, r *http.Request) {

	var (
		request  = &ObjectRequest{}
		response = &ObjectResponse{}
	)

	err := core.DecodeBodyJSON(r, request, h.Logger)
	if err != nil {
		core.RespondErro(w, r, http.StatusInternalServerError, &core.ErrMessage{Message: "Error in JSON format", Code: strconv.Itoa(http.StatusInternalServerError), Erro: err.Error()})
		return
	}

	object := request.Object

	response.CodResponse = "201"
	response.Message = "Process done..."
	response.Object = &object

	core.Respond(w, r, http.StatusOK, response)
	return
}
