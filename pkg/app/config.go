package app

import "os"

import "fmt"

type config struct {
	port string
}

// Config would be used to contain environment variables
var Config *config

// Port would return the port on which web server is listening the request to
func (c *config) Port() string {
	return c.port
}

// init is called automatically only once when the package is imported
func init() {

	port := os.Getenv("PORT")

	if port == "" {
		panicMissingEnvVariable("PORT")
	}

	Config = &config{
		port: port,
	}
}

func panicMissingEnvVariable(varName string) {
	panic(fmt.Sprintf("missing environment variable -> %s ", varName))
}
