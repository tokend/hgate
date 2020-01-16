/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type PaymentRequest struct {
	Amount Amount `json:"amount"`
	// Account address of the receiver
	DestinationAccount *string `json:"destination_account,omitempty"`
	// Receiving balance id
	DestinationBalance *string `json:"destination_balance,omitempty"`
	DestinationFee     Fee     `json:"destination_fee"`
	// Reference for the payment
	Reference     string `json:"reference"`
	SourceBalance string `json:"source_balance"`
	SourceFee     Fee    `json:"source_fee"`
	// Whether source of the payment should pay destination fee
	SourcePayForDestination bool `json:"source_pay_for_destination"`
	// Subject of the payment
	Subject string `json:"subject"`
}
