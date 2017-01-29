#!/bin/sh

echo "Removing Java 7.."
sudo yum remove  --assumeyes java-1.7.0-openjdk.x86_64
echo "Installing Java 8.."
sudo yum install --assumeyes ruby wget zip java-1.8.0-openjdk.x86_64 htop bzip2

# Setup Gatling
echo "Setting up Gatling.."
cd /tmp
/tmp/scripts/download_gatling.sh
rm -rf gatling/user-files/simulations/*
mkdir -p gatling/user-files/simulations/
cp -v -r /tmp/sim/com gatling/user-files/simulations/
sudo chown -R ec2-user:ec2-user /tmp/gatling*
