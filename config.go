package main

// Config for split processor.
type Config struct {
	Index       int    `config:"index"`
	FieldPath   string `config:"field_path"`
	Separator   string `config:"separator"`
	IgnoreError bool   `config:"ignore_error"`
	KeyName     string `config:"key_name"`
	EnvType     bool   `config:"env_type"`
}

func defaultConfig() Config {
	return Config{
		Index:       -1,
		FieldPath:   "",
		Separator:   "/",
		IgnoreError: true,
		KeyName:     "",
		EnvType:     false,
	}
}
