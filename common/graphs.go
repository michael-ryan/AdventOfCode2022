package common

import "math"

// Graph represents a graph as an adjacency list.
type Graph[T comparable] map[T][]Edge[T]

// Edge represents an Edge in the graph.
type Edge[T any] struct {
	To     T
	Weight int
}

// dijkstra finds the shortest path from the start node to all other nodes in the graph.
// It returns a map of node names to distances.
func (g Graph[T]) Dijkstra(start T) map[T]int {
	distances := make(map[T]int)
	for node := range g {
		distances[node] = math.MaxInt64
	}
	distances[start] = 0

	visited := make(map[T]bool)
	var previousVisited T
	first := true
	for len(visited) < len(g) {
		// Find the unvisited node with the smallest distance.
		var smallestNode T
		var smallestDistance int = math.MaxInt64

		for node, distance := range distances {
			if !visited[node] && distance < smallestDistance {
				smallestNode = node
				smallestDistance = distance
			}
		}
		if previousVisited == smallestNode && !first {
			break
		}
		// Visit the smallest node and update the distances of its neighbors.
		visited[smallestNode] = true
		previousVisited = smallestNode
		for _, e := range g[smallestNode] {
			if !visited[e.To] {
				newDistance := distances[smallestNode] + e.Weight
				if newDistance < distances[e.To] {
					distances[e.To] = newDistance
				}
			}
		}

		first = false
	}

	return distances
}
