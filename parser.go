package config

import (
	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"
)

// configurable is interface implemented by BaseConfig
type configurable interface {
	setConfigFunc(func(string) error)
}

// BaseConfig should be embed into your configuration struct
type BaseConfig struct {
	Config func(s string) error `long:"config" description:"The ini config file" no-ini:"true"`
	Global struct {
		Version string `long:"global-version" required:"true" ini-name:"version"`
	} `group:"global"`
}

func (b *BaseConfig) setConfigFunc(f func(s string) error) {
	b.Config = f
}

// Parse parses configuration file provided in --config flag
// cfgStruct should be a pointer
func Parse(cfgStruct interface{}) error {
	cfg, ok := cfgStruct.(configurable)
	if !ok {
		return errors.New("struct BaseConfig must be embed into cfgStruct")
	}
	p := flags.NewParser(cfg, flags.None)

	cfg.setConfigFunc(func(s string) error {
		ini := flags.NewIniParser(p)

		return ini.ParseFile(s)
	})

	_, err := p.Parse()
	return errors.Wrap(err, "failed to parse config file")
}
