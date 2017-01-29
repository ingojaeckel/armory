resource "aws_instance" "frontend" {
  ami = "${var.default_ami}"
	instance_type = "${var.ec2_instance_type}"
	key_name = "${var.aws_key_name}"
  iam_instance_profile = "${aws_iam_instance_profile.instance_profile.id}"
	vpc_security_group_ids = ["${aws_security_group.frontend-sg.id}"]
  subnet_id = "${var.ec2_subnet_id}"
  associate_public_ip_address = true

  tags {
    Name = "frontend"
    Environment = "dev"
  }
}

resource "aws_codedeploy_app" "app" {
    name = "${var.app_name}"
}

resource "aws_iam_role_policy" "role_policy" {
    name = "IamRolePolicy"
    role = "${aws_iam_role.codedeploy_role.id}"
    policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "autoscaling:CompleteLifecycleAction",
                "autoscaling:DeleteLifecycleHook",
                "autoscaling:DescribeAutoScalingGroups",
                "autoscaling:DescribeLifecycleHooks",
                "autoscaling:PutLifecycleHook",
                "autoscaling:RecordLifecycleActionHeartbeat",
                "ec2:DescribeInstances",
                "ec2:DescribeInstanceStatus",
                "tag:GetTags",
                "tag:GetResources",
                "s3:Get*",
                "s3:List*"
            ],
            "Resource": "*"
        }
    ]
}
EOF
}

resource "aws_iam_role" "codedeploy_role" {
    name = "IamCodeDeployRole"
    assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": [
          "codedeploy.amazonaws.com"
        ]
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

resource "aws_codedeploy_deployment_group" "deployment_group" {
    app_name = "${aws_codedeploy_app.app.name}"
    deployment_group_name = "dev"
    service_role_arn = "${aws_iam_role.codedeploy_role.arn}"
    ec2_tag_filter {
        key = "Name"
        type = "KEY_AND_VALUE"
        value = "frontend"
    }
}

resource "aws_security_group" "frontend-sg" {
  description = "Allow incoming SSH & HTTP, Allow outgoing connections."
  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port = 8080
    to_port = 8080
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port = -1
    to_port = -1
    protocol = "icmp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  vpc_id = "${var.ec2_vpc_id}"
  tags {
    Name = "FrontendSG"
  }
}

resource "aws_iam_instance_profile" "instance_profile" {
    name = "IamInstanceProfile"
    roles = ["${aws_iam_role.role.name}"]
}

resource "aws_iam_role" "role" {
    name = "IamRole"
    path = "/"
    assume_role_policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": [
              "sts:AssumeRole"
            ],
            "Principal": {
               "Service": "ec2.amazonaws.com"
            },
            "Effect": "Allow",
            "Sid": ""
        }
    ]
}
EOF
}

resource "aws_iam_policy" "policy" {
    name = "IamPolicy"
    path = "/"
    description = "IAM Policy to allow read-access to EC2 & S3 and autoscaling."
    policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "autoscaling:CompleteLifecycleAction",
        "autoscaling:DeleteLifecycleHook",
        "autoscaling:DescribeAutoScalingGroups",
        "autoscaling:DescribeLifecycleHooks",
        "autoscaling:PutLifecycleHook",
        "autoscaling:RecordLifecycleActionHeartbeat",
        "ec2:DescribeInstances",
        "ec2:DescribeInstanceStatus",
        "tag:GetTags",
        "tag:GetResources",
        "s3:Get*",
        "s3:List*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "policy_attachment" {
    role = "${aws_iam_role.role.name}"
    policy_arn = "${aws_iam_policy.policy.arn}"
}
