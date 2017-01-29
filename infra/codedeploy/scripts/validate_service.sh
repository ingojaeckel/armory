#!/bin/bash
while true
do
  ps aux |grep app
  response=$(curl --write-out %{http_code} --silent --output /dev/null http://localhost:8080/health)

	if [ ${response} -eq 200 ]; then
		echo "Successfully hit health check endpoint!"
		exit 0;
	fi

  sleep 1
done
