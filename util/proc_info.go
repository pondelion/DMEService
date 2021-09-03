package util

import (
	"github.com/mitchellh/go-ps"

	"dme_service/model"
)

func ProcessList() []model.Proc {
	processes, err := ps.Processes()

	if err != nil {
		panic(err)
	}

	var procs []model.Proc
	for _, p := range processes {
		procs = append(procs, model.Proc{
			PID:          p.Pid(),
			PPID:         p.PPid(),
			PROCESS_NAME: p.Executable(),
		})
	}

	return procs
}
