package gocouch

type Client struct {
    ClientBase
}

func NewClientByDSN(dsn string) *Client{
    transport := NewTransport(dsn)
    c := new(Client)
    c.transport = transport
    return c
}

func NewClientByTransport(transport ITransport) *Client {
    c := new(Client)
    c.transport = transport
    return c
}

func (cl *Client)ServerInfo() (j map[string]interface{}, e error) {
    return cl.request(GET, "")
}

func (cl *Client)CreateDatabase(dbName string) (j map[string]interface{}, e error) {

    return cl.request(PUT, dbName)
}

func (cl *Client)GetDatabase(dbName string) (db *Database, e error) {
    d := new(Database)
    d.Name = dbName
    d.transport = cl.transport
    return d, e
}
