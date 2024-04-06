DB_CONTID=  $(shell docker ps -qf "name=db_cont")
PG_ADMIN= $(shell docker ps -qf "name=pgadmin_container")

app-up-local:
	docker-compose --env-file ./pkg/env/.env -f ./deployment/docker-compose.yaml up -d

app-down-local:
	docker-compose -f ./deployment/docker-compose.yaml down 

app-up-docker:
	docker-compose --env-file ./pkg/env/docker.env -f docker-compose.yaml up -d

app-down-docker:
	docker-compose -f docker-compose.yaml down 

db-ip:
	docker inspect  ${DB_CONTID} | grep IPAddress

pg-ip:
	docker inspect  ${PG_ADMIN} | grep IPAddress

db-gen:
	gen --sqltype=postgres \
    --connstr "host=localhost port=5432 dbname=dirdatadb user=narendra password=123456 sslmode=disable" \
    --database=postgres \
    --table datafiles \
    --json \
    --gorm \
    --guregu \
    --out ./pkg \
    --json-fmt=snake \
    --generate-dao \
    --module="github.com/narendra121/data-dir-db/pkg" \
    --overwrite