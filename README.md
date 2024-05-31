
# Sistema de Tarefas Automáticas

## Descrição

O **Sistema de Tarefas Automáticas** é um projeto que tem como objetivo executar uma URL em um horário marcado diariamente, recolher as informações de data e hora, e salvar esses dados em um banco de dados. O sistema permite a configuração do horário e dos dias da semana em que as tarefas serão executadas, assim como a URL do site a ser acessado. Além disso, o sistema envia e-mails para os usuários cadastrados e fornece uma interface web para visualização de relatórios e gerenciamento da base de dados.

## Funcionalidades

- **Configuração de Horário e Dias da Semana**: Permite definir o horário e os dias da semana em que a URL será acessada.
- **Definição da URL**: Permite informar a URL do site a ser acessado.
- **Salvar Informações no Banco de Dados**: Armazena as informações de data e hora obtidas da URL no banco de dados.
- **Envio de E-mails**: Envia e-mails para os usuários cadastrados com as informações recolhidas.
- **Interface Web para Relatórios**: Interface web que permite aos usuários visualizar relatórios das informações coletadas.
- **Botão para Limpar a Base de Dados**: Funcionalidade para limpar os dados armazenados no banco de dados.
- **Imagem Docker**: Disponibiliza uma imagem Docker do sistema no DockerHub.

## Tecnologias Utilizadas

- **Backend**: Golang
- **Frontend**: Next.js
- **Banco de Dados**: PostgreSQL (ou outro banco de dados de sua escolha)
- **Servidor de E-mail**: Configuração de servidor SMTP para envio de e-mails
- **Docker**: Para containerização da aplicação

## Como Utilizar

### Pré-requisitos

- Docker e Docker Compose instalados
- Configuração de servidor SMTP para envio de e-mails

### Passos para Execução

1. **Clone o Repositório**

   ```bash
   git clone <URL_DO_REPOSITORIO>
   cd sistema-de-tarefas-automaticas
   ```
2. **Configuração do Ambiente**
   Crie um arquivo `.env` com as seguintes variáveis:

   ```
   DB_HOST=<host_do_banco_de_dados>
   DB_USER=<usuario_do_banco_de_dados>
   DB_PASSWORD=<senha_do_banco_de_dados>
   DB_NAME=<nome_do_banco_de_dados>
   SMTP_HOST=<host_do_servidor_smtp>
   SMTP_PORT=<porta_do_servidor_smtp>
   SMTP_USER=<usuario_smtp>
   SMTP_PASSWORD=<senha_smtp>
   ```
3. **Build e Execução com Docker**

   ```bash
   docker-compose up --build
   ```
4. **Acessar a Interface Web**
   Acesse `http://localhost:3000` para utilizar a interface web do sistema.

### Configuração do Sistema

Na interface web, você pode:

- Configurar a URL a ser acessada.
- Definir o horário e os dias da semana para execução das tarefas.
- Visualizar os relatórios das informações coletadas.
- Limpar a base de dados através do botão "Limpar Base de Dados".

### Publicação da Imagem Docker

A imagem Docker do sistema está disponível no DockerHub e pode ser obtida com o comando:

```bash
docker pull <USUARIO_DO_DOCKERHUB>/sistema-de-tarefas-automaticas
```

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests.

## Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).

## Contato

Para mais informações, entre em contato com [seu-email@example.com].

---

Espero que este README atenda às suas necessidades. Caso tenha mais alguma dúvida ou necessidade, sinta-se à vontade para perguntar!
