package AvlTree

func NewAvlEntry(k int, value string) *AvlEntry {

	return &AvlEntry{
		Key:   k,
		Value: value,
	}
}
