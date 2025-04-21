package monitor

type Status int

const (
	StatusDown = iota
	StatusUp
	StatusPending
)

type Monitor struct {
	Name string
}
