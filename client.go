package gocouchDB

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

// replicate database
func (cl *Client)Replicate(task *ReplicateTask) (j map[string]interface{}, err error) {

    body := map[string]interface{}{
        "continuous": task.Continuous,
        "create_target": task.CreateTarget,
        "source": task.Source,
        "target": task.Target,
    }

    if task.DocumentIDs != nil {
        body["doc_ids"] = task.DocumentIDs
    }

    if task.Proxy != "" {
        body["proxy"] = task.Proxy
    }

    if task.Cancel {
        body["cancel"] = task.Cancel
    }

    str, err := cl.handParams(body)

    if err != nil {
        return j, err
    }

    return cl.request(POST, "_replicate", str)
}

