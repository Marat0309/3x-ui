package job

import (
	"x-ui/logger"
	"x-ui/web/service"
)

type CheckNodeStatusJob struct {
	nodeService service.NodeService
}

func NewCheckNodeStatusJob() *CheckNodeStatusJob {
	return new(CheckNodeStatusJob)
}

func (j *CheckNodeStatusJob) Run() {
	nodes, err := j.nodeService.List()
	if err != nil {
		logger.Warning("list nodes failed:", err)
		return
	}
	for i := range nodes {
		if err := j.nodeService.Check(&nodes[i]); err != nil {
			logger.Warning("node check failed:", err)
		}
	}
}
