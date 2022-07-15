package myddlmaker

type indexes interface {
	Indexes() []*Index
}

type uniqueIndexes interface {
	UniqueIndexes() []*UniqueIndex
}

type foreignKeys interface {
	ForeignKeys() []*ForeignKey
}

type Index struct {
	Name    string
	Columns []string
}

func NewIndex(name string, col ...string) *Index {
	return &Index{
		Name:    name,
		Columns: col,
	}
}

type UniqueIndex struct {
	Name    string
	Columns []string
}

func NewUniqueIndex(name string, col ...string) *UniqueIndex {
	return &UniqueIndex{
		Name:    name,
		Columns: col,
	}
}

type ForeignKey struct {
	Name       string
	Columns    []string
	Table      string
	References []string
}

func NewForeignKey(name string, columns []string, table string, references []string) *ForeignKey {
	if name == "" {
		panic("name is missing")
	}
	if table == "" {
		panic("table is missing")
	}
	if len(columns) == 0 {
		panic("columns is missing")
	}
	if len(references) == 0 {
		panic("references is missing")
	}
	if len(columns) != len(references) {
		panic("columns and references must have same length")
	}
	return &ForeignKey{
		Name:       name,
		Columns:    columns,
		Table:      table,
		References: references,
	}
}