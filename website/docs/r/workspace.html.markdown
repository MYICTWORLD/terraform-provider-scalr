---
layout: "scalr"
page_title: "Scalr: scalr_workspace"
sidebar_current: "docs-resource-scalr-workspace"
description: |-
  Manages workspaces.
---

# scalr_workspace

Provides a workspace resource.

## Example Usage

Basic usage:

```hcl
resource "scalr_workspace" "test" {
  name         = "my-workspace-name"
  organization = "my-org-name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the workspace.
* `organization` - (Required) Name of the organization.
* `auto_apply` - (Optional) Whether to automatically apply changes when a
  Terraform plan is successful. Defaults to `false`.
* `operations` - (Optional) Whether to use remote execution mode. When set
  to `false`, the workspace will be used for state storage only.
  Defaults to `true`.
* `queue_all_runs` - (Optional) Whether all runs should be queued. When set
  to `false`, runs triggered by a VCS change will not be queued until at least
  one run is manually queued. Defaults to `true`.
* `terraform_version` - (Optional) The version of Terraform to use for this workspace. Defaults to the latest available version.
* `working_directory` - (Optional) A relative path that Terraform will execute
  within.  Defaults to the root of your repository.
* `vcs_repo` - (Optional) Settings for the workspace's VCS repository.

The `vcs_repo` block supports:

* `identifier` - (Required) A reference to your VCS repository in the format
  `:org/:repo` where `:org` and `:repo` refer to the organization and repository
  in your VCS provider.
* `branch` - (Optional) The repository branch that Terraform will execute from.
  Default to `master`.
* `ingress_submodules` - (Optional) Whether submodules should be fetched when
  cloning the VCS repository. Defaults to `false`.
* `oauth_token_id` - (Required) Token ID of the VCS Connection (OAuth Conection Token)
  to use.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The workspace's human-readable ID, which looks like
  `<ORGANIZATION>/<WORKSPACE>`.
* `external_id` - The workspace's opaque external ID, which looks like
  `ws-<RANDOM STRING>`.

## Import

Workspaces can be imported; use `<ORGANIZATION NAME>/<WORKSPACE NAME>` as the
import ID. For example:

```shell
terraform import scalr_workspace.test my-org-name/my-workspace-name
```
