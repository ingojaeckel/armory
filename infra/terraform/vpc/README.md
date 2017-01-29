# terraform/vpc

This can be used for the initial VPC setup. It likely only needs to be run once and likely does not have to be changed after.

# Requirements

* Create new `terraform` key. Download the `terraform.pem` file and store it within `~/.ssh/`. Make sure it is only readable by you: `chmod 400 terraform.pem`.
* Create a `terraform` user in `us-east-1`. Attach the `EC2Full` and `VPCFull` policies to the user account.

# Steps

```
terraform plan
terraform apply
```
