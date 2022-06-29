package main

import (
	"fmt"
	db "github.com/fsoto82/go-multi-tenant/database"
	"github.com/fsoto82/go-multi-tenant/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func displayProducts(tcp *db.TenantConnProvider, id int64) {
	if tenantConn, err := tcp.GetTenantConn(id); err != nil {
		//log.Println(err)

	} else {
		var products []models.Product
		result := tenantConn.Find(&products)
		if result.Error != nil {
			//log.Println("Error retrieving users", result.Error)
			//return result.Error
			return
		}
		for _, p := range products {
			log.Printf("Product %+v\n", p)
		}
	}
	//return nil
}

func main() {
	fmt.Println("Initializing database and gorm...")
	gormDB, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Panicln("Failed to open database")
	}

	tcp := db.NewTenantConnProvider(gormDB)

	displayProducts(tcp, 1)
	displayProducts(tcp, 1)
	//if displayProducts(tcp, 1) != nil {
	//	log.Println("Error reading products from tenant 1")
	//}

	displayProducts(tcp, 2)
	//if displayProducts(tcp, 2) != nil {
	//	log.Println("Error reading products from tenant 2")
	//}

	displayProducts(tcp, 3)
	//if displayProducts(tcp, 3) != nil {
	//	log.Println("Error reading products from tenant 3")
	//}

}
