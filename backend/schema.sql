-- Criar banco de dados (se não existir)
-- CREATE DATABASE IF NOT EXISTS onboardly_db;

-- Usar o banco
-- \c onboardly_db;

-- =============================================================================
-- SCRIPTS DE MIGRAÇÃO (Para quem já tem o banco antigo e quer atualizar)
-- =============================================================================

/* 
PASSO A PASSO DA MIGRAÇÃO:
1. Adicionar cargo 'Analista' para quem está na tabela analistas e não está na usuarios.
2. Atualizar a tabela de projetos para apontar para a tabela usuarios.
3. Remover a tabela analistas antiga.
*/

/*
-- 1. Migrar analistas para a tabela de usuários (caso ainda não existam lá)
-- Nota: Como analistas não tinham senha, definimos um hash padrão ou aleatório para bloqueio inicial
INSERT INTO usuarios (id, nome, email, senha_hash, cargo, created_at, updated_at)
SELECT id, nome, email, '$2a$10$YourDefaultHashHere', 'Analista', created_at, updated_at
FROM analistas
ON CONFLICT (email) DO UPDATE SET cargo = 'Analista';

-- 2. Atualizar Constraints de Projetos
-- Primeiro, remover a FK antiga
ALTER TABLE projetos_implantacao DROP CONSTRAINT IF EXISTS projetos_implantacao_analista_id_fkey;

-- Segundo, garantir que todos os analistas de projetos existam na tabela usuarios
-- (Já garantido pelo INSERT acima)

-- Terceiro, adicionar a nova FK apontando para usuarios
ALTER TABLE projetos_implantacao 
ADD CONSTRAINT projetos_implantacao_analista_id_fkey 
FOREIGN KEY (analista_id) REFERENCES usuarios(id);

-- 3. Remover tabela legada (OPCIONAL - FAÇA BACKUP ANTES)
-- DROP TABLE analistas;
-- =============================================================================
*/

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

-- Criar tabela usuarios para autenticação e gestão
CREATE TABLE IF NOT EXISTS usuarios (
    id UUID PRIMARY KEY,
    nome VARCHAR(150) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    senha_hash VARCHAR(255) NOT NULL,
    cargo VARCHAR(50) NOT NULL, -- 'Admin', 'Analista'
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Índice para email
CREATE INDEX IF NOT EXISTS idx_usuarios_email ON usuarios(email);

-- Criar tabela projetos_implantacao
CREATE TABLE IF NOT EXISTS projetos_implantacao (
    id UUID PRIMARY KEY,
    cliente_id UUID NOT NULL REFERENCES clientes(id),
    analista_id UUID NOT NULL REFERENCES usuarios(id),
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
