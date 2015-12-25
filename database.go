package gocouch

type Database struct {
    ClientBase
    Name string
}

func (cl *Database)Info() (j map[string]interface{}, e error) {
    return cl.request(GET, cl.Name)
}

func (cl *Database)Delete() (j map[string]interface{}, e error) {
    return cl.request(DELETE, cl.Name)
}
