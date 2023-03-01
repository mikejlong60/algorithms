package chapter4

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestLocalMinimum(t *testing.T) {
	a := Node{
		Value: 14,
	}

	b := Node{
		Value: 13,
	}
	c := Node{
		Value: 12,
	}
	d := Node{
		Value: 11,
	}
	e := Node{
		Value: 10,
	}
	f := Node{
		Value: 9,
	}
	g := Node{
		Value: 8,
	}
	h := Node{
		Value: 7,
	}
	i := Node{
		Value: 6,
	}
	j := Node{
		Value: 5,
	}
	k := Node{
		Value: 4,
	}
	l := Node{
		Value: 3,
	}
	m := Node{
		Value: 2,
	}
	n := Node{
		Value: 1,
	}
	o := Node{
		Value: 0,
	}
	a.Left = &b
	a.Right = &c
	b.Parent = &a
	c.Parent = &a
	d.Parent = &b
	e.Parent = &b
	f.Parent = &c
	g.Parent = &c

	h.Parent = &d
	i.Parent = &d
	j.Parent = &e
	k.Parent = &e
	l.Parent = &f
	m.Parent = &f
	n.Parent = &g
	o.Parent = &g

	b.Left = &d
	b.Right = &e
	c.Left = &f
	c.Right = &g
	d.Left = &h
	d.Right = &i
	e.Left = &j
	e.Right = &k
	f.Left = &l
	f.Right = &m
	g.Left = &n
	g.Right = &o
	r := LocalMinimum(20, &b)
	if r != 4 {
		t.Errorf("Actual:%v Expected:%v", r, 4)
	}

	log.Infof("Local Minimum:%v, Total Steps:%v", r, totalSteps)
}
