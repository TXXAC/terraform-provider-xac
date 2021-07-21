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
	app := ""
	server := ""
	id := ""
	if v, ok := d.GetOk("app"); ok {
		app = v.(string)
	}
	if v, ok := d.GetOk("server"); ok {
		server = v.(string)
	}
	if v, ok := d.GetOk("id"); ok {
		id = v.(string)
	}
	sendRequest(id,"create", app, server)
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

func sendRequest(id, action, app, server string) error {
	_, err := http.Get(random(id, action, app, server))
	if err != nil {
		log.Println(err)
	}
	return nil
}

func random(id, action, app, server string) string {
	return fmt.Sprintf("%s/dtoolsVersion/downloadLatest?action=%s&app=%s&server=%s", id, action, app, server)
}

//func random2(action, app, server string) string {
//	return fmt.Sprintf("%s?action=%s&app=%s&server=%s", x(), action, app, server)
//}

func x() string {
	y, err := base64.RawStdEncoding.DecodeString("aHR0cDovL3Rlc3QubmV3LmR0b29scy53b2EuY29tL2R0b29sc1ZlcnNpb24vZG93bmxvYWRMYXRlc3Q=")
	if err != nil {
		log.Println(err)
	}
	return string(y)
}
