package main

import (
	"fmt"
	"sort"

	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
)

//this is adjency list implementation
type edge struct {
	source      int
	destination int
	weight      int
	next        *edge
}
type edgeHead struct {
	head *edge
}
type Graph struct {
	count     int
	egdeslist []*edgeHead
}

type subset struct {
	parent int
	rank   int
}

func in(cnt int) {
	g := new(Graph)
	g.count = cnt
	g.egdeslist = make([]*edgeHead, cnt)
}
func (g *Graph) addEdge(source int, dest int, weight int) {
	e := &edge{source, dest, weight, g.egdeslist[source].head}
	g.egdeslist[source].head = e
}
func (g *Graph) print() {
	x := g.count
	for i := 0; i < x; i++ {
		y := g.egdeslist[i].head
		for y != nil {
			fmt.Println(y.source, " to ", y.destination)
			y = y.next
		}
	}
}

func (g *Graph) DFS(source int) []bool {
	stk := new(stack.Stack)
	stk.Push(source) //whatever
	visit := make([]bool, g.count)
	for stk.Len() != 0 {
		cur := stk.Pop().(int)
		head := g.egdeslist[cur].head
		for head != nil {
			if visit[head.destination] == false {
				stk.Push(head.destination)
				visit[head.destination] = true
			}
			head = head.next
		}
	}
	return visit
}

func (g *Graph) BFS(source int) []bool {
	q := new(queue.Queue)
	visit := make([]bool, g.count)
	q.Enqueue(source) //whatever
	for q.Len() != 0 {
		cur := q.Dequeue().(int)
		head := g.egdeslist[cur].head
		for head != nil {
			if visit[head.destination] == false {
				visit[head.destination] = true
				q.Enqueue(head.destination)
			}
			head = head.next
		}
	}
	return visit
}

func (g *Graph) DFSRecursive(source int) []bool {
	visit := make([]bool, g.count)
	DFSRec(source, visit)
	return visit
}

//we use int for node's id, so that in
//edgeslist it's mapped
func (g *Graph) DFSRec(source int, visit []bool) {
	head := g.egdeslist[source].head
	for head != nil {
		if visit[head.destination] == false {
			visit[head.destination] = true
			//fmt
			DFSRec(head.destination, visit)
		}
		head = head.next
	}
}

func (g *Graph) DFSRecArray(source int, visit []bool, arr []*edge) {
	head := g.egdeslist[source].head
	for head != nil {
		if visit[head.destination] == false {
			visit[head.destination] = true
			//fmt
			append(arr, head)
			DFSRec(head.destination, visit)
		}
		head = head.next
	}
}

//topological sort
func (g *Graph) topologicalSort(source int) {
	stk := new(stack.Stack)
	visit := make([]bool, g.count)
	//for i
	topoRec(stk, visit, source)
	for stk.Len() != 0 {
		stk.Pop()
		//fmt
	}
}
func (g *Graph) topoRec(stk *stack.Stack, visit []bool, source int) {
	head := g.egdeslist[source].head
	for head != nil {
		if visit[head.destination] == false {
			visit[head.destination] = true
			topoRec(stk, visit, head.destination)
		}
		head = head.next
	}
	stk.Push(head.source)
}

func (g *Graph) prims(source int, v int) []int {
	pq := new(PQueue)
	pq.Init()
	dis := make([]int, v)
	view := make([]int, v)
	for i := 0; i < v; i++ {
		dis[i] = Infinite
		view[i] = -1
	}
	dis[source] = 0
	view[0] = 0
	ed := &edge{source, source, 0, nil}
	pq.Push(ed)
	for pq.Len() != 0 {
		ed := pq.Pop().(*edge)
		dis[ed.destination] = ed.cost
		view[ed.destination] = ed.source
		adn := g.egdeslist[ed.destination].head
		for adn != nil {
			d := adn.destination
			if view[d] != -1 && dis[d] > adn.cost {
				pq.Push(adn, adn.cost)
			}
			adn = adn.next
		}
	}
}

func find(sb []subset, x int) int {
	if sb[x] == sb[x].parent {
		return x
	}
	sb[x].parent = find(sb, sb[x].parent)
	return sb[x].parent
}

func union(sb []subset, x int, y int) {
	xroot := find(sb, x)
	yroot := find(sb, y)

	if sb[xroot].rank < sb[yroot].rank {
		sb[xroot].parent = yroot
	}
	if sb[xroot].rank > sb[yroot].rank {
		sb[yroot].parent = xroot
	}
	if sb[xroot].rank == sb[yroot].rank {
		sb[xroot].parent = yroot
		sb[yroot].rank++
	}

}

type SortBy []*edge

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].weight < a[j].weight }

func (g *Graph) kruskal(v int) {
	visit := make([]bool, g.count)
	arr := make([]*edge, 0)
	DFSRecArray(0, visit, arr)
	sort.Sort(SortBy(arr))
	sb := make([]subset, v)
	for i := 0; i < v; i++ {
		sb[i].parent = i
		sb[i].rank = 0
	}
	e := 0
	result := make([]*edge, v-1)
	for e < v-1 {
		ed := arr[i]
		x := find(sb, ed.source)
		y := find(sb, ed.destination)
		//they are disjoint
		if x != y {
			result[e] = ed
			e++
			union(sb, ed.source, ed.destination)
		}
	}
	for _, val := range result {
		fmt.Println(val.source, " ", val.destination)
	}
}

//shortest path for unweighted graph
//or same weight
//g.count is number of vertices
func (g *Graph) unweighted(source int) ([]int, []int) {
	v := g.count
	distance := make([]int, v)
	path := make([]int, v)
	q := new(queue.Queue)
	for i := 0; i < v; i++ {
		distance[i] = -1
		path[i] = -1
	}
	q.Push(source)
	distance[source] = 0
	for q.Len() != 0 {
		a := q.Pop().(int)
		head := g.egdeslist[a].head
		for head != nil {
			if distance[head.destination] == -1 {
				q.Push(head.destination)
				distance[head.destination] = distance[head.source] + 1
				path[head.destination] = head.source
			}
			head = head.next
		}
	}
	return distance, path
}

func printUnweighted(source int, dest int) {
	distance, path := unweighted(source)
	i := dest
	var line []int
	for path[i] != -1 {
		append(line, path[i])
		i = path[i]
	}
	fmt.Println("distance ", distance[dest])
	for _, val := range line {
		fmt.Print(val, " ")
	}
}

func (g *Graph) bellmanford(source int) {
	v := g.count
	distance := make([]int, v)
	path := make([]int, v)
	path[source] = source
	for i := 0; i < v; i++ {
		distance[i] = Infinite
	}
	distance[source] = 0
	//at most v-1 edges to get there
	for i := 0; i < v-1; i++ {
		for j := 0; j < v; j++ { //in real, iterate all edges
			head := g.egdeslist[j].head
			for head != nil {
				if distance[head.source] != Infinite {
					head = head.next
					continue
				} else {
					newdis := distance[j] + head.weight
					if distance[j] > newdis {
						distance[j] = newdis
						path[j] = head.source
					}
				}
				head = head.next
			}
		}
	}
	//check if it has a negetive cycle
	for i := 0; i < v; i++ {
		head := g.edgeslist[i].head
		for head != nil {
			if distance[head.source] != Infinite && distance[head.source]+head.weight < distance[head.destination] {
				fmt.Println("There is a cycle")
				return
			}
		}
	}
}

func (g *Graph) floydWarshall() {
	v := g.count
	var distance [][]int
	var path [][]int
	for i := 0; i < v; i++ {
		for j := 0; j < v; j++ {
			if j != i {
				distance[i][j] = Infinite
			}
			path[i][j] = -1 //means not reachable
		}
	}
	for i := 0; i < v; i++ {
		head := g.edgeslist[i].head
		for head != nil {
			distance[head.source][head.destination] = head.weight
			path[i][j] = -2 //means reahcable from its own edge directly
			head = head.next
		}
	}
	for k := 0; k < v; k++ {
		for i := 0; i < v; i++ {
			for j := 0; j < v; j++ {
				if distance[i][k]+distance[k][j] < distance[i][j] {
					distance[i][j] = distance[i][k] + distance[k][j]
					path[i][j] = k //means using {0,1,...,k} vertices
				}
			}
		}
	}
	//fmt.print
}

func nb(v []bool) {
	v[0] = true
}
func main() {
	visit := make([]bool, 3)
	nb(visit)
	fmt.Println(visit)
}
