package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/filariow/bshop"
)

func (s *Server) handlerCreateBeer() http.HandlerFunc {
	type request struct {
		Name  string  `json:"name"`
		Cost  float64 `json:"cost"`
		Price float64 `json:"price"`
	}

	type response struct {
		ID int64 `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		d := request{}
		{
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				//TODO: log error
				s.error(w, r, http.StatusInternalServerError, "Can not read request body")
				return
			}

			if err := json.Unmarshal(b, &r); err != nil {
				//TODO: log error
				s.error(w, r, http.StatusInternalServerError, "Can not unmarshal body to JSON")
				return
			}
		}

		b := bshop.Beer{
			Product: bshop.Product{
				Name:  d.Name,
				Cost:  d.Cost,
				Price: d.Price,
			},
		}

		id, err := s.BeerRepo.CreateBeer(r.Context(), b)
		if err != nil {
			//TODO: log error
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
	}
}

func (s *Server) handlerReadBeer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (s *Server) handlerUpdateBeer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (s *Server) handlerListBeer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
