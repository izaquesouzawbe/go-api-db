package config

type Config struct {
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		Host     string `yaml:"host"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`

	Server struct {
		Port          string `yaml:"port"`
		Authorization string `yaml:"authorization"`
	} `yaml:"server"`
}

type InsertRequest struct {
	Table       string       `json:"table"`
	PrimaryKeys []string     `json:"primary_keys"`
	Records     []RecordType `json:"records"`
}

type RequestBodySQL struct {
	SQL string `json:"sql"`
}

type RecordType map[string]interface{}
