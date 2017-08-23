package ago

import (
    "log"
    "testing"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "github.com/facebookgo/inject"
)

func (ago *Ago) RunTestGet() string{

    httpTestServer := httptest.NewServer(ago.Router.NewHandler(ago.Routes))

    defer httpTestServer.Close()

    res, err := http.Get(httpTestServer.URL)
    if err != nil {
        log.Fatal(err)
    }
    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("%s", body)

    return string(body)
}

func (ago *Ago) RunTestPost() string{

    httpTestServer := httptest.NewServer(ago.Router.NewHandler(ago.Routes))

    defer httpTestServer.Close()

    res, err := http.Post(httpTestServer.URL, "", nil)
    if err != nil {
        log.Fatal(err)
    }
    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("%s", body)

    return string(body)
}

type TestController struct {
    Controller
}

func (this *TestController) Get() {
    this.Json("OK")
}

func Test_AgoRunGet(t *testing.T) {
    var g inject.Graph
    var ago Ago

    err := g.Provide(
        &inject.Object{Value: &ago},
    )

    if (err != nil) {
        t.Error(err)
    }

    if err := g.Populate(); err != nil {
        t.Error(err)
    }

    ago.AddRoute(&TestController{}, "/")
    body := ago.RunTestGet()

    if (body == "\"OK\"\n") {
        t.Log("It's OK")
    } else {
        t.Error("Body is not OK")
    }
}

func Test_AgoRunPost(t *testing.T) {
    var g inject.Graph
    var ago Ago

    err := g.Provide(
        &inject.Object{Value: &ago},
    )

    if (err != nil) {
        t.Error(err)
    }

    if err := g.Populate(); err != nil {
        t.Error(err)
    }

    ago.AddRoute(&TestController{}, "/")
    body := ago.RunTestPost()

    if (body == "Method Not Allowed\n") {
        t.Log("It's OK")
    } else {
        t.Error("Body is not OK")
    }
}
