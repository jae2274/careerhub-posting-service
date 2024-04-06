package repo

import (
	"context"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/common/domain/skill"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SkillNameRepo interface {
	GetSkills(ctx context.Context, isScanComplete bool) ([]string, error)
	SetScanComplete(ctx context.Context, skillNames []string) error
}

type SkillNameRepoImpl struct {
	col *mongo.Collection
}

func NewSkillNameRepo(db *mongo.Database) SkillNameRepo {
	col := db.Collection((&skill.SkillName{}).Collection())
	return &SkillNameRepoImpl{col: col}
}

func (r *SkillNameRepoImpl) GetSkills(ctx context.Context, isScanComplete bool) ([]string, error) {
	options := options.Find().SetProjection(bson.M{skill.SkillName_NameField: 1})

	cursor, err := r.col.Find(ctx, bson.M{skill.SkillName_IsScanCompleteField: isScanComplete}, options)
	if err != nil {
		return nil, err
	}

	var skillNames []string

	for cursor.Next(ctx) {
		var skillName skill.SkillName
		if err := cursor.Decode(&skillName); err != nil {
			return nil, err
		}
		skillNames = append(skillNames, skillName.Name)
	}

	return skillNames, nil
}

func (r *SkillNameRepoImpl) SetScanComplete(ctx context.Context, skillNames []string) error {
	_, err := r.col.UpdateMany(ctx, bson.M{skill.SkillName_NameField: bson.M{"$in": skillNames}}, bson.M{"$set": bson.M{skill.SkillName_IsScanCompleteField: true}})

	return err
}
