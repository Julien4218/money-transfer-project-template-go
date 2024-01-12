package shared

type WorkflowDefinition struct {
	ID string `yaml:"id"`

	Inputs []struct {
		Name    string `yaml:"name"`
		Type    string `yaml:"type"`
		Default string `yaml:"default"`
	} `yaml:"inputs"`

	Steps []struct {
		ID        string `yaml:"id"`
		Type      string `yaml:"type"`
		Version   string `yaml:"version"`
		NextStep  string `yaml:"nextStep,omitempty"`
		ErrorStep string `yaml:"errorStep,omitempty"`
	} `yaml:"steps"`
}
