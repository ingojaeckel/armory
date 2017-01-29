#!/bin/sh
wget https://aws-codedeploy-us-east-1.s3.amazonaws.com/latest/install
chmod +x ./install
sudo ./install auto
sudo yum install --assumeyes golang
sudo easy_install supervisor
