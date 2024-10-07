Observações:
1) Certifique-se de que as portas 6379 e 3306 estejam habilitadas no servidor docker 
2) Faça o Upload de todos os arquivos para o diretório escolhido no servidor docker
3) Certifique-se de aplicar os comandos dentro do diretório escolhido: cd [diretório]

Comandos:
1) Subir Docker: docker-compose up -d --build --force-recreate
2) Acessar o BD: docker-compose exec db bash
3) Acessar o DataBase: mysql -uroot -p dbmysql
                       Pass: root
4) Criar a tabela: Create table metrics (id varchar(25), dthr timestamp default current_timestamp, regtransp varchar(15), company varchar(255), price float);
5) Verificar a Tabela: Select * from metrics;
