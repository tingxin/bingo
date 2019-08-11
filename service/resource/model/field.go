package model

// FieldM used to
type FieldM struct {
	ID            int64  `json:"id"`
	IndicatorType int    `json:"indicator_type"`
	Order         int    `json:"order"`
	ResourceID    int64  `json:"resource_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Name          string `json:"name"`
	Table         string `json:"table"`
	Group         string `json:"group"`
	Selected      bool   `json:"selected"`
	CreateTime    string `json:"create_time"`
	UpdateTime    string `json:"update_time"`
}
