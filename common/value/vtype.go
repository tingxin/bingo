package value

type ValueType int8

const (
	Int ValueType = iota
	Int8
	Int16
	Int32
	Int64
	Float32
	Float64
	String
	Boolen
	DateTime
	Date
	TimeStamp
	Enum
)
