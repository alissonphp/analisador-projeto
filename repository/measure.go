package repository

import (
	"database/sql"
	"fmt"
	"github.com/alissonphp/analisador-projeto/domain/models"
)

type IMeasureRepository interface {
	Store(measure models.Measure, projectID int64) error
}
type MeasureRepository struct {
	DB *sql.DB
}

func (p MeasureRepository) Store(measure models.Measure, projectID int64) error {
	var bestValue interface{}
	if measure.BestValue != nil {
		bestValue = *measure.BestValue
	} else {
		bestValue = nil
	}
	_, err := p.DB.Exec("INSERT INTO metrics (project_id, metric, value, best_value) VALUES (?, ?, ?, ?)",
		projectID, measure.Metric, measure.Value, bestValue)
	if err != nil {
		return fmt.Errorf("erro ao inserir métrica %s: %v", measure.Metric, err)
	}
	return nil
}

func NewMeasureRepository(db *sql.DB) IMeasureRepository {
	return MeasureRepository{DB: db}
}
