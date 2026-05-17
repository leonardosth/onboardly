package models

import (
	"time"

	"github.com/google/uuid"
)

type StatusProjeto string

const (
	StatusBacklog      StatusProjeto = "Backlog"
	StatusEmAndamento  StatusProjeto = "Em_Andamento"
	StatusConcluido    StatusProjeto = "Concluido"
)

type Projeto struct {
	ID              uuid.UUID     `json:"id" db:"id"`
	ClienteID       uuid.UUID     `json:"cliente_id" db:"cliente_id"`
	AnalistaID      uuid.UUID     `json:"analista_id" db:"analista_id"`
	DataContratacao time.Time     `json:"data_contratacao" db:"data_contratacao"`
	DataAtivacao    *time.Time    `json:"data_ativacao,omitempty" db:"data_ativacao"`
	StatusAtivacao  bool          `json:"status_ativacao" db:"status_ativacao"`
	StatusProjeto   StatusProjeto `json:"status_projeto" db:"status_projeto"`
	CreatedAt       time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at" db:"updated_at"`
	DeletedAt       *time.Time    `json:"deleted_at,omitempty" db:"deleted_at"`
}

type ProjetoStatusHistorico struct {
	ID           uuid.UUID     `json:"id" db:"id"`
	ProjetoID    uuid.UUID     `json:"projeto_id" db:"projeto_id"`
	StatusAntigo *StatusProjeto `json:"status_antigo,omitempty" db:"status_antigo"`
	StatusNovo   StatusProjeto `json:"status_novo" db:"status_novo"`
	CreatedAt    time.Time     `json:"created_at" db:"created_at"`
}

type DashboardStats struct {
	TotalProjetos    int               `json:"total_projetos"`
	TotalClientes    int               `json:"total_clientes"`
	ReunioesHoje     int               `json:"reunioes_hoje"`
	PorStatus        map[string]int    `json:"por_status"`
	HistoricoMensal  []MonthlyStat     `json:"historico_mensal"`
	AtividadesRecent []RecentActivity  `json:"atividades_recentes"`
}

type MonthlyStat struct {
	Mes   string `json:"mes"`
	Total int    `json:"total"`
}

type RecentActivity struct {
	Tipo      string    `json:"tipo"` // 'Projeto', 'Cliente', 'Reuniao'
	Descricao string    `json:"descricao"`
	Status    string    `json:"status"`
	Data      time.Time `json:"data"`
}
