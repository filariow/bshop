package rest

import (
	"github.com/filariow/bshop/pkg/bill/storage"
	"github.com/filariow/bshop/pkg/bill/usecase"
	"github.com/filariow/bshop/pkg/http/rest/server"
)

//Server REST server structure
type controller struct {
	createBill usecase.CreateBillFunc
	readBill   usecase.ReadBillFunc
	updateBill usecase.UpdateBillFunc
	deleteBill usecase.DeleteBillFunc
	listBill   usecase.ListBillFunc
}

func New(repo storage.BillRepository) server.Controller {
	s := &controller{
		createBill: usecase.CreateBill(repo),
		readBill:   usecase.ReadBill(repo),
		updateBill: usecase.UpdateBill(repo),
		deleteBill: usecase.DeleteBill(repo),
		listBill:   usecase.ListBill(repo),
	}
	return s
}
