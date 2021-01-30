package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/filariow/bshop"
	"github.com/gorilla/mux"
)

func (s *Server) handlerCreateBeer() http.HandlerFunc {
	type request struct {
		Name   string  `json:"name"`
		Cost   float64 `json:"cost"`
		Price  float64 `json:"price"`
		Brewer string  `json:"brewer"`
		Size   float64 `json:"size"`
		Vol    float64 `json:"vol"`
	}

	type response struct {
		ID int64 `json:"id"`
	}

	isValid := func(r *request) (bool, map[string]string) {
		ee := map[string]string{}
		if r.Name == "" {
			ee["Name"] = "A Name is required"
		}

		if r.Cost < .0 {
			ee["Cost"] = "Cost must be bigger than or equal to 0"
		}

		if r.Price < .0 {
			ee["Price"] = "Price must be bigger than or equal to 0"
		}
		log.Printf("Validation errors: %v", ee)
		return len(ee) == 0, ee
	}

	return func(w http.ResponseWriter, r *http.Request) {
		d := request{}
		{
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println("can not read request body")
				s.error(w, r, http.StatusInternalServerError, "Can not read request body")
				return
			}

			if err := json.Unmarshal(b, &d); err != nil {
				log.Println("can not unmarshal request body as JSON")
				//TODO: log error
				s.error(w, r, http.StatusInternalServerError, "Can not unmarshal body as JSON")
				return
			}
		}
		log.Printf("request is: %v", d)

		if ok, ee := isValid(&d); !ok {
			s.respond(w, r, ee, http.StatusBadRequest)
			return
		}

		b := bshop.Beer{
			Product: bshop.Product{
				Name:  d.Name,
				Cost:  d.Cost,
				Price: d.Price,
			},
			Brewer: d.Brewer,
			Size:   d.Size,
			Vol:    d.Vol,
		}

		id, err := s.BeerRepo.Create(r.Context(), b)
		if err != nil {
			//TODO: handle InternalError, BadRequest, etc
			log.Println("Beer Repository returned error:", err)
			s.error(w, r, http.StatusBadRequest, fmt.Sprintf("Can not create beer: %v", err))
			return
		}
		b.ID = id

		rs := response{ID: id}
		s.respond(w, r, rs, http.StatusOK)
	}
}

func (s *Server) handlerDeleteBeer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)

		is, ok := v["id"]
		if !ok {
			log.Println("can not found id parameter for beer to delete")
			s.error(w, r, http.StatusBadRequest, `Delete beer needs the parameter "id" to be provided`)
			return
		}

		i, err := strconv.ParseInt(is, 10, 64)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Sprintf("Provided id (%v) is not a valid id", is))
			return
		}

		if err = s.BeerRepo.Delete(r.Context(), i); err != nil {
			//TODO: log error
			//TODO: handle InternalError, NotFound, etc
			s.error(w, r, http.StatusBadRequest, fmt.Sprintf("Can not delete Beer %v", i))
			return
		}
	}
}

func (s *Server) handlerReadBeer() http.HandlerFunc {
	type response struct {
		ID     int64   `json:"id"`
		Name   string  `json:"name"`
		Cost   float64 `json:"cost"`
		Price  float64 `json:"price"`
		Brewer string  `json:"brewer"`
		Size   float64 `json:"size"`
		Vol    float64 `json:"vol"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)

		is, ok := v["id"]
		if !ok {
			//TODO: log error
			log.Println("can not found id parameter for beer to read")
			s.error(w, r, http.StatusBadRequest, `Read beer needs the parameter "id" to be provided`)
			return
		}

		i, err := strconv.ParseInt(is, 10, 64)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Sprintf("Provided id (%v) is not a valid id", is))
			return
		}

		b, err := s.BeerRepo.Read(r.Context(), i)
		if err != nil {
			//TODO: log error
			//TODO: handle InternalError, NotFound, etc
			s.error(w, r, http.StatusBadRequest, fmt.Sprintf("Can not read beer %v", i))
			return
		}

		rs := response{
			ID:     b.ID,
			Name:   b.Name,
			Brewer: b.Brewer,
			Vol:    b.Vol,
			Size:   b.Size,
		}
		s.respond(w, r, rs, http.StatusOK)
	}
}

func (s *Server) handlerUpdateBeer() http.HandlerFunc {
	type request struct {
		Name   string  `json:"name"`
		Brewer string  `json:"brewer"`
		Cost   float64 `json:"cost"`
		Price  float64 `json:"price"`
		Vol    float64 `json:"vol"`
		Size   float64 `json:"size"`
	}

	isValid := func(r *request) (bool, map[string]string) {
		ee := map[string]string{}
		if r.Name == "" {
			ee["Name"] = "A Name is required"
		}

		if r.Cost < .0 {
			ee["Cost"] = "Cost must be bigger than or equal to 0"
		}

		if r.Price < .0 {
			ee["Price"] = "Price must be bigger than or equal to 0"
		}
		return len(ee) == 0, ee
	}

	return func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)

		is, ok := v["id"]
		if !ok {
			//TODO: log error
			s.error(w, r, http.StatusBadRequest, `Update beer needs the parameter "id" to be provided`)
			return
		}

		i, err := strconv.ParseInt(is, 10, 64)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, fmt.Sprintf("Provided id (%v) is not a valid id", is))
			return
		}

		d := request{}
		{
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				//TODO: log error
				s.error(w, r, http.StatusInternalServerError, "Can not read request body")
				return
			}

			if err := json.Unmarshal(b, &d); err != nil {
				//TODO: log error
				s.error(w, r, http.StatusInternalServerError, "Can not unmarshal body to JSON")
				return
			}
		}

		if ok, ee := isValid(&d); !ok {
			s.respond(w, r, ee, http.StatusBadRequest)
			return
		}

		b := bshop.Beer{
			Product: bshop.Product{
				ID:    i,
				Name:  d.Name,
				Cost:  d.Cost,
				Price: d.Price,
			},
			Brewer: d.Brewer,
			Size:   d.Size,
			Vol:    d.Vol,
		}

		if err := s.BeerRepo.Update(r.Context(), b); err != nil {
			//TODO: log error
			//TODO: handle InternalError, NotFound, etc
			s.error(w, r, http.StatusBadRequest, fmt.Sprintf("Can not read beer %v", i))
			return
		}
	}
}

func (s *Server) handlerListBeer() http.HandlerFunc {
	type beer struct {
		ID     int64   `json:"id"`
		Name   string  `json:"name"`
		Cost   float64 `json:"cost"`
		Price  float64 `json:"price"`
		Brewer string  `json:"brewer"`
		Size   float64 `json:"size"`
		Vol    float64 `json:"vol"`
	}

	type response struct {
		Beers []beer `json:"beers"`
	}

	mapResult := func(bb []bshop.Beer) []beer {
		rbb := make([]beer, len(bb))
		for i, rb := range bb {
			rbb[i] = beer{
				ID:     rb.ID,
				Name:   rb.Name,
				Cost:   rb.Cost,
				Price:  rb.Price,
				Brewer: rb.Brewer,
				Size:   rb.Size,
				Vol:    rb.Vol,
			}
		}
		return rbb
	}

	return func(w http.ResponseWriter, r *http.Request) {
		bb, err := s.BeerRepo.List(r.Context())
		if err != nil {
			//TODO: log error
			//TODO: handle InternalError, NotFound, etc
			s.error(w, r, http.StatusBadRequest, "Can not list beers ")
			return
		}

		bbr := mapResult(bb)
		s.respond(w, r, bbr, http.StatusOK)
	}
}
