package bruno

import (
	"github.com/thunkier/thunkmetrc/tools/pkg/schema"
)

type Request struct {
	Name         string
	Type         string
	Seq          int
	Method       string
	URL          string
	BodyType     string
	Body         string
	AuthType     string
	Docs         string
	Group        string
	FilePath     string
	Params       map[string]string
	BodySchema   *schema.Schema
	ReturnSchema *schema.Schema
	ResponseType string
	IsArray      bool
	IsPaginated  bool
	RequestType  string
	PathParams   []string
	Permissions  []string
	DocParams    []DocParam
}

type DocParam struct {
	Name        string
	Optional    bool
	Description string
}

type Collection struct {
	Requests []Request
}

func (r Request) Clone() Request {
	newReq := r
	if r.Params != nil {
		newReq.Params = make(map[string]string, len(r.Params))
		for k, v := range r.Params {
			newReq.Params[k] = v
		}
	}
	if r.BodySchema != nil {
		newReq.BodySchema = r.BodySchema.Clone()
	}
	if r.ReturnSchema != nil {
		newReq.ReturnSchema = r.ReturnSchema.Clone()
	}

	return newReq
}
