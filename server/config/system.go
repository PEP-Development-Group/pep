package config

type System struct {
	Env           string `mapstructure:"env" json:"env" yaml:"env"`
	Addr          int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType        string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	OssType       string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"useMultipoint" yaml:"use-multipoint"`

	FirstDay   string `mapstructure:"first_day" json:"first_day" yaml:"first_day"`
	CancelNums int    `mapstructure:"cancel_nums" json:"cancel_nums" yaml:"cancel_nums"`
}
