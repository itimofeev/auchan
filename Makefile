
GO_IMAGE=golang:1.10.4
SWAGGER_IMAGE=quay.io/goswagger/swagger:v0.17.2


deploy-db:
	docker stack deploy --compose-file tools/docker-stack.yml db

gen-server:
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
    		--header="Content-type: application/yaml" \
    		--header="Authorization: eyJUb2tlblR5cGUiOiJBUEkiLCJzYWx0IjoiMDY2NjM4MTQtODExMS00YTZkLWJmODEtMmFmOGFjNmU5ZDBiIiwiYWxnIjoiSFM1MTIifQ.eyJqdGkiOiIuZ2l0aHViLjE0Mzk4NCIsImlhdCI6MTUzMzU3MjY5NX0.x0DsS4uPuzA1eYxYOSdlGOzLXfdGyIPNXrRm8KYG1Xe82h2J48y_0AE0v1tsUZiyvwU6cZS3BBikjkiYbtJdtw" \
    		https://api.swaggerhub.com/apis/itimofeev/cityproject.auchan

rm:
	docker service rm $(shell docker service ls -q) || true

rm-containers:
	docker rm $(shell docker ps -a -f status=exited -q)