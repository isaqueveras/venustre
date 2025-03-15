package refx

import (
	"log"
	"math/rand"
	"time"

	"github.com/isaqueveras/venustre/v2"
)

type engx struct{}

func New() venustre.Worker {
	return &engx{}
}

func (*engx) Kind() string {
	return "change_ticket_state"
}

func (e *engx) Work(ctx *venustre.Context) {
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	log.Println("Olá mundo", e.Kind())
}
