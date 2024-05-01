# Chat App

## Running the app
`docker-compose up`

## Deploying to Minikube

1. Build, push images, and deploy (repeat for messages, servers, users)

   ```
   cd messages
   make deploy
   ```

2. Create tunnel to access load balancers

   ```
   sudo minikube tunnel
   ```

3. List services

   ```
   kubectl --context minikube --namespace main get svc
   ``` 
