# Sharky App Manifests

This repository contains the Kubernetes deployment manifests and application assets required for deploying and running the Sharky application.

## Repository Structure

```plaintext
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
``` cd assets
