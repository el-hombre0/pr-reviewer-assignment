Используется версия Golang 1.22.2.
Для запуска выполнить команды:
1. Поднять СУБД PostgreSQL (объединить приложение и БД в одном docker-compose не получилось): docker-compose up -d
2. Установить зависимости: go mod download
3. Запустить приложение: go run main.go
4. Создать тестового пользователя:
POST on http://127.0.0.1:8080/users/create
row JSON body:
{
    "user_id": "3",
	"username": "John",
	"team_name": "backend",
	"is_active": true
}
