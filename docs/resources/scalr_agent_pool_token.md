---
layout: "scalr"
page_title: "Scalr: scalr_agent_pool_token"
sidebar_current: "docs-resource-scalr-agent-pool-token"
description: |-
  Manages agent pool's tokens.
---

# scalr_agent_pool_token Resource

Manage the state of agent pool's tokens in Scalr. Create, update and destroy.

## Example Usage

Basic usage:

```hcl
resource "scalr_agent_pool_token" "default" {
  description   = "Some description"
  agent_pool_id = "apool-xxxxxxx"
}
```

## Argument Reference

* `description` - (Required) Description of the token.
* `agent_pool_id` - (Required) ID of the agent pool.

## Attribute Reference

All arguments plus:

* `id` - The ID of the token.
* `token` - The token of the agent pool.

## Import

To import agent pool's token use token ID as the import ID. For example:
```shell
terraform import scalr_agent_pool_token.default at-xxxxxxxxx
```