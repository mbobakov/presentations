package escape

import "testing"

// SLIDE0 OMIT
type click struct {
	id   string
	info *info
}

func (c *click) description() {}

func storeDescription(c *click) {
	c.description()
}

// SLIDE0 END OMIT

// SLIDE1 OMIT
type descriptioner interface {
	description()
}

func storeDescriptionIfaced(c descriptioner) {
	c.description()
}

// SLIDE1 END OMIT
// SLIDE2 OMIT
func BenchmarkExample1Ifaced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := &click{id: "1"}
		storeDescriptionIfaced(c)
	}
}

func BenchmarkExample1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := &click{id: "1"}
		storeDescription(c)
	}
}

// SLIDE2 END OMIT
