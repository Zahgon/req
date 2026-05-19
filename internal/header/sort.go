package header

type KeyValues struct {
	Key    string
	Values []string
}

type sorter struct {
	order map[string]int
	kvs   []KeyValues
}

func (s *sorter) Len() int           { _ = "STUB: not implemented"; return 0 }
func (s *sorter) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (s *sorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func SortKeyValues(kvs []KeyValues, orderedKeys []string) { _ = "STUB: not implemented"; return }
