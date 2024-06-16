package scheduler

import (
	"github.com/vvatanabe/cube/node"
	"github.com/vvatanabe/cube/task"
)

type Scheduler interface {
	SelectCandidateNodes()
	Score()
	Pick()
}

type RoundRobin struct {
	Name       string
	LastWorker int
}

func (r *RoundRobin) SelectCandidateNodes(t task.Task, nodes []*node.Node) []*node.Node {
	return nodes
}
