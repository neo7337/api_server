# API SERVER in GOLANG

<b>Commands</b>

```
docker build -t "<IMAGE_NAME>:<VERSION>" .
```
```
docker tag <IMAGE_NAME>:<VERSION> gcr.io/<PROJECT_ID>/<IMAGE_NAME>:latest
```
```
docker push gcr.io/<PROJECT_ID>/<IMAGE_NAME>:latest
```
```
docker run --publish 5000:8989 --name server <IMAGE_NAME>:<VERSION>
```
//Kill a running container
```
docker rm --force server
```
After pushing the image to Google Container Registry, we can deploy it on VM Instances, Google Kubernetes Engine, Google App Engine

Basic server test curl command on localhost:
```
curl --head localhost:8989
```