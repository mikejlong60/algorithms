package skiena_3

import (
	"bitbucket.org/pcastools/hash"
	"fmt"
	"math"
	"testing"
)

type KeyValuePair[K, V any] struct {
	key   K
	value V
}
type HashMap[K, V any] struct {
	underlying []KeyValuePair[K, V]
	zero       KeyValuePair[K, V]
	eq         func(k1, k2 K) bool
	emptyCells int
	hash       func(k K) uint32
}

func New[K, V any](zero KeyValuePair[K, V], eq func(k1, k2 K) bool, hash func(k K) uint32) HashMap[K, V] {
	zeroHashMap := func() []KeyValuePair[K, V] {
		a := make([]KeyValuePair[K, V], 10)
		for i := 0; i < len(a); i++ {
			a[i] = zero
		}
		return a
	}
	return HashMap[K, V]{
		underlying: zeroHashMap(),
		zero:       zero,
		eq:         eq,
		hash:       hash,
		emptyCells: 10,
	}
}

func Set[K, V any](m HashMap[K, V], kv KeyValuePair[K, V]) HashMap[K, V] {
	a := m.hash(kv.key)
	idx := int32(math.Mod(float64(a), float64(len(m.underlying)-1)))

	b := m.underlying[idx]
	assignToHashOrClosestEmpty := func() bool { //Assigns kv pair to either the hashed array index or the closest available spot in the underlying array,
		// moving to the right until you get to the end.
		// Returns true if successful. Otherwise false which indicates array needs resizing.
		if m.eq(b.key, m.zero.key) { //A hash collision. Look forward for an empty spot
			m.underlying[idx] = kv
			m.emptyCells = m.emptyCells - 1
			return true
		} else {
			underlyingSize := int32(len(m.underlying))
			for i := idx; i < underlyingSize; i++ {
				if m.eq(m.underlying[i].key, m.zero.key) {
					m.underlying[i] = kv
					m.emptyCells = m.emptyCells - 1
					return true
				}
			}
			return false
		}
	}

	resizeHashMap := func() HashMap[K, V] {
		return m //TODO Make this really do it, 1. Double the original size. 2. assign the new hash element as part of rehashing the whole old array. 3. Return the new HashMap
	}

	assigned := assignToHashOrClosestEmpty()
	if !assigned {
		resizeHashMap()
		m.underlying[a] = kv
	}

	return m
}

func TestGoHashMap(t *testing.T) {
	s := map[int]any{3: nil, 2: nil, 1: nil}
	s[12] = nil
	_, ok := s[12]
	if !ok {
		t.Errorf("Key 12 not added")
	}
	delete(s, 12)
	_, ok = s[12]
	if ok {
		t.Errorf("Key 12 not deleted")
	}
}

func TestYourOwnHashMap(t *testing.T) {
	fFNV32a := func(x int32) uint32 {
		return hash.Int32(x)
	}
	eq := func(k1, k2 int32) bool {
		if k1 == k2 {
			return true
		} else {
			return false
		}
	}
	m := New[int32, string](KeyValuePair[int32, string]{-1, ""}, eq, fFNV32a)
	m = Set(m, KeyValuePair[int32, string]{1234234, "fred"})
	fmt.Println(m)
}
