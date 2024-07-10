package grpc

import (
	"context"
	"fmt"
	"github.com/w3gop2p/elasticGrpc-proto/golang/data_ingest_worker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (a Adapter) GetData(ctx context.Context, empty *data_ingest_worker.Empty) (*data_ingest_worker.GetDataResponse, error) {
	log.Println("Data Retrieving...")
	result, err := a.api.GetData()
	if err != nil {
		return &data_ingest_worker.GetDataResponse{}, status.New(codes.Internal, fmt.Sprintf("failed to retrieve. %v ", err)).Err()
	}
	// Convert []domain.Ad to []*data_ingest_worker.Ad
	var ads []*data_ingest_worker.Ad
	for _, ad := range result {
		ads = append(ads, &data_ingest_worker.Ad{
			XId: ad.ID,
			Categories: &data_ingest_worker.Category{
				Subcategory: ad.Categories.Subcategory,
			},
			Title: &data_ingest_worker.Title{
				Ro: ad.Title.Ro,
				Ru: ad.Title.Ru,
			},
			Type:   ad.Type,
			Posted: ad.Posted,
		})
	}
	return &data_ingest_worker.GetDataResponse{Ads: ads}, nil
}
