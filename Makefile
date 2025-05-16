# Makefile for gptcli

build:
	go build -o gptcli cmd/gptcli/main.go

clean:
	rm -f gptcli 