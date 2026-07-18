DB_URL="postgres://postgres:Anvar2005!@localhost:5432/todolist?sslmode=disable"
MIGRATIONS_DIR=migrations

run:
	go run main.go

build:
	go build -o bin/todo_app main.go

clean:
	rm -rf bin/

migrate-new:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $$name

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) up

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) down 1

migrate-drop:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) down

migrate-force:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) force $(v)

migrate-version:
	migrate -path $(MIGRATIONS_DIR) -database $(DB_URL) version