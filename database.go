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

// get all docs
// http://docs.couchdb.org/en/1.6.1/api/database/bulk-api.html#db-bulk-docs
func (cl *Database)GetDocuments(params map[string]string) (j map[string]interface{}, err error){
    path := cl.Name + "/_all_docs"
    path = cl.joinParams(path, params)
    return cl.request(GET, path, nil)
}


// get docs by keys
func (cl *Database)GetDocumentsByKeys(body map[string]interface{}) (j map[string]interface{}, err error){
    path := cl.Name + "/_all_docs"
    str, err := cl.handParams(body)
    return cl.request(POST, path, str)
}

// update/insert/delete bulk documents
func (cl *Database)UpdateBulkDocuments(body map[string]interface{}) (j map[string]interface{}, err error){
    return cl.bulkDocuments(body)
}

func (cl *Database)InsertBulkDocuments(body map[string]interface{}) (j map[string]interface{}, err error){
    return cl.bulkDocuments(body)
}

func (cl *Database)DeleteBulkDocuments(body map[string]interface{}) (j map[string]interface{}, err error){
    return cl.bulkDocuments(body)
}

func (cl *Database)bulkDocuments(body map[string]interface{}) (j map[string]interface{}, err error){
    path := cl.Name + "/_bulk_docs"
    str, err := cl.handParams(body)

    if err != nil {
        return j, err
    }

    return cl.request(POST, path, str)
}
