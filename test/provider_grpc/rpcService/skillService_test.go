package rpcService

import (
	"context"
	"testing"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/provider_grpc/rpcService"
	"github.com/jae2274/Careerhub-dataProcessor/test/tinit"
	"github.com/stretchr/testify/require"
)

func TestSkillService(t *testing.T) {
	t.Run("RegisterSkill", func(t *testing.T) {
		skillRepo := tinit.InitProviderSkillRepo(t)
		skillNameRepo := tinit.InitProviderSkillNameRepo(t)
		skillService := rpcService.NewSkillService(skillRepo, skillNameRepo)

		sampleSkillGroups := [][]string{
			{"Python", "C++", "Linux", "aws    eks", "Linux", "linux", " java ", ""}, //공백 치환, 중복 제거, 빈 문자열 제거
			{"pyThon", "kubernetes", "aws", "java"},
			{"k8s", "aws eks", "eks"},
		}

		expectedSkillGroupResults := [][]string{
			{"python", "c++", "linux", "aws eks", "java"}, //대문자는 소문자로 변환되어 중복 제외
			{"python", "kubernetes", "aws", "java"},
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

		expectedAllSavedSkills := []string{"python", "c++", "linux", "aws eks", "java", "kubernetes", "aws", "k8s", "eks"} //대소문자는 구분되어 중복 제외되지 않는다.
		require.Equal(t, expectedAllSavedSkills, allSavedSkillNames)

		allSavedSkillNameStructs, err := skillNameRepo.FindAll()
		require.NoError(t, err)

		skillNameStructStrings := make([]string, len(allSavedSkillNameStructs))
		for i, skillNameStruct := range allSavedSkillNameStructs {
			skillNameStructStrings[i] = skillNameStruct.Name
		}

		require.Equal(t, expectedAllSavedSkills, skillNameStructStrings)
	})
}
