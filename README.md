O objetivo deste teste é avaliar seu desempenho em desenvolver uma solução de integração entre sistemas.

O problema consiste em receber 1 ou mais contatos de celulares através de uma API Rest e adicioná-los ao banco de dados do cliente Macapá ou do cliente VareJão.

Fluxo de Ações
- A API receberá um JSON via POST contendo o nome e celular;
- O cliente deverá estar autenticado para inserir o contato na base
- O contato deverá ser inserido no banco de dados do cliente seguindo as regras de cada cliente

Especificações da API:
- A autenticação será através de um token JWT no Authorization Header
- Cada cliente tem 1 uma chave única
- A lista de contatos que será inserido em cada cliente está no arquivo contato.json

Especificações do Cliente Macapá:
- Banco de dados Mysql
- Formato do Nome é somente maiúsculas
- O formato de telefone segue o padrão +55 (41) 93030-6905
- Em anexo está o sql de criação da tabela

Especificações do Cliente VareJão:
- Banco de dados Postgresql
- Formato do Nome é livre
- O formato de telefone segue o padrão 554130306905
- Em anexo está o sql de criação da tabela

A criação de um ambiente de testes usando Docker para simular o banco de dados do cliente é altamente recomendada. A solução poderá ser desenvolvida em Golang. Fique livre para desenhar a solução da maneira que achar mais conveniente e supor qualquer cenário que não foi abordado nas especificações acima. Se, por qualquer motivo, você não consiga completar este teste, recomendamos que nos encaminhe o que foi desenvolvido de qualquer maneira. A falta de cumprimento de alguns dos requisitos aqui descritos não implica necessariamente na desconsideração do candidato.


### Para baixar as requisições ja criadas 

[![Run in Insomnia}](https://insomnia.rest/images/run.svg)](https://insomnia.rest/run/?label=IntegrationAPI&uri=https%3A%2F%2Fgithub.com%2Fdiiineeei%2FIntegracaoJson%2Fblob%2Fmain%2Ftestdata%2FIntegrationAPIV1.json)

#### Step 1 Cadastro de Usuario 
```
curl --request POST \
    --url http://localhost:8080/api/user/register \
    --header 'Content-Type: application/json' \
    --data '{
    "name": "Rodinei Rodrigo de Lima",
    "email": "rodinei@gmail.com",
    "password": "123465789",
    "company": "macapa"
}'
```

#### Step 2 Geracao de token JWT
```
curl --request POST \
    --url http://localhost:8080/api/token \
    --header 'Content-Type: application/json' \
    --data '{
      "email": "rodinei@gmail.com",
      "password": "123465789"
}'
```

#### Step 3 Cadastro de contato
```
curl --request POST \
    --url http://localhost:8080/api/secured/contact/register \
    --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzIjoiUm9kaW5laSBSb2RyaWdvIGRlIExpbWEiLCJlbWFpbCI6InJvZGluZWlAZ21haWwuY29tIiwiY29tcGFueSI6Im1hY2FwYSIsImV4cCI6MTY4NTMxMzQ4M30.MZKjx53V2En5Oqtb6itpj7TeLGZkG0SC8bJhtOJOvG8' \
    --header 'Content-Type: application/json' \
    --data '{
    	"name": "Rodinei Lima",
    	"cellphone": "11966500000"
}'
```

#### Step 4 Cadastro de lista de contatos
```
curl --request POST \
  --url http://localhost:8080/api/secured/contact/register/list \
  --header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzIjoiUm9kaW5laSBSb2RyaWdvIGRlIExpbWEiLCJlbWFpbCI6InJvZGluZWlAZ21haWwuY29tIiwiY29tcGFueSI6Im1hY2FwYSIsImV4cCI6MTY4NTMxNTU2N30.m2ghoc0Wb6MjmH8P-5fntg_O371BQ-_W1k3LuAO3UtI' \
  --header 'Content-Type: application/json' \
  --data '{
	"contacts": [
		{
			"name": "Marina Rodrigues",
			"cellphone": "5541996941919"
		},
		{
			"name": "Nicolas Rodrigues",
			"cellphone": "5541954122723"
		},
		{
			"name": "Davi Lucca Rocha",
			"cellphone": "5541979210400"
		},
		{
			"name": "Lucas Barros",
			"cellphone": "5541944061868"
		},
		{
			"name": "Lucca Lima",
			"cellphone": "5541908727427"
		},
		{
			"name": "Benjamin Sales",
			"cellphone": "5541914106998"
		},
		{
			"name": "Ana Laura Pereira",
			"cellphone": "5541919883324"
		},
		{
			"name": "Henrique da Conceição",
			"cellphone": "5541936250777"
		},
		{
			"name": "Ana Clara Porto",
			"cellphone": "5541988315943"
		},
		{
			"name": "Ryan Souza",
			"cellphone": "5541937901945"
		},
		{
			"name": "Nina da Cruz",
			"cellphone": "5541930283480"
		},
		{
			"name": "Esther Costa",
			"cellphone": "5541976620331"
		},
		{
			"name": "Felipe Ribeiro",
			"cellphone": "5541987077478"
		},
		{
			"name": "Sra. Amanda Cavalcanti",
			"cellphone": "5541905979333"
		},
		{
			"name": "Bruno Farias",
			"cellphone": "5541945698108"
		},
		{
			"name": "Sra. Mirella Aragão",
			"cellphone": "5541932294266"
		},
		{
			"name": "Nicole Ramos",
			"cellphone": "5541972713343"
		},
		{
			"name": "Augusto Novaes",
			"cellphone": "5541934105240"
		},
		{
			"name": "Vitor Hugo Cunha",
			"cellphone": "5541981870128"
		},
		{
			"name": "Luiz Miguel Monteiro",
			"cellphone": "5541994140336"
		},
		{
			"name": "Alícia Santos",
			"cellphone": "5541978861534"
		},
		{
			"name": "Sr. Marcos Vinicius Duarte",
			"cellphone": "5541979767374"
		},
		{
			"name": "Luiz Gustavo da Cunha",
			"cellphone": "5541971153240"
		},
		{
			"name": "Murilo Moraes",
			"cellphone": "5541971140196"
		},
		{
			"name": "Julia Costela",
			"cellphone": "5541938133697"
		},
		{
			"name": "André Cardoso",
			"cellphone": "5541925935033"
		},
		{
			"name": "Luiza Campos",
			"cellphone": "5541942303433"
		},
		{
			"name": "Nathan da Mata",
			"cellphone": "5541971551888"
		},
		{
			"name": "Nina Sales",
			"cellphone": "5541952206418"
		},
		{
			"name": "Natália Caldeira",
			"cellphone": "5541977897055"
		},
		{
			"name": "Thiago Cardoso",
			"cellphone": "5541928464916"
		},
		{
			"name": "Yago Ferreira",
			"cellphone": "5541962241042"
		},
		{
			"name": "Isabel Moura",
			"cellphone": "5541917433481"
		},
		{
			"name": "Francisco Nogueira",
			"cellphone": "5541916159692"
		},
		{
			"name": "André Monteiro",
			"cellphone": "5541987830622"
		},
		{
			"name": "Ana Julia Barbosa",
			"cellphone": "5541965101052"
		},
		{
			"name": "Emanuella Vieira",
			"cellphone": "5541937084511"
		},
		{
			"name": "Rebeca Correia",
			"cellphone": "5541944003092"
		},
		{
			"name": "Luiza Lopes",
			"cellphone": "5541950648087"
		},
		{
			"name": "Sr. Vitor Hugo Jesus",
			"cellphone": "5541966777904"
		},
		{
			"name": "Bárbara da Rosa",
			"cellphone": "5541984242583"
		},
		{
			"name": "Fernando Melo",
			"cellphone": "5541993278935"
		},
		{
			"name": "Melissa Pereira",
			"cellphone": "5541961982198"
		},
		{
			"name": "Luigi Melo",
			"cellphone": "5541937316914"
		},
		{
			"name": "Davi Lucas Silva",
			"cellphone": "5541919590454"
		},
		{
			"name": "Sr. Marcos Vinicius Alves",
			"cellphone": "5541940281564"
		},
		{
			"name": "Igor da Rocha",
			"cellphone": "5541981148730"
		},
		{
			"name": "Noah Ramos",
			"cellphone": "5541943467544"
		},
		{
			"name": "Sofia Pinto",
			"cellphone": "5541942446391"
		},
		{
			"name": "Joaquim Vieira",
			"cellphone": "5541988403129"
		}
	]
}'
```
