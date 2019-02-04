package graph

import (
	"container/list"
	// "fmt"
	// "reflect"
	"errors"
)

// *FYI: https://www.geeksforgeeks.org/graph-and-its-representations/
// *FYI: https://github.com/jesand/numgo/tree/master/graph
// *FYI: https://github.com/arnauddri/algorithms/blob/master/data-structures/graph/
//  FYI: https://github.com/maximelamure/algorithms/blob/master/datastructure
//  FYI: https://github.com/starwander/goraph/
//  FYI: https://github.com/Workiva/go-datastructures/tree/master/graph

// TODO: make better graph realization

type SearchQueue interface {
	BatchEnqueue(values []string)
	Dequeue() (string, error)
	Len() int
}

type ListQueue struct {
	list.List
}

func (q *ListQueue) Enqueue(value string) {
	q.PushBack(value)
}

func (q *ListQueue) BatchEnqueue(values []string) {
	// fmt.Printf("Add %v in queue\n", values)
	for i := 0; i < len(values); i++ {
		q.Enqueue(values[i])
	}
}

func (q *ListQueue) Dequeue() (string, error) {
	if q.Len() > 0 {
		e := q.Front()
		value := e.Value
		q.Remove(e)
		return value.(string), nil
	}
	return "", errors.New("Queue is empty")
}


// TODO: universal interface!

type MapGraph struct {
	m map[string][]string
}

func NewMapGraph() *MapGraph {
	graph := new(MapGraph)
	graph.m = make(map[string][]string)
	return graph
}

// TODO: from, to params
func (g *MapGraph) AddEdge(from, to string) {
	// fmt.Printf("AddEdge %s -> %s\n", from, to)
	// fmt.Printf("MapGraph state: %v\n", g)
	g.m[from] = append(g.m[from], to)
}

func (g *MapGraph) Children(of string) []string {
	return g.m[of]
}

func (g *MapGraph) BFS(from string, check func(string) bool) (string, bool) {
// func MapBFS(graph map[string][]string, key func(string) bool) bool {
	// fmt.Printf("Start BFS in %v\n", graph)
	var queue SearchQueue
	// fmt.Printf("1 type(queue=%s", reflect.TypeOf(queue))
	queue = new(ListQueue)
	// fmt.Printf("2 type(queue=%s", reflect.TypeOf(queue))

	// TODO: need better way to determine root!
	queue.BatchEnqueue(g.Children(from))

	for queue.Len() > 0  {
		vertex, _ := queue.Dequeue()
		// fmt.Printf("Check vertex: %s\n", vertex)
		if check(vertex) {
			return vertex, true
		} else {
			// fmt.Printf("Check failed\n")
			queue.BatchEnqueue(g.Children(vertex))
		}
		// fmt.Printf("\n")
	}
	return "", false

}
