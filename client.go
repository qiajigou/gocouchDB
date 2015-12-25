package gocouch

type Client struct {
    ClientBase
}

func NewClientByDSN(dsn string) *Client{
    transport := NewTransport(dsn)
    c := new(Client)
    c.transport = transport
    c.Headers = map[string]string {}
    return c
}

func NewClientByTransport(transport ITransport) *Client {
    c := new(Client)
    c.transport = transport
    c.Headers = map[string]string {}
    return c
}

// get server info
func (cl *Client)ServerInfo() (j map[string]interface{}, err error) {
    return cl.request(GET, "", nil)
}

// create database with dbName
func (cl *Client)CreateDatabase(dbName string) (j map[string]interface{}, err error) {

    return cl.request(PUT, dbName, nil)
}

// this is a tricky function
// couchdb only this interface return a list
// not a key-value map json
func (cl *Client)ListAllDatabases() (j []string, err error) {
    return cl.requestList(GET, "_all_dbs", nil)
}

// get database
func (cl *Client)GetDatabase(dbName string) (d *Database, err error) {
    d = NewDatabase(dbName, cl.transport)
    d.SetAuth(cl.Username, cl.Password)
    return d, err
}

