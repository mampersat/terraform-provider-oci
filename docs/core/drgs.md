# oci_core_drg

## Drg Resource

### Drg Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the DRG.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `id` - The DRG's Oracle ID (OCID).
* `state` - The DRG's current state.
* `time_created` - The date and time the DRG was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 



### Create Operation
Creates a new Dynamic Routing Gateway (DRG) in the specified compartment. For more information,
see [Dynamic Routing Gateways (DRGs)](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingDRGs.htm).

For the purposes of access control, you must provide the OCID of the compartment where you want
the DRG to reside. Notice that the DRG doesn't have to be in the same compartment as the VCN,
the DRG attachment, or other Networking Service components. If you're not sure which compartment
to use, put the DRG in the same compartment as the VCN. For more information about compartments
and access control, see [Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).
For information about OCIDs, see [Resource Identifiers](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).

You may optionally specify a *display name* for the DRG, otherwise a default is provided.
It does not have to be unique, and you can change it. Avoid entering confidential information.


The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to contain the DRG.
* `display_name` - (Optional) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.


### Update Operation
Updates the specified DRG's display name. Avoid entering confidential information.


The following arguments support updates:
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

### Example Usage

```
resource "oci_core_drg" "test_drg" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.drg_display_name}"
}
```

# oci_core_drgs

## Drg DataSource

Gets a list of drgs.

### List Operation
Lists the DRGs in the specified compartment.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.


The following attributes are exported:

* `drgs` - The list of drgs.

### Example Usage

```
data "oci_core_drgs" "test_drgs" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
}
```