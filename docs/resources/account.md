---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "aiven_account Resource - terraform-provider-aiven"
subcategory: ""
description: |-
  The Account resource allows the creation and management of an Aiven Account.
---

# aiven_account (Resource)

The Account resource allows the creation and management of an Aiven Account.

## Example Usage

```terraform
resource "aiven_account" "account1" {
  name = "<ACCOUNT_NAME>"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Account name

### Optional

- `primary_billing_group_id` (String) Billing group id

### Read-Only

- `account_id` (String) Account id
- `create_time` (String) Time of creation
- `id` (String) The ID of this resource.
- `is_account_owner` (Boolean) If true, user is part of the owners team for this account
- `owner_team_id` (String) Owner team id
- `tenant_id` (String) Tenant id
- `update_time` (String) Time of last update

## Import

Import is supported using the following syntax:

```shell
terraform import aiven_account.account1 account_id
```
