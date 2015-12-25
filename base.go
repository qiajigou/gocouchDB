package gocouch

import (
    "io"
    "strings"
    "encoding/json"
    "encoding/base64"
)

type ClientBase struct {
    Headers map[string]string
    Username string
    Password string
    transport ITransport
}

func (cl *ClientBase)SetAuth(username, password string) {
    cl.Username = username
    cl.Password = password
}

func (cl *ClientBase)beforeRequest() {
    if cl.Username != "" && cl.Password != "" {
        auth := []byte(cl.Username + ":" + cl.Password)
        hashed := "Basic " + base64.StdEncoding.EncodeToString(auth)
        cl.Headers["Authorization"] = string(hashed)
    }
}

func (cl *ClientBase)request(method, path string, body io.Reader)(j map[string]interface{}, err error) {
    cl.beforeRequest()

    j, err = cl.transport.request(method, path, body, cl.Headers)

    defer cl.ClearHeaders()

    if err != nil {
        return j, err
    }

    return j, err
}

func (cl *ClientBase)requestList(method, path string, body io.Reader)(j []string, err error) {
    cl.beforeRequest()

    j, err = cl.transport.requestList(method, path, body, cl.Headers)

    defer cl.ClearHeaders()

    if err != nil {
        return j, err
    }

    return j, err
}

func (cl *ClientBase)handParams(body map[string]interface{}) (ret *strings.Reader, err error) {
    json, err := json.Marshal(body)

    defer cl.ClearHeaders()

    if err != nil {
        return ret, err
    }

    str := strings.NewReader(string(json))

    return str, err

}

func (cl *ClientBase)SetHeaders(headers map[string]string) {
    cl.Headers = headers
}

func (cl *ClientBase)ClearHeaders(){
    cl.Headers = map[string]string {}
}