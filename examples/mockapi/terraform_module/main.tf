module "get_request" {
  source = "./modules/restful"

  method = "GET"
  uri = "http://localhost:8080/events/1"
  expected_response_code = 200
  expected_response_body = "{\"ID\":\"1\",\"Title\":\"The title\",\"Description\":\"A short description\"}"
  json_key_outputs = ["Title", "Description"]
}

module "post_request" {
  source = "./modules/restful"

  method = "POST"
  uri = "http://localhost:8080/event"
  headers = ["Content-Type:application/json"]
  request_body = "{\"ID\":\"2\",\"Title\":\"title\",\"Description\":\"A description for title 2\"}"
  expected_response_body = "{\"ID\":\"2\",\"Title\":\"title\",\"Description\":\"A description for title 2\"}"
  expected_response_code = 201
}
