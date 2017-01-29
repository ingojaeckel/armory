package main

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type AppConfig struct {
	// Environment Name of the environment mostly for logging purposes.
	Environment string `json:"env"`
	// LogglyToken Used for authenticating against loggly.
	LogglyToken string `json:"loggly.token"`
	// FlushDuration The time after which logs will be flushed to send them to loggly.
	FlushDurationSeconds int32 `json:"loggly.flush.seconds"`

	// AmiID Ids used to control which worker instance will be launched and where. This should be the latest worker AMI ID.
	AmiID string `json:"aws.ami.id"`
	// VpcSubnet The subnet in which the instance should be launched.
	VpcSubnet string `json:"aws.vpc.subnet"`
	// IamInstanceProfile Instance profile to be used for the worker instances.
	IamInstanceProfile string `json:"aws.instance.profile"`

	// SecurityGroup The EC2 security group used for the worker instances.
	SecurityGroup string `json:"aws.security.group"`
	// Keyname EC2 key name of the SSH key to be used for worker instances.
	Keyname string `json:"aws.keyname"`
	// AwsAccessKey AWS access key that allows to launch & terminate worker instances.
	AwsAccessKey string `json:"aws.access.key"`
	// AwsSecretKey AWS secret key that allows to launch & terminate worker instances.
	AwsSecretKey string `json:"aws.secret.key"`
	// WorkerUsername The user name used to login to the worker instance. This will depend on the AMI used on the worker. For Amazon Linux it will likely be ec2-user.
	WorkerUsername string `json:"aws.username"`

	// GatlingInitialMemory Inital JVM heap memory available to Gatling.
	GatlingInitialMemory string `json:"gatling.memory.initial"`
	// GatingMaxMemory Max JVM heap memory available to Gatling. Needs at least a t2.micro instance.
	GatingMaxMemory string `json:"gatling.memory.max"`
	// GatlingRoot Location of gatling installation.
	GatlingRoot string `json:"gatling.root"`
	// SkipInstanceCreation True if no worker instance should be created. False otherwise. If true, the fallback worker IP address will be used.
	SkipInstanceCreation bool `json:"worker.startup.skip"`
	// FallbackWorkerIPAddress The private IP of the worker instance that will be used if the instance creation is skipped.
	FallbackWorkerIPAddress string `json:"worker.fallback.ip"`
}

type Config struct {
	Base               AppConfig
	FlushDuration      time.Duration
	IamInstanceProfile *ec2.IamInstanceProfileSpecification
}

var conf Config

func initConfiguration() error {
	// TODO Validate appConfig configuration & prevent startup if necessary.
	conf = Config{
		Base:               appConfig,
		FlushDuration:      time.Duration(appConfig.FlushDurationSeconds) * time.Second,
		IamInstanceProfile: &ec2.IamInstanceProfileSpecification{Arn: aws.String(appConfig.IamInstanceProfile)},
	}
	return nil
}
