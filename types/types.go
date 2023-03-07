package types

type IJson interface {
	JsonEncode() string
	JsonDecode(v any) error
}

type ITypeStringAny interface {
	Set(k string, v any)
	SetForce(k string, v any, force bool)
	Exist(k string) bool
	Delete(k string)
	GetDefault(k string, d any) any
	Keys() []string
	IJson
}
type ITypeStrings interface {
	Set(k, v string)
	SetForce(k, v string, force bool)
	Exist(k string) bool
	Delete(k string)
	GetDefault(k, d string) string
	Keys() []string
	Values() []string
	Implode(sep string) string
	IJson
}
