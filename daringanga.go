package main

import (
  "fmt"
  "math/rand"
  "github.com/marcelcorso/nuruu"
)

func IN(h *nuruu.Graph) []string {

  marked := make(map[string]bool)

  for _, v := range(h.Vertices()) {
    if h.Degree(v) == 0 {
      marked[v] = true
    } else {
      if (rand.Float64() < (1 / h.Degree(v))) {
        marked[v] = true
      }
    }
  }

  for _, edge := range(h.Edges()) {
    if marked[edge.row] && edge.column[edge.column] {
      if rand.Float64() < (h.Degree(edge.column) / (h.Degree(edge.row + h.Degree(edge.column))) ) {
        delete(marked, edge.row)
      } else {
        delete(marked, edge.column)
      }
    }
  }

  // S = set of vertices that remain marked
  vertices := make([]string, 0, len(marked))
  for key := range marked {
    vertices = append(vertices, key)
  }
  return vertices
}

func main() {

  h := nuruu.NewGraph()
  h.Add("a", "b")

  i := nuruu.NewGraph()
  for ; len(h.Vertices()) > 0; {
    s := IN(h)
    h.RemoveVertices(s.Vertices() + s.Neighbors(s.Vertices()))

    i.AddGraph(s)
  }

  // i has the independent set

  fmt.Println("Hello, dude")
}

