Используется версия Golang 1.22.2.
Для запуска выполнить команды:
1. Поднять СУБД PostgreSQL и Golang приложение через docker-compose: docker-compose up --build
2. Создать тестового пользователя:
POST on http://127.0.0.1:8080/users/create
row JSON body:
{
	"user_id": "3",
	"username": "John",
	"team_name": "backend",
	"is_active": true
}

Реализованы эндпоинты:
* pullRequest/create
* pullRequest/merge
* team/add
* team/get
* users/setIsActive
