package event

import (
	"context"
	"k8s-lx1036/k8s-ui/dashboard/controllers/resource/common"
	"k8s-lx1036/k8s-ui/dashboard/controllers/resource/common/dataselect"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"strings"
)

var FailedReasonPartials = []string{"failed", "err", "exceeded", "invalid", "unhealthy",
	"mismatch", "insufficient", "conflict", "outof", "nil", "backoff"}

func GetPodsEventWarnings(events []corev1.Event, pods []corev1.Pod) []Event {

}

func ListNamespaceEventsByQuery(
	k8sClient kubernetes.Interface,
	namespaceName string,
	dataSelect *dataselect.DataSelectQuery) (EventList, error) {
	rawEventList, err := k8sClient.CoreV1().Events(namespaceName).List(context.TODO(), common.ListEverything)
	if err != nil {
		return EventList{}, err
	}
	
	eventList := toEventList(FillEventsType(rawEventList.Items))
	
	return eventList, nil
}

func FillEventsType(events []corev1.Event) []corev1.Event {
	for i := range events {
		if len(events[i].Type) == 0 { // type is empty
			if isFailedReason(events[i].Reason, FailedReasonPartials...) {
				events[i].Type = corev1.EventTypeWarning
			} else {
				events[i].Type = corev1.EventTypeNormal
			}
		}
	}
	
	return events
}

func toEventList(events []corev1.Event) EventList {
	eventList := EventList{
		ListMeta: common.ListMeta{
			TotalItems: len(events),
		},
	}
	
	
}

func isFailedReason(reason string, partials ...string) bool {
	for _, partial := range partials {
		if strings.Contains(strings.ToLower(reason), partial) {
			return true
		}
	}
	
	return false
}
