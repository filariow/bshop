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

func (c *controller) handlerCreate{{.ent}}() http.HandlerFunc {
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
		log.Println("Handling create {{.ent | lower }} request:", r.URL.String())
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

		b := bshop.{{.ent}}{
			Product: bshop.Product{
				Name:  d.Name,
				Cost:  d.Cost,
				Price: d.Price,
			},
			Brewer: d.Brewer,
			Size:   d.Size,
			Vol:    d.Vol,
		}

		id, err := c.create{{.ent}}(r.Context(), b)
		if err != nil {
			//TODO: handle InternalError, BadRequest, etc
			log.Println("{{.ent}} Repository returned error:", err)
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Can not create {{.ent | lower }}: %v", err))
			return
		}
		b.ID = id

		rs := response{ID: id}
		helpers.Respond(w, r, http.StatusOK, rs)
	}
}

func (c *controller) handlerDelete{{.ent}}() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling delete {{.ent | lower }} request:", r.URL.String())
		is := chi.URLParam(r, "id")
		if is == "" {
			log.Println("can not found id parameter for {{.ent | lower }} to delete")
			helpers.Error(w, r, http.StatusBadRequest, `Delete {{.ent | lower }} needs the parameter "id" to be provided`)
			return
		}

		i, err := strconv.ParseUint(is, 10, 64)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Provided id (%v) is not a valid id", is))
			return
		}

		if err = c.delete{{.ent}}(r.Context(), i); err != nil {
			//TODO: log error
			//TODO: handle InternalError, NotFound, etc
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Can not delete {{.ent}} %v", i))
			return
		}
	}
}

func (c *controller) handlerRead{{.ent}}() http.HandlerFunc {
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
		log.Println("Handling read {{.ent | lower }} request:", r.URL.String())
		is := chi.URLParam(r, "id")
		if is == "" {
			//TODO: log error
			log.Println("can not found id parameter for {{.ent | lower }} to read")
			helpers.Error(w, r, http.StatusBadRequest, `Read {{.ent | lower }} needs the parameter "id" to be provided`)
			return
		}

		i, err := strconv.ParseUint(is, 10, 64)
		if err != nil {
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Provided id (%v) is not a valid id", is))
			return
		}

		b, err := c.read{{.ent}}(r.Context(), i)
		if err != nil {
			//TODO: log error
			//TODO: handle InternalError, NotFound, etc
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Can not read {{.ent | lower }} %v", i))
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

func (c *controller) handlerUpdate{{.ent}}() http.HandlerFunc {
	type request struct {
		Name   string  `json:"name"`
		Brewer string  `json:"brewer"`
		Cost   float64 `json:"cost"`
		Price  float64 `json:"price"`
		Vol    float64 `json:"vol"`
		Size   float64 `json:"size"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling update {{.ent | lower }} request:", r.URL.String())
		is := chi.URLParam(r, "id")
		if is == "" {
			//TODO: log error
			helpers.Error(w, r, http.StatusBadRequest, `Update {{.ent | lower }} needs the parameter "id" to be provided`)
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

		b := bshop.{{.ent}}{
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

		if err := c.update{{.ent}}(r.Context(), b); err != nil {
			//TODO: log error
			//TODO: handle InternalError, NotFound, etc
			helpers.Error(w, r, http.StatusBadRequest, fmt.Sprintf("Can not read {{.ent | lower }} %v", i))
			return
		}
	}
}

func (c *controller) handlerList{{.ent}}() http.HandlerFunc {
	type {{.ent | lower }} struct {
		ID     uint64  `json:"id"`
		Name   string  `json:"name"`
		Cost   float64 `json:"cost"`
		Price  float64 `json:"price"`
		Brewer string  `json:"brewer"`
		Size   float64 `json:"size"`
		Vol    float64 `json:"vol"`
	}

	type response struct {
		{{.ent}}s []{{.ent | lower }} `json:"{{.ent | lower }}s"`
	}

	mapResult := func(bb []bshop.{{.ent}}) []{{.ent | lower }} {
		rbb := make([]{{.ent | lower }}, len(bb))
		for i, rb := range bb {
			rbb[i] = {{.ent | lower }}{
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
		log.Println("Handling list {{.ent | lower }}s request:", r.URL.String())
		bb, err := c.list{{.ent}}(r.Context())
		if err != nil {
			//TODO: log error
			//TODO: handle InternalError, NotFound, etc
			helpers.Error(w, r, http.StatusBadRequest, "Can not list {{.ent | lower }}s ")
			return
		}

		bbr := mapResult(bb)
		helpers.Respond(w, r, http.StatusOK, bbr)
	}
}
