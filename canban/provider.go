package canban

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"canban_project": resourceProject(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}
