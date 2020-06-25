DC            := docker-compose -f build/docker-compose.yml -p sqlboiler_inner_join_test
TEST_USER     := test_user
TEST_PASSWORD := test_pass
TEST_DB       := test_db
TEST_PORT     := 3232
DBA           := db mysql -u${TEST_USER} -p${TEST_PASSWORD} ${TEST_DB}
DBA_ROOT      := db mysql -uroot -proot ${TEST_DB}

.PHONY: up
up:
	${DC} up -d db

.PHONY: down
down:
	${DC} down

.PHONY: clean
clean:
	${DC} down --volumes
	rm -rf log

.PHONY: wait
wait:
	${DC} run --rm wait

.PHONY: ps
ps:
	${DC} ps

.PHONY: exec
exec:
	${DC} exec ${DBA_ROOT}
