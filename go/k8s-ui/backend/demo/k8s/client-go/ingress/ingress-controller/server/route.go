package server

import (
	"crypto/tls"
	"k8s-lx1036/k8s-ui/backend/demo/k8s/client-go/ingress/ingress-controller/watcher"
	"net/url"
	"regexp"
)

// A RoutingTable contains the information needed to route a request.
type RoutingTable struct {
	certificatesByHost map[string]map[string]*tls.Certificate
	backendsByHost     map[string][]routingTableBackend
}

type routingTableBackend struct {
	pathRE *regexp.Regexp
	url    *url.URL
}

func NewRoutingTable(payload *watcher.Payload) *RoutingTable {
	rt := &RoutingTable{
		certificatesByHost: make(map[string]map[string]*tls.Certificate),
		backendsByHost:     make(map[string][]routingTableBackend),
	}
	rt.init(payload)
	return rt
}

func (rt *RoutingTable) init(payload *watcher.Payload) {

}
