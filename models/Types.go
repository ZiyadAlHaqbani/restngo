package models

type MatchType string

const (
	NOTYPE     MatchType = "NOTYPE"
	TypeString MatchType = "string"
	TypeFloat  MatchType = "float64"
	TypeBool   MatchType = "bool"
	TypeObject MatchType = "map[string]interface{}"
	TypeArray  MatchType = "[]interface{}"
)

type HTTPMethod string

const (
	GET  HTTPMethod = "GET"
	POST HTTPMethod = "POST"
)
