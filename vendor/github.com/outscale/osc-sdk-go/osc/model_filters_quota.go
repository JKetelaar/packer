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
// FiltersQuota One or more filters.
type FiltersQuota struct {
	// The group names of the quotas.
	Collections []string `json:"Collections,omitempty"`
	// The names of the quotas.
	QuotaNames []string `json:"QuotaNames,omitempty"`
	// The resource IDs if these are resource-specific quotas, `global` if they are not.
	QuotaTypes []string `json:"QuotaTypes,omitempty"`
	// The description of the quotas.
	ShortDescriptions []string `json:"ShortDescriptions,omitempty"`
}
