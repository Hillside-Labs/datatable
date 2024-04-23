package datatable

import "github.com/a-h/templ"

type TableModel struct {
	Config Config
	Rows   []Row
}

type Header struct {
	Title string
	Abbr  string
}

type Config struct {
	Headers []Header
	ID      string
}

type Value struct {
	Data interface{}
	// Style
	Formula Formula
}

type Formula struct {
	Text string
}

type Row struct {
	ID      string
	Values  map[string]string
	Hrefs   map[string]string
	Actions map[string][]RowAction
}

func (row Row) Value(key string) string {
	v, _ := row.Values[key]
	return v
}

func (row Row) Href(key string) templ.SafeURL {
	v, _ := row.Hrefs[key]
	return templ.URL(v)
}

func (row Row) Action(key string) []RowAction {
	v, _ := row.Actions[key]
	return v
}

type RowAction struct {
	Name       string
	Icon       templ.Component
	Attributes templ.Attributes
}
