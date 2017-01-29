package main

// Insert your configuration here.
var appConfig = AppConfig{
	AwsAccessKey:            "your_access_key",
	AwsSecretKey:            "your_secret_key",
	LogglyToken:             "your_loggly_token",
	VpcSubnet:               "your_vpc_subnet",
	SecurityGroup:           "your_security_group",
	Keyname:                 "your_ec2_keyname",
	AmiID:                   "your_worker_ami_id",
	IamInstanceProfile:      "your_iam_instance_profile",
	Environment:             "dev",
	FlushDurationSeconds:    5,
	GatlingInitialMemory:    "256M",
	GatingMaxMemory:         "512M",
	WorkerUsername:          "ec2-user",
	GatlingRoot:             "/tmp/gatling",
	SkipInstanceCreation:    false,
	FallbackWorkerIPAddress: "",
}
