package domain

type Failure struct {
	Equipment    string `gorm:"column:equipment;primary_key'"`
	FailureTime  string `gorm:"column:failure_time;primary_key'"`
	Reason       string `gorm:"column:reason'"`
	ContinueTime string `gorm:"column:continue_time'"`
}

type FailureRepository interface {
	Create(*Failure) error
	FetchAll() ([]Failure, error)
}
