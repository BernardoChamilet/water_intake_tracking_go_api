# Recorte API Pro Health - controle de água.

Este é um pequeno recorte da API Pro Health (API completa de dieta e treino por mim desenvolvida).

## Estrutura da api:
* Routes: definição de rotas (nome, método http e função)
* Controllers: funções das rotas (recebe requisição e chama outros pacotes para enviar resposta)
* Repositories: funções de interação com banco de dados
* Models: classes para validar dados
* Middlewares: funções a serem executadas entre a requisição e chamar funções das rotas de fato
* Auth: funções que envolvem autorização/jwt
* Config: pacote de inicialização de variáveis de ambiente (.env)
* Database: abertura da conexão com banco de dados
* Response: formatação de respostas a serem devolvidas
* Secutiry: funções de segurança/hash
* Utils: funções de utilidades diversas que não se encaixam em nenhum do pacotes