package utils

// Node represents a node in a Graph.
type Node struct {
	name string
}

// SetName sets the name of the node.
func (n *Node) SetName(name string) {
	n.name = name
}

// Name returns the name of the node.
func (n *Node) Name() string {
	return n.name
}

// Edge represents an edge between two Nodes in a Graph.
type Edge struct {
	weight     int
	start, end *Node
}

// Weight returns the weight of the edge.
func (e *Edge) Weight() int {
	return e.weight
}

// Graph represents a collection of Nodes connected by Edges.
type Graph struct {
	nodes []*Node
	edges map[Node][]*Edge
}

// Contains returns true if the Graph contains the Node, or false if it does not.
func (g *Graph) Contains(node *Node) bool {
	for i := range g.nodes {
		if g.nodes[i] == node {
			return true
		}
	}
	return false
}

// GetNodes returns all the nodes in the Graph.
func (g *Graph) GetNodes() []*Node {
	nodes := make([]*Node, len(g.nodes))
	copy(nodes, g.nodes)
	return nodes
}

// GetNode returns the Node with the specified name, or nil if it does not exist in the Graph.
func (g *Graph) GetNode(name string) *Node {
	for i := range g.nodes {
		if g.nodes[i].name == name {
			return g.nodes[i]
		}
	}
	return nil
}

// AddNode adds a node to the Graph.
func (g *Graph) AddNode(node *Node) {
	g.nodes = append(g.nodes, node)
}

// AddEdge adds an Edge to the graph between two nodes.
func (g *Graph) AddEdge(start, end *Node, weight int) {
	if g.edges == nil {
		g.edges = make(map[Node][]*Edge)
	}
	edge := new(Edge)
	edge.weight = weight
	edge.start = start
	edge.end = end
	g.edges[*start] = append(g.edges[*start], edge)
}

// Neighbors returns the Nodes that are directly connected to a Node.
func (g *Graph) Neighbors(node *Node) []*Node {
	if edges, ok := g.edges[*node]; ok {
		nodes := make([]*Node, 0)
		for _, edge := range edges {
			nodes = append(nodes, edge.end)
		}
		return nodes
	}
	return []*Node{}
}

// IsNeighbor returns true if end is a neighbor of start, or false if it is not.
func (g *Graph) IsNeighbor(start, end *Node) bool {
	for _, node := range g.Neighbors(start) {
		if node == end {
			return true
		}
	}
	return false
}

// Weight returns the weight of the Edge that connects two Nodes, or panics if the Nodes are not connected.
func (g *Graph) Weight(start, end *Node) int {
	if edges, ok := g.edges[*start]; ok {
		for _, edge := range edges {
			if edge.end == end {
				return edge.Weight()
			}
		}
	}
	panic("No edge between the two provided nodes.")
}
