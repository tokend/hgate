/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type UpdateAssetRequest struct {
	// Tasks to create request with
	AllTasks *uint32 `json:"all_tasks,omitempty"`
	// Unique asset identifier
	Code           string  `json:"code"`
	CreatorDetails Details `json:"creator_details"`
	// Policies specified for the asset
	Policies uint32 `json:"policies"`
	// omit if creating new request
	RequestId *uint64 `json:"request_id,omitempty"`
}
