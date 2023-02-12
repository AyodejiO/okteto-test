# How To
Kindly go through the following steps to deploy this application:

1. You need access to https://okteto.assessment.jdm.okteto.net to deploy this application. Run all the required go commands to get the dependencies.

2. Run the command `cp .env.example .env` to create a clone of your local environment variables. In the new [.env file](./.env), kindly provide the name of your K8 namespace. You can only leave this blank if you have elevated permission in the cluster.

3. You should run `okteto deploy --build` to create the deployment, or if you modify the [K8s.yml](./k8s.yml) or the [okteto](./okteto.yml) file. 

4. Once your deployment has been set up, you can connect to the Okteto development platform using the command `okteto up`, and once you're connected, you should start the application using the command `go run main.go`. You can run the test routes in the [pods.http](./http/pods.http) file to test the respective routes.

5. As an added advantage, you can manage the K8 cluster by hooking to it locally. Connect your kubectl using the `okteto context use https://okteto.assessment.jdm.okteto.net`, and view your deployment using the `kubectl get deploy -n <your namespace>`. You can also manually scale the deployment using the `kubectl scale --replicas=<your desired number> deployment/tech-assessment` command.

6. To run the tests, simply run the following command `go test ./...` and you're good to go. Please feel free to make changes and send in your PR. I look forward to the collab. 