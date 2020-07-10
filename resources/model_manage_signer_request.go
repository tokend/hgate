/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type ManageSignerRequest struct {
	// * 0 - create * 1 - update * 2 - remove
	Action           int32             `json:"action"`
	RemoveSignerData *RemoveSignerData `json:"remove_signer_data,omitempty"`
	SignerData       *SignerData       `json:"signer_data,omitempty"`
}
