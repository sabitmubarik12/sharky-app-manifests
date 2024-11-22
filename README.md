# Sharky App Manifests

This repository contains the Kubernetes deployment manifests and application assets required for deploying and running the Sharky application.

## Repository Structure

```
.
├── manifests/              # Kubernetes manifests for deploying the Sharky application
│   ├── deployment.yaml     # Deployment file for application services
│   ├── service.yaml        # Service definition for the application
│   ├── ingress.yaml        # Ingress configuration for traffic management
│   └── configmap.yaml      # ConfigMap for application settings
├── assets/                 # Dockerfile and application code written in Go
│   ├── Dockerfile          # Dockerfile for building the application image
│   ├── main.go             # Main Go application source code
│   └── handlers/           # Directory for Go handlers and supporting logic
└── README.md               # Project documentation
```
Key Features
Manifests Directory
Contains Kubernetes manifests for deploying the Sharky application, including:
Deployments
Services
Ingress
ConfigMaps
Assets Directory
Houses the application's Go source code and Dockerfile for containerizing the application.
Prerequisites
Kubernetes Cluster (Google Kubernetes Engine or equivalent)
kubectl CLI configured for your cluster
Docker CLI for building application images
Traefik or any other Ingress controller for traffic management (optional)
Deployment Instructions
Step 1: Build the Docker Image
Navigate to the assets directory:

``` cd assets```

Build the Docker image:

```docker build -t sharky-app:v1 .```
Push the Docker image to a container registry (e.g., Docker Hub or ECR):

```docker tag sharky-app:v1 <your-repo>/sharky-app:v1
docker push <your-repo>/sharky-app:v1```
Step 2: Deploy the Application to Kubernetes
Navigate to the manifests directory:

```cd ../manifests```
Apply the manifests:

```kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl apply -f ingress.yaml
kubectl apply -f configmap.yaml```
Step 3: Verify the Deployment
Check the status of pods:

```kubectl get pods```
Verify the service is accessible:

```kubectl get svc```
If Ingress is configured, access the application using the specified DNS or IP.

Notes
Update the deployment.yaml file with the correct Docker image tag from your container registry.
Ensure the DNS settings are properly configured for Ingress to route traffic to the application.
Modify configmap.yaml to adjust application-specific configurations.

Application Features
The Sharky application, written in Go, is designed for scalability and performance. Key features include:
Modular architecture with extensible handlers.
Optimized Dockerfile for containerization.

Feel free to reach out if you have questions or need further assistance!
