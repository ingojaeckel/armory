variable "aws_access_key" {}
variable "aws_secret_key" {}
variable "aws_key_path" {}
variable "aws_key_name" {}

variable "app_name" {
  default = "armory"
}

variable "default_ami" {
  default = "insert_worker_ami_here"
}

variable "availability_zone" {
  default = "us-east-1a"
}

variable "aws_region" {
  description = "EC2 Region for the VPC"
  default = "us-east-1"
}

variable "ec2_instance_type" {
  description = "EC2 Instance Type"
  default = "t2.nano"
}

variable "vpc_cidr" {
  description = "CIDR for the whole VPC"
  default = "10.0.0.0/16"
}

variable "public_subnet_cidr" {
  description = "CIDR for the Public Subnet"
  default = "10.0.0.0/24"
}

variable "private_subnet_cidr" {
  description = "CIDR for the Private Subnet"
  default = "10.0.1.0/24"
}

variable "ec2_vpc_id" {
  description = "EC2 VPC ID"
  default = "vpc-1c459c7a"
}

variable "ec2_subnet_id" {
  description = "EC2 Subnet ID"
  default = "subnet-57ec727a"
}

variable "application_name" {
  default = "worker"
}

variable "environment" {
  default = "dev"
}

variable "deployment_group" {
  default = "blue"
}

variable "min_instance_count" {
  default = 1
}

variable "desired_capacity" {
  default = 1
}

variable "max_instance_count" {
  default = 1
}
