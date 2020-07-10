/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ChangeRoleRequest struct {
	// ID of the account role to set for destination account
	AccountRole    uint64  `json:"account_role"`
	AllTasks       *uint32 `json:"all_tasks,omitempty"`
	CreatorDetails Details `json:"creator_details"`
	// ID of account to change roles
	Destination string `json:"destination"`
}
