package database

import (
	"errors"
	"github.com/fsoto82/go-multi-tenant/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type TenantConnProvider struct {
	db    *gorm.DB
	cache map[int64]*gorm.DB
}

func NewTenantConnProvider(db *gorm.DB) *TenantConnProvider {
	return &TenantConnProvider{
		db:    db,
		cache: map[int64]*gorm.DB{},
	}
}

func (t *TenantConnProvider) GetTenantConn(id int64) (*gorm.DB, error) {
	db, exists := t.cache[id]
	if exists {
		return db, nil
	}
	var tenant *models.Tenant
	result := t.db.First(&tenant, id)
	if result.Error != nil {
		log.Println("Error getting tenant data: ", result.Error)
	}
	if result.RowsAffected == 0 {
		log.Printf("Tenant %v not found\n", id)
		return nil, errors.New("tenant not found")
	}
	//var tenantDB *gorm.DB
	tenantDB, err := gorm.Open(postgres.Open(tenant.Data.String()))
	if err != nil {
		log.Println("Can not connect to tenant", id, err)
		return nil, errors.New("can not connect to tenant database")
	}
	t.cache[id] = tenantDB
	return tenantDB, nil
}
