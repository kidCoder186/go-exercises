package account

import "sync"

// Account structure
type Account struct {
	balance int64
	closed  bool
	mux     sync.Mutex
}

// Open opens a new account.
func Open(initialDeposit int64) *Account {
	if initialDeposit >= 0 {
		return &Account{initialDeposit, false, sync.Mutex{}}
	}
	return nil
}

// Close closes the current account and returns all balance in the account.
func (t *Account) Close() (payout int64, ok bool) {
	t.mux.Lock()
	defer t.mux.Unlock()
	if !t.closed {
		t.closed = true
		return t.balance, true
	}
	return 0, false
}

// Balance returns current balance of an account.
func (t *Account) Balance() (balance int64, ok bool) {
	t.mux.Lock()
	defer t.mux.Unlock()
	if !t.closed {
		return t.balance, true
	}
	return 0, false
}

// Deposit adds an 'amount' into current balance and returns the added balance.
func (t *Account) Deposit(amount int64) (newBlance int64, ok bool) {
	t.mux.Lock()
	defer t.mux.Unlock()
	if !t.closed {
		if t.balance+amount >= 0 {
			t.balance += amount
			return t.balance, true
		}
		return 0, false
	}
	return 0, false
}
