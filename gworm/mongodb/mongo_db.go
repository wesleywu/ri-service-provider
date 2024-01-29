package mongodb

import (
	"github.com/gogf/gf/v2/container/gvar"
	"go.mongodb.org/mongo-driver/bson"
)

type (
	Raw    bson.Raw                 // Raw is a raw sql that will not be treated as argument but as a direct sql part.
	Value  = *gvar.Var              // Value is the field value type.
	Record map[string]Value         // Record is the row record of the table.
	Result []Record                 // Result is the row record array.
	Map    = map[string]interface{} // Map is alias of map[string]interface{}, which is the most common usage map type.
	List   = []Map                  // List is type of map array.
)

const (
	modelForDaoSuffix = `ForDao`
)
