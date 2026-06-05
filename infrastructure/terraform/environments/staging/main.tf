# Staging Environment Terraform Config

terraform {
  required_version = ">= 1.7.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
  
  # Remote state configuration placeholder
  # backend "s3" {
  #   bucket         = "paisa-terraform-state"
  #   key            = "staging/terraform.tfstate"
  #   region         = "ap-south-1"
  #   dynamodb_table = "paisa-terraform-locks"
  # }
}

provider "aws" {
  region = "ap-south-1" # Mumbai region standard for Indian users
}

module "vpc" {
  source     = "../../modules/vpc"
  env        = "staging"
  cidr_block = "10.0.0.0/16"
}
