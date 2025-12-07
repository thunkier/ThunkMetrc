package bruno

// Request represents a parsed Bruno request.
type Request struct {
	Name       string
	Type       string
	Seq        int
	Method     string
	URL        string
	BodyType   string
	Body       string
	AuthType   string
	Docs       string
	Group      string
	FilePath   string
	Params     map[string]string
	BodySchema *Schema
}

// SchemaType represents a JSON type.
type SchemaType string

const (
	TypeString  SchemaType = "string"
	TypeInt     SchemaType = "int"
	TypeFloat   SchemaType = "float"
	TypeBool    SchemaType = "bool"
	TypeObject  SchemaType = "object"
	TypeArray   SchemaType = "array"
	TypeUnknown SchemaType = "unknown"
)

// Schema represents the structure of a JSON body.
type Schema struct {
	Type       SchemaType
	Properties map[string]*Schema
	ItemType   *Schema
	IsNullable bool
}

// Collection represents a set of parsed Bruno requests.
type Collection struct {
	Requests []Request
}

// Clone creates a deep copy of the Request.
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
	return newReq
}

// Clone creates a deep copy of the Schema.
func (s *Schema) Clone() *Schema {
	if s == nil {
		return nil
	}
	newSchema := &Schema{
		Type:       s.Type,
		IsNullable: s.IsNullable,
	}
	if s.Properties != nil {
		newSchema.Properties = make(map[string]*Schema, len(s.Properties))
		for k, v := range s.Properties {
			newSchema.Properties[k] = v.Clone()
		}
	}
	if s.ItemType != nil {
		newSchema.ItemType = s.ItemType.Clone()
	}
	return newSchema
}
