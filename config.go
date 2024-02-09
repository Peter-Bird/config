/*
	peter-bird.com/config

	usage:

		const (
			LOAD_CFG_ERR_FMT = "Failed to load configuration: %s"
		)

		configFile := config.GetConfigFilePath()

		cfg, err := config.LoadConfig[Config](configFile)
		if err != nil {
			log.Fatalf(LOAD_CFG_ERR_FMT, err)
		}
*/

package config

import (
	"encoding/json"
	"flag"
	"os"
)

const (
	CONFIG_FLAG  = "config"
	CONFIG_HELP  = "Path to configuration file"
	ENV_VAR      = "CONFIG_FILE"
	DEFAULT_FILE = "./config/config.json"
)

// LoadConfig is a generic function to load a
// config file into a provided struct type.
func LoadConfig[T any](filePath string) (*T, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config T
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// GetConfigFilePath retrieves the configuration file path.
// It first checks for a command-line argument, then an environment variable,
// and falls back to a default if neither are provided.
func GetConfigFilePath() string {
	// Command-line flag
	cmdConfigFile := flag.String(CONFIG_FLAG, "", CONFIG_HELP)
	flag.Parse()

	if *cmdConfigFile != "" {
		return *cmdConfigFile
	}

	// Environment variable
	envConfigFile := os.Getenv(ENV_VAR)
	if envConfigFile != "" {
		return envConfigFile
	}

	// Default value
	return DEFAULT_FILE
}

/*
	Note:
	-----
	flag.Parse() will parse the command line arguments
	each time GetConfigFilePath is called.

	If this function is called multiple times,
	you might consider parsing flags only once or
	using a different mechanism to avoid repeated parsing.

	Flag Parsing in Libraries:
	--------------------------
	If this package is intended to be used as a library
	in other applications, be cautious with flag.Parse()
	inside the library.

	This can interfere with the command-line parsing of the
	main application.

	Typically, command-line parsing is expected to be done
	in the main package of the application.
*/
