package netpoll

import "k8s-lx1036/routing-go/app/framework/network/internal"

// Poller ...
type Poller struct {
	fd            int
	asyncJobQueue internal.AsyncJobQueue
}

