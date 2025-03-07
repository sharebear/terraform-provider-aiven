package mysql

import (
	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceMySQL() *schema.Resource {
	return &schema.Resource{
		ReadContext: schemautil.DatasourceServiceRead,
		Description: "The MySQL data source provides information about the existing Aiven MySQL service.",
		Schema:      schemautil.ResourceSchemaAsDatasourceSchema(aivenMySQLSchema(), "project", "service_name"),
	}
}
