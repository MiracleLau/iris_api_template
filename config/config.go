package config

type SiteConfig struct {
	SiteAddress string      `yaml:"site_address"`
	DbString    string      `yaml:"db_string"`
	LogPath     string      `yaml:"log_path"`
	KeyLookup	string      `yaml:"key_lookup"`
	Redis       RedisConfig `yaml:"redis"`
}

type RedisConfig struct {
	Network  string `yaml:"network"`
	Addr     string `yaml:"addr"`
	Port     int	 `yaml:"port"`
	Prefix   string `yaml:"prefix"`
	User	 string	`yaml:"user"`
	Password string `yaml:"password"`
	DB		int		`yaml:"db"`
}
