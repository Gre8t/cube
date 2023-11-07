package scheduler

type scheduler interface {
	SelectCandidateNodes()
	Score()
	Pick()
}