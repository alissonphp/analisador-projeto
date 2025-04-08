package main

import (
	"github.com/alissonphp/analisador-projeto/domain/usecase"
	"github.com/alissonphp/analisador-projeto/infrastructure/db"
	"github.com/alissonphp/analisador-projeto/repository"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Initialize database
	DB := db.InitSQLiteDB("infrastructure/db/report.db")

	// Initialize repositories
	projectRepo := repository.NewProjectRepository(DB)
	measureRepo := repository.NewMeasureRepository(DB)

	// Initialize use case
	collectUseCase := usecase.NewCollectUseCase(projectRepo, measureRepo)

	// Get project metrics
	err := collectUseCase.GetProjectMetrics("input.json")
	if err != nil {
		panic(err)
	}
}
