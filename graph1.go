package main 

import (
	"fmt"
)
type Vertex struct{
	key int
	adjacent []*Vertex
}
type Graph struct{
	vertices []*Vertex
}
func (g *Graph) getVertex(vertex int) *Vertex{
	for i,v := range g.vertices{
		if v.key == vertex{
			return g.vertices[i]
		}
	}
	return nil
}
func (g *Graph)addVertex(vertex int)error{
	if contains(g.vertices, vertex){
		err := fmt.Errorf("Vertex %d already present",vertex)
		return err
	}else{
		v := &Vertex{
			key: vertex,
		}	
		g.vertices = append(g.vertices,v)

	}	
	return nil
}
func (g *Graph)addEdge(to int,from int)error{
	toVertex := g.getVertex(to)
	fromVertex := g.getVertex(from)
	if toVertex == nil || fromVertex == nil{
		return fmt.Errorf("Not a valid edge from %d ---> %d",from,to)
	} else if contains(fromVertex.adjacent, toVertex.key){
		return fmt.Errorf("Edge from vertex %d ---> %d already exists", fromVertex.key)
	} else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
		return nil
	}
}
func contains(v []*Vertex,key int)bool{
	for _,v := range v{
		if v.key == key{
			return true	
		}
	}
	return false
}

func (g *Graph) Print(){
	for _, v := range g.vertices {
		fmt.Printf("%d:", v.key)
		for _, v := range v.adjacent{
			fmt.Printf("%d",v.key)
		}
		fmt.Println()
	}
}
func main(){
	g := &Graph{}
	g.addVertex(1)
	g.addVertex(2)
	g.addVertex(3)
	g.addEdge(1,2)
	g.addEdge(2,3)
	g.addEdge(1,3)
	g.addEdge(3,1)
	g.Print()
}
