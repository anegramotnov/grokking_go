package graph

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "fmt"
)

func TestListQueue(t *testing.T) {
    // var q SearchQueue
    q := new(ListQueue)

    _, err := q.Dequeue()
    assert.NotEqual(t, nil, err)

    q.Enqueue("first")

    e, _ := q.Dequeue()
    assert.Equal(t, "first", e)

    _, err = q.Dequeue()
    assert.NotEqual(t, nil, err)

    q.Enqueue("first")
    q.BatchEnqueue([]string{"second", "third"})
    q.Enqueue("fourth")

    e, _ = q.Dequeue()
    assert.Equal(t, "first", e)

    e, _ = q.Dequeue()
    assert.Equal(t, "second", e)

}

type MapGraphTestSuite struct {
    suite.Suite
    graph *MapGraph
}

func isMockfn(s string) bool {
    return true
}

func (suite *MapGraphTestSuite) SetupTest() {
    graph := NewMapGraph()
    graph.AddEdge("you", "alice")
    graph.AddEdge("you", "bob")
    graph.AddEdge("you", "claire")
    graph.AddEdge("bob", "anuj")
    graph.AddEdge("bob", "peggy")
    graph.AddEdge("alice", "peggy")
    suite.graph = graph

}

func (suite *MapGraphTestSuite) TestEmptyGraph() {
    graph := NewMapGraph()

    assert.Equal(suite.T(), []string(nil), graph.Children("mock"))

    _, found := graph.BFS("mock", isMockfn)
    assert.Equal(suite.T(), false, found)
}

func (suite *MapGraphTestSuite) TestAddEdge() {
    referenceMap := make(map[string][]string)
    referenceMap["you"] = []string{"alice", "bob", "claire"}
    referenceMap["bob"] = []string{"anuj", "peggy"}
    referenceMap["alice"] = []string{"peggy"}

    graph := NewMapGraph()

    graph.AddEdge("you", "alice")
    graph.AddEdge("you", "bob")
    graph.AddEdge("you", "claire")
    graph.AddEdge("bob", "anuj")
    graph.AddEdge("bob", "peggy")
    graph.AddEdge("alice", "peggy")

    assert.Equal(suite.T(), referenceMap, graph.m)

}

func (suite *MapGraphTestSuite) TestChildren() {
    assert.Equal(suite.T(), []string{"alice", "bob", "claire"},
                 suite.graph.Children("you"))
}


func TestMapGrapTestSuite(t *testing.T) {
    suite.Run(t, new(MapGraphTestSuite))
}

func TestMapGraph(t *testing.T) {
    graph := NewMapGraph()

    assert.Equal(t, []string(nil), graph.Children("you"))

    graph.AddEdge("you", "alice")
    graph.AddEdge("you", "bob")
    graph.AddEdge("you", "claire")

    assert.Equal(t, []string{"alice", "bob", "claire"}, graph.Children("you"))

    assert.Equal(t, []string(nil), graph.Children("alice"))

}


func isMangoSeller(s string) bool {
    if s[len(s)-1:] == "m" {
        fmt.Printf("'%s' is Mango Seller!\n", s)
        return true
    }
    return false
}

func TestMapGraphBFS(t *testing.T) {
    graph := NewMapGraph()

    graph.AddEdge("you", "alice")
    graph.AddEdge("you", "bob")
    graph.AddEdge("you", "claire")

    graph.AddEdge("bob", "anuj")
    graph.AddEdge("bob", "peggy")

    graph.AddEdge("alice", "peggy")


    mangoSeller, found := graph.BFS("you", isMangoSeller)
    assert.Equal(t, false, found)

    graph.AddEdge("claire", "jonny")
    graph.AddEdge("claire", "thom")

    mangoSeller, found = graph.BFS("you", isMangoSeller)
    assert.Equal(t, true, found)
    assert.Equal(t, "thom", mangoSeller)

    graph.AddEdge("jonny", "mom")

    mangoSeller, found = graph.BFS("you", isMangoSeller)
    assert.Equal(t, true, found)
    assert.Equal(t, "thom", mangoSeller)




}