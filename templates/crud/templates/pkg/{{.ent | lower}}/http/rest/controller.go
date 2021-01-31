package rest

import (
	"github.com/filariow/bshop/pkg/{{.ent | lower }}/storage"
	"github.com/filariow/bshop/pkg/{{.ent | lower }}/usecase"
	"github.com/filariow/bshop/pkg/http/rest/server"
)

//Server REST server structure
type controller struct {
	create{{.ent}} usecase.Create{{.ent}}Func
	read{{.ent}}   usecase.Read{{.ent}}Func
	update{{.ent}} usecase.Update{{.ent}}Func
	delete{{.ent}} usecase.Delete{{.ent}}Func
	list{{.ent}}   usecase.List{{.ent}}Func
}

func New(repo storage.{{.ent}}Repository) server.Controller {
	s := &controller{
		create{{.ent}}: usecase.Create{{.ent}}(repo),
		read{{.ent}}:   usecase.Read{{.ent}}(repo),
		update{{.ent}}: usecase.Update{{.ent}}(repo),
		delete{{.ent}}: usecase.Delete{{.ent}}(repo),
		list{{.ent}}:   usecase.List{{.ent}}(repo),
	}
	return s
}
