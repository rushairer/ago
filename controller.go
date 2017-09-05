package ago

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//ControllerInterface   Base controller interface
type ControllerInterface interface {
	Init(w http.ResponseWriter, r *http.Request, Params map[string]string)
}

//Controller Base controller
type Controller struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
	Params         map[string]string
}

//Init Init controller
func (c *Controller) Init(w http.ResponseWriter, r *http.Request, params map[string]string) {
	c.ResponseWriter = w
	c.Request = r
	c.Params = params
}

//Prepare Run this before all actions
func (c *Controller) Prepare() {

}

//Finish Run this after all actions
func (c *Controller) Finish() {

}

//Get Get method action
func (c *Controller) Get() {
	http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}

//Post Post method action
func (c *Controller) Post() {
	http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}

//Delete Delete method action
func (c *Controller) Delete() {
	http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}

//Put Put method action
func (c *Controller) Put() {
	http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}

//Head Head method action
func (c *Controller) Head() {
	http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}

//Patch Patch method action
func (c *Controller) Patch() {
	http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}

//Options Options method action
func (c *Controller) Options() {
	http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}

//JSON Output a json string
func (c *Controller) JSON(value interface{}) {
	c.ResponseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	c.ResponseWriter.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(c.ResponseWriter).Encode(value); err != nil {
		panic(err)
	}
}

//JSONNotFound Output a json string: not found
func (c *Controller) JSONNotFound() {
	c.ResponseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	c.ResponseWriter.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(c.ResponseWriter).Encode(Result404); err != nil {
		panic(err)
	}
}

//JSONForbiddon Output a json string: forbiddon
func (c *Controller) JSONForbiddon() {
	c.ResponseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	c.ResponseWriter.WriteHeader(http.StatusForbidden)
	if err := json.NewEncoder(c.ResponseWriter).Encode(Result403); err != nil {
		panic(err)
	}
}

//HTML Output an html string
func (c *Controller) HTML(html string) {
	c.ResponseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")
	c.ResponseWriter.WriteHeader(http.StatusOK)
	if _, err := io.WriteString(c.ResponseWriter, html); err != nil {
		panic(err)
	}
}

//HTMLNotFound Output an html string: not found
func (c *Controller) HTMLNotFound() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")
	c.ResponseWriter.WriteHeader(http.StatusNotFound)
	if _, err := io.WriteString(c.ResponseWriter, fmt.Sprintf("%d %s", Result404.Code, Result404.Msg)); err != nil {
		panic(err)
	}
}

//HTMLForbiddon Output an html string: forbiddon
func (c *Controller) HTMLForbiddon() {
	c.ResponseWriter.Header().Set("Content-Type", "text/html; charset=UTF-8")
	c.ResponseWriter.WriteHeader(http.StatusForbidden)
	if _, err := io.WriteString(c.ResponseWriter, fmt.Sprintf("%d %s", Result403.Code, Result403.Msg)); err != nil {
		panic(err)
	}
}
