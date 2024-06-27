# Multi-Container Application Deployment with Docker Compose and Kubernetes

## Overview

This project demonstrates a multi-container application deployment consisting of a frontend, backend, and a database using Docker Compose for local development and Kubernetes for production deployment. It also includes automated infrastructure scaling with Terraform, unit tests for Terraform code, and a rollback strategy for infrastructure changes.

## Project Structure

- `frontend/`: Contains the React frontend code.
- `backend/`: Contains the Node.js backend code.
- `database/`: Contains the database setup (e.g., PostgreSQL).
- `kubernetes/`: Contains Kubernetes deployment and service YAML files.
- `terraform/`: Contains Terraform configuration for AWS infrastructure.
- `tests/`: Contains Terratest unit tests for Terraform code.

## Prerequisites

- Docker
- Docker Compose
- Minikube
- Kubectl
- Terraform
- Go (for Terratest)
- AWS CLI (if deploying on AWS)

## Getting Started :

### 1. Clone the Repositories
//```sh

      git clone https://github.com/Anand-1432/Techdome-backend.git backend
      git clone https://github.com/Anand-1432/Techdome-frontend.git frontend

### 2. Build and Push Docker Images
  Ensure Docker is running and then build and push the images to Docker Hub.

Build Docker Images
  Navigate to each directory and build the Docker images:

cd backend
      docker build -t your-dockerhub-username/backend:latest .
      docker push your-dockerhub-username/backend:latest

cd ../frontend
      docker build -t your-dockerhub-username/frontend:latest .
      docker push your-dockerhub-username/frontend:latest

### 3. Stop Docker Compose Services
  If running, stop the Docker Compose services:

    docker-compose down

### 4. Start Minikube
  Start Minikube to create a local Kubernetes cluster:

    minikube start

### 5. Deploy to Kubernetes
  Apply the Kubernetes deployment and service files:

    kubectl apply -f kubernetes/frontend-deployment.yaml
    kubectl apply -f kubernetes/frontend-service.yaml
    kubectl apply -f kubernetes/backend-deployment.yaml
    kubectl apply -f kubernetes/backend-service.yaml

### 6. Access the Services
  Use Minikube to get the service URLs:
    
    minikube service frontend --url
    minikube service backend --url

### 7. Infrastructure Automation with Terraform
  Navigate to the terraform directory and apply the configuration:

    cd terraform
    terraform init
    terraform apply

  This will create an AWS Auto Scaling Group with policies to scale based on CPU utilization.

### 8. Unit Testing with Terratest
  Ensure you have Go installed and run the tests:

    cd tests
    go test -v

### 9. Rollback Strategy
  To rollback infrastructure changes, use the backup state file:

    cp terraform.tfstate.backup terraform.tfstate
    terraform apply

  ### DIRECTORY STRUCTURE :
.
├── backend/
│   ├── Dockerfile
│   ├── package.json
│   ├── server.js
│   └── ...
├── frontend/
│   ├── Dockerfile
│   ├── package.json
│   ├── public/
│   ├── src/
│   └── ...
├── kubernetes/
│   ├── frontend-deployment.yaml
│   ├── frontend-service.yaml
│   ├── backend-deployment.yaml
│   ├── backend-service.yaml
│   └── ...
├── terraform/
│   ├── main.tf
│   ├── variables.tf
│   ├── outputs.tf
│   └── ...
└── tests/
    ├── main_test.go
    └── ...

 ### Setup Terraform :
Step 1: Create a Terraform configuration file 'main.tf'
Step 2: Apply the Terraform configuration

            terraform init
            terraform apply
            
### Implementing Unit Tests for Terraform Code:
      To implement unit tests for Terraform code, you can use terratest, a Go library that makes it easier to write automated tests for your infrastructure code.

Step 1: Install Terratest
Step 2: Write Terratest Test Cases 'main_test.go'
Step 3: Run the Tests

      go test -v

### Demonstrating a Rollback Strategy for Infrastructure Changes:
Step 1: Create a Backup of Your State File
      Before making changes, ensure you have a backup of your current state file.

      cp terraform.tfstate terraform.tfstate.backup

Step 2: Make Changes and Apply Them
      Make changes to your Terraform configuration and apply them.
      
      terraform apply

Step 3: Rollback to the Previous State
      If something goes wrong, you can rollback to the previous state using the backup.

      cp terraform.tfstate.backup terraform.tfstate
      terraform apply

Alternatively, you can use terraform destroy to tear down the new infrastructure and re-apply the original configuration.
