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
			{"Python", "C++", "Linux", "aws    eks", "Linux", "linux", " java ", ""}, //공백 치환, 중복 제거, 빈 문자열 제거
			{"pyThon", "kubernetes", "aws", "java"},
			{"k8s", "aws eks", "eks"},
		}

		expectedSkillGroupResults := [][]string{
			{"Python", "C++", "Linux", "aws eks", "linux", "java"}, //대소문자가 다른 Linux, linux는 별도의 skill로 간주된다.
			{"pyThon", "kubernetes", "aws", "java"},
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

		expectedAllSavedSkills := []string{"Python", "C++", "Linux", "aws eks", "linux", "java", "pyThon", "kubernetes", "aws", "k8s", "eks"} //대소문자는 구분되어 중복 제외되지 않는다.
		require.Equal(t, expectedAllSavedSkills, allSavedSkillNames)
	})
}