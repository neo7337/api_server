# API SERVER

<b>Commands</b>

docker build -t "<IMAGE_NAME>:<VERSION>" .

docker tag <IMAGE_NAME>:<VERSION> gcr.io/<PROJECT_ID>/<IMAGE_NAME>:latest

docker push gcr.io/<PROJECT_ID>/<IMAGE_NAME>:latest

After pushing the image to Google Container Registry, we can deploy it on VM Instances, Google Kubernetes Engine, Google App Engine