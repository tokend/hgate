/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type SignerData struct {
	Details Details `json:"details"`
	// Identity of the signer
	Identity uint32 `json:"identity"`
	// Public key of the signer
	PublicKey string `json:"public_key"`
	// ID of the role that will be attached to a signer
	SignerRole uint64 `json:"signer_role"`
	// weight that signer will have, threshold for all SignerRequirements equals 1000
	Weight uint32 `json:"weight"`
}
