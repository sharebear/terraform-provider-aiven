---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "aiven_kafka Data Source - terraform-provider-aiven"
subcategory: ""
description: |-
  The Kafka data source provides information about the existing Aiven Kafka services.
---

# aiven_kafka (Data Source)

The Kafka data source provides information about the existing Aiven Kafka services.

## Example Usage

```terraform
data "aiven_kafka" "kafka1" {
  project      = data.aiven_project.pr1.project
  service_name = "my-kafka1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `project` (String) Identifies the project this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
- `service_name` (String) Specifies the actual name of the service. The name cannot be changed later without destroying and re-creating the service so name should be picked based on intended service usage rather than current attributes.

### Read-Only

- `additional_disk_space` (String) Additional disk space. Possible values depend on the service type, the cloud provider and the project. Therefore, reducing will result in the service rebalancing.
- `cloud_name` (String) Defines where the cloud provider and region where the service is hosted in. This can be changed freely after service is created. Changing the value will trigger a potentially lengthy migration process for the service. Format is cloud provider name (`aws`, `azure`, `do` `google`, `upcloud`, etc.), dash, and the cloud provider specific region name. These are documented on each Cloud provider's own support articles, like [here for Google](https://cloud.google.com/compute/docs/regions-zones/) and [here for AWS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html).
- `components` (List of Object) Service component information objects (see [below for nested schema](#nestedatt--components))
- `default_acl` (Boolean) Create default wildcard Kafka ACL
- `disk_space` (String) Service disk space. Possible values depend on the service type, the cloud provider and the project. Therefore, reducing will result in the service rebalancing.
- `disk_space_cap` (String) The maximum disk space of the service, possible values depend on the service type, the cloud provider and the project.
- `disk_space_default` (String) The default disk space of the service, possible values depend on the service type, the cloud provider and the project. Its also the minimum value for `disk_space`
- `disk_space_step` (String) The default disk space step of the service, possible values depend on the service type, the cloud provider and the project. `disk_space` needs to increment from `disk_space_default` by increments of this size.
- `disk_space_used` (String) Disk space that service is currently using
- `id` (String) The ID of this resource.
- `kafka` (List of Object) Kafka server provided values (see [below for nested schema](#nestedatt--kafka))
- `kafka_user_config` (List of Object) Kafka user configurable settings (see [below for nested schema](#nestedatt--kafka_user_config))
- `karapace` (Boolean) Switch the service to use Karapace for schema registry and REST proxy
- `maintenance_window_dow` (String) Day of week when maintenance operations should be performed. One monday, tuesday, wednesday, etc.
- `maintenance_window_time` (String) Time of day when maintenance operations should be performed. UTC time in HH:mm:ss format.
- `plan` (String) Defines what kind of computing resources are allocated for the service. It can be changed after creation, though there are some restrictions when going to a smaller plan such as the new plan must have sufficient amount of disk space to store all current data and switching to a plan with fewer nodes might not be supported. The basic plan names are `hobbyist`, `startup-x`, `business-x` and `premium-x` where `x` is (roughly) the amount of memory on each node (also other attributes like number of CPUs and amount of disk space varies but naming is based on memory). The available options can be seem from the [Aiven pricing page](https://aiven.io/pricing).
- `project_vpc_id` (String) Specifies the VPC the service should run in. If the value is not set the service is not run inside a VPC. When set, the value should be given as a reference to set up dependencies correctly and the VPC must be in the same cloud and region as the service itself. Project can be freely moved to and from VPC after creation but doing so triggers migration to new servers so the operation can take significant amount of time to complete if the service has a lot of data.
- `service_host` (String) The hostname of the service.
- `service_integrations` (List of Object) Service integrations to specify when creating a service. Not applied after initial service creation (see [below for nested schema](#nestedatt--service_integrations))
- `service_password` (String, Sensitive) Password used for connecting to the service, if applicable
- `service_port` (Number) The port of the service
- `service_type` (String) Aiven internal service type code
- `service_uri` (String, Sensitive) URI for connecting to the service. Service specific info is under "kafka", "pg", etc.
- `service_username` (String) Username used for connecting to the service, if applicable
- `state` (String) Service state. One of `POWEROFF`, `REBALANCING`, `REBUILDING` or `RUNNING`
- `static_ips` (Set of String) Static IPs that are going to be associated with this service. Please assign a value using the 'toset' function. Once a static ip resource is in the 'assigned' state it cannot be unbound from the node again
- `tag` (Set of Object) Tags are key-value pairs that allow you to categorize services. (see [below for nested schema](#nestedatt--tag))
- `termination_protection` (Boolean) Prevents the service from being deleted. It is recommended to set this to `true` for all production services to prevent unintentional service deletion. This does not shield against deleting databases or topics but for services with backups much of the content can at least be restored from backup in case accidental deletion is done.

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


<a id="nestedatt--kafka"></a>
### Nested Schema for `kafka`

Read-Only:

- `access_cert` (String)
- `access_key` (String)
- `connect_uri` (String)
- `rest_uri` (String)
- `schema_registry_uri` (String)


<a id="nestedatt--kafka_user_config"></a>
### Nested Schema for `kafka_user_config`

Read-Only:

- `custom_domain` (String)
- `ip_filter` (List of String)
- `kafka` (List of Object) (see [below for nested schema](#nestedobjatt--kafka_user_config--kafka))
- `kafka_authentication_methods` (List of Object) (see [below for nested schema](#nestedobjatt--kafka_user_config--kafka_authentication_methods))
- `kafka_connect` (String)
- `kafka_connect_config` (List of Object) (see [below for nested schema](#nestedobjatt--kafka_user_config--kafka_connect_config))
- `kafka_rest` (String)
- `kafka_rest_config` (List of Object) (see [below for nested schema](#nestedobjatt--kafka_user_config--kafka_rest_config))
- `kafka_version` (String)
- `private_access` (List of Object) (see [below for nested schema](#nestedobjatt--kafka_user_config--private_access))
- `privatelink_access` (List of Object) (see [below for nested schema](#nestedobjatt--kafka_user_config--privatelink_access))
- `public_access` (List of Object) (see [below for nested schema](#nestedobjatt--kafka_user_config--public_access))
- `schema_registry` (String)
- `schema_registry_config` (List of Object) (see [below for nested schema](#nestedobjatt--kafka_user_config--schema_registry_config))
- `static_ips` (String)

<a id="nestedobjatt--kafka_user_config--kafka"></a>
### Nested Schema for `kafka_user_config.kafka`

Read-Only:

- `auto_create_topics_enable` (String)
- `compression_type` (String)
- `connections_max_idle_ms` (String)
- `default_replication_factor` (String)
- `group_initial_rebalance_delay_ms` (String)
- `group_max_session_timeout_ms` (String)
- `group_min_session_timeout_ms` (String)
- `log_cleaner_delete_retention_ms` (String)
- `log_cleaner_max_compaction_lag_ms` (String)
- `log_cleaner_min_cleanable_ratio` (String)
- `log_cleaner_min_compaction_lag_ms` (String)
- `log_cleanup_policy` (String)
- `log_flush_interval_messages` (String)
- `log_flush_interval_ms` (String)
- `log_index_interval_bytes` (String)
- `log_index_size_max_bytes` (String)
- `log_message_downconversion_enable` (String)
- `log_message_timestamp_difference_max_ms` (String)
- `log_message_timestamp_type` (String)
- `log_preallocate` (String)
- `log_retention_bytes` (String)
- `log_retention_hours` (String)
- `log_retention_ms` (String)
- `log_roll_jitter_ms` (String)
- `log_roll_ms` (String)
- `log_segment_bytes` (String)
- `log_segment_delete_delay_ms` (String)
- `max_connections_per_ip` (String)
- `max_incremental_fetch_session_cache_slots` (String)
- `message_max_bytes` (String)
- `min_insync_replicas` (String)
- `num_partitions` (String)
- `offsets_retention_minutes` (String)
- `producer_purgatory_purge_interval_requests` (String)
- `replica_fetch_max_bytes` (String)
- `replica_fetch_response_max_bytes` (String)
- `socket_request_max_bytes` (String)
- `transaction_remove_expired_transaction_cleanup_interval_ms` (String)
- `transaction_state_log_segment_bytes` (String)


<a id="nestedobjatt--kafka_user_config--kafka_authentication_methods"></a>
### Nested Schema for `kafka_user_config.kafka_authentication_methods`

Read-Only:

- `certificate` (String)
- `sasl` (String)


<a id="nestedobjatt--kafka_user_config--kafka_connect_config"></a>
### Nested Schema for `kafka_user_config.kafka_connect_config`

Read-Only:

- `connector_client_config_override_policy` (String)
- `consumer_auto_offset_reset` (String)
- `consumer_fetch_max_bytes` (String)
- `consumer_isolation_level` (String)
- `consumer_max_partition_fetch_bytes` (String)
- `consumer_max_poll_interval_ms` (String)
- `consumer_max_poll_records` (String)
- `offset_flush_interval_ms` (String)
- `offset_flush_timeout_ms` (String)
- `producer_compression_type` (String)
- `producer_max_request_size` (String)
- `session_timeout_ms` (String)


<a id="nestedobjatt--kafka_user_config--kafka_rest_config"></a>
### Nested Schema for `kafka_user_config.kafka_rest_config`

Read-Only:

- `consumer_enable_auto_commit` (String)
- `consumer_request_max_bytes` (String)
- `consumer_request_timeout_ms` (String)
- `producer_acks` (String)
- `producer_linger_ms` (String)
- `simpleconsumer_pool_size_max` (String)


<a id="nestedobjatt--kafka_user_config--private_access"></a>
### Nested Schema for `kafka_user_config.private_access`

Read-Only:

- `prometheus` (String)


<a id="nestedobjatt--kafka_user_config--privatelink_access"></a>
### Nested Schema for `kafka_user_config.privatelink_access`

Read-Only:

- `jolokia` (String)
- `kafka` (String)
- `kafka_connect` (String)
- `kafka_rest` (String)
- `prometheus` (String)
- `schema_registry` (String)


<a id="nestedobjatt--kafka_user_config--public_access"></a>
### Nested Schema for `kafka_user_config.public_access`

Read-Only:

- `kafka` (String)
- `kafka_connect` (String)
- `kafka_rest` (String)
- `prometheus` (String)
- `schema_registry` (String)


<a id="nestedobjatt--kafka_user_config--schema_registry_config"></a>
### Nested Schema for `kafka_user_config.schema_registry_config`

Read-Only:

- `leader_eligibility` (String)
- `topic_name` (String)



<a id="nestedatt--service_integrations"></a>
### Nested Schema for `service_integrations`

Read-Only:

- `integration_type` (String)
- `source_service_name` (String)


<a id="nestedatt--tag"></a>
### Nested Schema for `tag`

Read-Only:

- `key` (String)
- `value` (String)


