package toposort

type Edge struct {
    start, end *Node
}

type Node struct {  
    val int
}

type DAG struct {
    nodes []*Node
    edges []*Edge
}

func reverse(ns []*Node) {
    c := len(ns) / 2
    for i := 0; i < c; i++ {
        is := len(ns) - i - 1 
        ns[i], ns[is] = ns[is], ns[i]
    }
}

func TopoSort(g *DAG) []*Node {
    res := make([]*Node,0)
    q := make([]*Node,0)
    
    // O(1) lookup => O(|N| + |E|)
    outs := make(map[*Node][]*Edge)
    inCounts := make(map[*Node]int)

    for _, n := range g.nodes {
        outs[n] = make([]*Edge,0)
        inCounts[n] = 0
    }

    for _, e := range g.edges {
        outs[e.start] = append(outs[e.start], e)
        inCounts[e.end] += 1
    }

    // init q  => O(|N|)
    for k, v := range inCounts {
        if v == 0 {
            q = append(q, k)
        }
    }

    // O(|N| + |E|)
    for len(q) != 0 {
        qt := make([]*Node,0)
        for _, n := range q {
            res = append(res, n)

            for _, e := range outs[n] {
                inCounts[e.end] -= 1

                if inCounts[e.end] == 0 {
                    qt = append(qt, e.end)
                }
            }
        }
        q = qt
    }

    return res
}