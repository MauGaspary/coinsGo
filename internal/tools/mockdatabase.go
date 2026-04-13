package tools

import (
	"time"
)

type MockDatabase struct {}

var mockLoginDetails = map[string]LoginDetails{
	"rafaela": {
		AuthToken: "123",
		AccountID: "rafaela",
	},
	"mariaclara": {
		AuthToken: "456",
		AccountID: "anaclara",
	},
	"mauricio": {
		AuthToken: "789",
		AccountID: "mauricio",
	},
}

var mockBalanceDetails = map[string]AccountDetails{
	"rafaela": {
		Balance:   1000.50,
		AccountID: "rafaela",
	},
	"mariaclara": {
		Balance:   250.75,
		AccountID: "mariaclara",
	},
	"mauricio": {
		Balance:   300.00,
		AccountID: "mauricio",
	},
}

func (db *MockDatabase) GetUserLoginDetails(accountID string) *LoginDetails {
	time.Sleep(100 * time.Millisecond)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[accountID]
	if !ok {
		return nil
	}

	return &clientData
}

func (db *MockDatabase) GetUserAccount(accountID string) *AccountDetails {
	time.Sleep(100 * time.Millisecond)

	var accountData = AccountDetails{}
	accountData, ok := mockBalanceDetails[accountID]
	if !ok {
		return nil
	}

	return &accountData
}

func (db *MockDatabase) SetupDatabase() error {
	return nil
}
