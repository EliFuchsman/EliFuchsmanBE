#ecr.tf

terraform {
  required_providers {
    aws = ">= 5.36.0, <= 5.36.0"
  }
}

provider "aws" {
  region = var.AWS_REGION
}
