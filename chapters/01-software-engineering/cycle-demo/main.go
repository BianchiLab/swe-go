package main

import (
	"context"
	"example.com/cycle-demo/entity"
	"example.com/cycle-demo/warehouse"
)

func wireComponents() {
	entity.AcquirePacker = func(ctx context.Context) entity.Packer {
		return warehouse.AcquireRobot(ctx)
	}
}

func main() {
	wireComponents()
}
