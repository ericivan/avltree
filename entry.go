package AvlTree

type AvlEntry struct {
	Key    int
	Value  string
	Left   *AvlEntry
	Right  *AvlEntry
	Parent *AvlEntry
	Height int
}

func (avl *AvlEntry) GetValue() {

}

func (avl *AvlEntry) SetValue(value string) string {

	avl.Value = value

	return value
}



