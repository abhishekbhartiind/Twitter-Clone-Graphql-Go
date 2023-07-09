


mock:
	mockery --all --keeptree

# migrate:
# 	migrate -source file://postgres/migrations \
# 			-database postgres://postgres:@127.0.0.1:5432/twitter?sslmode=disable up

# rollback:
# 	migrate -source file://postgres/migrations \
# 			-database postgres://postgres:postgres@127.0.0.1:5432/twitter?sslmode=disable down 1

# drop:
# 	migrate -source file://postgres/migrations \
# 			-database postgres://postgres:postgres@127.0.0.1:5432/twitter?sslmode=disable drop




migrate:
	migrate -source file://postgres/migrations \
			-database postgres://dreamer:YNlGXN1IC70jfe0phRMPMAedilAjUovB@dpg-cikppjh5rnuvtgr1guq0-a.singapore-postgres.render.com/twitter_fxn0 up

rollback:
	migrate -source file://postgres/migrations \
			-database postgres://dreamer:YNlGXN1IC70jfe0phRMPMAedilAjUovB@dpg-cikppjh5rnuvtgr1guq0-a.singapore-postgres.render.com/twitter_fxn0 down 1

drop:
	migrate -source file://postgres/migrations \
			-database postgres://dreamer:YNlGXN1IC70jfe0phRMPMAedilAjUovB@dpg-cikppjh5rnuvtgr1guq0-a.singapore-postgres.render.com/twitter_fxn0 drop


migration:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir postgres/migrations $$name

run:
	go run cmd/graphqlserver/main.go
