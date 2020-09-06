package errors

import "github.com/pkg/errors"

var (
	// ErrSystem : system error
	ErrSystem = errors.New("system error")
)

var (
	// ErrMySQLDBEmpty : empty DB
	ErrMySQLDBEmpty = errors.New("empty mysql db")
	// ErrMySQLConnErrtion : mysql connection error
	ErrMySQLConnErrtion = errors.New("mysql connection error")
	// ErrMySQLDBAutoMigrate : mysql db auto migrate error
	ErrMySQLDBAutoMigrate = errors.New("mysql db auto migrate error")
	// ErrMySQLCreate : mysql create model error
	ErrMySQLCreate = errors.New("mysql create model error")
	// ErrMySQLRead : mysql read model error
	ErrMySQLRead = errors.New("mysql read model error")
	// ErrMySQLUpdate : mysql update model error
	ErrMySQLUpdate = errors.New("mysql update model error")
	// ErrMySQLDelete : mysql delete model error
	ErrMySQLDelete = errors.New("mysql delete model error")
)

var (
	// ErrInvalidSellingPercentage : invalid selling percentage
	ErrInvalidSellingPercentage = errors.New("invalid selling percentage")

	// ErrInvalidSellingPriceFormat : invalid selling price format
	ErrInvalidSellingPriceFormat = errors.New("Invalid selling price format. It must be specified as a positive decimal value to two decimal places")
	// ErrInvalidSellingPrice : invalid selling price
	ErrInvalidSellingPrice = errors.New("invalid selling price")

	// ErrInvalidBuyingPrice : invalid buying price
	ErrInvalidBuyingPrice = errors.New("invalid buying price")

	// ErrWagerNotExist : wager not exist
	ErrWagerNotExist = errors.New("wager not exist")
)

// Wrap : wrap error with a message
func Wrap(err error, messages ...string) error {
	for _, msg := range messages {
		err = errors.Wrap(err, msg)
	}
	return err
}
