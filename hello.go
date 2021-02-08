package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)  	

type ipscoring struct {
	ID          string `json:"ID"`
	IP       string `json:"IP"`
}

type ip_list []ipscoring

var ip_address = ip_list{
	{
		ID:          "1",
		IP:       "12.23.34.54",
	},
	{
		ID:          "2",
		IP:       "11.11.11.11",
	},
	{
		ID:          "3",
		IP:       "12.22.22.22",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func createIp(w http.ResponseWriter, r *http.Request) {
	var newIp ipscoring
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, " enter ip")
	}

	json.Unmarshal(reqBody, &newIp)
	ip_address = append(ip_address, newIp)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newIp)
}

func get_ip(w http.ResponseWriter, r *http.Request) {
	ip_id := mux.Vars(r)["id"]

	for _, singleEvent := range ip_address {
		if singleEvent.ID == ip_id {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func getAllIp(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ip_address)
}

func update_ip(w http.ResponseWriter, r *http.Request) {
	ip_id := mux.Vars(r)["id"]
	var updated_ip ipscoring

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "enter ip : ")
	}
	json.Unmarshal(reqBody, &updated_ip)

	for i, singleEvent := range ip_address {
		if singleEvent.ID == ip_id {
			singleEvent.IP = updated_ip.IP
			ip_address = append(ip_address[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func delete_ip(w http.ResponseWriter, r *http.Request) {
	ip_id := mux.Vars(r)["id"]

	for i, singleEvent := range ip_address {
		if singleEvent.ID == ip_id {
			ip_address = append(ip_address[:i], ip_address[i+1:]...)
			fmt.Fprintf(w, "The IP with ID %v has been deleted successfully", ip_id)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/ipscore", createIp).Methods("POST")
	router.HandleFunc("/ipscore", getAllIp).Methods("GET")
	router.HandleFunc("/ipscore/{id}", get_ip).Methods("GET")
	router.HandleFunc("/ipscore/{id}", update_ip).Methods("PUT")
	router.HandleFunc("/ipscore/{id}", delete_ip).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}