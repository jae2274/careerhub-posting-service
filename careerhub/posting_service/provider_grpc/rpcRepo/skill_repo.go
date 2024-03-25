package rpcRepo

import (
	"context"
	"time"

	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/common/domain/skill"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SkillRepo struct {
	col *mongo.Collection
}

func NewSkillRepo(col *mongo.Collection) *SkillRepo {
	return &SkillRepo{
		col: col,
	}
}

// func (cRepo *SkillRepo) FindIDByName(ctx context.Context, skillName string) (*primitive.ObjectID, error) {
// 	var result struct {
// 		ID primitive.ObjectID `bson:"_id"`
// 	}

// 	opts := options.FindOne().SetProjection(bson.D{{Key: skill.Skill_IdField, Value: 1}})
// 	err := cRepo.col.FindOne(ctx, bson.M{skill.Skill_SkillNamesField: skillName}, opts).Decode(&result)

// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	return &result.ID, nil
// }

func (cRepo *SkillRepo) SaveSkills(ctx context.Context, skillNames []string) error {
	now := time.Now()

	skills := make([]any, len(skillNames))
	for i, skillName := range skillNames {
		skills[i] = &skill.Skill{
			DefaultName: skillName,
			SkillNames:  []string{skillName},
			InsertedAt:  now,
			UpdatedAt:   now,
		}
	}

	opts := options.InsertMany().SetOrdered(false) // 중복되는 데이터가 있어도 에러를 내지 않고 나머지 데이터를 저장하도록 설정
	_, err := cRepo.col.InsertMany(ctx, skills, opts)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil
		}
		return err
	}

	return nil
}

// func (cRepo *SkillRepo) AppendSiteSkill(ctx context.Context, skillId primitive.ObjectID, skillName *skill.SkillName) (bool, error) {
// 	now := time.Now()
// 	skillName.InsertedAt = now
// 	skillName.UpdatedAt = now

// 	result, err := cRepo.col.UpdateByID(ctx, skillId, bson.M{
// 		"$push": bson.M{
// 			skill.SkillName_NameField: skillName,
// 		},
// 	})

// 	if err != nil {
// 		return false, err
// 	}

// 	if result.ModifiedCount == 0 {
// 		return false, fmt.Errorf("no document was modified. SkillId: %s", skillId.Hex())
// 	}

// 	return true, nil
// }

func (cRepo *SkillRepo) FindAll() ([]*skill.Skill, error) {
	var companies []*skill.Skill

	cursor, err := cRepo.col.Find(context.Background(), bson.D{})
	if err != nil {
		if mongo.ErrNilDocument == err {
			return []*skill.Skill{}, nil
		}
		return nil, err
	}

	if err := cursor.All(context.Background(), &companies); err != nil {
		return nil, err
	}

	return companies, nil
}
