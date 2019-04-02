
.PHONY: run
run: gallery db
	./gallery

gallery: *.go
	go build -o gallery

./PHONY: clean
clean: dbclean
	[[ -f gallery ]] && rm -f gallery || echo "Already clean"

.PHONY: db
db: dbdir dbimage
	[[ `docker ps | egrep 'pg-docker'|wc -l` -eq "0" ]] && docker run --rm   --name pg-docker \
		-e POSTGRES_PASSWORD=docker -d -p 5432:5432 \
		-v postgres_data:/var/lib/postgresql/data  postgres || echo "DB running"

.PHONY: dbclean
dbclean: dbstop dbdirclean
#	[[ `docker volume ls | egrep postgres_data | wc -l` -gt "0" ]] &&  docker volume rm postgres_data || echo "volume already removed"


.PHONY: dbimage
dbimage:
	docker pull postgres

.PHONY: dbdir
dbdir:
	[[ ! -d ./postgres_data ]] && mkdir postgres_data || echo "postgres_data folder already exists"

.PHONY: dbstop
dbstop:
	[[ `docker ps | egrep pg-docker | wc -l` -gt "0" ]] && docker stop pg-docker || echo "DB stopped"

.PHONY: dbdirclean
dbdirclean:
	[[  -d ./postgres_data ]] && rm -rf ./postgres_data || echo "Data cleared"

.PHONY: test
test: db
	go test -v

.PHONY: fresh
fresh:
	fresh -c fresh.conf
