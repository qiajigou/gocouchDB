package gocouchDB

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
    return cl.do(GET, cl.Name, nil, nil)
}

// delete a database
func (cl *Database)Delete() (j map[string]interface{}, err error) {
    return cl.do(DELETE, cl.Name, nil, nil)
}

// create document
func (cl *Database)CreateDocument(key string, body map[string]interface{}) (j map[string]interface{}, err error){
    path := cl.Name + "/" + key

    return cl.do(PUT, path, body, nil)
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
    return cl.do(GET, cl.Name + "/_all_docs", nil, params)
}


// get docs by keys
func (cl *Database)GetDocumentsByKeys(body map[string]interface{}) (j map[string]interface{}, err error){
    return cl.do(GET, cl.Name + "/_all_docs", body, nil)
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
    return cl.do(POST, cl.Name + "/_bulk_docs", body, nil)
}

// compact database
func (cl *Database)Compact() (j map[string]interface{}, err error) {
    return cl.do(POST, cl.Name + "/_compact", nil, nil)
}

// compact design doc
func (cl *Database)CompactDesignDoc(designDoc string) (j map[string]interface{}, err error) {
    return cl.do(POST, cl.Name + "/_compact/" + designDoc, nil, nil)
}

// ensure full commit
func (cl *Database)EnsureFullCommit()(j map[string]interface{}, err error) {
    return cl.do(POST, cl.Name + "/_ensure_full_commit", nil, nil)
}

// clean db view
func (cl *Database)ViewCleanUp() (j map[string]interface{}, err error) {
    return cl.do(POST, cl.Name + "/_view_cleanup", nil, nil)
}

// get db security
func (cl *Database)GetSecurity() (j map[string]interface{}, err error) {
    return cl.do(GET, cl.Name + "/_security", nil, nil)
}

// set db security
func (cl *Database)SetSecurity(body map[string]interface{}) (j map[string]interface{}, err error) {
    return cl.do(POST, cl.Name + "/_security", body, nil)
}
