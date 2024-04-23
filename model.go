package datatable

type TableData struct {
	Headings []string
	ColTypes []interface{}
	Cells    []Cell
}

type Cell struct {
	Row   int
	Col   int
	Value Value
}

type Value struct {
	Data interface{}
	// Style
	Formula Formula
}

type Formula struct {
	Text string
}
