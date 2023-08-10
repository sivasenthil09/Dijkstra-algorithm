package main

import (
	"fmt"
	"math"
)

type Graph map[string]map[string]int

func Dijkstra(graph Graph, start string) map[string]int {
	distances := make(map[string]int)
	visited := make(map[string]bool)

	// Initialize distances and visited
	for vertex := range graph {
		distances[vertex] = math.MaxInt32
		visited[vertex] = false
	}
	distances[start] = 0

	// Dijkstra's algorithm
	for i := 0; i < len(graph); i++ {
		// Find the vertex with the smallest distance
		minDist := math.MaxInt32
		var minVertex string
		for vertex, distance := range distances {
			if !visited[vertex] && distance < minDist {
				minDist = distance
				minVertex = vertex
			}
		}

		// Mark the current vertex as visited
		visited[minVertex] = true

		// Update distances for neighboring vertices
		for neighbor, weight := range graph[minVertex] {
			if !visited[neighbor] {
				if dist := distances[minVertex] + weight; dist < distances[neighbor] {
					distances[neighbor] = dist
				}
			}
		}
	}

	return distances
}

func main() {
	graph := Graph{
		"A": {"B": 1, "C": 4},
		"B": {"A": 1, "C": 2, "D": 5},
		"C": {"A": 4, "B": 2, "D": 1},
		"D": {"B": 5, "C": 1},
	}

	start := "A"
	distances := Dijkstra(graph, start)

	fmt.Printf("Shortest distances from vertex %s:\n", start)
	for vertex, distance := range distances {
		fmt.Printf("%s: %d\n", vertex, distance)
	}
}
