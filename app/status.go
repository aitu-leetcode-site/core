package app

import "sync/atomic"

type runStatus struct {
	v atomic.Int32
}

const (
	initialStatus = 0
	runningStatus = 1
	closedStatus  = 2
)

func (r *runStatus) isClosed() bool {
	return r.v.Load() == closedStatus
}

func (r *runStatus) sedClosed() {
	r.v.Store(int32(closedStatus))
}

func (r *runStatus) isRunning() bool {
	return r.v.Load() == runningStatus
}

func (r *runStatus) setRunning() {
	r.v.Store(runningStatus)
}

func (r *runStatus) setInitialStatus() {
	r.v.Store(initialStatus)
}
