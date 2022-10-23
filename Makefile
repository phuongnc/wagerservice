# ----- prepare database test -----
test-renew-seed:
	$(eval TESTDB_ID := $(shell docker-compose ps -q testdb))
	docker exec -i $(TESTDB_ID) mysql -uadmin -pWagerService123 wager_test < ./test/testdb/seed.sql

test:
	@./test/test.sh

test-all: docker-up test-renew-seed test

# ----- Docker -----
docker-up:
	@docker-compose up -d

docker-down:
	@docker-compose down

docker-clear:
	@echo "Are you ok to remove volumes? [y/N]" && read ans && [ $${ans:-N} = y ]
	@docker-compose down -v
