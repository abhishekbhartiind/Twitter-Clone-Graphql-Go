mock:
	mockery --all --keeptree

migrate:
	migrate -source file://postgres/migrations \
			-database postgres//postgres:@127.0.0.1:5432/twitter?sslmode=disable up
		
rollback:
	migrate -source file://postgres/migrations \
			-database postgres//postgres:@127.0.0.1:5432/twitter?sslmode=disable down


drop:
	migrate -source file://postgres/migrations \
			-database postgres//postgres:@127.0.0.1:5432/twitter?sslmode=disable drop

