# API SERVER in GOLANG

<b>Go Commands</b>

The below command will generate the golang executable build for linux platforms if your current system is linux
```
go build server.go
```

Once the build is generated we can run that build directly using below command
```
./server
```

<b>Docker Commands</b>

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
Kill a running container
```
docker rm --force server
```
After pushing the image to Google Container Registry, we can deploy it on VM Instances, Google Kubernetes Engine, Google App Engine

<b>Testing the server</b>

Basic server test curl command on localhost:
```
curl --head localhost:8989
```

<b>Kubernetes Commands</b>

Deploying Deployment Server
```
kubectl apply -f kubernetes/deployments/server.yaml
kubectl describe deployment server-deployment
```

Deploying Service
```
kubectl apply -f kubernetes/services/server.yaml
kubectl describe service server-deployment
minikube service server-deployment --url
```