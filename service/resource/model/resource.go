package model

// ResourceM used to
type ResourceM struct {
	ID          int64     `json:"id,omitempty"`
	Name        string    `json:"name"`
	Desc        string    `json:"desc"`
	Creator     string    `json:"creator"`
	Editor      string    `json:"editor"`
	CreateTime  string    `json:"create_time"`
	UpdateTime  string    `json:"update_time"`
	VisibleTime string    `json:"visible_time"`
	Visible     string    `json:"visible"`
	Order       int       `json:"order"`
	Kind        string    `json:"kind"`
	Fields      []*FieldM `json:"fields"`
}
