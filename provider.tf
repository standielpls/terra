provider "aws" {
  profile    = "default"
  region     = "us-east-1"
}

resource "aws_instance" "web1" {
  ami           = "ami-2757f631"
  instance_type = "t2.micro"
}
