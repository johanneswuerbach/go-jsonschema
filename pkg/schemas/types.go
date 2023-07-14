package schemas

const (
	TypeNameString  = "string"
	TypeNameArray   = "array"
	TypeNameNumber  = "number"
	TypeNameInteger = "integer"
	TypeNameObject  = "object"
	TypeNameBoolean = "boolean"
	TypeNameNull    = "null"
)

func IsPrimitiveType(t string) bool {
	switch t {
	case TypeNameString, TypeNameNumber, TypeNameInteger, TypeNameBoolean, TypeNameNull:
		return true

	default:
		return false
	}
}

func isPrimitiveTypeList(types []*Type) bool {
	for _, typ := range types {
		if len(typ.Type) == 0 {
			continue
		}

		if !IsPrimitiveType(typ.Type[0]) {
			return false
		}
	}

	return true
}
