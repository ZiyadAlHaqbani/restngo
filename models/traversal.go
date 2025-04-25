package models

type Traversal struct {
	By_index bool
	By_field bool
	Index    int
	Field    string
}

func IndexTraversal(index int) Traversal {
	return Traversal{
		By_index: true,
		Index:    index,
	}
}

func FieldTraversal(field string) Traversal {
	return Traversal{
		By_field: true,
		Field:    field,
	}
}
