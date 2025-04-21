package probe

type Status = bool

const (
	StatusSucceed = true
	StatusFailed  = false
)

type Result struct {
	Status Status
}

type Probe interface {
	Execute() (*Result, error)
}
