package config

type DbConnection struct {
	Driver    string
	Host      string
	Port      string
	Database  string
	Username  string
	Password  string
	Charset   string
	Collation string
	Prefix    string
	Strict    bool
	Schema    string
}

var (
	DefaultConnect string
	Connections    map[string]*DbConnection
)

func initDatabase() {
	DefaultConnect = getEnvValueWithDefault("DB_CONNECTION", "mysql")
	Connections = map[string]*DbConnection{
		"mysql": &DbConnection{
			Driver:    "mysql",
			Host:      getEnvValueWithDefault("DB_HOST", "localhost"),
			Port:      getEnvValueWithDefault("DB_PORT", "3306"),
			Database:  getEnvValueWithDefault("DB_DATABASE", "forge"),
			Username:  getEnvValueWithDefault("DB_USERNAME", "forge"),
			Password:  getEnvValueWithDefault("DB_PASSWORD", ""),
			Charset:   "utf8",
			Collation: "utf8_general_ci",
			Prefix:    "",
			Strict:    false,
		},
	}
}
