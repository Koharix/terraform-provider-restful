provider "restful" {}

resource "restful_rest_call" "GET1" {
  method = var.method
  uri = var.uri
  headers = var.headers
  request_body = var.request_body
  expected_response_code = var.expected_response_code
  expected_response_body = var.expected_response_body
}