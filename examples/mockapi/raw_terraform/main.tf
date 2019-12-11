provider "restful" {}

//Get all events, expected_response_body will be equal, response code expected to be 200 
resource "restful_rest_call" "GET1" {
  method = "GET"
  uri = "http://localhost:8080/events/1"
  expected_response_code = 200
  expected_response_body = "{\"ID\":\"1\",\"Title\":\"Introduction to Golang\",\"Description\":\"Come join us for a chance to learn how golang works and get to eventually try it out\"}"
  json_key_outputs = ["Title", "Description"]
}
//GET expected_response_body is in the response body but response body contains more data
resource "restful_rest_call" "GET2" {
  method = "GET"
  uri = "http://localhost:8080/events"
  expected_response_body = "[{\"Title\":\"Introduction to Golang\"}]"
}
//GET a specific item

resource "restful_rest_call" "POST1" {
  method = "POST"
  uri = "http://localhost:8080/event"
  headers = ["Content-Type:application/json"]
  request_body = "{\"ID\":\"2\",\"Title\":\"title\",\"Description\":\"desc2\"}"
  expected_response_body = "{\"ID\":\"2\",\"Title\":\"title\",\"Description\":\"desc2\"}"
  expected_response_code = 201
}
resource "restful_rest_call" "DELETE1" {
  method = "DELETE"
  uri = "http://localhost:8080/events/2"
  expected_response_code = 200
  depends_on = [restful_rest_call.GET3]
}
resource "restful_rest_call" "PATCH1" {
  method = "PATCH"
  uri = "http://localhost:8080/events/2"
  headers = ["Content-Type:application/json"]
  request_body = "{\"ID\":\"2\",\"Title\":\"title\",\"Description\":\"desc3\"}"
  expected_response_code = 200
  depends_on = [restful_rest_call.POST1]
}
resource "restful_rest_call" "GET3" {
  method = "GET"
  uri = "http://localhost:8080/events/2"
  expected_response_code = 200
  expected_response_body = "{\"ID\":\"2\",\"Title\":\"title\",\"Description\":\"desc3\"}"
  depends_on = [restful_rest_call.PATCH1]
}



