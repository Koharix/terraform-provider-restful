# Terraform provider for any REST api

The Restful Provider supports Terraform 0.12.x and later.

* [Terraform Website](https://www.terraform.io)
* [Restful Provider Usage Examples](https://github.com/Koharix/terraform-provider-restful/tree/master/examples)

## Usage Example
```
# No configuration is requierd for the proivder
provider "restful" {}

# GET request that returns object information
resource "restful_rest_call" "get_object" {
  method = "GET"
  uri = "http://localhost:8080/events/1"
  expected_response_code = 200
  expected_response_body = "{\"ID\":\"1\",\"Title\":\"Introduction to Golang\",\"Description\":\"Come join us for a chance to learn how golang works and get to eventually try it out\"}"
  json_key_outputs = ["Title", "Description"]
}

# POST request to create a new object
resource "restful_rest_call" "post_object" {
  method = "POST"
  uri = "http://localhost:8080/event"
  headers = ["Content-Type:application/json"]
  request_body = "{\"ID\":\"2\",\"Title\":\"title\",\"Description\":\"desc2\"}"
  expected_response_body = "{\"ID\":\"2\",\"Title\":\"title\",\"Description\":\"desc2\"}"
  expected_response_code = 201
}

# PATCH request to update an object
resource "restful_rest_call" "patch_object" {
  method = "PATCH"
  uri = "http://localhost:8080/events/2"
  headers = ["Content-Type:application/json"]
  request_body = "{\"ID\":\"2\",\"Title\":\"title\",\"Description\":\"desc3\"}"
  expected_response_code = 200
  depends_on = [restful_rest_call.post_object]
}

# DELETE request to delete an object
resource "restful_rest_call" "delete_object" {
  method = "DELETE"
  uri = "http://localhost:8080/events/2"
  expected_response_code = 200
  depends_on = [restful_rest_call.patch_object]
}
```

## Provider configuration
No configuration is used by the provider block.

## `rest_call` resource configuration
- `id` (string, computed): The id is set equal to the inputed uri.
- `method` (string, required): The HTTP method being applied to the call.
- `uri` (string, required): The endpoint to hit by the REST call.
- `headers` (list(string), optional): List of headers applied to the call.
- `request_body` (string, optional): The body of the call, required to be valid json.
- `expected_response_body` (string, optional): Validates if the inputed json here is "in" the json of the response body
- `expected_response_code` (number, optional): Validates that the if the response code is equal to this value.
- `json_key_outputs` (list(string), optional): List of json keys that you want as an output variable by the resource
- `outputs` (list(string),  computed): List of the values to the keys in the response body specified in the json_key_outputs input.

## Installation
There are two standard methods of installing this provider detailed [in Terraform's documentation](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins). You can place the file in the directory of your .tf file in `terraform.d/plugins/{OS}_{ARCH}/` or place it in your home directory at `~/.terraform.d/plugins/{OS}_{ARCH}/`

The released binaries are named `terraform-provider-restapi_vX.Y.Z-{OS}-{ARCH}` so you know which binary to install. You *may* need to rename the binary you use during installation to just `terraform-provider-restapi_vX.Y.Z`.

Once downloaded, be sure to make the plugin executable by running `chmod +x terraform-provider-restapi_vX.Y.Z-{OS}-{ARCH}`.

&nbsp;

## Contributing
Pull requests are always welcome! Please be sure the following things are taken care of with your pull request:
* `go fmt` is run before pushing
* Be sure to add a test case for new functionality (or explain why this cannot be done)
* Run the `scripts/test.sh` script to be sure everything works
* Ensure new attributes can also be set by environment variables

#### Development environment requirements
* [Golang](https://golang.org/dl/) is installed and `go` is in your path
* [Terraform](https://www.terraform.io/downloads.html) is installed and `terraform` is in your path 