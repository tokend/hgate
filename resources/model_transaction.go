/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Transaction struct {
	Key
	Attributes TransactionAttributes `json:"attributes"`
}
type TransactionResponse struct {
	Data     Transaction `json:"data"`
	Included Included    `json:"included"`
}

type TransactionListResponse struct {
	Data     []Transaction `json:"data"`
	Included Included      `json:"included"`
	Links    *Links        `json:"links"`
}

// MustTransaction - returns Transaction from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustTransaction(key Key) *Transaction {
	var transaction Transaction
	if c.tryFindEntry(key, &transaction) {
		return &transaction
	}
	return nil
}
