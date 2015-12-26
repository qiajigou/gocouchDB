package gocouchDB

import (
    "io"
)

type Document struct {
    Name string
    DatabaseName string
    ClientBase
}

func NewDocument(dbName, docName string, transport ITransport) *Document {
    d := new(Document)
    d.Name = docName
    d.DatabaseName = dbName
    d.transport = transport
    d.Headers = map[string]string {}
    return d
}

// get path of the document
func (cl *Document)path() (url string) {
    return cl.DatabaseName + "/" + cl.Name
}

// get info of the document
func (cl *Document)GetInfo() (j map[string]interface{}, err error) {
    return cl.do(GET, cl.path(), nil, nil)
}

// get leaf node of the couchdb B-tree
func (cl *Document)Leaf() (j map[string]interface{}, err error) {
    params := map[string]string {
        "leaf": "true",
    }

    return cl.GetInfoByParams(params)
}

// get info by reversion
func (cl *Document)GetInfoByReversion(rev string) (j map[string]interface{}, err error) {

    params := map[string]string{
        "rev": rev,
    }

    return cl.GetInfoByParams(params)
}

// get info by params
// you can do anything you want with this function
// like GetInfoByReversion
func (cl *Document)GetInfoByParams(params map[string]string) (j map[string]interface{}, err error) {
    path := cl.joinParams(cl.path(), params)
    return cl.do(GET, path, nil, nil)
}

// get reversion of the document
func (cl *Document)GetReversion()(r string, err error) {
    body, err := cl.Leaf()

    if err != nil {
        return "", err
    }

    return body["_rev"].(string), err
}

// update document
func (cl *Document)Update(body map[string]interface{}) (j map[string]interface{}, err error) {

    rev, err := cl.GetReversion()

    if err != nil {
        return j, err
    }

    body["_rev"] = rev

    return cl.do(PUT, cl.path(), body, nil)
}

// delete document by update _deleted
func (cl *Document)Delete() (j map[string]interface{}, err error) {

    body := map[string]interface{}{
        "_deleted": true,
    }

    return cl.Update(body)
}

// get attachment path
func (cl *Document)attachmentPath(attname string) (url string) {
    return cl.DatabaseName + "/" + cl.Name + "/" + attname
}

// create attachement
func (cl *Document)CreateAttachment(attname string, data io.Reader, headers map[string]string) (j map[string]interface{}, err error){
    rev, err := cl.GetReversion()

    if err != nil {
        return j, err
    }

    return cl.CreateAttachmentByReversion(attname, rev, data, headers)
}

// get attachment
func (cl *Document)GetAttachment(attname string) (j map[string]interface{}, err error) {
    return cl.do(GET, cl.attachmentPath(attname), nil, nil)
}

// delete attachment
func (cl *Document)DeleteAttachment(attname string) (j map[string]interface{}, err error) {
    rev, err := cl.GetReversion()

    if err != nil {
        return j, err
    }

    return cl.DeleteAttachmentByReversion(attname, rev)
}

// create attachment
func (cl *Document)CreateAttachmentByReversion(attname, rev string, data io.Reader, headers map[string]string) (j map[string]interface{}, err error) {
    path := cl.attachmentPath(attname) + "?rev=" + rev
    return cl.request(PUT, path, data)
}

// get attachment
func (cl *Document)GetAttachmentByReversion(attname, rev string) (j map[string]interface{}, err error) {
    path := cl.attachmentPath(attname)
    rev, err = cl.GetReversion()

    if err != nil {
        return j, err
    }

    path = path + "?rev=" + rev
    return cl.do(GET, cl.attachmentPath(attname), nil, nil)
}

// delete attachment
func (cl *Document)DeleteAttachmentByReversion(attname, rev string) (j map[string]interface{}, err error) {
    path := cl.attachmentPath(attname) + "?rev=" + rev
    return cl.do(DELETE, path, nil, nil)
}

