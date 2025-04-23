package models

type MatchType string

const (
	TypeString MatchType = "string"
	TypeFloat  MatchType = "float64"
	TypeBool   MatchType = "bool"
	TypeObject MatchType = "map[string]interface{}"
	TypeArray  MatchType = "[]interface{}"
)
