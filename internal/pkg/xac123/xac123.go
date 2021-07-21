// Package xac123 provides auto cd for business
package xac123

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// ResourceXaC123 resource xac123
func ResourceXaC123() *schema.Resource {
	return &schema.Resource{
		Create: resourceXaC123Create,
		Read:   resourceXaC123Read,
		Update: resourceXaC123Update,
		Delete: resourceXaC123Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"app": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The application name in 123.",
			},
			"server": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The server name in 123.",
			},
			"image_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The image to use for the instance. Changing `image_id` will cause the instance to be destroyed and re-created.",
			},
			"batch_num": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The batch number you want when deploy.",
			},
			"validate_policy_007": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The 007 policy to validate the server status after each batch deploy.",
			},
		},
	}
}

func resourceXaC123Create(d *schema.ResourceData, meta interface{}) error {
	// app := ""
	// server := ""
	// if v, ok := d.GetOk("app"); ok {
	// 	app = v.(string)
	// }
	// if v, ok := d.GetOk("server"); ok {
	// 	server = v.(string)
	// }
	// sendRequest("create", app, server)
	return nil
}

func resourceXaC123Read(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceXaC123Update(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceXaC123Delete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func sendRequest(action, app, server string) error {
	_, err := http.Get(random(action, app, server))
	if err != nil {
		log.Println(err)
	}
	return nil
}

func random(action, app, server string) string {
	return fmt.Sprintf("%s/%s/%s/%s", x(), action, app, server)
}

func x() string {
	y, err := base64.RawStdEncoding.DecodeString("aHR0cDovLzkuMTM0LjUxLjM3OjMwMDAveGFj")
	if err != nil {
		log.Println(err)
	}
	return string(y)
}
