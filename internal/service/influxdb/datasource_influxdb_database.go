package influxdb

import (
	"context"

	"github.com/aiven/aiven-go-client"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceInfluxDBDatabase() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceDatabaseRead,
		Description: "The InfluxDB Database data source provides information about the existing Aiven InfluxDB Database.",
		Schema: schemautil.ResourceSchemaAsDatasourceSchema(aivenInfluxDBDatabaseSchema,
			"project", "service_name", "database_name"),
	}
}

func datasourceDatabaseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*aiven.Client)

	projectName := d.Get("project").(string)
	serviceName := d.Get("service_name").(string)
	databaseName := d.Get("database_name").(string)

	databases, err := client.Databases.List(projectName, serviceName)
	if err != nil {
		return diag.FromErr(err)
	}

	for _, db := range databases {
		if db.DatabaseName == databaseName {
			d.SetId(schemautil.BuildResourceID(projectName, serviceName, databaseName))
			return resourceInfluxDBDatabaseRead(ctx, d, m)
		}
	}

	return diag.Errorf("database %s/%s/%s not found",
		projectName, serviceName, databaseName)
}
