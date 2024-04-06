package apirepo

import (
	"context"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/skill"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SkillRepo interface {
	GetAllSkills(context.Context) ([]skill.Skill, error)
}

type SkillRepoImpl struct {
	col *mongo.Collection
}

func NewSkillRepo(col *mongo.Collection) SkillRepo {
	return &SkillRepoImpl{
		col: col,
	}
}

func (repo *SkillRepoImpl) GetAllSkills(ctx context.Context) ([]skill.Skill, error) {
	option := options.Find().SetProjection(bson.M{"_id": 0, skill.Skill_DefaultNameField: 1, skill.Skill_SkillNamesField: 1})
	cursor, err := repo.col.Find(ctx, bson.M{}, option)
	if err != nil {
		return nil, err
	}

	var skills []skill.Skill
	if err := cursor.All(ctx, &skills); err != nil {
		return nil, err
	}

	return skills, nil
}
