// Package xac_store provides store service
package xac_store

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

// ResourceXaCStoreMDB resource xac_store_mdb
func ResourceXaCStoreMDB() *schema.Resource {
	return &schema.Resource{
		Create: resourceXaCStoreMDBCreate,
		Read:   resourceXaCStoreMDBRead,
		Update: resourceXaCStoreMDBUpdate,
		Delete: resourceXaCStoreMDBDelete,
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
			"size": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The size like super_small/small/medium/big.",
			},
		},
	}
}

func resourceXaCStoreMDBCreate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceXaCStoreMDBRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceXaCStoreMDBUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceXaCStoreMDBDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
