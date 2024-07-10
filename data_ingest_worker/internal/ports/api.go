package ports

import (
	"github.com/w3gop2p/elasticGrpc/data_ingest_worker/internal/application/domain"
)

type APIPort interface {
	GetData() ([]domain.Ad, error)
}
