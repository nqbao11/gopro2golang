# DevOps Capstone Requirement

|CRITERIA|MEETS SPECIFICATIONS|
|-|-|
|Create Github repository with project code.|All project code is stored in a GitHub repository and a link to the repository has been provided for reviewers.|
|Use image repository to store Docker images|The project uses a centralized image repository to manage images built in the project. After a clean build, images are pushed to the repository.|
|Execute linting step in code pipeline|Code is checked against a linter as part of a Continuous Integration step (demonstrated w/ two screenshots)|
|Build a Docker container in a pipeline|The project takes a Dockerfile and creates a Docker container in the pipeline.|
|The Docker container is deployed to a Kubernetes cluster|The cluster is deployed with CloudFormation or Ansible. This should be in the source code of the student’s submission.|
|Use Blue/Green Deployment or a Rolling Deployment successfully|The project performs the correct steps to do a blue/green or rolling deployment into the environment selected. Submit the following screenshots as evidence of the successful completion of chosen deployment methodology|

- Screenshot of the Circle CI or Jenkins pipeline showing all stages passed successfully.
- Screenshot of your AWS EC2 page showing the newly created (for blue/green) or modified (for rolling) instances running as the EKS cluster nodes.
- Screenshot of the kubectl command output showing that the deployment is successful, pods are running, and the service can be accessed via an external IP or port forwarding.
- Screenshot showing that you can access the application after deployment.