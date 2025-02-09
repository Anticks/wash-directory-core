package services

import (
	"context"

	"go.uber.org/zap"
)

type CarWashService struct {
	logger *zap.SugaredLogger
	// db, client, or other dependencies as needed
}

func NewCarWashService(logger *zap.SugaredLogger) *CarWashService {
	return &CarWashService{
		logger: logger,
	}
}

func (s *CarWashService) StartWash(ctx context.Context) error {
	s.logger.Infow("Starting car wash!")
	// Possibly insert row in DB, etc.
	return nil
}

func (s *CarWashService) GetWashStatus(ctx context.Context, washID string) (string, error) {
	s.logger.Infow("Starting car wash!")
	// Possibly query DB for wash status
	return "in progress", nil
}
