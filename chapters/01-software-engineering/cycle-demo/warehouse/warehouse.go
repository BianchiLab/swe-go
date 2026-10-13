package warehouse

import (
	"context"

	"example.com/cycle-demo/entity"
)

// Robot navigates the warehouse floor and fetches items for packing.
type Robot struct {
	// Various fields.
}

// AcquireRobot blocks until a Robot becomes available or until the
// context expires.
func AcquireRobot(ctx context.Context) *Robot {
	// TODO: Implement robot acquisition.
	return &Robot{}
}

// Pack instructs the robot to pick up an item from its shelf and place
// it into a box that will be shipped to the customer.
func (r *Robot) Pack(item *entity.Item, to *entity.Box) error {
	// TODO: Implement the packing operation.
	return nil
}
