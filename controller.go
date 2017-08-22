package ago

import (
    "net/http"
    "encoding/json"
)

type ControllerInterface interface {
    Init(w http.ResponseWriter, r *http.Request, Params map[string]string)
}

type Controller struct {
    ResponseWriter      http.ResponseWriter
    Request             *http.Request
    Params              map[string]string
}

func (c *Controller) Init(w http.ResponseWriter, r *http.Request, params map[string]string) {
    c.ResponseWriter = w
    c.Request = r
    c.Params = params
}

func (c *Controller) Prepare() {

}

func (c *Controller) Finish() {

}

func (c *Controller) Get() {
    http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Post() {
    http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Delete() {
    http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Put() {
    http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Head() {
    http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Patch() {
    http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Options() {
    http.Error(c.ResponseWriter, "Method Not Allowed", 405)
}

func (c *Controller) Json(value interface{}) {
    c.ResponseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
    c.ResponseWriter.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(c.ResponseWriter).Encode(value); err != nil {
        panic(err)
    }
}

func (c *Controller) JsonNotFound(value interface{}) {
    c.ResponseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
    c.ResponseWriter.WriteHeader(http.StatusNotFound)
    if err := json.NewEncoder(c.ResponseWriter).Encode(Result404); err != nil {
        panic(err)
    }
}

func (c *Controller) JsonForbiddon(value interface{}) {
    c.ResponseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")
    c.ResponseWriter.WriteHeader(http.StatusForbidden)
    if err := json.NewEncoder(c.ResponseWriter).Encode(Result403); err != nil {
        panic(err)
    }
}
