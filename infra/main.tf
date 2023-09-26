provider "aws" {
  region = "us-west-2"  # Modify as needed
}

module "eks_cluster" {
  source          = "terraform-aws-modules/eks/aws"
  cluster_name    = "harrycorp-forex"
  cluster_version = "1.21"
  subnets         = ["subnet-12345678", "subnet-23456789"]  # Modify with actual subnet IDs
  vpc_id          = "vpc-01234567"  # Modify with actual VPC ID
  worker_groups = {
    forex_workers = {
      instance_type = "m5.large"  # Modify instance type as needed
      asg_max_size  = 5
    }
  }
}

output "kubeconfig" {
  value = module.eks_cluster.kubeconfig
}
