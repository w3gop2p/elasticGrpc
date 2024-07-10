package api

import (
	"github.com/w3gop2p/elasticGrpc/data_ingest_worker/internal/application/domain"
	"github.com/w3gop2p/elasticGrpc/data_ingest_worker/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) GetData() ([]domain.Ad, error) {
	dataAds, err := a.db.Get()
	if err != nil {
		return []domain.Ad{}, err
	}
	return dataAds, nil
}
