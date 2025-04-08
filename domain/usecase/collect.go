package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/alissonphp/analisador-projeto/domain/models"
	adapter "github.com/alissonphp/analisador-projeto/infrastructure/http"
	"github.com/alissonphp/analisador-projeto/repository"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

type ICollectUseCase interface {
	GetProjectMetrics(filePath string) error
}

type CollectUseCaseImpl struct {
	projectRepo repository.IProjectRepository
	measureRepo repository.IMeasureRepository
}

func (c CollectUseCaseImpl) GetProjectMetrics(filePath string) error {
	data, err := readFile(filePath)
	if err != nil {
		return err
	}
	projects := readProjects(data)

	var wg sync.WaitGroup
	errChan := make(chan error, len(projects))

	for _, project := range projects {
		wg.Add(1)

		go func(project models.Project) {
			defer wg.Done()

			apiCall := project.GetApiCall()
			url := fmt.Sprintf("%sapi/measures/component?component=%s&metricKeys=bugs,vulnerabilities,code_smells,coverage,complexity,line_coverage,duplicated_blocks,sqale_index,sqale_debt_ratio,sqale_rating,duplicated_lines,duplicated_lines_density,cognitive_complexity,alert_status,security_rating,reliability_rating,ncloc,functions,branch_coverage,uncovered_lines,test_success_density,reliability_rating,reliability_remediation_effort", apiCall.Host, project.Identifier)
			log.Printf("Consultando API para o projeto %s na URL: %s", project.Name, url)
			client := adapter.NewHttpClient(url, "GET")
			client.Init()
			client.SetHeader("Authorization", "Basic "+apiCall.Key)
			resp, err := client.Execute()
			if err != nil {
				errChan <- err
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				log.Printf("Resposta da API não OK para %s: %s", project.Name, resp.Status)
				return
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("Erro ao ler o corpo da resposta: %v", err)
				return
			}
			var apiResp models.APIResponse
			err = json.Unmarshal(body, &apiResp)
			if err != nil {
				log.Printf("Erro ao ler o corpo da resposta: %v", err)
				return
			}

			id, err := c.projectRepo.Store(project)
			if err != nil {
				log.Printf("Erro ao armazenar dados para %s: %v", project.Name, err)
				return
			}
			for _, measure := range apiResp.Component.Measures {
				err = c.measureRepo.Store(measure, id)
				if err != nil {
					errChan <- fmt.Errorf("erro ao inserir métrica %s: %v", measure.Metric, err)
					return
				}
			}
		}(project)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	return nil
}

func NewCollectUseCase(projectRepo repository.IProjectRepository, measureRepo repository.IMeasureRepository) ICollectUseCase {
	return CollectUseCaseImpl{
		projectRepo: projectRepo,
		measureRepo: measureRepo,
	}
}

func readFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}
	return bytes, nil
}

func readProjects(data []byte) []models.Project {
	var projects []models.Project
	err := json.Unmarshal(data, &projects)
	if err != nil {
		log.Fatalf("Erro ao parsear o arquivo JSON para os projetos: %v", err)
	}
	return projects
}
