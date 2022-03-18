package controller

import (
	"fmt"
	"strings"
)

// Should I use "clean-up" or "cleanup"?
// * cleanup is a noun - e.g "The cleanup"
// * clean-up is a verb - e.g. "I clean-up"

type (
	podCleanupKey    = string // describes the pod to cleanup + the cleanup action to take
	podCleanupAction = string
)

const (
	deletePod           podCleanupAction = "deletePod"
	shutdownPod         podCleanupAction = "shutdownPod"
	labelPodCompleted   podCleanupAction = "labelPodCompleted"
	terminateContainers podCleanupAction = "terminateContainers"
	killContainers      podCleanupAction = "killContainers"
)

func newPodCleanupKey(workflowNamespace, cluster, namespace, podName string, action podCleanupAction) podCleanupKey {
	return fmt.Sprintf("%s/%s/%s/%s/%v", workflowNamespace, cluster, namespace, podName, action)
}

func parsePodCleanupKey(k podCleanupKey) (workflowNamespace, cluster, namespace, podName string, action podCleanupAction) {
	parts := strings.Split(k, "/")
	if len(parts) != 5 {
		return "", "", "", "", ""
	}
	return parts[0], parts[1], parts[2], parts[3], parts[4]
}