// resource_server.go
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

//data source: must not implement Create, Update or Delete
func dataServer() *schema.Resource {
	return &schema.Resource{
		Read: resourceServerRead,
		Schema: map[string]*schema.Schema{
			"outputv": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	outputstr := "this is a sample data"

	d.SetId(outputstr)
	d.Set("outputv", outputstr)
	return nil
}
