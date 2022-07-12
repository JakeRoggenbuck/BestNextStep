package col

type Col struct {
	ID    int64  `json:"_id"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Owner int64  `json:"-"`
}
