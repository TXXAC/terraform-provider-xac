// Package xac provides auto and legal ci/cd for business
package xac

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"TXXAC/terraform-provider-xac/internal/pkg/xac007"
	"TXXAC/terraform-provider-xac/internal/pkg/xac123"
	"TXXAC/terraform-provider-xac/internal/pkg/xac_paas"
	"TXXAC/terraform-provider-xac/internal/pkg/xac_store"
)

const (
	ProviderOperator = "XAC_OPERATOR"
)

// Provider provides auto and legal ci/cd for business
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"operator": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(ProviderOperator, nil),
				Description: "The operator, please provides your English name, which named rtx ago",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"xac_123_images":       &schema.Resource{},
			"xac_obs_ops_products": &schema.Resource{},
			"xac_cmdb_modules":     &schema.Resource{},
		},

		ResourcesMap: map[string]*schema.Resource{
			"xac_123":          xac123.ResourceXaC123(),
			"xac_007":          xac007.ResourceXaC007(),
			"xac_store_mdb":    xac_store.ResourceXaCStoreMDB(),
			"xac_store_bdb":    xac_store.ResourceXaCStoreMDB(),
			"xac_store_dcache": xac_store.ResourceXaCStoreMDB(),
			"xac_store_redis":  xac_store.ResourceXaCStoreMDB(),
			"xac_paas_cos":     xac_paas.ResourceXaCPaaSCOS(),
			"xac_paas_cvm":     xac_paas.ResourceXaCPaaSCOS(),
			"xac_paas_es":      xac_paas.ResourceXaCPaaSCOS(),
			"xac_paas_ckafka":  xac_paas.ResourceXaCPaaSCOS(),
		},
	}
}
