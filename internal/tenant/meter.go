package tenant

import "sync"

var usage = map[string]int{}
var muMeter sync.Mutex

func Inc(id string) {

	muMeter.Lock()
	defer muMeter.Unlock()

	usage[id]++
}

func Get(id string) int {

	muMeter.Lock()
	defer muMeter.Unlock()

	return usage[id]
}
