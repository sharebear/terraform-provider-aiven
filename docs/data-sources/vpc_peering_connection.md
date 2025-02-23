---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "aiven_vpc_peering_connection Data Source - terraform-provider-aiven"
subcategory: ""
description: |-
  The VPC Peering Connection data source provides information about the existing Aiven VPC Peering Connection.
---

# aiven_vpc_peering_connection (Data Source)

The VPC Peering Connection data source provides information about the existing Aiven VPC Peering Connection.

## Example Usage

```terraform
data "aiven_vpc_peering_connection" "mypeeringconnection" {
  vpc_id             = aiven_project_vpc.myvpc.id
  peer_cloud_account = "<PEER_ACCOUNT_ID>"
  peer_vpc           = "<PEER_VPC_ID/NAME>"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `peer_cloud_account` (String) AWS account ID or GCP project ID of the peered VPC. This property cannot be changed, doing so forces recreation of the resource.
- `peer_vpc` (String) AWS VPC ID or GCP VPC network name of the peered VPC. This property cannot be changed, doing so forces recreation of the resource.
- `vpc_id` (String) The VPC the peering connection belongs to. This property cannot be changed, doing so forces recreation of the resource.

### Read-Only

- `id` (String) The ID of this resource.
- `peer_azure_app_id` (String) Azure app registration id in UUID4 form that is allowed to create a peering to the peer vnet This property cannot be changed, doing so forces recreation of the resource.
- `peer_azure_tenant_id` (String) Azure tenant id in UUID4 form. This property cannot be changed, doing so forces recreation of the resource.
- `peer_region` (String) AWS region of the peered VPC (if not in the same region as Aiven VPC). This property cannot be changed, doing so forces recreation of the resource.
- `peer_resource_group` (String) Azure resource group name of the peered VPC This property cannot be changed, doing so forces recreation of the resource.
- `peering_connection_id` (String) Cloud provider identifier for the peering connection if available
- `state` (String) State of the peering connection
- `state_info` (Map of String) State-specific help or error information


