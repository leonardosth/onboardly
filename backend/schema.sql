-- Criar banco de dados (se não existir)
-- CREATE DATABASE IF NOT EXISTS onboardly_db;

-- Usar o banco
-- \c onboardly_db;

-- Criar tabela clientes
CREATE TABLE IF NOT EXISTS clientes (
    id UUID PRIMARY KEY,
    nome VARCHAR(150) NOT NULL,
    cnpj VARCHAR(14) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Índice para CNPJ (já único, mas para performance)
CREATE INDEX IF NOT EXISTS idx_clientes_cnpj ON clientes(cnpj);

-- Criar tabela analistas
CREATE TABLE IF NOT EXISTS analistas (
    id UUID PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Criar tabela projetos_implantacao
CREATE TABLE IF NOT EXISTS projetos_implantacao (
    id UUID PRIMARY KEY,
    cliente_id UUID NOT NULL REFERENCES clientes(id),
    analista_id UUID NOT NULL REFERENCES analistas(id),
    data_contratacao DATE NOT NULL,
    data_ativacao DATE,
    status_ativacao BOOLEAN DEFAULT FALSE,
    status_projeto VARCHAR(20) NOT NULL CHECK (status_projeto IN ('Backlog', 'Em_Andamento', 'Concluido')),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Criar tabela reunioes
CREATE TABLE IF NOT EXISTS reunioes (
    id UUID PRIMARY KEY,
    projeto_id UUID NOT NULL REFERENCES projetos_implantacao(id),
    data_agendada TIMESTAMP NOT NULL,
    status VARCHAR(20) NOT NULL CHECK (status IN ('Agendada', 'Realizada', 'Remarcada', 'No_Show')),
    observacoes TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Tabela de histórico de status de projetos
CREATE TABLE IF NOT EXISTS projeto_status_historico (
    id UUID PRIMARY KEY,
    projeto_id UUID NOT NULL REFERENCES projetos_implantacao(id),
    status_antigo VARCHAR(20),
    status_novo VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Índices adicionais
CREATE INDEX IF NOT EXISTS idx_projetos_cliente_id ON projetos_implantacao(cliente_id);
CREATE INDEX IF NOT EXISTS idx_projetos_analista_id ON projetos_implantacao(analista_id);
CREATE INDEX IF NOT EXISTS idx_projetos_data_contratacao ON projetos_implantacao(data_contratacao);
CREATE INDEX IF NOT EXISTS idx_reunioes_projeto_id ON reunioes(projeto_id);
CREATE INDEX IF NOT EXISTS idx_reunioes_data_agendada ON reunioes(data_agendada);
CREATE INDEX IF NOT EXISTS idx_status_hist_projeto_id ON projeto_status_historico(projeto_id);
