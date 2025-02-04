package config

type Config struct {
	// Host is the hostname to listen on.
	Host string `json:"host"`
	// Port is the port to listen on.
	Port string `json:"port"`
}
