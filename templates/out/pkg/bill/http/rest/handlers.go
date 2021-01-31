package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/filariow/bshop"
	"github.com/filariow/bshop/pkg/http/rest/helpers"
	"github.com/go-chi/chi"
)

func (c *controller) handlerCreateBill() http.HandlerFunc {
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
		log.Println("Handling create bill request:", r.URL.String())
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

		b := bshop.Bill{
			Product: bshop.Product{
				Name:  d.Name,
				Cost:  d.Cost,
				Price: d.Price,
			},
			Brewer: d.Brewer,
			Size:   d.Size,
			Vol:    d.Vol,
		}

		id, err := c.createBill(r.Context(), b)
		if err != nil {
			//TODO: handle InternalError, BadRequest, etc
			log.Println("Bill Repository returned error:", err)
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Can not create bill: %v", err))
			return
		}
		b.ID = id

		rs := response{ID: id}
		helpers.Respond(w, r, http.StatusOK, rs)
	}
}

func (c *controller) handlerDeleteBill() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling delete bill request:", r.URL.String())
		is := chi.URLParam(r, "id")
		if is == "" {
			log.Println("can not found id parameter for bill to delete")
			helpers.Error(w, r, http.StatusBadRequest, `Delete bill needs the parameter "id" to be provided`)
			return
		}

		i, err := strconv.ParseUint(is, 10, 64)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Provided id (%v) is not a valid id", is))
			return
		}

		if err = c.deleteBill(r.Context(), i); err != nil {
			//TODO: log error
			//TODO: handle InternalError, NotFound, etc
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Can not delete Bill %v", i))
			return
		}
	}
}

func (c *controller) handlerReadBill() http.HandlerFunc {
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
		log.Println("Handling read bill request:", r.URL.String())
		is := chi.URLParam(r, "id")
		if is == "" {
			//TODO: log error
			log.Println("can not found id parameter for bill to read")
			helpers.Error(w, r, http.StatusBadRequest, `Read bill needs the parameter "id" to be provided`)
			return
		}

		i, err := strconv.ParseUint(is, 10, 64)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Provided id (%v) is not a valid id", is))
			return
		}

		b, err := c.readBill(r.Context(), i)
		if err != nil {
			//TODO: log error
			//TODO: handle InternalError, NotFound, etc
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Can not read bill %v", i))
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

func (c *controller) handlerUpdateBill() http.HandlerFunc {
	type request struct {
		Name   string  `json:"name"`
		Brewer string  `json:"brewer"`
		Cost   float64 `json:"cost"`
		Price  float64 `json:"price"`
		Vol    float64 `json:"vol"`
		Size   float64 `json:"size"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling update bill request:", r.URL.String())
		is := chi.URLParam(r, "id")
		if is == "" {
			//TODO: log error
			helpers.Error(w, r, http.StatusBadRequest, `Update bill needs the parameter "id" to be provided`)
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

		b := bshop.Bill{
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

		if err := c.updateBill(r.Context(), b); err != nil {
			//TODO: log error
			//TODO: handle InternalError, NotFound, etc
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Can not read bill %v", i))
			return
		}
	}
}

func (c *controller) handlerListBill() http.HandlerFunc {
	type bill struct {
		ID     uint64  `json:"id"`
		Name   string  `json:"name"`
		Cost   float64 `json:"cost"`
		Price  float64 `json:"price"`
		Brewer string  `json:"brewer"`
		Size   float64 `json:"size"`
		Vol    float64 `json:"vol"`
	}

	type response struct {
		Bills []bill `json:"bills"`
	}

	mapResult := func(bb []bshop.Bill) []bill {
		rbb := make([]bill, len(bb))
		for i, rb := range bb {
			rbb[i] = bill{
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
		log.Println("Handling list bills request:", r.URL.String())
		bb, err := c.listBill(r.Context())
		if err != nil {
			//TODO: log error
			//TODO: handle InternalError, NotFound, etc
			helpers.Error(w, r, http.StatusBadRequest, "Can not list bills ")
			return
		}

		bbr := mapResult(bb)
		helpers.Respond(w, r, http.StatusOK, bbr)
	}
}
