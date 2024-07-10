package dataworker

import (
	"context"
	"github.com/w3gop2p/elasticGrpc-proto/golang/data_ingest_worker"
	"github.com/w3gop2p/elasticGrpc/data_store_service/internal/application/domain"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Adapter struct {
	dataworker data_ingest_worker.RetrieveDataClient
}

func NewAdapter(dataWorkerServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	conn, err := grpc.Dial(dataWorkerServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	client := data_ingest_worker.NewRetrieveDataClient(conn)
	return &Adapter{dataworker: client}, nil
}

func (a *Adapter) GetDataFromWorker(ctx context.Context) ([]domain.Adv, error) {
	data, err := a.dataworker.GetData(ctx, &data_ingest_worker.Empty{})
	if err != nil {
		log.Printf("error getting data from dataworker: %v", err)
	}
	var domainAds []domain.Adv
	for _, ad := range data.Ads {
		domainAd := domain.Adv{
			ID: ad.XId,
			Categories: domain.Category{
				Subcategory: ad.Categories.Subcategory,
			},
			Title: domain.Title{
				Ro: ad.Title.Ro,
				Ru: ad.Title.Ru,
			},
			Type:   ad.Type,
			Posted: ad.Posted,
		}
		domainAds = append(domainAds, domainAd)
	}
	return domainAds, nil
}
