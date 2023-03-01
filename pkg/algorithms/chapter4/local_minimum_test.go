package chapter4

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestLocalMinimumFromInnerNode(t *testing.T) {
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
	r := LocalMinimum(&b)
	if r != 4 {
		t.Errorf("Actual:%v Expected:%v", r, 4)
	}

	log.Infof("Local Minimum:%v, Total Steps:%v", r, totalSteps)
}

func TestLocalMinimumFromTopNode(t *testing.T) {
	a := Node{
		Value: 140,
	}

	b := Node{
		Value: 130,
	}
	c := Node{
		Value: 120,
	}
	d := Node{
		Value: 110,
	}
	e := Node{
		Value: 100,
	}
	f := Node{
		Value: 90,
	}
	g := Node{
		Value: 80,
	}
	h := Node{
		Value: 70,
	}
	i := Node{
		Value: 60,
	}
	j := Node{
		Value: 50,
	}
	k := Node{
		Value: 40,
	}
	l := Node{
		Value: 3,
	}
	m := Node{
		Value: 20,
	}
	n := Node{
		Value: 10,
	}
	o := Node{
		Value: 1800,
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
	r := LocalMinimum(&a)
	if r != 3 {
		t.Errorf("Actual:%v Expected:%v", r, 3)
	}

	log.Infof("Local Minimum:%v, Total Steps:%v", r, totalSteps)
}

func TestLocalMinimumFromLeafNode(t *testing.T) {
	a := Node{
		Value: 140,
	}

	b := Node{
		Value: 130,
	}
	c := Node{
		Value: 120,
	}
	d := Node{
		Value: 110,
	}
	e := Node{
		Value: 100,
	}
	f := Node{
		Value: 90,
	}
	g := Node{
		Value: 80,
	}
	h := Node{
		Value: 70,
	}
	i := Node{
		Value: 60,
	}
	j := Node{
		Value: 50,
	}
	k := Node{
		Value: 40,
	}
	l := Node{
		Value: 3,
	}
	m := Node{
		Value: 20,
	}
	n := Node{
		Value: 10,
	}
	o := Node{
		Value: 1800,
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
	r := LocalMinimum(&o)
	if r != 80 {
		t.Errorf("Actual:%v Expected:%v", r, 80)
	}

	log.Infof("Local Minimum:%v, Total Steps:%v", r, totalSteps)
}
