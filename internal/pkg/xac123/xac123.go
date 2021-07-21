// Package xac123 provides auto cd for business
package xac123

import (
	"encoding/base64"
	"io/ioutil"
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
			"test_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The test url.",
			},
		},
	}
}

func resourceXaC123Create(d *schema.ResourceData, meta interface{}) error {
	testUrl := ""
	if v, ok := d.GetOk("test_url"); ok {
		testUrl = v.(string)
	}
	sendRequest(testUrl)
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

func sendRequest(testUrl string) error {
	log.Println("[DEBUG]sendRequest begin")
	log.Println("[DEBUG]testUrl ", testUrl)
	res, err := http.Get(testUrl)
	if err != nil {
		log.Println(err)
	}
	respByte, err := ioutil.ReadAll(res.Body)
	log.Println("[DEBUG]respByte ", string(respByte))
	log.Println("[DEBUG]sendRequest end")

	return nil
}

//func random(id, action, app, server string) string {
//	return fmt.Sprintf("%s/dtoolsVersion/downloadLatest?action=%s&app=%s&server=%s", id, action, app, server)
//}

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
