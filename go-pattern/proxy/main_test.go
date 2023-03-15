package proxy

import "testing"

func TestStationProxy(t *testing.T) {
	obj := StationProxy{
		station: &Station{
			stock: 1,
		},
	}
	obj.sell("lucy")
	obj.sell("lucy")
}
