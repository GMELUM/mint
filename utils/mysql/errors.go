package mysql

func NewError(err error) *MySQLError {
	return &MySQLError{
		Number:   45000,
		SQLState: [5]byte{0, 0, 0, 0, 0},
		Message:  err.Error(),
	}
}
