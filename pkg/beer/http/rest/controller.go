package rest

import (
	"github.com/filariow/bshop/pkg/beer/storage"
	"github.com/filariow/bshop/pkg/beer/usecase"
	"github.com/filariow/bshop/pkg/http/rest/server"
)

//Server REST server structure
type controller struct {
	createBeer usecase.CreateBeerFunc
	readBeer   usecase.ReadBeerFunc
	updateBeer usecase.UpdateBeerFunc
	deleteBeer usecase.DeleteBeerFunc
	listBeer   usecase.ListBeerFunc
}

func New(repo storage.BeerRepository) server.Controller {
	s := &controller{
		createBeer: usecase.CreateBeer(repo),
		readBeer:   usecase.ReadBeer(repo),
		updateBeer: usecase.UpdateBeer(repo),
		deleteBeer: usecase.DeleteBeer(repo),
		listBeer:   usecase.ListBeer(repo),
	}
	return s
}
