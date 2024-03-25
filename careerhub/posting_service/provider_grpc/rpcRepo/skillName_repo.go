package rpcRepo

import (
	"context"
	"time"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/skill"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SkillNameRepo struct {
	col *mongo.Collection
}

func NewSkillNameRepo(col *mongo.Collection) *SkillNameRepo {
	return &SkillNameRepo{
		col: col,
	}
}

func (cRepo *SkillNameRepo) SaveSkillNames(ctx context.Context, skillNames []string) error {
	now := time.Now()
	skillNameStructs := make([]any, len(skillNames))
	for i, skillName := range skillNames {
		skillNameStructs[i] = skill.SkillName{
			Name:           skillName,
			InsertedAt:     now,
			UpdatedAt:      now,
			IsScanComplete: false,
		}
	}

	opts := options.InsertMany().SetOrdered(false) // 중복되는 데이터가 있어도 에러를 내지 않고 나머지 데이터를 저장하도록 설정
	_, err := cRepo.col.InsertMany(ctx, skillNameStructs, opts)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil
		}
		return err
	}

	return nil
}

func (cRepo *SkillNameRepo) FindAll() ([]*skill.SkillName, error) {
	var skillNames []*skill.SkillName

	cursor, err := cRepo.col.Find(context.Background(), bson.D{})
	if err != nil {
		if mongo.ErrNilDocument == err {
			return []*skill.SkillName{}, nil
		}
		return nil, err
	}

	if err := cursor.All(context.Background(), &skillNames); err != nil {
		return nil, err
	}

	return skillNames, nil
}
