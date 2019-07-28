package model

// Condition struct
type Condition struct {
	Table   string         `json:"table"`
	DB      string         `json:"db"`
	Fields  []*FieldDetail `json:"fields"`
	Filters []*Filter      `json:"filters"`
}

// FieldDetail struct
type FieldDetail struct {
	Field
	IndicatorType int8 `json:"indicator_type"`
	ValueType     int8 `json:"value_type"`
}

// Field struct
type Field struct {
	Key   string `json:"key"`
	Table string `json:"table"`
}

// Filter struct
type Filter struct {
	Field
	Operator  int8        `json:"operator"`
	ValueType int8        `json:"value_type"`
	Group     string      `json:"group"`
	Value     interface{} `json:"value"`
}
