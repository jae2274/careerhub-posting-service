package repo

import (
	"context"
)

type SkillNameRepo interface {
	GetSkills(ctx context.Context, isScanComplete bool) ([]string, error)
	SetScanComplete(ctx context.Context, skillNames []string) error
}
