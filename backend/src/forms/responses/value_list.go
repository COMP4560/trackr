package responses

type ValueList struct {
	Values      []Value `json:"values"`
	TotalValues int64   `json:"totalValues"`
}
