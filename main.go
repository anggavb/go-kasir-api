package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONNECTION"`
}

func main() {
	config := viper.New()

	// 1. ENV selalu dibaca
	config.AutomaticEnv()

	// 2. Optional: load .env untuk local
	if _, err := os.Stat(".env"); err == nil {
		config.SetConfigFile(".env")
		if err := config.ReadInConfig(); err != nil {
			log.Fatal(err)
		}
	}
	c := Config{
		Port:   config.GetString("PORT"),
		DBConn: config.GetString("DB_CONNECTION"),
	}

	// setup database
	db, err := database.InitDB(c.DBConn)
	if err != nil {
		fmt.Println(config.Get("DB_CONNECTION"))
		log.Fatal(err)
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// setup routes
	http.HandleFunc("/api/product", productHandler.HandleProducts)
	http.HandleFunc("/api/product/", productHandler.HandleProductByID)

	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			GetCategories(w, r)
		} else if r.Method == "POST" {
			CreateCategory(w, r)
		}
	})

	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			GetCategoryById(w, r)
		} else if r.Method == "PUT" {
			UpdateCategory(w, r)
		} else if r.Method == "DELETE" {
			DeleteCategory(w, r)
		}
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "API is running",
		})
	})
	fmt.Println("Server running di localhost:" + c.Port)

	err = http.ListenAndServe(":"+c.Port, nil)
	if err != nil {
		fmt.Println("Gagal runing server")
	}
}
