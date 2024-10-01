package postgres

import (
	"context"
	"time"

	"gorm.io/gorm/clause"

	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/model"
	"github.com/TimeleapLabs/unchained/internal/repository"
	"github.com/TimeleapLabs/unchained/internal/transport/database"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

type EventLogRepo struct {
	client database.Database
}

func (r EventLogRepo) Find(ctx context.Context, block uint64, hash []byte, index uint64) ([]model.EventLog, error) {
	currentRecords := []model.EventLog{}
	err := r.client.
		GetConnection().
		WithContext(ctx).
		Table("event_log").
		Where("block", block).
		Where("transaction", hash).
		Where("index", index).
		Preload("Signers").
		Find(&currentRecords)

	if err != nil {
		utils.Logger.With("err", err).Error("Cant fetch event log records from database")
		return nil, consts.ErrInternalError
	}

	return currentRecords, nil
}

func (r EventLogRepo) Upsert(ctx context.Context, data model.EventLog) error {
	err := r.client.
		GetConnection().
		WithContext(ctx).
		Table("event_log").
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "block"}, {Name: "transaction"}, {Name: "index"}},
			UpdateAll: true,
		}).
		Create(&model.DataFrame{
			Hash:      nil,
			Timestamp: time.Now(),
			Data:      data,
		})

	if err != nil {
		utils.Logger.With("err", err).Error("Cant upsert event log record to database")
		return consts.ErrInternalError
	}

	return nil
}

func NewEventLog(client database.Database) repository.EventLog {
	return &EventLogRepo{
		client: client,
	}
}
