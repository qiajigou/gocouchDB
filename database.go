package gocouch

type Database struct {
    ClientBase
    Name string
}

func NewDatabase(dbName string, transport ITransport) *Database {
    d := new(Database)
    d.Name = dbName
    d.transport = transport
    d.Headers = map[string]string {}
    return d
}

// show database info
func (cl *Database)GetInfo() (j map[string]interface{}, err error) {
    return cl.request(GET, cl.Name, nil)
}

// delete a database
func (cl *Database)Delete() (j map[string]interface{}, err error) {
    return cl.request(DELETE, cl.Name, nil)
}

// create document
func (cl *Database)CreateDocument(key string, body map[string]interface{}) (j map[string]interface{}, err error){
    str, err := cl.handParams(body)

    if err != nil {
        return j, err
    }

    path := cl.Name + "/" + key

    return cl.request(PUT, path, str)
}

// get document
func (cl *Database)GetDocument(key string) (d *Document, err error) {
    d = NewDocument(cl.Name, key, cl.transport)
    d.SetAuth(cl.Username, cl.Password)
    return d, err
}