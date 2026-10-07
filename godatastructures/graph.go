package godatastructures

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"sort"
)

/*
This Graph definitions:
  - More probable to be Sparse than Dense, hence not a complete Graph
  - Undirected (Insert links a node and its neighbors both ways)
  - Unweighted
  - Acyclic (could be Cyclic)
  - represented as Adjacency List, not Adjacency Matrix
  - V is the number of nodes, E is the number of edges (used in the complexities below)
*/
type GraphInterface[T constraints.Ordered] interface {
	Empty() bool                       // returns whether graph is empty. O(1)
	Size() int                         // returns number of items in graph.
	Insert(val Pair[T], neighbors []T) // helper function to build/update graph. inserts T into graph if not present, adds neighbors. O(neighbors)
	Remove(val T)                      // helper function to build/update graph. removes T from graph, removes T from other nodes' neighbors. O(V)
	DepthFirstTraversal() []T          // returns values in graph ordered by DFS processing. O(V + ElogV)
	BreadthFirstTraversal() []T        // returns values in graph ordered by BFS processing. O(V + ElogV)
}

// Graph should implement GraphInterface; write a constructor & methods to complete it
type Graph[T constraints.Ordered] struct {
	nodes     []Pair[T]
	neighbors Map[T, *Set[T]]
}

// /*      Methods: Graph struct      */
func NewGraph[T constraints.Ordered]() *Graph[T] {
	return &Graph[T]{
		nodes:     []Pair[T]{},
		neighbors: *NewMap[T, *Set[T]](),
	}
}

// helper function to add node (and its empty list of neighbors) if doesn't exist in Graph
func (g *Graph[T]) addNode(val Pair[T]) {
	// // if there's a node with same key, update its value
	// for _, node := range g.nodes {
	//   if node.Key == val.Key {
	//     node.Value = val.Value
	//     return
	//   }
	// }
	if g.neighbors.Contains(val.Key) {
		fmt.Printf("Pair.Key %v already exists in the Graph\n", val.Key)
		return
	}
	g.nodes = append(g.nodes, val)
	g.neighbors.Set(val.Key, NewSet[T]())
}

// helper function to add edge FROM -> TO node
func (g *Graph[T]) addEdge(from, to T) {
	//get the list of neighbors from node that we want
	adjacencyList, _ := g.neighbors.Get(from).(*Set[T])
	//if it has connection already, print error msg
	if adjacencyList.Contains(to) {
		fmt.Printf("There's already a connection between %v and %v\n", from, to)
		return
	}
	//insert TO the list
	adjacencyList.Insert(to)
}

// returns whether graph is empty. O(1)
func (g *Graph[T]) Empty() bool {
	return len(g.nodes) == 0
}

// returns number of items in graph.
func (g *Graph[T]) Size() int {
	return len(g.nodes)
}

// helper function to build/update graph. inserts T into graph if not present, adds neighbors. O(neighbors)
func (g *Graph[T]) Insert(val Pair[T], neighbors []T) {
	// if the node already exists, only its neighbors get updated
	g.addNode(val)

	for _, neighbor := range neighbors {
		// neighbors must already exist in the Graph
		if !g.neighbors.Contains(neighbor) {
			fmt.Printf("Neighbor %v doesn't exist in Graph\n", neighbor)
			continue
		}
		g.addEdge(val.Key, neighbor)
		if neighbor != val.Key {
			g.addEdge(neighbor, val.Key)
		}
	}
}

// helper function to build/update graph. removes T from graph, removes T from other nodes' neighbors. O(V)
func (g *Graph[T]) Remove(val T) {
	index := -1
	for i, node := range g.nodes {
		if node.Key == val {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Printf("Node with Val %v doesn't exist in Graph\n", val)
		return
	}

	// remove == ignore that specific val in Graph
	// 1. removing from list of nodes
	g.nodes = append(g.nodes[:index], g.nodes[index+1:]...)

	// 2. removing from list of neighbors
	// remove the Key, so it won't be able to access the val.(*Set[T])
	// thus being removed by garbage collector
	g.neighbors.Remove(val)

	// 3. removing all connections from neighbors to this val
	for _, key := range g.neighbors.Keys() {
		adjacencyList, _ := g.neighbors.Get(key).(*Set[T])

		if adjacencyList.Contains(val) {
			adjacencyList.Remove(val)
		}
	}
}

// helper function to get a node's neighbors in ascending order, so traversals are deterministic
func (g *Graph[T]) sortedNeighbors(val T) []T {
	adjacencyList, _ := g.neighbors.Get(val).(*Set[T])
	neighbors := adjacencyList.Values()
	sort.Slice(neighbors, func(i, j int) bool { return neighbors[i] < neighbors[j] })
	return neighbors
}

// DepthFirstTraversal func helper
func (g *Graph[T]) dfs(val T, visited *map[T]bool, result *[]T) {
	if !(*visited)[val] {
		*result = append(*result, val)
		(*visited)[val] = true

		for _, neighborVal := range g.sortedNeighbors(val) {
			g.dfs(neighborVal, visited, result)
		}
	}
}

// returns values in graph ordered by DFS processing. O(V + ElogV)
// we could use Stack
func (g *Graph[T]) DepthFirstTraversal() []T {
	if len(g.nodes) == 0 {
		fmt.Printf("No nodes present in Graph: %v\n", g.nodes)
		return []T{}
	}

	result := []T{}
	visited := make(map[T]bool)
	for _, node := range g.nodes {
		g.dfs(node.Key, &visited, &result)
	}
	return result
}

// BreadthFirstTraversal func helper
func (g *Graph[T]) bfs(val T, queue *[]T, visited *map[T]bool, result *[]T) {
	for len(*queue) > 0 {
		currVal := (*queue)[0]
		*queue = (*queue)[1:]

		if !(*visited)[currVal] {
			*result = append(*result, currVal)
			(*visited)[currVal] = true

			for _, neighborVal := range g.sortedNeighbors(currVal) {
				*queue = append(*queue, neighborVal)
			}
		}
	}
}

// returns values in graph ordered by BFS processing. O(V + ElogV)
func (g *Graph[T]) BreadthFirstTraversal() []T {
	if len(g.nodes) == 0 {
		fmt.Printf("No nodes present in Graph: %v\n", g.nodes)
		return []T{}
	}

	result := []T{}
	visited := make(map[T]bool)
	queue := []T{}
	for _, node := range g.nodes {
		currVal := node.Key
		if !visited[currVal] {
			queue = append(queue, currVal)
			g.bfs(currVal, &queue, &visited, &result)
		}
	}
	return result
}
