package ollamaplugin

import (
	"context"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

type OllamaPlugin struct {
	handle framework.Handle
}

func (pl *OllamaPlugin) Name() string {
	return "OllamaPlugin"
}

func (pl *OllamaPlugin) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	if pod.Labels["app"] == "ollama" {
		if node := nodeInfo.Node(); node != nil {
			if val, exists := node.Labels["gpu"]; exists && val == "true" {
				return framework.NewStatus(framework.Success)
			}
		}
		return framework.NewStatus(framework.Unschedulable, "Ollama pods must run on GPU nodes")
	}
	return framework.NewStatus(framework.Success)
}

func New(configuration *runtime.Unknown, handle framework.Handle) (framework.Plugin, error) {
	return &OllamaPlugin{handle: handle}, nil
}
