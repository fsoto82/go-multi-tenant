package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Tenant struct {
	//gorm.Model
	ID   int
	Code string
	Name string
	Data TenantConnData
}

type TenantConnData struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (Tenant) TableName() string {
	return "tenant"
}

func (tcd TenantConnData) Value() (driver.Value, error) {
	return json.Marshal(tcd)
}

func (tcd TenantConnData) String() string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
		tcd.Host, tcd.Port, tcd.Database, tcd.User, tcd.Password)
}

func (tcd *TenantConnData) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &tcd)
}
