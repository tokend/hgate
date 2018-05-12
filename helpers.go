package hgate

import (
	"gitlab.com/tokend/go/keypair"
	"gitlab.com/tokend/go/xdr"
	"errors"
)

func IsValidAccountId(accountId string) bool {
	_, err := keypair.Parse(accountId)
	return err == nil
}

func ToXDRString64(s string) (xdr.String64, error) {
	if len(s) > 64 {
		return "", errors.New("invalid length - must be lower then 64")
	}
	return xdr.String64(s), nil
}
func ToXDRString256(s string) (xdr.String256, error) {
	if len(s) > 256 {
		return "", errors.New("invalid length - must be lower then 256")
	}
	return xdr.String256(s), nil
}
