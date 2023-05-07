package commons

type Dependency interface {
	DependencyName() string
	OnLoad() bool
	OnExit() bool
}