GOPATH := /home/isucon/isubata/webapp/go
export GOPATH

build:
	go build -v isubata
	sudo systemctl restart isubata.golang.service

vet:
	go vet ./src/isubata/...

