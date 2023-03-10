package schema

import (
	"github.com/tavroi/ssi-sdk/credential/schema"
)

// Resolution is an interface that defines a generic method of resolving a schema
type Resolution interface {
	Resolve(id string) (*schema.VCJSONSchema, error)
}
