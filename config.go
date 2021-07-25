package main

// Config for split processor.
type Config struct {
	Index           int    `config:"index"`
	FieldPath       string `config:"field_path"`
	Separator       string `config:"separator"`
	IgnoreError     bool   `config:"ignore_error"`
	KeyName         string `config:"key_name"`
	EnvTypeEnable   bool   `config:"env_type_enable"`
	Mode            string `config:"mode"`
	EnableTimeStamp bool   `config:"enable_timestamp"`
	EnableSendTime  bool   `config:"enable_send_time"`
}

func defaultConfig() Config {
	return Config{
		Index:           -1,
		FieldPath:       "",
		Separator:       "/",
		IgnoreError:     true,
		KeyName:         "",
		EnvTypeEnable:   false,
		EnableTimeStamp: false,
		EnableSendTime:  false,
	}
}
