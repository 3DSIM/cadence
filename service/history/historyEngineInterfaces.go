package history

import (
	h "code.uber.internal/devexp/minions/.gen/go/history"
	workflow "code.uber.internal/devexp/minions/.gen/go/shared"
	"code.uber.internal/devexp/minions/common"
)

type (
	// Engine represents an interface for managing workflow execution history.
	Engine interface {
		common.Daemon
		// TODO: Convert workflow.WorkflowExecution to pointer all over the place
		StartWorkflowExecution(request *workflow.StartWorkflowExecutionRequest) (*workflow.StartWorkflowExecutionResponse, error)
		GetWorkflowExecutionHistory(
			request *workflow.GetWorkflowExecutionHistoryRequest) (*workflow.GetWorkflowExecutionHistoryResponse, error)
		RecordDecisionTaskStarted(request *h.RecordDecisionTaskStartedRequest) (*h.RecordDecisionTaskStartedResponse, error)
		RecordActivityTaskStarted(request *h.RecordActivityTaskStartedRequest) (*h.RecordActivityTaskStartedResponse, error)
		RespondDecisionTaskCompleted(request *workflow.RespondDecisionTaskCompletedRequest) error
		RespondActivityTaskCompleted(request *workflow.RespondActivityTaskCompletedRequest) error
		RespondActivityTaskFailed(request *workflow.RespondActivityTaskFailedRequest) error
		RecordActivityTaskHeartbeat(
			request *workflow.RecordActivityTaskHeartbeatRequest) (*workflow.RecordActivityTaskHeartbeatResponse, error)
	}

	// EngineFactory is used to create an instance of sharded history engine
	EngineFactory interface {
		CreateEngine(context ShardContext) Engine
	}

	historySerializer interface {
		Serialize(history []*workflow.HistoryEvent) ([]byte, error)
		Deserialize(data []byte) ([]*workflow.HistoryEvent, error)
	}

	transferQueueProcessor interface {
		common.Daemon
		UpdateMaxAllowedReadLevel(maxAllowedReadLevel int64)
	}

	timerQueueProcessor interface {
		common.Daemon
		NotifyNewTimer(taskID int64)
	}
)