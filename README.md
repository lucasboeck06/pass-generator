Gerador de Senhas Seguras com Compartilhamento Temporário
    📌 O que o sistema faz?
        Gera senhas seguras conforme critérios definidos pelo usuário (tamanho, caracteres especiais, etc.).
        Permite armazenar senhas temporariamente de forma segura.
        Gera um link temporário para acessar a senha sem precisar fazer login.
        O link expira automaticamente após um tempo definido.

    🛠 Tecnologias & Ferramentas
        Golang para a API.
        PostgreSQL ou SQLite para armazenamento das senhas.
        Criptografia AES para proteger as senhas.
        Redis para gerenciar a expiração dos links temporários.
        Gin ou Fiber para estruturar a API REST.

    ⚙️ Fluxo do sistema
        O usuário solicita a geração de uma senha.
        O sistema gera uma senha segura conforme os critérios.
        A senha é criptografada e armazenada no banco de dados.
        Um link temporário é gerado para acessar a senha.
        Se o usuário acessar o link antes da expiração, ele vê a senha.
        Após a expiração, a senha é removida automaticamente.

    🚀 Diferenciais que impressionam no portfólio
        ✅ Implementação de criptografia (mostra preocupação com segurança).
        ✅ Uso de Goroutines para deletar senhas expiradas.
        ✅ Redis para armazenamento rápido e expiração automática.
        ✅ Um CRUD funcional e seguro com API RESTful.

=========================================================================================

📍 Roadmap de Desenvolvimento
    🟢 Fase 1: Configuração do Ambiente com Docker
        🔹 Instale Docker e Docker Compose no seu sistema.
        🔹 Crie um Dockerfile para containerizar a API Golang.
        🔹 Configure um docker-compose.yml para rodar o banco de dados em um container separado.
        🔹 Use PostgreSQL como banco e exponha a porta para acessar via DBeaver.
        🔹 Teste a conexão entre os containers.

    ✅ Output esperado: API e banco rodando em containers separados, podendo ser gerenciados via DBeaver.

    🟡 Fase 2: Geração de Senhas Seguras
        🔹 Crie uma função para gerar senhas aleatórias seguras.
        🔹 Permita que o usuário defina regras personalizadas (tamanho, caracteres especiais, etc.).
        🔹 Teste diferentes abordagens para geração de senhas (math/rand vs crypto/rand).
        🔹 Crie uma API RESTful que recebe parâmetros e retorna uma senha segura.
        🔹 Crie um endpoint para testar a geração de senhas.

    ✅ Output esperado: API funcional rodando em Docker e retornando senhas seguras.

    🟠 Fase 3: Armazenamento Seguro das Senhas
        🔹 Configure GORM para conectar a API ao PostgreSQL.
        🔹 Criptografe as senhas geradas antes de armazená-las (AES ou bcrypt).
        🔹 Salve as senhas criptografadas no banco de dados.
        🔹 Teste inserções e consultas no banco usando DBeaver.

    ✅ Output esperado: Senhas sendo armazenadas de forma criptografada no PostgreSQL.

    🔴 Fase 4: Geração e Expiração de Links Temporários
        🔹 Crie um UUID único para cada senha gerada.
        🔹 Gere um link temporário (/senha/{uuid}).
        🔹 Adicione um container Redis no docker-compose.yml para gerenciar expiração.
        🔹 Armazene os links no Redis com tempo de expiração.
        🔹 Se o usuário acessar o link antes de expirar, descriptografe e exiba a senha.
        🔹 Após expirar, o link deixa de funcionar e a senha é removida.

    ✅ Output esperado: Links temporários que expiram automaticamente após um tempo definido, usando Redis.

    🟣 Fase 5: Expansão & Segurança
        🔹 Adicione rate limit (evitar abusos de requisições na API).
        🔹 Implemente JWT (se quiser salvar senhas associadas a usuários).
        🔹 Registre logs para monitorar acessos às senhas.
        🔹 Adicione testes unitários para validar a geração e recuperação de senhas.

    ✅ Output esperado: API segura, escalável e pronta para produção.

    📌 Tecnologias Utilizadas
        ✅ Golang - Backend
        ✅ Docker & Docker Compose - Containerização
        ✅ PostgreSQL - Banco de dados (acessado via DBeaver)
        ✅ Redis - Expiração de links temporários
        ✅ Gin ou Fiber - Framework para API REST
        ✅ GORM - ORM para interação com o banco
        ✅ AES ou bcrypt - Criptografia das senhas

    🛠 Setup do Ambiente com Docker
        1️⃣ Crie o Dockerfile para a API
        2️⃣ Configure o docker-compose.yml para rodar PostgreSQL e Redis
        3️⃣ Acesse o banco via DBeaver
        4️⃣ Inicie os containers e rode a API


Arquitetura:

/projeto-senhas/
│── /backend/                  # Código-fonte do backend
│    ├── /cmd/                 # Ponto de entrada da aplicação
│    │    ├── main.go          # Arquivo principal da API
│    │    ├── server.go        # Inicialização do servidor e rotas
│    ├── /controllers/         # Controladores (lógica de negócio)
│    │    ├── password.go      # Controle de endpoint
│    │    ├── auth.go          # Futuro: autenticação
│    ├── /models/              # Modelos de banco de dados
│    │    ├── password.go      # Estrutura e métodos do modelo Password
│    ├── /services/            # Regras de negócio e segurança
│    │    ├── crypto.go        # Criptografia AES-256
│    │    ├── password.go      # Geração e validação de senhas
│    ├── /routes/              # Definição de rotas da API
│    │    ├── routes.go        # Agrupa todas as rotas
│    ├── /db/                  # Conexão e migrações do banco
│    │    ├── database.go      # Inicialização do PostgreSQL
│    │    ├── migrations.sql   # Script de migração do banco
│── /frontend/                 # Código-fonte do frontend
│    ├── index.html            # Página inicial
│    ├── style.css             # Estilização
│    ├── script.js             # Requisições para a API
│── /deploy/                   # Infraestrutura
│    ├── docker-compose.yml    # Configuração do PostgreSQL e Go
│    ├── .env                  # Variáveis de ambiente
│── README.md                  # Documentação do projeto