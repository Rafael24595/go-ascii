package commons

type Dependency interface {
	OnLoad() bool
	OnExit() bool
}