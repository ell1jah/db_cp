.PHONY: run genData runPostgres buildPostgres

genData:
	python3 scripts/genInitData.py

runPostgres:
	docker run -d --rm -p 5432:5432 clothshop

buildPostgres:
	docker build --tag=clothshop:latest . -f build/package/postgres/Dockerfile

fillPostgres:
	docker exec -it $(shell docker container ls --latest --quiet) psql -U postgres -d clothshop -f "mnt/copy.sql"
