Нужно будет создать базу и сделать миграцию
Команды миграции -
up:  migrate -path ./schema -database 'DB_URL' up
down: migrate -path ./schema -database 'DB_URL' down
Затем можете запустить команду docker-compose build, затем docker-compose up
