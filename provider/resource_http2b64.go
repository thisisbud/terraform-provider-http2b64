package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/ackers-bud/terraform-provider-http2b64/client"
)

func resourcehttp2b64() *schema.Resource {
	return &schema.Resource{
		Create: Create,
		Update: Update,
		Read:   ReadUrl,
		Delete: Delete,

		Schema: map[string]*schema.Schema{
			"url": {
				Type:     schema.TypeString,
				Required: true,
			},

			"id": {
				Description: "The ID of this resource.",
				Type:        schema.TypeString,
				Computed:    true,
			},

			"status_code": {
				Description: "the returned http status code",
				Type:        schema.TypeInt,
				Computed:    true,
			},

			"response_body_base64": {
				Description: "The Returned body base64 encoded",
				Type:        schema.TypeString,
				Computed:    true,
			},
		},
	}
}

//nolint:ineffassign
func Create(d *schema.ResourceData, meta interface{}) error {

	url := d.Get("url").(string)

	responseBody, statusCode, err := client.GetFile(url)
	if err != nil {
		return fmt.Errorf("error Getting resource '%v'", err)
	}

	err = d.Set("status_code", statusCode)
	if err != nil {
		return fmt.Errorf("error Setting status_code '%v'", err)
	}

	err = d.Set("response_body_base64", responseBody)
	if err != nil {
		return fmt.Errorf("error Setting response_body_base64 '%v'", err)
	}

	d.SetId(url)
	return nil
}

func Update(d *schema.ResourceData, meta interface{}) error {

	url := d.Get("url").(string)

	responseBody, statusCode, err := client.GetFile(url)
	if err != nil {
		return fmt.Errorf("error Getting resource '%v'", err)
	}

	err = d.Set("status_code", statusCode)
	if err != nil {
		return fmt.Errorf("error Setting status_code '%v'", err)
	}

	err = d.Set("response_body_base64", responseBody)
	if err != nil {
		return fmt.Errorf("error Setting response_body_base64 '%v'", err)
	}

	d.SetId(url)

	return nil
}

func ReadUrl(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func Delete(d *schema.ResourceData, meta interface{}) error {

	d.SetId("")

	return nil
}
