package models

type Traversal struct {
	by_index bool
	by_field bool
	index    int
	field    string
}

func IndexTraversal(index int) Traversal {
	return Traversal{
		by_index: true,
		index:    index,
	}
}

func FieldTraversal(field string) Traversal {
	return Traversal{
		by_field: true,
		field:    field,
	}
}
