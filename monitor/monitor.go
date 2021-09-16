package monitor

import (
	monitor_controller "github.com/ProjectAthenaa/sonic-core/protos/monitorController"
	"github.com/ProjectAthenaa/sonic-core/sonic/base"
	"github.com/ProjectAthenaa/sonic-core/sonic/face"
	config "github.com/ProjectAthenaa/supreme/config"
	"github.com/prometheus/common/log"
)

var _ face.MonitorCallback = (*Task)(nil)

type Task struct {
	*base.BMonitor

	keywords *monitor_controller.Keywords

	PID         string
	Size        string
	sizeId      string
	category    string
	channelName string
}

func NewTask(data *monitor_controller.Task) (*Task, error) {
	task := &Task{
		BMonitor:    &base.BMonitor{Data: data},
		category:    data.Metadata[*config.Module.Fields[0].FieldKey],
		Size:        data.Metadata[*config.Module.Fields[3].FieldKey],
		keywords:    data.GetKeywords(),
		channelName: data.RedisChannel,
	}

	task.Callback = task

	return task, nil
}

func (tk *Task) TaskLoop() {
	var err error
	for {
		select {
		case <-tk.Ctx.Done():
			return
		default:
			if err = tk.iteration(); err != nil {
				log.Error("error completing iteration", err)
			}
		}
	}
}

func (tk *Task) OnStarting() {

}

func (tk *Task) OnStopping() {

}
