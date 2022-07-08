package step

type Step struct {
	ID int64
	Name string

	// Item IDs to the left and right
	// -1 if nil
	Left int64
	Right int64
}
