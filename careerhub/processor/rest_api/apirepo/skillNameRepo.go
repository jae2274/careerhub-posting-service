package apirepo

import (
	"context"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/domain/skill"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SkillNameRepo interface {
	GetAllSkills(context.Context) ([]string, error)
}

type SkillNameRepoImpl struct {
	col *mongo.Collection
}

func NewSkillNameRepo(col *mongo.Collection) SkillNameRepo {
	return &SkillNameRepoImpl{
		col: col,
	}
}

func (repo *SkillNameRepoImpl) GetAllSkills(ctx context.Context) ([]string, error) {
	option := options.Find().SetProjection(bson.M{"_id": 0, skill.SkillName_NameField: 1})
	cursor, err := repo.col.Find(ctx, bson.M{}, option)
	if err != nil {
		return nil, err
	}

	var skills []skill.SkillName
	if err := cursor.All(ctx, &skills); err != nil {
		return nil, err
	}

	skillNames := make([]string, len(skills))
	for i, skill := range skills {
		skillNames[i] = skill.Name
	}

	return skillNames, nil
}
