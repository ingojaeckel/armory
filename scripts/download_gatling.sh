#!/bin/sh
VERSION=2.2.2
curl https://repo1.maven.org/maven2/io/gatling/highcharts/gatling-charts-highcharts-bundle/${VERSION}/gatling-charts-highcharts-bundle-${VERSION}-bundle.zip -o gatling-charts-highcharts-bundle-${VERSION}-bundle.zip
unzip -q gatling-charts-highcharts-bundle-${VERSION}-bundle.zip
mv gatling-charts-highcharts-bundle-${VERSION} gatling
rm -rf gatling/user-files/simulations/*
