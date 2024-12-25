package main

import (
	"github.com/zahaev/cloud-technologies-and-backend/internal/worker"
)

func main() {
	worker.RunWorkers(3)
}
