package reset

// import (
// 	"testing"

// 	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/vars"
// 	"github.com/jae2274/Careerhub-dataProcessor/test/tinit"
// 	"github.com/stretchr/testify/require"
// )

// func TestReset(t *testing.T) {
// 	tinit.InitDB(t)
// 	envVars, err := vars.Variables()
// 	require.NoError(t, err)

// 	tinit.InitSQS(t, envVars.JobPostingQueue)
// 	tinit.InitSQS(t, envVars.CompanyQueue)
// 	tinit.InitSQS(t, envVars.ClosedQueue)
// }
