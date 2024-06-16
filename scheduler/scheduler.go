package scheduler

type Scheduler interface {
	SelectCandidateNodes()
	Score()
	Pick()
}

type RoundRobin struct {
	Name       string
	LastWorker int
}
