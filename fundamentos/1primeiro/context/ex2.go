package main

import (
	"context"
	"fmt"
	"time"
)

func operationOne(ctx context.Context) {
	n := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context canceled for op-1")
			return // returning not to leak the goroutine
		default:
			fmt.Printf("OperationOne: %d\n", n)
			time.Sleep(500 * time.Millisecond)
			n++
		}
	}
}

func operationTwo(ctx context.Context) {
	n := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context canceled for op-2")
			return // returning not to leak the goroutine
		default:
			fmt.Printf("OperationTwo: %d\n", n)
			time.Sleep(250 * time.Millisecond)
			n++
		}
	}
}

func exemploContext(ctx context.Context) {
	fmt.Println(ctx.Value(propriedade("propriedadeQualquer")))
}

type propriedade string

func main() {
	// context.Background() -> contexto padrao, utilizado para criar outros contextos em cima dele
	ctx, cancel := context.WithCancel(context.Background()) // inicializador de contexto
	cancel()

	// evitar usar withValue
	// serve para definir um valor no context
	ctx = context.WithValue(ctx, propriedade("propriedadeQualquer"), "valor")
	exemploContext(ctx)

	d := time.Now().Add(15000 * time.Millisecond) // calculando 5 seg após a hora atual
	ctx, cancel1 := context.WithDeadline(ctx, d)  // tempo q expiraria o contexto
	defer cancel1()
	go operationOne(ctx)

	d = time.Now().Add(10000 * time.Millisecond)
	ctx, cancel2 := context.WithDeadline(context.Background(), d)
	defer cancel2()
	go operationTwo(ctx)

	time.Sleep(3 * time.Second)
	cancel1()
	time.Sleep(20 * time.Second)
}
