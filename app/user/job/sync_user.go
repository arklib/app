package job

import (
	"fmt"

	"github.com/arklib/ark/job"
)

func (it *Job) SyncUser(p job.Payload[uint]) error {
	fmt.Println("sync_user", *p.Data)
	return nil
}
