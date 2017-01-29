#!/bin/sh
rm -rf package
mkdir -p package/scripts

cp -v -r ../infra/codedeploy/*	package
cp -v app			package
cd package
tar -cvzf app.tar.gz *

# leave package here until next build to allow debugging
