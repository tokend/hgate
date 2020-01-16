/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type CreateAssetRequest struct {
	// Tasks to create request with
	AllTasks *uint32 `json:"all_tasks,omitempty"`
	// Unique asset identifier
	Code           string  `json:"code"`
	CreatorDetails Details `json:"creator_details"`
	// Amount to be issued automatically right after the asset creation
	InitialPreissuedAmount Amount `json:"initial_preissued_amount"`
	// Maximal amount to be issued
	MaxIssuanceAmount Amount `json:"max_issuance_amount"`
	// Policies specified for the asset creation
	Policies uint32 `json:"policies"`
	// Address of an account that performs pre issuance
	PreIssuanceAssetSigner string `json:"pre_issuance_asset_signer"`
	// omit if creating new request
	RequestId *uint64 `json:"request_id,omitempty"`
	// Number of digits after point (comma)
	TrailingDigitsCount uint32 `json:"trailing_digits_count"`
	// Numeric type of asset
	Type uint64 `json:"type"`
}
