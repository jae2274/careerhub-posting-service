package repo

import (
	"context"
	"testing"

	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/provider_grpc/rpcRepo"
	"github.com/jae2274/careerhub-posting-service/careerhub/posting_service/scanner_grpc/repo"
	"github.com/jae2274/careerhub-posting-service/test/tinit"
	"github.com/stretchr/testify/require"
)

func TestSkillNameRepo(t *testing.T) {
	db := tinit.InitDB(t)
	providerRepo := rpcRepo.NewSkillNameRepo(db)
	skillNameRepo := repo.NewSkillNameRepo(db)

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
