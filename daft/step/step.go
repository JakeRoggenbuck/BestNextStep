package step

type Step struct {
	ID   int64  `json:"_id"`
	Name string `json:"name"`
	Desc string `json:"desc"`

	Collection int64 `json:"collection"`
	Owner      int64 `json:"-"`
}
