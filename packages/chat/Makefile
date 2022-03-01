include ./packages/chat/configs/configs.env

export

generate:
	sqlc generate

migrate/up:
	migrate -database ${database_url} -path ./packages/chat/configs/db/migrations up

migrate/down:
	migrate -database ${database_url} -path ./packages/chat/configs/db/migrations down


