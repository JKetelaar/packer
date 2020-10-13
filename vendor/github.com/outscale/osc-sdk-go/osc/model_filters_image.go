/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.2
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc
// FiltersImage One or more filters.
type FiltersImage struct {
	// The account aliases of the owners of the OMIs.
	AccountAliases []string `json:"AccountAliases,omitempty"`
	// The account IDs of the owners of the OMIs. By default, all the OMIs for which you have launch permissions are described.
	AccountIds []string `json:"AccountIds,omitempty"`
	// The architectures of the OMIs (`i386` \\| `x86_64`).
	Architectures []string `json:"Architectures,omitempty"`
	// Indicates whether the block device mapping is deleted when terminating the VM.
	BlockDeviceMappingDeleteOnVmDeletion bool `json:"BlockDeviceMappingDeleteOnVmDeletion,omitempty"`
	// The device names for the volumes.
	BlockDeviceMappingDeviceNames []string `json:"BlockDeviceMappingDeviceNames,omitempty"`
	// The IDs of the snapshots used to create the volumes.
	BlockDeviceMappingSnapshotIds []string `json:"BlockDeviceMappingSnapshotIds,omitempty"`
	// The sizes of the volumes, in gibibytes (GiB).
	BlockDeviceMappingVolumeSizes []int32 `json:"BlockDeviceMappingVolumeSizes,omitempty"`
	// The types of volumes (`standard` \\| `gp2` \\| `io1`).
	BlockDeviceMappingVolumeTypes []string `json:"BlockDeviceMappingVolumeTypes,omitempty"`
	// The descriptions of the OMIs, provided when they were created.
	Descriptions []string `json:"Descriptions,omitempty"`
	// The locations where the OMI files are stored on Object Storage Unit (OSU).
	FileLocations []string `json:"FileLocations,omitempty"`
	// The IDs of the OMIs.
	ImageIds []string `json:"ImageIds,omitempty"`
	// The names of the OMIs, provided when they were created.
	ImageNames []string `json:"ImageNames,omitempty"`
	// The account IDs of the users who have launch permissions for the OMIs.
	PermissionsToLaunchAccountIds []string `json:"PermissionsToLaunchAccountIds,omitempty"`
	// If `true`, lists all public OMIs. If `false`, lists all private OMIs.
	PermissionsToLaunchGlobalPermission bool `json:"PermissionsToLaunchGlobalPermission,omitempty"`
	// The device names of the root devices (for example, `/dev/sda1`).
	RootDeviceNames []string `json:"RootDeviceNames,omitempty"`
	// The types of root device used by the OMIs (always `bsu`).
	RootDeviceTypes []string `json:"RootDeviceTypes,omitempty"`
	// The states of the OMIs (`pending` \\| `available` \\| `failed`).
	States []string `json:"States,omitempty"`
	// The keys of the tags associated with the OMIs.
	TagKeys []string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the OMIs.
	TagValues []string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the OMIs, in the following format: \"Filters\":{\"Tags\":[\"TAGKEY=TAGVALUE\"]}.
	Tags []string `json:"Tags,omitempty"`
	// The virtualization types (always `hvm`).
	VirtualizationTypes []string `json:"VirtualizationTypes,omitempty"`
}
