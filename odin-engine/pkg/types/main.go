package types

// EngineConfig is a type to be used for accessing the engine configuration file information
type EngineConfig struct {
	OdinVars struct {
		Master string `yaml:"master"`
		Port   string `yaml:"port"`
	} `yaml:"odin"`
	Mongo struct {
		Address string `yaml:"address"`
	} `yaml:"mongo"`
	Storage struct {
		Name    string `yaml:"name"`
		Address string `yaml:"address"`
	} `yaml:"storage"`
}

// JobConfig is a type to be used for accessing job configuration file information
type JobConfig struct {
	Provider struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"provider"`
	Job struct {
		Name        string `yaml:"name"`
		ID          string `yaml:"id"`
		Description string `yaml:"description"`
		Language    string `yaml:"language"`
		File        string `yaml:"file"`
		Schedule    string `yaml:"schedule"`
	} `yaml:"job"`
}

// StringFormat is a type to be used for storing schedule information in the cron format
type StringFormat struct {
	Minute string
	Hour   string
	Dom    string
	Mon    string
	Dow    string
}
