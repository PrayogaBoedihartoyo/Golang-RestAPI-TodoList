package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"main/model"
	"net/http"
	"strconv"
)

type response struct {
	Id          int64  `json:"id_todo,omitempty"`
	Status      string `json:"status,omitempty"`
	Description string `json:"description"`
	Message     string `json:"message"`
}

type Response struct {
	Status      string       `json:"status"`
	Message     string       `json:"message"`
	Description string       `json:"description"`
	Data        []model.Todo `json:"todo"`
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo

	// decode data json request ke todo
	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Fatalf("Tidak bisa mendecode dari request body.  %v", err)
	}

	// panggil modelsnya lalu insert buku
	insertID := model.CreateTodo(todo)

	// format response objectnya
	res := response{
		Id:      insertID,
		Message: "Data buku telah ditambahkan",
	}

	// kirim response
	json.NewEncoder(w).Encode(res)
}

// AmbilBuku mengambil single data dengan parameter id
func FindTodo(w http.ResponseWriter, r *http.Request) {
	// kita set headernya
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// dapatkan idbuku dari parameter request, keynya adalah "id"
	params := mux.Vars(r)

	// konversi id dari tring ke int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	// memanggil models ambilsatubuku dengan parameter id yg nantinya akan mengambil single data
	todo, err := model.FindTodo(int64(id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data buku. %v", err)
	}

	// kirim response
	json.NewEncoder(w).Encode(todo)
}

// Ambil semua data buku
func FindAllTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// memanggil models AmbilSemuaBuku
	todos, err := model.FindAllTodo()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var response Response
	response.Status = "Successfully"
	response.Message = "Success"
	response.Data = todos

	// kirim semua response
	json.NewEncoder(w).Encode(response)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {

	// kita ambil request parameter idnya
	params := mux.Vars(r)

	// konversikan ke int yang sebelumnya adalah string
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	// buat variable buku dengan type models.Buku
	var todo model.Todo

	// decode json request ke variable buku
	err = json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Fatalf("Tidak bisa decode request body.  %v", err)
	}

	// panggil updatebuku untuk mengupdate data
	updatedRows := model.UpdateTodo(int64(id), todo)

	// ini adalah format message berupa string
	msg := fmt.Sprintf("Buku telah berhasil diupdate. Jumlah yang diupdate %v rows/record", updatedRows)

	// ini adalah format response message
	res := response{
		Id:      int64(id),
		Message: msg,
	}

	// kirim berupa response
	json.NewEncoder(w).Encode(res)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {

	// kita ambil request parameter idnya
	params := mux.Vars(r)

	// konversikan ke int yang sebelumnya adalah string
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	// panggil fungsi hapusbuku , dan convert int ke int64
	deletedRows := model.DeleteTodo(int64(id))

	// ini adalah format message berupa string
	msg := fmt.Sprintf("buku sukses di hapus. Total data yang dihapus %v", deletedRows)

	// ini adalah format reponse message
	res := response{
		Id:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}
