package rpcService

import (
	"context"
	"testing"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/grpc/rpcService"
	"github.com/jae2274/Careerhub-dataProcessor/test/tinit"
	"github.com/stretchr/testify/require"
)

func TestSkillService(t *testing.T) {
	t.Run("RegisterSkill", func(t *testing.T) {
		skillRepo := tinit.InitSkillRepo(t)
		skillService := rpcService.NewSkillService(skillRepo)

		sampleSkillGroups := [][]string{
			{"Python", "C++", "Linux", "aws    eks", "linux", " java ", ""}, //소문자 통일, 공백 제거, 중복 제거, 빈 문자열 제거
			{"pyThon", "kubernetes", "aws"},                                 //소문자 통일
			{"k8s", "aws eks", "eks"},
		}

		expectedSkillGroupResults := [][]string{
			{"python", "c++", "linux", "aws eks", "java"},
			{"python", "kubernetes", "aws"},
			{"k8s", "aws eks", "eks"},
		}

		for i, sampleSkillGroup := range sampleSkillGroups {
			skillGroupResult, err := skillService.RegisterSkill(context.TODO(), sampleSkillGroup)
			require.NoError(t, err)

			require.Equal(t, expectedSkillGroupResults[i], skillGroupResult)
		}

		allSavedSkills, err := skillRepo.FindAll()
		require.NoError(t, err)

		allSavedSkillNames := make([]string, len(allSavedSkills))
		for i, savedSkill := range allSavedSkills {
			allSavedSkillNames[i] = savedSkill.DefaultName
		}

		expectedAllSavedSkills := []string{"python", "c++", "linux", "aws eks", "java", "kubernetes", "aws", "k8s", "eks"}
		require.Equal(t, expectedAllSavedSkills, allSavedSkillNames)
	})
}
