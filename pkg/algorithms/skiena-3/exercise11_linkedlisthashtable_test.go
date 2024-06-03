package skiena_3

import (
	"bitbucket.org/pcastools/hash"
	"fmt"
	"github.com/greymatter-io/golangz/linked_list"
	"github.com/greymatter-io/golangz/option"
	"math"
	"testing"
)

type HashMap2[K, V any] struct {
	underlying []*linked_list.LinkedList[KeyValuePair[K, V]]
	eq         func(k1, k2 KeyValuePair[K, V]) bool
	hash       func(k K) uint32
}

func New2[K, V any](eq func(k1, k2 KeyValuePair[K, V]) bool, hash func(k K) uint32, capacity int32) HashMap2[K, V] {
	return HashMap2[K, V]{
		eq:         eq,
		hash:       hash,
		underlying: make([]*linked_list.LinkedList[KeyValuePair[K, V]], capacity),
	}
}

func Get2[K, V any](m HashMap2[K, V], k K, f func(a KeyValuePair[K, V]) bool) option.Option[KeyValuePair[K, V]] {
	a := m.hash(k)
	idx := int32(math.Mod(float64(a), float64(len(m.underlying))))

	b := m.underlying[idx]
	c := linked_list.Filter[KeyValuePair[K, V]](b, f)

	if c != nil {
		return option.Some[KeyValuePair[K, V]]{linked_list.Head[KeyValuePair[K, V]](c)}
	} else {
		return option.None[KeyValuePair[K, V]]{}
	}
}

func Set2[K, V any](m HashMap2[K, V], kv KeyValuePair[K, V]) HashMap2[K, V] {
	a := m.hash(kv.key)
	idx := int32(math.Mod(float64(a), float64(len(m.underlying))))

	b := m.underlying[idx]
	putKeyValuePair2InRightBucket := func() *linked_list.LinkedList[KeyValuePair[K, V]] {
		//1.  Get the array element containing the bucket(linked list) for the key.
		//TODO 2. Look in that bucket to see if that KeyValuePair2 is there already.
		//TODO 3. If it is, replace its value.
		//TODO 4. If it's not, add it to the head of the linked list in the bucket.

		return linked_list.Push[KeyValuePair[K, V]](kv, b)
	}
	d := putKeyValuePair2InRightBucket()
	m.underlying[idx] = d
	return m
}

func TestYourOwnHashMap2(t *testing.T) {
	fFNV32a := func(x int32) uint32 {
		return hash.Int32(x)
	}
	eq := func(k1, k2 KeyValuePair[int32, string]) bool {
		if k1.key == k2.key {
			return true
		} else {
			return false
		}
	}
	m := New2[int32, string](eq, fFNV32a, 13)
	k := int32(1234234)
	m = Set2(m, KeyValuePair[int32, string]{k, "fred"})

	f := func(a KeyValuePair[int32, string]) bool {
		eq := func(k1, k2 KeyValuePair[int32, string]) bool {
			if k1.key == k2.key {
				return true
			} else {
				return false
			}
		}

		if eq(a, KeyValuePair[int32, string]{k, ""}) {
			return true
		} else {
			return false
		}
	}

	//The value should be in the map
	err := option.GetOrElse(Get2(m, k, f), KeyValuePair[int32, string]{k, fmt.Sprintf("Should not have found:%v in HashMap", 188)}) //fmt.Sprintf("Should not have found:%v in HashMap", 188))
	if err.key != k {
		t.Errorf(err.value)
	}

	//The value should NOT be in the map
	err = option.GetOrElse(Get2(m, 188, f), KeyValuePair[int32, string]{188, fmt.Sprintf("Should not have found:%v in HashMap", 188)}) //fmt.Sprintf("Should not have found:%v in HashMap", 188))
	if err.value != fmt.Sprintf("Should not have found:%v in HashMap", 188) {
		t.Errorf(err.value)
	}
}
