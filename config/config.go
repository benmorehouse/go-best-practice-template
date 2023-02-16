package configs

import (
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	session "github.com/aws/aws-sdk-go/aws/session"
)

// Configuration struct will embody the initial configuration file
// which is unexported: no other entities should alter this beyond
// our surfer package
type Configuration struct {
	URL        string
	Port       int64
	Local      bool
	Verbose    bool
	ZoneConfig *aws.Config
	DynamoHost string
	AWSRegion  string
}

// LocalEnvVar to tell if local run
const LocalEnvVar = "LOCAL_RUN"

// New will return a new configuration
func New() *Configuration {
	localstr := os.Getenv(LocalEnvVar)
	if strings.ToLower(localstr) == "true" {
		return parseLocalEnvironment()
	}
	return parseLambdaEnvironment()
}

// NewAWSSession will spin up a new session with AWS
func (c *Configuration) NewAWSSession() (*session.Session, error) {
	return session.NewSession(c.ZoneConfig)
}

func parseLocalEnvironment() *Configuration {
	// availability zone must be this for local dynamo
	const zone = "eu-west-2"

	return &Configuration{
		Local:      true,
		Verbose:    true,
		Port:       0,
		URL:        "",
		ZoneConfig: aws.NewConfig().WithRegion(zone),
		DynamoHost: "http://localhost:8000",
		AWSRegion:  zone,
	}
}

// parseLambdaEnvironment will parse the environment variables set by
// the aws lambda application. References : https://docs.aws.amazon.com/lambda/latest/dg/golang-envvars.html
func parseLambdaEnvironment() *Configuration {
	// zone stands for the availability zone that we are running this function from
	const zone = "us-west-2"

	return &Configuration{
		Local:      false,
		Verbose:    false,
		Port:       0,
		URL:        "",
		ZoneConfig: aws.NewConfig().WithRegion(zone),
		AWSRegion:  zone,
	}
}
