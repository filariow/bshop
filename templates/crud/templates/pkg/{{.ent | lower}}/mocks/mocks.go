package mocks

//go:generate mockgen -source pkg/storage/storage.go -package=mocks -destination=pkg/{{.ent | lower }}s/mocks/{{.ent | lower }}s_repo.go github.com/filariow/bshop/pkg/{{.ent | lower }}/storage/storage.go {{.ent}}Repository
