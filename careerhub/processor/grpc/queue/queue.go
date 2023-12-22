package queue

import (
	"fmt"
)

type Queue interface {
	Send(message []byte) error
	Recv() ([][]byte, error)
}

type FakeQueue struct {
}

func (fq *FakeQueue) Send(message []byte) error {

	fmt.Printf("%s\n\n", string(message))
	return nil
}

func (fq *FakeQueue) Recv() ([][]byte, error) {
	return nil, nil
}
