package customers

import (
	"context"
	"database/sql"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

//ErrNotFound for
var ErrNotFound = errors.New("item not found")

//ErrInternal for
var ErrInternal = errors.New("internal error")


// Service npenctasnset co6oi cepsuc no ynpasnenwo OaHHepamn.
type Service struct {
	db *sql.DB
}

// NewService co3qaét cepsuc.
func NewService(db *sql.DB) *Service {
	return &Service{db: db}
	
}

// Customer npenctasnaet codoi GaHHep.
type Customer struct {
	ID int64			'json:"id"'
	Name string			'json:"name"'
	Phone string		'json:"phone"'	
	Active bool			'json:"active"'
	Created time.Time 	'json:"created"'
}

// ByID Bo3BpawaeT OaHHep no upeHTHOuKaTopy.
func (s *Service) ByID(ctx context.Context, id int64) (*Customer, error) {
	item := &Customer{}

	err := s.db.QueryRowContext(ctx, '
		SELECT id, name, phone, active, created FROM customers WHERE id = $1
	', id).Scan(&item.Id, &item.Name, &item.Phone, &item.Active, &item.Created)
	
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}

	if err != nil {
		log.Print(err)
		return nil, ErrInternal
	}

	return item, nil
}

// All for
func (s *Service) All(ctx context.Context) ([]*Customer, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	// for _, Customer := range s.items {
	// 	if Customer.ID == id {
	// 		return Customer, nil
	// 	}
	// }
	// Customers := s.items
	// if len(s.items) == 0 {
	// 	return nil, errors.New("no items found")
	// }
	return s.items, nil
	//panic("not implemented")
}

// Save for
func (s *Service) Save(ctx context.Context, item *Customer, file multipart.File) (*Customer, error) {
	//var lastID int64
	s.mu.RLock()
	defer s.mu.RUnlock()
	//	for _, Customer := range s.items {
	if item.ID == 0 {
		// lenCustomers := len(s.items) - 1
		// for i, Customer := range s.items {
		// 	if i == lenCustomers {
		// 		lastID = Customer.ID
		// 	}
		// }
		s.nextAccountID++
		item.ID = s.nextAccountID
		//		item.ID = int64(len(s.items)) + 1
		nameImage := item.Image
		if nameImage != "" {
			extenIndex := strings.Index(nameImage, ".")
			fileExtension := nameImage[extenIndex:]
			item.Image = strconv.FormatInt(item.ID, 10) + fileExtension
		}
		saveFile(file, item)
		s.items = append(s.items, item)

		return item, nil
	}
	if item.ID != 0 {
		for _, Customer := range s.items {
			if Customer.ID == item.ID {
				Customer.Button = item.Button
				Customer.Content = item.Content
				Customer.Link = item.Link
				Customer.Title = item.Title
				if item.Image != "" {
					
					Customer.Image = item.Image
					saveFile(file, item)
				}
				if item.Image == "" {
					item.Image = Customer.Image
					//Customer.Image = ""
					// nameImage := item.Image
					// extenIndex := strings.Index(nameImage, ".")
					// fileExtension := nameImage[extenIndex:]
					// item.Image = strconv.FormatInt(item.ID, 10) + fileExtension
					// Customer.Image = item.Image
				}
				
				return item, nil
			}
		}
	}
	return nil, errors.New("item not found")

}

// RemoveByID for
func (s *Service) RemoveByID(ctx context.Context, id int64) (*Customer, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for i, Customer := range s.items {
		if Customer.ID == id {

			s.items = append(s.items[:i], s.items[i+1:]...)

			return Customer, nil
		}
	}

	return nil, errors.New("item not found")
}

// Initial for
func (s *Service) Initial(request *http.Request) Customer {

	idParam := request.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		log.Print(err)

	}

	titleParam := request.URL.Query().Get("title")
	contentParam := request.URL.Query().Get("content")
	buttonParam := request.URL.Query().Get("button")
	linkParam := request.URL.Query().Get("link")

	Customer := Customer{
		ID: id,

		Title: titleParam,

		Content: contentParam,

		Button: buttonParam,

		Link: linkParam,

		Image: "image1",
	}

	// Customer2 := Customer{
	// 	ID: 2,

	// 	Title: "Title New",

	// 	Content: "Content New",

	// 	Button: "Button New",

	// 	Link: "Link New",
	// }

	//item := s.items
	//	s.items = append(s.items, &Customer)
	//s.items = append(s.items, &Customer2)
	//item[1] = &Customer
	//	panic("not implemented")

	return Customer
}

func saveFile(fileA multipart.File, item *Customer) {
	content := make([]byte, 0)
	buf := make([]byte, 4)
		for {
			read, err := fileA.Read(buf)
			if err == io.EOF {
				break
			}
			content = append(content, buf[:read]...)
		}

		fileNameNew := item.Image
		if fileNameNew != "" {
			wdd1 := "web/Customers" + "/" + fileNameNew
			//wdd1 := "c:/projects/http/web/Customers" + "/" + fileNameNew
			//log.Print(wdd)
			err := ioutil.WriteFile(wdd1, content, 0600)
			if err != nil {
				log.Print(err)
	
			}
		}
		
}