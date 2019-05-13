package txnbuild

import (
	"strconv"

	"github.com/stellar/go/support/errors"
	"github.com/stellar/go/xdr"
)

// SimpleAccount implements the Account interface.
type SimpleAccount struct {
	AccountID string
	Sequence  string
}

// GetAccountID returns the Account ID.
func (sa *SimpleAccount) GetAccountID() string {
	return sa.AccountID
}

// IncrementSequenceNumber increments the internal record of the
// account's sequence number by 1.
func (sa *SimpleAccount) IncrementSequenceNumber() (xdr.SequenceNumber, error) {
	seqNum, err := sa.GetSequenceNumber()
	if err != nil {
		return xdr.SequenceNumber(0), err
	}
	seqNum++
	sa.Sequence = strconv.FormatInt(int64(seqNum), 10)
	return sa.GetSequenceNumber()
}

// GetSequenceNumber returns the sequence number of the account.
func (sa *SimpleAccount) GetSequenceNumber() (xdr.SequenceNumber, error) {
	seqNum, err := strconv.ParseUint(sa.Sequence, 10, 64)
	if err != nil {
		return 0, errors.Wrap(err, "Failed to parse account sequence number")
	}

	return xdr.SequenceNumber(seqNum), nil
}

// NewSimpleAccount is a factory method that creates a SimpleAccount from "accountID" and "sequence".
// If "sequence" is not set, it defaults to 0.
func NewSimpleAccount(accountID string, sequence string) SimpleAccount {
	if sequence == "" {
		sequence = "0"
	}
	return SimpleAccount{accountID, sequence}
}

// ensure that SimpleAccount implements Account interface.
var _ Account = &SimpleAccount{}
