package builder

import "testing"

func TestBuilder(t *testing.T) {
	d1 := &Director{}
	mac := &MacBookBuilder{}
	d1.MakeProduct(mac)
	mac.GetProduct()
}
