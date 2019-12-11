variable "method" {
    description = "The method of the REST call, ie GET, POST, PUT, DELETE, etc..."
    type = string
}

variable "uri" {
    description = "Endpoint to hit."
    type = string
}

variable "headers" {
    description = "Headers to pass in."
    type = list(string)
    default = null
}

variable "request_body" {
    description = "The json body in the request."
    type = string
    default = null
}

variable "expected_response_code" {
    description = "The expected status code in the response."
    type = number
    default = null
}

variable "expected_response_body" {
    description = "The part or all of the expected json response."
    type = string
    default = null
}

variable "json_key_outputs" {
    description = "List of json keys that you want the value of it stored as an output."
    type = list(string)
    default = null
}


