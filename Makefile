SRC=gopong.go computer.go game.go server.go

build:
	go build $(SRC)

run:
	./gopong $(ARGS)
