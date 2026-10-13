package entity

import "context"

// Item represents an item that will be shipped to a customer.
type Item struct {
	// Various fields.
}

// Box contains items that are shipped to the customer.
type Box struct {
	// Various fields.
}

// Packer is implemented by objects that can pack an Item into a Box.
type Packer interface {
	Pack(*Item, *Box) error
}

// AcquirePacker returns a Packer instance.
var AcquirePacker func(context.Context) Packer

// Pack packs qty items of type i into the box.
func (b *Box) Pack(i *Item, qty int) error {
	p := AcquirePacker(context.Background())

	for j := 0; j < qty; j++ {
		if err := p.Pack(i, b); err != nil {
			return err
		}
	}

	return nil
}
