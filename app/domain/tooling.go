package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Tooling struct {
	Marking    string      `gorm:"primaryKey;column:marking"`
	TypeName   string      `gorm:"column:type"`
	Type       ToolingType `gorm:"column:type;foreignKey:TypeName"`
	Properties string      `gorm:"column:properties"`
}

type ToolingRequest struct {
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Type        *ToolingType `json:"type,omitempty"`
	DecayLevel  string       `json:"decayLevel,omitempty"`
	Supplier    *Supplier    `json:"supplier,omitempty"`
	AcquireTime *time.Time   `json:"acquireTime,omitempty"`
	Amount      int          `json:"amount,omitempty"`
}

type ToolingResponse struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Amount int    `json:"amount"`
}

type ToolingUsecase interface {
	Create(tooling *Tooling) error
	HydrateProperties(request ToolingRequest) (*Tooling, error)
	GetAll() ([]Tooling, error)
	//GetByConditions(conditions ToolingRequest) ([]Tooling, error)
	Update(tooling *Tooling) error
	Delete(marking string) error
}

type ToolingRepository interface {
	Create(tooling *Tooling) error
	Fetch(conditions string) ([]Tooling, error)
	Edit(tooling *Tooling) error
	Remove(marking string) error
}
