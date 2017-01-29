#!/bin/bash
S3_BUCKET=insert_your_s3_bucket
BUNDLE_TYPE=tgz
CODEDEPLOY_APP=armory
DEPLOY_GROUP=dev

BRANCH=`git rev-parse --abbrev-ref HEAD`
REVISION=`git log --pretty=format:'%h' -n 1`
TIMESTAMP=$(date +%Y.%m.%d-%H.%M)

BUILD_NUMBER=${TIMESTAMP}-${BRANCH}-${REVISION}
S3_KEY=app-${BUILD_NUMBER}.tar.gz

export AWS_DEFAULT_REGION=us-east-1
export AWS_ACCESS_KEY_ID=insert_your_access_key
export AWS_SECRET_ACCESS_KEY=insert_your_secret_key

aws s3 cp package/app.tar.gz s3://${S3_BUCKET}/app-${BUILD_NUMBER}.tar.gz

DEPLOYMENT_ID=`aws deploy create-deployment              \
  --application-name ${CODEDEPLOY_APP}    \
  --deployment-group-name ${DEPLOY_GROUP} \
  --description "deployment"              \
  --s3-location bucket=${S3_BUCKET},bundleType=${BUNDLE_TYPE},key=${S3_KEY} | jq -r .deploymentId`

echo "Deployment ID: ${DEPLOYMENT_ID}"

for i in `seq 60`;
do
	DEPLOYMENT_STATUS=`aws deploy get-deployment --deployment-id ${DEPLOYMENT_ID} | jq -r .deploymentInfo.status`

	if [ "${DEPLOYMENT_STATUS}" = "Succeeded" ]; then
		echo "Deployment successful!"
		exit 0;
	fi
	if [ "${DEPLOYMENT_STATUS}" = "Failed" ]; then
		echo "Deployment failed!"
		exit 1;
	fi

	echo "Deployment status: ${DEPLOYMENT_STATUS}"
	sleep 1
done

echo "Deployment timed out. Check Code Deploy."
