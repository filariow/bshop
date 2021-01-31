package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/filariow/bshop"
	"github.com/filariow/bshop/internal/http/rest/helpers"
	"github.com/gorilla/mux"
)

func (c *controller) handlerCreateBeer() http.HandlerFunc {
	type request struct {
		Name   string  `json:"name"`
		Cost   float64 `json:"cost"`
		Price  float64 `json:"price"`
		Brewer string  `json:"brewer"`
		Size   float64 `json:"size"`
		Vol    float64 `json:"vol"`
	}

	type response struct {
		ID uint64 `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling create beer request:", r.URL.String())
		d := request{}
		{
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Println("can not read request body")
				helpers.Error(w, r, http.StatusInternalServerError, "Can not read request body")
				return
			}

			if err := json.Unmarshal(b, &d); err != nil {
				log.Println("can not unmarshal request body as JSON")
				//TODO: log error
				helpers.Error(w, r, http.StatusInternalServerError, "Can not unmarshal body as JSON")
				return
			}
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

		id, err := c.createBeer(r.Context(), b)
		if err != nil {
			//TODO: handle InternalError, BadRequest, etc
			log.Println("Beer Repository returned error:", err)
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Can not create beer: %v", err))
			return
		}
		b.ID = id

		rs := response{ID: id}
		helpers.Respond(w, r, http.StatusOK, rs)
	}
}

func (c *controller) handlerDeleteBeer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling delete beer request:", r.URL.String())
		v := mux.Vars(r)

		is, ok := v["id"]
		if !ok {
			log.Println("can not found id parameter for beer to delete")
			helpers.Error(w, r, http.StatusBadRequest, `Delete beer needs the parameter "id" to be provided`)
			return
		}

		i, err := strconv.ParseUint(is, 10, 64)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Provided id (%v) is not a valid id", is))
			return
		}

		if err = c.deleteBeer(r.Context(), i); err != nil {
			//TODO: log error
			//TODO: handle InternalError, NotFound, etc
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Can not delete Beer %v", i))
			return
		}
	}
}

func (c *controller) handlerReadBeer() http.HandlerFunc {
	type response struct {
		ID     uint64  `json:"id"`
		Name   string  `json:"name"`
		Cost   float64 `json:"cost"`
		Price  float64 `json:"price"`
		Brewer string  `json:"brewer"`
		Size   float64 `json:"size"`
		Vol    float64 `json:"vol"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling read beer request:", r.URL.String())
		v := mux.Vars(r)

		is, ok := v["id"]
		if !ok {
			//TODO: log error
			log.Println("can not found id parameter for beer to read")
			helpers.Error(w, r, http.StatusBadRequest, `Read beer needs the parameter "id" to be provided`)
			return
		}

		i, err := strconv.ParseUint(is, 10, 64)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Provided id (%v) is not a valid id", is))
			return
		}

		b, err := c.readBeer(r.Context(), i)
		if err != nil {
			//TODO: log error
			//TODO: handle InternalError, NotFound, etc
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Can not read beer %v", i))
			return
		}

		rs := response{
			ID:     b.ID,
			Name:   b.Name,
			Brewer: b.Brewer,
			Cost:   b.Cost,
			Price:  b.Price,
			Vol:    b.Vol,
			Size:   b.Size,
		}
		helpers.Respond(w, r, http.StatusOK, rs)
	}
}

func (c *controller) handlerUpdateBeer() http.HandlerFunc {
	type request struct {
		Name   string  `json:"name"`
		Brewer string  `json:"brewer"`
		Cost   float64 `json:"cost"`
		Price  float64 `json:"price"`
		Vol    float64 `json:"vol"`
		Size   float64 `json:"size"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling update beer request:", r.URL.String())
		v := mux.Vars(r)

		is, ok := v["id"]
		if !ok {
			//TODO: log error
			helpers.Error(w, r, http.StatusBadRequest, `Update beer needs the parameter "id" to be provided`)
			return
		}

		i, err := strconv.ParseUint(is, 10, 64)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Provided id (%v) is not a valid id", is))
			return
		}

		d := request{}
		{
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				//TODO: log error
				helpers.Error(w, r, http.StatusInternalServerError, "Can not read request body")
				return
			}

			if err := json.Unmarshal(b, &d); err != nil {
				//TODO: log error
				helpers.Error(w, r, http.StatusInternalServerError, "Can not unmarshal body to JSON")
				return
			}
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

		if err := c.updateBeer(r.Context(), b); err != nil {
			//TODO: log error
			//TODO: handle InternalError, NotFound, etc
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Can not read beer %v", i))
			return
		}
	}
}

func (c *controller) handlerListBeer() http.HandlerFunc {
	type beer struct {
		ID     uint64  `json:"id"`
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
		log.Println("Handling list beers request:", r.URL.String())
		bb, err := c.listBeer(r.Context())
		if err != nil {
			//TODO: log error
			//TODO: handle InternalError, NotFound, etc
			helpers.Error(w, r, http.StatusBadRequest, "Can not list beers ")
			return
		}

		bbr := mapResult(bb)
		helpers.Respond(w, r, http.StatusOK, bbr)
	}
}
