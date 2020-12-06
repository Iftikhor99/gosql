package app

import (
	"encoding/json"
	//	"io"
//	"io/ioutil"
	"log"
	"net/http"

	//	"os"
	"strconv"
	//	"strings"

	"github.com/Iftikhor99/gosql/pkg/customers"
	"github.com/gorilla/mux"
)

const (
	//GET for
	GET = "GET"
	//POST for
	POST = "POST"
	//DELETE for
	DELETE = "DELETE"
)

// Server npegctasnseT coOow normyeckwi CepBep Hawero npunomeHna.
type Server struct {
	mux *mux.Router

	customersSvc *customers.Service
}

// NewServer - OyHKUMA-KOHCTpykTOp pina co3maHna cepsepa.
func NewServer(mux *mux.Router, customersSvc *customers.Service) *Server {

	return &Server{mux: mux, customersSvc: customersSvc}

}

//ServeHTTP for
func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	s.mux.ServeHTTP(writer, request)

}

// Init wHmunannsupyet cepsep (permctpupyet sce Handler's)
func (s *Server) Init() {
	//s.mux.HandleFunc("/customers.getAll", s.handleGetAllcustomers)
	s.mux.HandleFunc("/customers", s.handleGetAllcustomers).Methods(GET)
	//s.mux.HandleFunc("/customers.getAllActive", s.handleGetAllActivecustomers)
	s.mux.HandleFunc("/customers/active", s.handleGetAllActivecustomers).Methods(GET)
	//s.mux.HandleFunc("/customers.getById", s.handleGetCustomerByID)
	s.mux.HandleFunc("/customers/{id}", s.handleGetCustomerByID).Methods(GET)
	//s.mux.HandleFunc("/customers.blockById", s.handleCustomerblockByID)
	s.mux.HandleFunc("/customers/{id}/block", s.handleCustomerblockByID).Methods(POST)
	//s.mux.HandleFunc("/customers.unblockById", s.handleCustomerunblockByID)
	s.mux.HandleFunc("/customers/{id}/unblock", s.handleCustomerunblockByID).Methods(DELETE)
	//s.mux.HandleFunc("/customers.removeById", s.handleCustomerremoveByID)
	s.mux.HandleFunc("/customers", s.handleCustomerRemoveByID).Methods(DELETE)
	//s.mux.HandleFunc("/customers.save", s.handleSaveCustomer)
	s.mux.HandleFunc("/customers", s.handleSaveCustomer).Methods(POST)
	

}

func (s *Server) handleGetCustomerByID(writer http.ResponseWriter, request *http.Request) {

	//idParam := request.URL.Query().Get("id")
	idParam, ok := mux.Vars(request)["id"]
	if !ok {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	item, err := s.customersSvc.ByID(request.Context(), id)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(item)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(data)

	if err != nil {
		log.Print(err)
	}

}

func (s *Server) handleCustomerblockByID(writer http.ResponseWriter, request *http.Request) {

	//idParam := request.URL.Query().Get("id")
	idParam, ok := mux.Vars(request)["id"]
	if !ok {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	item, err := s.customersSvc.BlockByID(request.Context(), id)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(item)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(data)

	if err != nil {
		log.Print(err)
	}

}


func (s *Server) handleCustomerunblockByID(writer http.ResponseWriter, request *http.Request) {

	//idParam := request.URL.Query().Get("id")

	idParam, ok := mux.Vars(request)["id"]
	if !ok {
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	item, err := s.customersSvc.UnBlockByID(request.Context(), id)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(item)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(data)

	if err != nil {
		log.Print(err)
	}

}


func (s *Server) handleCustomerRemoveByID(writer http.ResponseWriter, request *http.Request) {

	idParam := request.URL.Query().Get("id")

	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	item, err := s.customersSvc.RemoveByID(request.Context(), id)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(item)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(data)

	if err != nil {
		log.Print(err)
	}

}


func (s *Server) handleSaveCustomer(writer http.ResponseWriter, request *http.Request) {
	
	log.Print(request.RequestURI)
	log.Print(request.Method)
	log.Print(request.Header)
	log.Print(request.Header.Get("Content-Type"))

	log.Print(request.FormValue("id"))
	log.Print(request.FormValue("name"))
	log.Print(request.FormValue("phone"))
	
	
	// body, err := ioutil.ReadAll(request.Body)
	// if err != nil {
	// 	log.Print(err)
	// }
	// log.Printf("%s", body)

	// err = request.ParseMultipartForm(10 * 1024 * 1024)
	// if err != nil {
	// 	log.Print(err)
	// }

	// log.Print(request.Form)
	// log.Print(request.PostForm)
	// idParam := request.FormValue("id")
	
	// id, err := strconv.ParseInt(idParam, 10, 64)
	// if err != nil {
	// 	log.Print(err)

	// }

	// nameParam := request.FormValue("name")
	// phoneParam := request.FormValue("phone")
	
	// customer := customers.Customer{
	// 	ID: id,

	// 	Name: nameParam,

	// 	Phone: phoneParam,
		
	// }

	var customer *customers.Customer
	err := json.NewDecoder(request.Body).Decode(&customer)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	item, err := s.customersSvc.Save(request.Context(), *customer)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	
	data, err := json.Marshal(item)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(data)

	if err != nil {
		log.Print(err)
	}

	log.Print(item)

}


func (s *Server) handleGetAllcustomers(writer http.ResponseWriter, request *http.Request) {
	log.Print(request)
	log.Print(request.Header)
	log.Print(request.Body)
	item, err := s.customersSvc.All(request.Context())

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(item)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(data)

	if err != nil {
		log.Print(err)
	}

	log.Printf("%#v", item)
}

func (s *Server) handleGetAllActivecustomers(writer http.ResponseWriter, request *http.Request) {
	log.Print(request)
	log.Print(request.Header)
	log.Print(request.Body)
	item, err := s.customersSvc.AllActive(request.Context())

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(item)

	if err != nil {
		log.Print(err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(data)

	if err != nil {
		log.Print(err)
	}

	log.Printf("%#v", item)
}
