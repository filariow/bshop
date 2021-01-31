package mocks

//go:generate mockgen -source pkg/storage/storage.go -package=mocks -destination=pkg/bills/mocks/bills_repo.go github.com/filariow/bshop/pkg/bill/storage/storage.go BillRepository
