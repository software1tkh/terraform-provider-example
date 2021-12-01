data "example_server" "my-server-name" {	
}

output "output" {
  value = ["${data.example_server.my-server-name.outputv}"]
}
