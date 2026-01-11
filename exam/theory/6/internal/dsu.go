package internal

type DSU interface {
	MakeSet(v int)
	FindSet(v int) int
	UnionSets(u, v int)
}

// CompleteDSU -- rank heuristic + path shortage heuristic
type CompleteDSU struct {
	parent []int
	size   []int
}

func (c *CompleteDSU) MakeSet(v int) {
	c.parent[v] = v
	c.size[v]++
}

func (c *CompleteDSU) FindSet(v int) int {
	if v == c.parent[v] {
		return v
	}
	c.parent[v] = c.FindSet(c.parent[v])
	return c.parent[v]
}

func (c *CompleteDSU) UnionSets(u, v int) {
	u = c.FindSet(u)
	v = c.FindSet(v)

	if u != v {
		if c.size[u] < c.size[v] {
			u, v = v, u
		}
		c.parent[v] = u
		c.size[u] += c.size[v]
	}
}
