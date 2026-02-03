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
		AccountID: "mariaclara",
	},
	"laura": {
		AuthToken: "789",
		AccountID: "laura",
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
	"laura": {
		Balance:   300.00,
		AccountID: "laura",
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

func (db *MockDatabase) GetUserAccountDetails(accountID string) *AccountDetails {
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
