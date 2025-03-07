---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "aiven_kafka_connector Data Source - terraform-provider-aiven"
subcategory: ""
description: |-
  The Kafka connector data source provides information about the existing Aiven Kafka connector.
---

# aiven_kafka_connector (Data Source)

The Kafka connector data source provides information about the existing Aiven Kafka connector.

## Example Usage

```terraform
data "aiven_kafka_connector" "kafka-es-con1" {
  project        = aiven_project.kafka-con-project1.project
  service_name   = aiven_service.kafka-service1.service_name
  connector_name = "kafka-es-con1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `connector_name` (String) The kafka connector name. This property cannot be changed, doing so forces recreation of the resource.
- `project` (String) Identifies the project this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.
- `service_name` (String) Specifies the name of the service that this resource belongs to. To set up proper dependencies please refer to this variable as a reference. This property cannot be changed, doing so forces recreation of the resource.

### Read-Only

- `config` (Map of String, Sensitive) The Kafka Connector configuration parameters.
- `id` (String) The ID of this resource.
- `plugin_author` (String) The Kafka connector author.
- `plugin_class` (String) The Kafka connector Java class.
- `plugin_doc_url` (String) The Kafka connector documentation URL.
- `plugin_title` (String) The Kafka connector title.
- `plugin_type` (String) The Kafka connector type.
- `plugin_version` (String) The version of the kafka connector.
- `task` (Set of Object) List of tasks of a connector. (see [below for nested schema](#nestedatt--task))

<a id="nestedatt--task"></a>
### Nested Schema for `task`

Read-Only:

- `connector` (String)
- `task` (Number)


