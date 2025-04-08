# Analisador de Projeto

[![Go Report Card](https://goreportcard.com/badge/github.com/alissonphp/analisador-projeto)](https://goreportcard.com/report/github.com/alissonphp/analisador-projeto)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/alissonphp/analisador-projeto)](go.mod)

## Descrição

O **Analisador de Projeto** é uma aplicação em Go que coleta métricas de projetos utilizando a API de análise de código. Ele processa um arquivo JSON contendo informações dos projetos, consulta a API para obter métricas e armazena os dados em um repositório.

## Funcionalidades

- Leitura de um arquivo JSON com informações dos projetos.
- Consulta a uma API para obter métricas como bugs, vulnerabilidades, cobertura de código, entre outras.
- Armazenamento das métricas e informações dos projetos em repositórios.

## Estrutura do Projeto

- **`domain/models`**: Contém os modelos de dados utilizados no projeto.
- **`domain/usecase`**: Implementa os casos de uso, como a coleta de métricas.
- **`infrastructure/http`**: Contém a implementação do cliente HTTP para comunicação com a API.
- **`repository`**: Define e implementa os repositórios para armazenamento de dados.

## Pré-requisitos

- **Go**: Versão 1.20 ou superior.
- **Arquivo de entrada**: Um arquivo JSON contendo os projetos a serem analisados (exemplo: `input-example.json`).

## Instalação

1. Clone o repositório:
   ```bash
   git clone https://github.com/alissonphp/analisador-projeto.git
   cd analisador-projeto
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

3. Configure o arquivo de entrada (`input.json`) com os projetos a serem analisados.

## Uso

Execute o comando abaixo para iniciar a coleta de métricas:

```bash
go run main.go <caminho-do-arquivo-json>
```

Exemplo:

```bash
go run main.go input.json
```

## Exemplo de Arquivo de Entrada

```json
[
  {
    "name": "test",
    "squad": "my-squad",
    "identifier": "service-project",
    "source": "cloud"
  }
]
```

## Contribuição

Contribuições são bem-vindas! Siga os passos abaixo:

1. Faça um fork do repositório.
2. Crie uma branch para sua feature ou correção de bug:
   ```bash
   git checkout -b minha-feature
   ```
3. Faça o commit das suas alterações:
   ```bash
   git commit -m "Minha nova feature"
   ```
4. Envie para o repositório remoto:
   ```bash
   git push origin minha-feature
   ```
5. Abra um Pull Request.

## Licença

Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.