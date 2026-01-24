package main

import (
	"encoding/json"
	"kasir-api/models"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

var categories = []models.Category{
	{ID: 1, Name: "Baju Lebaran", Description: "Deskripsi Baju Lebaran"},
	{ID: 2, Name: "Peralatan Dapur", Description: "Deskripsi Peralatan Dapur"},
	{ID: 3, Name: "Sembako", Description: "Deskripsi Sembako"},
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var newCategory models.Category
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newCategory.ID = len(categories) + 1
	categories = append(categories, newCategory)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(newCategory)
}

func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, c := range categories {
		if c.ID == id {
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updateCategory models.Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range categories {
		if categories[i].ID == id {
			updateCategory.ID = id
			categories[i] = updateCategory

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(categories)
			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, c := range categories {
		if c.ID == id {
			// nyobain pake package slices
			categories = slices.Concat(categories[:i], categories[i+1:])

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Category deleted",
			})

			return
		}
	}

	http.Error(w, "Category not found", http.StatusNotFound)
}
