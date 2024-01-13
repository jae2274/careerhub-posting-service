package rpcService

import (
	"context"
	"regexp"
	"strings"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/rpcRepo"
)

type SkillService struct {
	skillRepo *rpcRepo.SkillRepo
}

func NewSkillService(skillRepo *rpcRepo.SkillRepo) *SkillService {
	return &SkillService{
		skillRepo: skillRepo,
	}
}

func (s *SkillService) RegisterSkill(ctx context.Context, skillNames []string) ([]string, error) {
	skillNames = preprocessSkillNames(skillNames)
	err := s.skillRepo.SaveSkills(ctx, skillNames)

	return skillNames, err
}

func preprocessSkillNames(skillNames []string) []string {

	for i, skillName := range skillNames {
		skillNames[i] = preprocessSkillName(skillName)
	}

	skillNames = removeEmpty(skillNames)
	skillNames = removeDuplicate(skillNames)
	return skillNames
}

func preprocessSkillName(skillName string) string {
	skillName = strings.TrimSpace(skillName)
	skillName = strings.ToLower(skillName)
	skillName = regexp.MustCompile(`\s+`).ReplaceAllString(skillName, " ")
	return skillName
}

func removeEmpty(skillNames []string) []string {
	nonEmptySkillNames := make([]string, 0, len(skillNames))
	for _, skillName := range skillNames {
		if skillName != "" {
			nonEmptySkillNames = append(nonEmptySkillNames, skillName)
		}
	}
	return nonEmptySkillNames
}

func removeDuplicate(skillNames []string) []string {
	uniqueSkillNames := make(map[string]bool)
	resultSkillNames := make([]string, 0, len(skillNames))

	for _, skillName := range skillNames {
		if _, ok := uniqueSkillNames[skillName]; !ok {
			uniqueSkillNames[skillName] = true
			resultSkillNames = append(resultSkillNames, skillName)
		}
	}

	return resultSkillNames
}
