package repository

import "context"

type SettingsRepository struct {
}

func NewSettingsRepository(ctx context.Context) *SettingsRepository {
	return &SettingsRepository{}
}
