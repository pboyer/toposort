package toposort

import ("testing")

func assert(t *testing.T, ns []*Node, i int, n *Node){
    if ns[i] != n {
        t.Fatalf("wrong node at pos %d, expected %v, got %v", i, n, ns[i])
    }
}

func TestTopoSortSimple(t *testing.T){
    a := &Node{0}
    b := &Node{1}
    c := &Node{2}
    d := &Node{3}

    ab := &Edge{ a, b }
    bc := &Edge{ b, c }
    cd := &Edge{ c, d }

    dag := &DAG{
        nodes : []*Node{ a, b, c, d },
        edges : []*Edge{ ab, bc, cd },
    }

    ns := TopoSort(dag)

    assert(t, ns, 0, a)
    assert(t, ns, 1, b)
    assert(t, ns, 2, c)
    assert(t, ns, 3, d)
}

func TestTopoSortSimple2(t *testing.T){
    a := &Node{0}
    b := &Node{1}
    c := &Node{2}
    d := &Node{3}
    e := &Node{4}

    ab := &Edge{ a, b }
    bc := &Edge{ b, c }
    bd := &Edge{ b, d }
    ae := &Edge{ a, e }

    dag := &DAG{
        nodes : []*Node{ a, b, c, d, e },
        edges : []*Edge{ ab, bc, bd, ae },
    }

    ns := TopoSort(dag)

    assert(t, ns, 0, a)
    assert(t, ns, 1, b)
    assert(t, ns, 2, e)
    assert(t, ns, 3, c)
    assert(t, ns, 4, d)
}