/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateAccountRequest struct {
	// ID of the role that will be attached to an account
	AccountRole uint64 `json:"account_role"`
	// ID of account to be created
	Destination string `json:"destination"`
	// ID of an another account that introduced this account into the system
	Referrer *string      `json:"referrer,omitempty"`
	Signers  []SignerData `json:"signers"`
}
