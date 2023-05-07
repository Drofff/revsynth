package cli

import (
	"os"
	"strconv"
	"strings"

	"github.com/Drofff/revsynth/aco"
	"github.com/Drofff/revsynth/circuit"
	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func sliceToStr(s []int) []string {
	res := make([]string, 0)
	for _, el := range s {
		res = append(res, strconv.Itoa(el))
	}
	return res
}

func gateToString(gate circuit.Gate) string {
	return "g(" +
		strings.Join(sliceToStr(gate.TargetBits()), ",") +
		", [" +
		strings.Join(sliceToStr(gate.ControlBits()), ",") +
		"])"
}

func pheromoneAmountToInt(pheromoneAmount float64) int {
	return int(pheromoneAmount * 1000)
}

func SaveGraph(pheromones aco.Pheromones, filePath string) error {
	g := graph.New(graph.StringHash, graph.Directed())
	addedVertexes := map[string]bool{}

	for _, pheromoneDeposit := range pheromones {
		fromVertex := pheromoneDeposit.FromState.Key()
		_, fromAdded := addedVertexes[fromVertex]
		if !fromAdded {
			err := g.AddVertex(fromVertex)
			if err != nil {
				return err
			}
			addedVertexes[fromVertex] = true
		}

		toVertex := pheromoneDeposit.ToState.Key()
		_, toAdded := addedVertexes[toVertex]
		if !toAdded {
			err := g.AddVertex(toVertex)
			if err != nil {
				return err
			}
			addedVertexes[toVertex] = true
		}

		err := g.AddEdge(fromVertex, toVertex,
			graph.EdgeWeight(pheromoneAmountToInt(pheromoneDeposit.PheromoneAmount)),
			graph.EdgeAttribute("gate", gateToString(pheromoneDeposit.UsedGate)))
		if err != nil {
			return err
		}
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	return draw.DOT(g, file)
}
