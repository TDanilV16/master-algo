package internal

type NaiveDSU struct {
	parent []int
}

func (n *NaiveDSU) MakeSet(v int) {
	n.parent[v] = v
}

func (n *NaiveDSU) FindSet(v int) int {
	if v == n.parent[v] {
		return v
	}
	return n.FindSet(n.parent[v])
}

func (n *NaiveDSU) UnionSets(u, v int) {
	u = n.FindSet(u)
	v = n.FindSet(v)
	if u != v {
		n.parent[v] = u
	}
}
