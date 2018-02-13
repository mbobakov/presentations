package escape

import "testing"

type info struct {
	clientID int
}

func BenchmarkIndirectAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {

		// c := new(click)
		// c.info = &info{} // BAD

		c := &click{
			info: &info{}, // GOOD
		}

		storeDescription(c)
	}
}

//go:noinline
func (c *click) add(f int) { _ = f }

func BenchmarkIndirectCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := 6
		c := new(click)

		// foo := c.add //BAD
		// foo(t)

		c.add(t)

		storeDescription(c)
	}
}

func BenchmarkClosure(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := new(click)
		// func() { //BAD
		// 	c.description()
		// }()

		c.uberDescription()

	}
}
func (c *click) uberDescription() { c.description() }
func BenchmarkIfaced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := new(click)
		// storeDescriptionIfaced(c) // BAD
		storeDescription(c)
	}
}
