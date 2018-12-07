
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

rm-containers:
	docker rm $(shell docker ps -a -f status=exited -q)