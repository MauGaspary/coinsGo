package tools

import(
	log "github.com/sirupsen/logrus"
)
type LoginDetails struct {
	AuthToken string
	AccountID string
}

type AccountDetails struct {
	Balance float64
	AccountID string
}

type DatabaseInterface interface {
	GetUserLoginDetails(accountID string) *LoginDetails
	GetUserAccount(accountID string) *AccountDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	
	var database DatabaseInterface = &MockDatabase{}

	var err error = database.SetupDatabase()
	if err != nil{
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
