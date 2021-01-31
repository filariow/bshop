package bshop

//go:generate mockgen -source pkg/storage/storage.go -package=mocks -destination=pkg/beers/mocks/beers_repo.go github.com/filariow/bshop/pkg/storage/storage.go BeerRepository
