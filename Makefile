
GO_IMAGE=golang:1.10.4
SWAGGER_IMAGE=quay.io/goswagger/swagger:v0.17.2


deploy-db:
	docker stack deploy --compose-file tools/docker-stack.yml db

gen-server: download
	docker run --rm -v $(GOPATH):/go/ -w /go/src/github.com/itimofeev/auchan -t $(SWAGGER_IMAGE) \
		generate server \
		--target=server \
		-f tools/swagger.yml

gen-client:
	docker run --rm -v $(GOPATH):/go/ -w /go/src/github.com/itimofeev/auchan -t $(SWAGGER_IMAGE) \
		generate client \
		--existing-models=github.com/itimofeev/auchan/models \
		--target=auchan \
		-f tools/swagger.yml

download:
	wget -O tools/swagger.yml\
    		--header="Accept: application/yaml" \
    		https://api.swaggerhub.com/apis/itimofeev/cityproject.auchan/1.0.0

rm:
	docker service rm $(shell docker service ls -q) || true

rm-full: rm rm-containers
	docker volume rm db_pg_data

rm-containers:
	docker rm $(shell docker ps -a -f status=exited -q)

release: build-docker build-image upload clean run-remote

upload:
	scp -r auchan.img root@159.69.121.222:/root/auchan
	scp -r tools/docker-stack.yml root@159.69.121.222:/root/auchan/tools
	scp -r tools/cmd.sql root@159.69.121.222:/root/auchan/tools
	scp -r Makefile root@159.69.121.222:/root/auchan
	ssh root@159.69.121.222 "cd auchan; ls"

run-remote:
	ssh root@159.69.121.222 "cd auchan; \
		docker load -i auchan.img; \
		make rm; \
		make deploy-db \
		"

clean:
	rm auchan auchan.img

build-docker:
	docker run \
		-v $$GOPATH:/go \
		-w /go/src/github.com/itimofeev/auchan \
		-t $(GO_IMAGE) \
		make build
	docker cp `docker ps -q -n=1`:/go/bin/linux_amd64/auchan ./

build-image:
	docker build  \
		--force-rm=true \
		-t auchan \
		-f tools/auchan.Dockerfile `pwd`
	docker image save auchan -o auchan.img

import-products:
	docker cp cmd.sql `docker ps -q -f name=db_db`:/cmd.sql \
		&& docker exec `docker ps -q -f name=db_db` psql postgres postgres -f /cmd.sql

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o $$GOPATH/bin/linux_amd64/auchan github.com/itimofeev/auchan/server/cmd/city-project-for-auchan-server