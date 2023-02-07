package config

import "PACKAGENAME/utilities"

// Base Configures

var ServerPort string = utilities.GetEnv(
	"SERVER_PORT",
	"8080",
)
