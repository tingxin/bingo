package command

// SignUpCmd used to
type SignUpCmd struct {
	SignCmd
	Name string `json:"name"`
	Role string `json:"role"`
}
