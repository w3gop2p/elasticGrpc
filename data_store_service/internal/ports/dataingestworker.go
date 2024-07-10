package ports

import (
	"context"
	"github.com/w3gop2p/elasticGrpc/data_store_service/internal/application/domain"
)

type DataIngestWorkerPort interface {
	GetDataFromWorker(context.Context) ([]domain.Adv, error)
}
