package rest

import (
	"net/http"

	"github.com/filariow/bshop/pkg/beer/storage"
	"github.com/filariow/bshop/pkg/beer/usecase"
)

//Server REST server structure
type controller struct {
	mux http.Handler

	createBeer usecase.CreateBeerFunc
	readBeer   usecase.ReadBeerFunc
	updateBeer usecase.UpdateBeerFunc
	deleteBeer usecase.DeleteBeerFunc
	listBeer   usecase.ListBeerFunc
}

func New(repo storage.BeerRepository, pathPrefix string) http.Handler {
	s := &controller{
		createBeer: usecase.CreateBeer(repo),
		readBeer:   usecase.ReadBeer(repo),
		updateBeer: usecase.UpdateBeer(repo),
		deleteBeer: usecase.DeleteBeer(repo),
		listBeer:   usecase.ListBeer(repo),
	}
	s.registerRoutes(pathPrefix)
	return s
}

func (s *controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
