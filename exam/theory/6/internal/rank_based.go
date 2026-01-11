package internal

type RankDSU struct {
	parent []int
	size   []int
}

func (r *RankDSU) MakeSet(v int) {
	r.parent[v] = v
	r.size[v] = 1
}

func (r *RankDSU) FindSet(v int) int {
	if v == r.parent[v] {
		return v
	}

	return r.FindSet(r.parent[v])
}

func (r *RankDSU) UnionSets(u, v int) {
	u = r.FindSet(u)
	v = r.FindSet(v)

	if u != v {
		if r.size[u] < r.size[v] {
			u, v = v, u
		}
		r.parent[v] = u
		r.size[u] += r.size[v]
	}
}
