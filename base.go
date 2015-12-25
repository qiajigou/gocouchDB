package gocouch

type ClientBase struct {
    Headers map[string]string
    transport ITransport
}

func (cl *ClientBase)request(method, path string)(j map[string]interface{}, e error) {
    j, err := cl.transport.request(method, path, nil, cl.Headers)

    defer cl.ClearHeaders()

    if err != nil {
        return j, err
    }

    return j, err
}

func (cl *ClientBase)SetHeaders(headers map[string]string) {
    cl.Headers = headers
}

func (cl *ClientBase)ClearHeaders(){
    cl.Headers = map[string]string {}
}