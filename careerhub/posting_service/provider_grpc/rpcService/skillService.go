package rpcService

import (
	"context"
	"regexp"
	"strings"

	"github.com/jae2274/Careerhub-posting-service/careerhub/posting_service/provider_grpc/rpcRepo"
)

type SkillService struct {
	skillRepo     *rpcRepo.SkillRepo
	skillNameRepo *rpcRepo.SkillNameRepo
}

func NewSkillService(skillRepo *rpcRepo.SkillRepo, skillNameRepo *rpcRepo.SkillNameRepo) *SkillService {
	return &SkillService{
		skillRepo:     skillRepo,
		skillNameRepo: skillNameRepo,
	}
}

func (s *SkillService) RegisterSkill(ctx context.Context, skillNames []string) ([]string, error) {
	if len(skillNames) == 0 {
		return []string{}, nil
	}

	skillNames = preprocessSkillNames(skillNames)

	err := s.skillNameRepo.SaveSkillNames(ctx, skillNames)
	if err != nil {
		return skillNames, err
	}

	err = s.skillRepo.SaveSkills(ctx, skillNames)
	if err != nil {
		return skillNames, err
	}

	return skillNames, err
}

func preprocessSkillNames(skillNames []string) []string {

	for i, skillName := range skillNames {
		skillNames[i] = preprocessSkillName(skillName)
	}

	skillNames = removeEmpty(skillNames)
	skillNames = lowerCase(skillNames)
	skillNames = removeDuplicate(skillNames)
	return skillNames
}

func preprocessSkillName(skillName string) string {
	skillName = strings.TrimSpace(skillName)
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

func lowerCase(skillNames []string) []string {
	for i, skillName := range skillNames {
		skillNames[i] = strings.ToLower(skillName)
	}
	return skillNames
}
