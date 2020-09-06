package errors

import "github.com/pkg/errors"

var (
	// ErrMySQLDBEmpty : empty DB
	ErrMySQLDBEmpty = errors.New("empty mysql db")
	// ErrMySQLConnErrtion : mysql connErrtion error
	ErrMySQLConnErrtion = errors.New("mysql connErrtion error")
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

// Wrap : wrap error with a message
func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}
