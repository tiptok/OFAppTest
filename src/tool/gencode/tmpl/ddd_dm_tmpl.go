package tmpl

var ProtocolDomainModel = `
package domain

type {{.Model}} struct {
{{.Items}}
}

type {{.Model}}Repository interface {
	Save(dm *{{.Model}}) (*{{.Model}}, error)
	Remove(dm *{{.Model}}) (*{{.Model}}, error)
	FindOne(queryOptions map[string]interface{}) (*{{.Model}}, error)
	Find(queryOptions map[string]interface{}) (int64, []*{{.Model}}, error)
}

func (m *{{.Model}}) Identify() interface{} {
	if m.Id == 0 {
		return nil
	}
	return m.Id
}
`
