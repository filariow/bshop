package rest

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/filariow/bshop"
)

func (s *Server) handlerCreateBill() http.HandlerFunc {
	type request struct {
		name string `json:"name"`
	}

	type response struct {
		id uint64 `json:"id"`
	}

	isValid := func(r *request) (bool, map[string]string) {
		ee := map[string]string{}
		if r.Name == "" {
			ee["Name"] = "A name is required"
		}
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
				s.error(w, r, http.StatusInternalServerError, "Can not unmarshal body as JSON")
				return
			}
		}
		if ok, ee := isValid(&d); !ok {
			s.respond(w, r, ee, http.StatusBadRequest)
			return
		}

		b := bshop.Bill{
			Name:      d.name,
			CreatedAt: time.Now(),
		}

	}
}
