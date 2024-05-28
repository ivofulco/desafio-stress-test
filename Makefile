build:
	go build -o loadtester ./cmd/loadtester ;

use:
	./loadtester --url=http://google.com --requests=100 --concurrency=10	

docker-clean-up:
	docker rm $(docker ps -aq) -f

desafio: build use
