package cfg

var config *Config = nil

func NewConfig(env Env) *Config {
	if (config == nil) {
		setup()
		tmp := newMetaConfig().GetConfig(env)
		config = &tmp
	}

	return config
}

func setup() {
	newMetaConfig().SetSlice(Production, Config{
		Env: Production,
	}).InheritSlice(
		Testing, 
		Production, 
		func(cfg Config) Config {
			cfg.Env = Testing
			return cfg
		},
	).InheritSlice(
		Local, 
		Testing,
		func(cfg Config) Config {
			cfg.Env = Local
			return cfg
		},
	)
}
