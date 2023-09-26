
```markdown
# HarryCorp Forex Pricing Web Application

This project demonstrates the setup and deployment of a Forex pricing web application on Amazon Elastic Kubernetes Service (EKS) using CI/CD with GitHub Actions, Helm charts, Terraform for infrastructure, and Prometheus/Grafana for monitoring.

## Prerequisites

Before you begin, make sure you have the following prerequisites installed and configured:

1. **Git**: Install Git on your local machine.
   - [Git Installation Guide](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

2. **Docker**: Install Docker to build and run container images.
   - [Docker Installation Guide](https://docs.docker.com/get-docker/)

3. **kubectl**: Install `kubectl` to interact with your Kubernetes cluster.
   - [kubectl Installation Guide](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

4. **Terraform**: Install Terraform for managing infrastructure as code.
   - [Terraform Installation Guide](https://learn.hashicorp.com/tutorials/terraform/install-cli)

5. **AWS CLI**: Configure AWS CLI with your AWS credentials.
   - [AWS CLI Configuration Guide](https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html)

6. **Helm**: Install Helm for managing Kubernetes applications.
   - [Helm Installation Guide](https://helm.sh/docs/intro/install/)

## Project Setup

1. Clone the project repository:
   ```bash
   git clone github.com/<>/harrycorp-forex-app
   cd harrycorp-forex-app
   ```

2. Build the Docker image for the Forex pricing web application:
   ```bash
   cd forex-app
   docker build -t harrycorp/forex-pricing:latest .
   ```

3. Push the Docker image to a container registry (e.g., Docker Hub, Amazon ECR):
   ```bash
   docker push harrycorp/forex-pricing:latest
   ```

4. Create an Amazon EKS cluster using Terraform:
   ```bash
   cd ../infra
   terraform init
   terraform apply
   ```

5. Deploy the Forex pricing web application to the EKS cluster using Helm:
   ```bash
   kubectl config use-context harrycorp-forex  # Use the EKS cluster context
   helm upgrade --install forex-app ../forex-app --namespace default --values ../forex-app/values.yaml
   ```

6. Set up monitoring with Prometheus and Grafana (assuming you have Prometheus and Grafana deployments in your cluster).

## Accessing the Application

Once the deployment is complete, you can access the Forex pricing web application:

- Obtain the external IP of the application:
  ```bash
  kubectl get svc -n default
  ```

- Open a web browser and navigate to the application using the external IP.

## Monitoring Metrics

You can access Prometheus and Grafana for monitoring:

- Prometheus: http://<prometheus_external_ip>:9090
- Grafana: http://<grafana_external_ip>:3000

## Cleaning Up

To clean up and remove resources:

1. Delete the EKS cluster using Terraform:
   ```bash
   cd ../infra
   terraform destroy
   ```

2. Delete the Docker image from the container registry (if applicable).

3. Optionally, delete the Helm release:
   ```bash
   helm uninstall forex-app -n default
   ```

4. Delete any remaining resources, such as ConfigMaps and Services.

Remember to adjust configuration files and settings to match your specific requirements and preferences.
```

Make sure to replace `<repository_url>` with the actual URL of your project repository if it's hosted on a version control platform. This `instructions.md` file provides a step-by-step guide for setting up and running the project.