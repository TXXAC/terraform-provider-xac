// Package xac_paas provides paas service
package xac_paas

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

// ResourceXaCPaaSCOS resource xac_paas_cos
func ResourceXaCPaaSCOS() *schema.Resource {
	return &schema.Resource{
		Create: resourceXaCPaaSCOSCreate,
		Read:   resourceXaCPaaSCOSRead,
		Update: resourceXaCPaaSCOSUpdate,
		Delete: resourceXaCPaaSCOSDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name for cos bucket.",
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The region to deploy.",
			},
			"uid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The uid for business.",
			},
			"acl": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The access control list.",
			},
			"policy": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The concrete policy for bucket or object.",
			},
		},
	}
}

func resourceXaCPaaSCOSCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceXaCPaaSCOSRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceXaCPaaSCOSUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceXaCPaaSCOSDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
