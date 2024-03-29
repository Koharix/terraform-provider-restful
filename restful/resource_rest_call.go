package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceRestCall() *schema.Resource {
	return &schema.Resource{
		Create: resourceRestCallCreate,
		Read:   resourceRestCallRead,
		Update: resourceRestCallUpdate,
		Delete: resourceRestCallDelete,

		Schema: map[string]*schema.Schema{
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"uri": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"request_body": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"headers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
				  Type: schema.TypeString,
				},
			},
			"expected_response_body": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"expected_response_code": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"json_key_outputs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
				  Type: schema.TypeString,
				},
			},
			"outputs": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
				  Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceRestCallCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("function resourceRestCallCreate() started")

	uri := d.Get("uri").(string)
	d.SetId(uri) //Can be changed, filler value

	resBody, resCode, err := restCall(d.Get("method").(string), uri, d.Get("headers").([]interface{}), []byte(d.Get("request_body").(string)))
	if err != nil {
		return err
	}

	err = resChecker(resBody, []byte(d.Get("expected_response_body").(string)), resCode, d.Get("expected_response_code").(int))
	if err != nil {
		return err
	}
	
	if d.Get("json_key_outputs") != nil {
		err = setOutputs(d, resBody, (d.Get("json_key_outputs").([]interface{})))
	}

	return err
}

func resourceRestCallRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceRestCallUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceRestCallCreate(d, m)
}

func resourceRestCallDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}