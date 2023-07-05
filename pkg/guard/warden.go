package guard

import (
	"fmt"
	"github.com/justtrackio/gosoline/pkg/cfg"
	"github.com/justtrackio/gosoline/pkg/db"
	"github.com/justtrackio/gosoline/pkg/log"
	"github.com/ory/ladon"
)

type AuditLogger struct {
	db db.Client
}

func (a AuditLogger) LogRejectedAccessRequest(request *ladon.Request, pool ladon.Policies, deciders ladon.Policies) {
	fmt.Printf(" Access Rejected request: %v, poll: %v, deciders: %v\n")
}

func (a AuditLogger) LogGrantedAccessRequest(request *ladon.Request, pool ladon.Policies, deciders ladon.Policies) {
	fmt.Printf(" Access Rejected request: %v, poll: %v, deciders: %v\n")
}

type WardenOption func(warden *ladon.Ladon) error

// todo add comment
func WithBaseAuditLog(config cfg.Config, logger log.Logger) WardenOption {
	return func(warden *ladon.Ladon) error {
		warden.AuditLogger = AuditLogger{}
		return nil
	}
}

// WithBaseSqlManager  This method will create a sqlBase manager for warden from config and attach it to ladon.Ladon.
func WithBaseSqlManager(config cfg.Config, logger log.Logger) WardenOption {
	return func(warden *ladon.Ladon) error {
		sqlManager, err := NewSqlManager(config, logger)
		if err != nil {
			return fmt.Errorf("can not create sqlManager: %w", err)
		}
		warden.Manager = sqlManager

		return nil
	}
}

//todo need add new option for sql manager (just provide you sql manager)

func NewWarden(config cfg.Config, logger log.Logger, opt ...WardenOption) (*ladon.Ladon, error) {
	var err error
	warden := &ladon.Ladon{}

	for _, v := range opt {
		err = v(warden)
		if err != nil {
			return nil, fmt.Errorf("error do options for NewWarden: %w", err)
		}
	}

	return warden, nil
}
