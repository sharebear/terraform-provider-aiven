---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "aiven_m3db Resource - terraform-provider-aiven"
subcategory: ""
description: |-
  The M3 DB resource allows the creation and management of Aiven M3 services.
---

# aiven_m3db (Resource)

The M3 DB resource allows the creation and management of Aiven M3 services.

## Example Usage

```terraform
resource "aiven_m3db" "m3" {
  project                 = data.aiven_project.foo.project
  cloud_name              = "google-europe-west1"
  plan                    = "business-8"
  service_name            = "my-m3db"
  maintenance_window_dow  = "monday"
  maintenance_window_time = "10:00:00"

  m3db_user_config {
    m3db_version = 1.1

    namespaces {
      name = "my_ns1"
      type = "unaggregated"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `project` (String) Identifies the project this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
- `service_name` (String) Specifies the actual name of the service. The name cannot be changed later without destroying and re-creating the service so name should be picked based on intended service usage rather than current attributes.

### Optional

- `additional_disk_space` (String) Additional disk space. Possible values depend on the service type, the cloud provider and the project. Therefore, reducing will result in the service rebalancing.
- `cloud_name` (String) Defines where the cloud provider and region where the service is hosted in. This can be changed freely after service is created. Changing the value will trigger a potentially lengthy migration process for the service. Format is cloud provider name (`aws`, `azure`, `do` `google`, `upcloud`, etc.), dash, and the cloud provider specific region name. These are documented on each Cloud provider's own support articles, like [here for Google](https://cloud.google.com/compute/docs/regions-zones/) and [here for AWS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html).
- `disk_space` (String) Service disk space. Possible values depend on the service type, the cloud provider and the project. Therefore, reducing will result in the service rebalancing.
- `m3db_user_config` (Block List, Max: 1) M3db user configurable settings (see [below for nested schema](#nestedblock--m3db_user_config))
- `maintenance_window_dow` (String) Day of week when maintenance operations should be performed. One monday, tuesday, wednesday, etc.
- `maintenance_window_time` (String) Time of day when maintenance operations should be performed. UTC time in HH:mm:ss format.
- `plan` (String) Defines what kind of computing resources are allocated for the service. It can be changed after creation, though there are some restrictions when going to a smaller plan such as the new plan must have sufficient amount of disk space to store all current data and switching to a plan with fewer nodes might not be supported. The basic plan names are `hobbyist`, `startup-x`, `business-x` and `premium-x` where `x` is (roughly) the amount of memory on each node (also other attributes like number of CPUs and amount of disk space varies but naming is based on memory). The available options can be seem from the [Aiven pricing page](https://aiven.io/pricing).
- `project_vpc_id` (String) Specifies the VPC the service should run in. If the value is not set the service is not run inside a VPC. When set, the value should be given as a reference to set up dependencies correctly and the VPC must be in the same cloud and region as the service itself. Project can be freely moved to and from VPC after creation but doing so triggers migration to new servers so the operation can take significant amount of time to complete if the service has a lot of data.
- `service_integrations` (Block List) Service integrations to specify when creating a service. Not applied after initial service creation (see [below for nested schema](#nestedblock--service_integrations))
- `static_ips` (Set of String) Static IPs that are going to be associated with this service. Please assign a value using the 'toset' function. Once a static ip resource is in the 'assigned' state it cannot be unbound from the node again
- `tag` (Block Set) Tags are key-value pairs that allow you to categorize services. (see [below for nested schema](#nestedblock--tag))
- `termination_protection` (Boolean) Prevents the service from being deleted. It is recommended to set this to `true` for all production services to prevent unintentional service deletion. This does not shield against deleting databases or topics but for services with backups much of the content can at least be restored from backup in case accidental deletion is done.
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `components` (List of Object) Service component information objects (see [below for nested schema](#nestedatt--components))
- `disk_space_cap` (String) The maximum disk space of the service, possible values depend on the service type, the cloud provider and the project.
- `disk_space_default` (String) The default disk space of the service, possible values depend on the service type, the cloud provider and the project. Its also the minimum value for `disk_space`
- `disk_space_step` (String) The default disk space step of the service, possible values depend on the service type, the cloud provider and the project. `disk_space` needs to increment from `disk_space_default` by increments of this size.
- `disk_space_used` (String) Disk space that service is currently using
- `id` (String) The ID of this resource.
- `m3db` (List of Object) M3 specific server provided values (see [below for nested schema](#nestedatt--m3db))
- `service_host` (String) The hostname of the service.
- `service_password` (String, Sensitive) Password used for connecting to the service, if applicable
- `service_port` (Number) The port of the service
- `service_type` (String) Aiven internal service type code
- `service_uri` (String, Sensitive) URI for connecting to the service. Service specific info is under "kafka", "pg", etc.
- `service_username` (String) Username used for connecting to the service, if applicable
- `state` (String) Service state. One of `POWEROFF`, `REBALANCING`, `REBUILDING` or `RUNNING`

<a id="nestedblock--m3db_user_config"></a>
### Nested Schema for `m3db_user_config`

Optional:

- `custom_domain` (String) Custom domain
- `ip_filter` (List of String) IP filter
- `limits` (Block List, Max: 1) M3 limits (see [below for nested schema](#nestedblock--m3db_user_config--limits))
- `m3_version` (String) M3 major version (deprecated, use m3db_version)
- `m3coordinator_enable_graphite_carbon_ingest` (String) Enable Graphite ingestion using Carbon plaintext protocol
- `m3db_version` (String) M3 major version (the minimum compatible version)
- `namespaces` (Block List, Max: 2147483647) List of M3 namespaces (see [below for nested schema](#nestedblock--m3db_user_config--namespaces))
- `private_access` (Block List, Max: 1) Allow access to selected service ports from private networks (see [below for nested schema](#nestedblock--m3db_user_config--private_access))
- `project_to_fork_from` (String) Name of another project to fork a service from. This has effect only when a new service is being created.
- `public_access` (Block List, Max: 1) Allow access to selected service ports from the public Internet (see [below for nested schema](#nestedblock--m3db_user_config--public_access))
- `rules` (Block List, Max: 1) M3 rules (see [below for nested schema](#nestedblock--m3db_user_config--rules))
- `service_to_fork_from` (String) Name of another service to fork from. This has effect only when a new service is being created.
- `static_ips` (String) Static IP addresses

<a id="nestedblock--m3db_user_config--limits"></a>
### Nested Schema for `m3db_user_config.limits`

Optional:

- `max_recently_queried_series_blocks` (String) The maximum number of blocks that can be read in a given lookback period.
- `max_recently_queried_series_disk_bytes_read` (String) The maximum number of disk bytes that can be read in a given lookback period.
- `max_recently_queried_series_lookback` (String) The lookback period for 'max_recently_queried_series_blocks' and 'max_recently_queried_series_disk_bytes_read'.
- `query_docs` (String) The maximum number of docs fetched in single query.
- `query_require_exhaustive` (String) Require exhaustive result
- `query_series` (String) The maximum number of series fetched in single query.


<a id="nestedblock--m3db_user_config--namespaces"></a>
### Nested Schema for `m3db_user_config.namespaces`

Optional:

- `name` (String) The name of the namespace
- `options` (Block List, Max: 1) Namespace options (see [below for nested schema](#nestedblock--m3db_user_config--namespaces--options))
- `resolution` (String) The resolution for an aggregated namespace
- `type` (String) The type of aggregation (aggregated/unaggregated)

<a id="nestedblock--m3db_user_config--namespaces--options"></a>
### Nested Schema for `m3db_user_config.namespaces.options`

Optional:

- `retention_options` (Block List, Max: 1) Retention options (see [below for nested schema](#nestedblock--m3db_user_config--namespaces--options--retention_options))
- `snapshot_enabled` (String) Controls whether M3DB will create snapshot files for this namespace
- `writes_to_commitlog` (String) Controls whether M3DB will include writes to this namespace in the commitlog

<a id="nestedblock--m3db_user_config--namespaces--options--retention_options"></a>
### Nested Schema for `m3db_user_config.namespaces.options.retention_options`

Optional:

- `block_data_expiry_duration` (String) Controls how long we wait before expiring stale data
- `blocksize_duration` (String) Controls how long to keep a block in memory before flushing to a fileset on disk
- `buffer_future_duration` (String) Controls how far into the future writes to the namespace will be accepted
- `buffer_past_duration` (String) Controls how far into the past writes to the namespace will be accepted
- `retention_period_duration` (String) Controls the duration of time that M3DB will retain data for the namespace




<a id="nestedblock--m3db_user_config--private_access"></a>
### Nested Schema for `m3db_user_config.private_access`

Optional:

- `m3coordinator` (String) Allow clients to connect to m3coordinator with a DNS name that always resolves to the service's private IP addresses. Only available in certain network locations


<a id="nestedblock--m3db_user_config--public_access"></a>
### Nested Schema for `m3db_user_config.public_access`

Optional:

- `m3coordinator` (String) Allow clients to connect to m3coordinator from the public internet for service nodes that are in a project VPC or another type of private network


<a id="nestedblock--m3db_user_config--rules"></a>
### Nested Schema for `m3db_user_config.rules`

Optional:

- `mapping` (Block List, Max: 10) List of M3 mapping rules (see [below for nested schema](#nestedblock--m3db_user_config--rules--mapping))

<a id="nestedblock--m3db_user_config--rules--mapping"></a>
### Nested Schema for `m3db_user_config.rules.mapping`

Optional:

- `aggregations` (List of String) List of aggregations to be applied
- `drop` (String) Drop the matching metric
- `filter` (String) The metrics to be used with this particular rule
- `name` (String) The (optional) name of the rule
- `namespaces` (List of String) Namespace filters for this particular rule
- `tags` (Block List, Max: 10) List of tags to be appended to matching metrics (see [below for nested schema](#nestedblock--m3db_user_config--rules--mapping--tags))

<a id="nestedblock--m3db_user_config--rules--mapping--tags"></a>
### Nested Schema for `m3db_user_config.rules.mapping.tags`

Optional:

- `name` (String) Name of the tag
- `value` (String) Value of the tag





<a id="nestedblock--service_integrations"></a>
### Nested Schema for `service_integrations`

Required:

- `integration_type` (String) Type of the service integration. The only supported value at the moment is `read_replica`
- `source_service_name` (String) Name of the source service


<a id="nestedblock--tag"></a>
### Nested Schema for `tag`

Required:

- `key` (String) Service tag key
- `value` (String) Service tag value


<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `update` (String)


<a id="nestedatt--components"></a>
### Nested Schema for `components`

Read-Only:

- `component` (String)
- `host` (String)
- `kafka_authentication_method` (String)
- `port` (Number)
- `route` (String)
- `ssl` (Boolean)
- `usage` (String)


<a id="nestedatt--m3db"></a>
### Nested Schema for `m3db`

Read-Only:

## Import

Import is supported using the following syntax:

```shell
terraform import aiven_m3db.m3 project/service_name
```
