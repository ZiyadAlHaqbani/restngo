package nodes

import (
	httphandler "htestp/http_handler"
	"htestp/models"
)

type Get_Node struct {
	next        models.Node
	request     httphandler.Request
	response    httphandler.HTTPResponse
	constraints models.Constraint
}

type Post_Node struct {
	next        models.Node
	request     httphandler.Request
	response    httphandler.HTTPResponse
	constraints models.Constraint
}
