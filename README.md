# Cloud Run example
In this repo, you could find a simple example of a golang microservice that has been deployed on GCP Cloud Run. Cloud Run is built upon Knative and provides a high level of abstraction for deploying a container into either GKE or managed infrastructure by Google.

![image](https://user-images.githubusercontent.com/12199867/106071077-ee8ed180-6105-11eb-885e-93d08d2c88f5.png)
![image](https://user-images.githubusercontent.com/12199867/106071094-f9e1fd00-6105-11eb-87e0-aa21253a56e1.png)


# How-to
Setup the project on the GCP
```bash
$ export PROJECT_ID=project-example
$ export ACCOUNT_NAME:=account-example
$ export SERVICE_NAME:=service-name

$ gcloud auth login

$ gcloud projects create $PROJECT_ID
$ gcloud config set project $PROJECT_ID
```

Enable Google APIs
```bash
$ gcloud services enable cloudbuild.googleapis.com run.googleapis.com containerregistry.googleapis.com
```

Create Service Accounts
```bash

$ gcloud iam service-accounts create dp_$ACCOUNT_NAME \
  --description="Cloud Run deploy account" \
  --display-name="Cloud-Run-Deploy"

$ gcloud iam service-accounts create rt_$ACCOUNT_NAME \
  --description="Cloud Run runtime account" \
  --display-name="Cloud-Run-Runtime"


$ gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member=serviceAccount:dp_$ACCOUNT_NAME@$PROJECT_ID.iam.gserviceaccount.com \
  --role=roles/run.admin

$ gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member=serviceAccount:dp_$ACCOUNT_NAME@$PROJECT_ID.iam.gserviceaccount.com \
  --role=roles/storage.admin

$ gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member=serviceAccount:dp_$ACCOUNT_NAME@$PROJECT_ID.iam.gserviceaccount.com \
  --role=roles/iam.serviceAccountUser

$ gcloud projects add-iam-policy-binding $PROJECT_ID \
  --member=serviceAccount:rt_$ACCOUNT_NAME@$PROJECT_ID.iam.gserviceaccount.com \
  --role=roles/run.invoker

$ gcloud iam service-accounts keys create key.json \
    --iam-account dp_$ACCOUNT_NAME@$PROJECT_ID.iam.gserviceaccount.com
```

Deploy the container via CloudBuild
```bash
$ gcloud builds submit --config cloudbuild.yml --substitutions=_SERVICE_NAME=$SERVICE_NAME,_ACCOUNT_NAME=rt_$ACCOUNT_NAME
```

Deploy the container with buildpack
```bash
$ gcloud beta run deploy $PROJECT_ID --source .
```

## Links
https://github.com/ahmetb/cloud-run-faq
https://github.com/GoogleContainerTools/distroless
https://cloud.google.com/run/docs/how-to