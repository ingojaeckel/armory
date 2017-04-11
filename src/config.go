package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type AppConfig struct {
	// Environment Name of the environment mostly for logging purposes.
	Environment string
	// LogglyEnabled True if logging via Loggly should be enabled. False otherwise.
	LogglyEnabled bool
	// LogglyToken Used for authenticating against loggly.
	LogglyToken string
	// FlushDuration The time after which logs will be flushed to send them to loggly.
	FlushDurationSeconds int32

	// AmiID Ids used to control which worker instance will be launched and where. This should be the latest worker AMI ID.
	AmiID string
	// VpcSubnet The subnet in which the instance should be launched.
	VpcSubnet string
	// IamInstanceProfile Instance profile to be used for the worker instances.
	IamInstanceProfile string

	// SecurityGroup The EC2 security group used for the worker instances.
	SecurityGroup string
	// Keyname EC2 key name of the SSH key to be used for worker instances.
	Keyname string
	// AwsAccessKey AWS access key that allows to launch & terminate worker instances.
	AwsAccessKey string
	// AwsSecretKey AWS secret key that allows to launch & terminate worker instances.
	AwsSecretKey string
	// WorkerUsername The user name used to login to the worker instance. This will depend on the AMI used on the worker. For Amazon Linux it will likely be ec2-user.
	WorkerUsername string

	// GatlingInitialMemory Inital JVM heap memory available to Gatling.
	GatlingInitialMemory string
	// GatingMaxMemory Max JVM heap memory available to Gatling. Needs at least a t2.micro instance.
	GatingMaxMemory string
	// GatlingRoot Location of gatling installation.
	GatlingRoot string
	// SkipInstanceCreation True if no worker instance should be created. False otherwise. If true, the fallback worker IP address will be used.
	SkipInstanceCreation bool
	// FallbackWorkerIPAddress The private IP of the worker instance that will be used if the instance creation is skipped.
	FallbackWorkerIPAddress string
}

type Config struct {
	Base                         AppConfig
	FlushDuration                time.Duration
	IamInstanceProfile           *ec2.IamInstanceProfileSpecification
	EnvVars                      []string
	GatlingBinary, ResultsFolder string
}

var conf *Config

func initConfiguration() {
	if conf != nil {
		return
	}

	appConfig := getAppConfig()

	// TODO Validate appConfig configuration & prevent startup if necessary.
	conf = &Config{
		Base:               *appConfig,
		FlushDuration:      time.Duration(appConfig.FlushDurationSeconds) * time.Second,
		IamInstanceProfile: &ec2.IamInstanceProfileSpecification{Arn: aws.String(appConfig.IamInstanceProfile)},
		EnvVars:            []string{fmt.Sprintf("JAVA_OPTS=\"-Xms%s -Xmx%s\"", appConfig.GatlingInitialMemory, appConfig.GatingMaxMemory)},
		GatlingBinary:      fmt.Sprintf("%s/bin/gatling.sh", appConfig.GatlingRoot),
		ResultsFolder:      fmt.Sprintf("%s/results/testname", appConfig.GatlingRoot),
	}

	fmt.Printf("initialized config: %s\n", conf.Base.GatlingRoot)
}
