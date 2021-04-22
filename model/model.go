package model

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Indica-se o tipo de conteúdo desejado, posteriormente inicializa-se a variável "book" sendo uma slice de Book (estrutura principal da aplicação) e atribui ao endereço da mesma o que for encontrado no banco de dados, encodando e demonstrando para o cliente

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book []Book
	DB.Find(&book)
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(http.StatusInternalServerError)
	}
}

// Verifica-se caso exista o id requisitado, se ouver ele é retornado ao cliente, do contrário é retornado um 404

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book Book
	err := DB.First(&book, params["id"]).Error
	errors.Is(err, gorm.ErrRecordNotFound)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(http.StatusNotFound)
	} else {
		json.NewEncoder(w).Encode(book)
	}

}

// Verifica-se se os inputs do cliente são condizentes, caso seja, as informações são armazenadas no banco de dados

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(http.StatusInternalServerError)
	} else {
		err02 := DB.Create(&book).Error
		errors.Is(err, gorm.ErrRecordNotFound)

		if err02 != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(http.StatusBadRequest)

		} else {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(book)

		}
	}
}

// Verifica-se se o id é existente, caso seja é atualizado no banco de dados

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book Book
	err := DB.First(&book, params["id"]).Error
	errors.Is(err, gorm.ErrRecordNotFound)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(http.StatusBadRequest)
	} else {
		json.NewDecoder(r.Body).Decode(&book)
		DB.Save(&book)
		json.NewEncoder(w).Encode(book)

	}

}

// Verifica-se se o id é existente, caso seja ele é deletado do banco de dados

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var book Book
	err := DB.First(&book, params["id"]).Error
	errors.Is(err, gorm.ErrRecordNotFound)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(http.StatusNotFound)

	} else {
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			DB.Delete(&book)
			json.NewEncoder(w).Encode("the book is deleted successfully")

		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(http.StatusInternalServerError)

		}
	}

}
