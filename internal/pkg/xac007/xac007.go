// Package xac007 ...
package xac007

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

// ResourceXaC007 resource xac007
func ResourceXaC007() *schema.Resource {
	return &schema.Resource{
		Create: resourceXaC007Create,
		Read:   resourceXaC007Read,
		Update: resourceXaC007Update,
		Delete: resourceXaC007Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"app": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The application name in 123.",
			},
			"server": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The server name in 123.",
			},
			"monitor": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The monitor name.",
			},
			"dimensions": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The dimension list.",
				MaxItems:    100,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Dimension name.",
						},
					},
				},
			},
			"values": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The feature list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy": {
							Type:        schema.TypeInt,
							Required:    true,
							ForceNew:    true,
							Description: "Policy like sum/count/avg/max/min/set.",
						},
					},
				},
			},
		},
	}
}

func resourceXaC007Create(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceXaC007Read(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceXaC007Update(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceXaC007Delete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
