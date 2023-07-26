provider "aws" {
  region = "us-west-2"
}

terraform {
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "benmorehouse"

    workspaces {
      name = "mono"
    }
  }
}

locals {
  region      = "us-west-2"
  account_id  = "721252812367"
}

resource "random_password" "example_dev" {
  length           = 16
  special          = true
  override_special = "!#$%&*()-_=+[]{}<>:?"
}
