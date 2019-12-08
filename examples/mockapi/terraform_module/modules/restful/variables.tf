variable "method" {
    type = string
}

variable "uri" {
    type = string
}

variable "headers" {
    type = list(string)
    default = null
}

variable "request_body" {
    type = string
    default = null
}

variable "expected_response_code" {
    type = number
    default = null
}

variable "expected_response_body" {
    type = string
    default = null
}


