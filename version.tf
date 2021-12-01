terraform {
  required_providers {
    example = {
      version = "~> 1.0.2"
      //source  = "software1tkh/example"
      source  = "terraform-example.com/exampleprovider2/example"
    }
  }
}
