provider "restful" {}

resource "restful_rest_call" "Call" {
  method = var.method
  uri = var.uri
  headers = var.headers
  request_body = var.request_body
  expected_response_code = var.expected_response_code
  expected_response_body = var.expected_response_body
  json_key_outputs = var.json_key_outputs
}