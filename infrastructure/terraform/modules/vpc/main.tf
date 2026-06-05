# VPC Module Skeleton (Blueprint Section 10.3)
# Defines public subnets, app subnets, database subnets, and compute subnets.

variable "env" {
  type        = string
  description = "Deployment environment"
}

variable "cidr_block" {
  type        = string
  default     = "10.0.0.0/16"
  description = "VPC CIDR block"
}

resource "aws_vpc" "main" {
  cidr_block           = var.cidr_block
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name        = "paisa-${var.env}-vpc"
    Environment = var.env
    Project     = "paisa"
    ManagedBy   = "terraform"
  }
}
