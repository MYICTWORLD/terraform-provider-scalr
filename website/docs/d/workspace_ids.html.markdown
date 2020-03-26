---
layout: "scalr"
page_title: "Scalr: scalr_workspace_ids"
sidebar_current: "docs-datasource-scalr-workspace-ids"
description: |-
  Get information on (external) workspace IDs.
---

# Data Source: scalr_workspace_ids

Use this data source to get a map of (external) workspace IDs.

## Example Usage

```hcl
data "scalr_workspace_ids" "app-frontend" {
  names        = ["app-frontend-prod", "app-frontend-dev1", "app-frontend-staging"]
  organization = "my-org-name"
}

data "scalr_workspace_ids" "all" {
  names        = ["*"]
  organization = "my-org-name"
}
```

## Argument Reference

The following arguments are supported:

* `names` - (Required) A list of workspace names to search for. Names that don't
  match a real workspace will be omitted from the results, but are not an error.

    To select _all_ workspaces for an organization, provide a list with a single
    asterisk, like `["*"]`. No other use of wildcards is supported.
* `organization` - (Required) Name of the organization.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `ids` - A map of workspace names and their human-readable IDs, which look like
  `<ORGANIZATION>/<WORKSPACE>`.
* `external_ids` - A map of workspace names and their opaque external IDs, which
  look like `ws-<RANDOM STRING>`.
