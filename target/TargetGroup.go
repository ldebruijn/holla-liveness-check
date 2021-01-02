package target

type Group struct {
	Name     string   `yaml:"name"`
	Schedule string   `yaml:"schedule"`
	Targets  []Target `yaml:"targets"`
}
