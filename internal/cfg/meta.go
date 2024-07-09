package cfg

type Meta map[Env]Config

type ConfigUpdater func(Config) Config

var metaConfig *Meta = nil

func makeMetaConfig() *Meta {
	cfg := make(Meta)
	for _, env := range []Env{Local, Testing, Production} {
		cfg[env] = Config{}
	}
	return &cfg
}

func newMetaConfig() *Meta {
	if metaConfig == nil {
		metaConfig = makeMetaConfig()
	}
	return metaConfig
}

func (mcfg *Meta) SetSlice(env Env, cfg Config) *Meta {
	(*mcfg)[env] = cfg
	return mcfg
}

func (mcfg *Meta) InheritSlice(to Env, from Env, updater ConfigUpdater) *Meta {
	(*mcfg)[to] = updater((*mcfg)[from])
	return mcfg
}

func (mcfg *Meta) GetConfig(env Env) Config {
	return (*mcfg)[env]
}
