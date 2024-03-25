package repo

import (
	"context"
	"testing"

	"github.com/jae2274/careerhub-posting-service/test/tinit"
	"github.com/stretchr/testify/require"
)

func TestSkillNameRepo(t *testing.T) {
	providerRepo := tinit.InitProviderSkillNameRepo(t)
	skillNameRepo := tinit.InitScannerSkillNameRepo(t)

	savedSkillNames := []string{"java", "python", "go"}
	err := providerRepo.SaveSkillNames(context.Background(), savedSkillNames)
	require.NoError(t, err)

	skillNames, err := skillNameRepo.GetSkills(context.Background(), false)
	require.NoError(t, err)
	require.Equal(t, savedSkillNames, skillNames)

	skillNames, err = skillNameRepo.GetSkills(context.Background(), true)
	require.NoError(t, err)
	require.Empty(t, skillNames)

	err = skillNameRepo.SetScanComplete(context.Background(), []string{"java", "go"})
	require.NoError(t, err)

	skillNames, err = skillNameRepo.GetSkills(context.Background(), false)
	require.NoError(t, err)
	require.Equal(t, []string{"python"}, skillNames)

	skillNames, err = skillNameRepo.GetSkills(context.Background(), true)
	require.NoError(t, err)
	require.Equal(t, []string{"java", "go"}, skillNames)

	err = skillNameRepo.SetScanComplete(context.Background(), []string{"python"})
	require.NoError(t, err)

	skillNames, err = skillNameRepo.GetSkills(context.Background(), false)
	require.NoError(t, err)
	require.Empty(t, skillNames)

	skillNames, err = skillNameRepo.GetSkills(context.Background(), true)
	require.NoError(t, err)
	require.Equal(t, savedSkillNames, skillNames)

	err = providerRepo.SaveSkillNames(context.Background(), []string{"java", "python", "go", "c++"}) //이미 저장된 스킬이라서 3개는 저장되지 않음
	require.NoError(t, err)

	skillNames, err = skillNameRepo.GetSkills(context.Background(), false)
	require.NoError(t, err)
	require.Equal(t, []string{"c++"}, skillNames)

	skillNames, err = skillNameRepo.GetSkills(context.Background(), true)
	require.NoError(t, err)
	require.Equal(t, savedSkillNames, skillNames)
}
