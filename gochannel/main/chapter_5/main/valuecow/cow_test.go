package valuecow

import (
	"math/rand"
	"sync"
	"testing"
)

func TestConcurrentArray(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		t.Run("New", testNew)
		array := NewConcurrentArray(uint32(rand.Int31n(100)))
		maxI := uint32(5)
		//t.Run("Set", func(t *testing.T) {
			testSet(array, maxI, t)
		//})
		//t.Run("Get", func(t *testing.T) {
		testGet(array, maxI, t)
		//})
	})
}

func testGet(array ConcurrentArray, maxI uint32, t *testing.T) {
	arrayLen := array.Len()
	intMax := int((maxI - 1) * (arrayLen - 1))
	for i := uint32(0); i < arrayLen; i++ {
		elem, err := array.Get(i)
		if err != nil {
			t.Fatalf("Unexpected error: %s (index: %d)", err, i)

		}

		if elem < 0 || elem > intMax {
			t.Fatalf("Incorect element: %d! (index: %d, expect max: %d)",
				elem, i, intMax)
		}
		t.Logf("Get [%d] elem: %d", i, elem)
	}
}

func testSet(array ConcurrentArray, maxI uint32, t *testing.T) {
	arrayLen := array.Len()
	var wg sync.WaitGroup
	wg.Add(int(maxI))
	for i := uint32(0); i < maxI; i++ {
		go func(i uint32) {
			defer wg.Done()
			for j := uint32(0); j < arrayLen; j++ {
				err := array.Set(j, int(j * i))
				t.Logf("Set [%d] elem: %d", j, j * i)
				if uint32(j) >= arrayLen && err == nil {
					t.Fatalf("Unexpected nil error! (index: %d)", j)
				} else {
					if err != nil {
						t.Fatalf("Unexpected error: %s (index: %d)", err, j)
					}
				}
			}
		}(i)
	}
	wg.Wait()
	
}

func testNew(t *testing.T) {
	
}