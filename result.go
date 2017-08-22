package ago

import (
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	Result200 = Result{Code: 200, Msg: "Success."}
	Result201 = Result{Code: 201, Msg: "Created."}
	Result202 = Result{Code: 202, Msg: "Accepted."}
	Result204 = Result{Code: 204, Msg: "No content."}

	Result400 = Result{Code: 400, Msg: "Invalid request."}
	Result401 = Result{Code: 401, Msg: "Unauthorized."}
	Result403 = Result{Code: 403, Msg: "Forbidden."}
	Result404 = Result{Code: 404, Msg: "Not found."}
	Result406 = Result{Code: 406, Msg: "Not acceptable."}
	Result410 = Result{Code: 410, Msg: "Gone."}
	Result422 = Result{Code: 422, Msg: "Unprocesable entity."}

	Result500 = Result{Code: 500, Msg: "Internal server error."}
)
