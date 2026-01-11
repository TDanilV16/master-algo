package internal

type PathShortageDSU struct {
	parent []int
}

func (p *PathShortageDSU) MakeSet(v int) {
	p.parent[v] = v
}

func (p *PathShortageDSU) FindSet(v int) int {
	if p.parent[v] == v {
		return v
	}
	p.parent[v] = p.FindSet(p.parent[v])
	return p.parent[v]
}

func (p *PathShortageDSU) UnionSets(u, v int) {
	u = p.FindSet(u)
	v = p.FindSet(v)
	if u != v {
		p.parent[v] = u
	}
}
