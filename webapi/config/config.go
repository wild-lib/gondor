package config

import (
	"fmt"
	"os"

	"github.com/astro-bug/gondor/webapi/config/dialect"
	"gopkg.in/yaml.v2"
)

var cfg *Settings

type Settings struct {
	IsEmpty        bool                  `json:"-" yaml:"-"`
	Application    AppConfig             `json:"application" yaml:"application"`
	Connections    map[string]ConnConfig `json:"connections" yaml:"connections"`
	MicroServices  []ServConfig          `json:"micro_services" yaml:"micro_services"`
	ReverseTargets []ReverseTarget       `json:"reverse_targets" yaml:"reverse_targets"`
}

type AppConfig struct {
	Debug       bool `json:"debug" yaml:"debug"`
	PluralTable bool `json:"plural_table" yaml:"plural_table"`
}

type ConnConfig struct {
	DriverName  string `json:"driver_name" yaml:"driver_name"`
	TablePrefix string `json:"table_prefix" yaml:"table_prefix"`
	ReadOnly    string `json:"read_only" yaml:"read_only"`
	Params      dialect.ConnParams
}

type ServConfig struct {
	Protocol string `json:"protocol" yaml:"protocol"`
	Params   dialect.ConnParams
}

func GetSettings() *Settings {
	if cfg == nil {
		cfg = new(Settings)
		cfg.IsEmpty = true
	}
	return cfg
}

func ReadSettings(file string) (*Settings, error) {
	cfg = new(Settings)
	rd, err := os.Open(file)
	if err == nil {
		err = yaml.NewDecoder(rd).Decode(&cfg)
		if err == nil {
			cfg.IsEmpty = false
		}
	}
	return cfg, err
}

func SaveSettings(file string) error {
	if cfg = GetSettings(); cfg.IsEmpty {
		return fmt.Errorf("the settings is not exists")
	}
	wt, err := os.Open(file)
	if err == nil {
		err = yaml.NewEncoder(wt).Encode(cfg)
	}
	return err
}

func (cfg Settings) GetSource(name string) (ReverseSource, ConnConfig) {
	var ok bool
	src, c := ReverseSource{}, ConnConfig{}
	if c, ok = cfg.Connections[name]; !ok {
		return src, c
	}
	d := dialect.GetDialectByName(c.DriverName)
	if d != nil {
		src.Database = d.Name()
		src.ConnStr = d.GetDSN(c.Params)
	}
	return src, c
}
