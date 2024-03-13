package apirepo

import (
	"context"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/skill"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SkillRepo interface {
	GetAllSkill(ctx context.Context, skillName []string) ([]*skill.Skill, error)
}

type SkillRepoImpl struct {
	col *mongo.Collection
}

func NewSkillRepo(skillCollection *mongo.Collection) SkillRepo {
	return &SkillRepoImpl{
		col: skillCollection,
	}
}

func (repo *SkillRepoImpl) GetAllSkill(ctx context.Context, skillName []string) ([]*skill.Skill, error) {
	var result []*skill.Skill
	cursor, err := repo.col.Find(ctx, bson.M{skill.Skill_SkillNamesField: bson.M{"$in": skillName}})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}
