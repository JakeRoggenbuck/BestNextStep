package step

type Step struct {
	ID   int64 `json:"_id"`
	Name string `json:"name"`
	Desc string `json:"desc"`

	// Item IDs to the left and right
	// -1 if nil
	Left  int64 `json:"left"`
	Right int64 `json:"right"`
	Owner int64 `json:"-"`
}
