package config

type IConfig interface {

}

// verify Config implements IConfig
var _ IConfig = (*Config)(nil)
