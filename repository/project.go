package repository

import (
	"database/sql"
	"fmt"
	"github.com/alissonphp/analisador-projeto/domain/models"
	"sync"
	"time"
)

var dbMutex sync.Mutex

type IProjectRepository interface {
	Store(project models.Project) (int64, error)
}
type ProjectRepository struct {
	DB *sql.DB
}

func (p ProjectRepository) Store(project models.Project) (int64, error) {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	
	dateNow := time.Now().Format(time.RFC3339)
	res, err := p.DB.Exec("INSERT INTO projects (name, squad, identifier, source, consultation_date) VALUES (?, ?, ?, ?, ?)",
		project.Name, project.Squad, project.Identifier, project.Source, dateNow)
	if err != nil {
		return 0, fmt.Errorf("erro ao inserir projeto: %v", err)
	}
	projectID, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("erro ao obter o ID do projeto: %v", err)
	}
	return projectID, nil
}

func NewProjectRepository(db *sql.DB) IProjectRepository {
	return ProjectRepository{DB: db}
}
