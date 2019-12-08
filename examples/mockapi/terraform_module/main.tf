module "get_request" {
  source = "./modules/restful"

  method = "GET"
  uri = "http://localhost:8080/events"
  expected_response_code = 200
  expected_response_body = "[{\"ID\":\"1\",\"Title\":\"Introduction to Golang\",\"Description\":\"Come join us for a chance to learn how golang works and get to eventually try it out\"}]"

}

module "post_request" {
  source = "./modules/restful"

  method = "POST"
  uri = "http://localhost:8080/event"
  headers = ["Content-Type:application/json"]
  request_body = "{\"ID\":\"2\",\"Title\":\"title\",\"Description\":\"desc2\"}"
  expected_response_body = "{\"ID\":\"2\",\"Title\":\"title\",\"Description\":\"desc2\"}"
  expected_response_code = 201
}
