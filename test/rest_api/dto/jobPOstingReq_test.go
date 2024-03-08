package dto

import (
	"testing"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/rest_api/dto"
	"github.com/stretchr/testify/require"
)

func TestDecodeQuery(t *testing.T) {

	_, err := dto.GetQuery("eyJwYWdlIjozLCJzaXplIjoxNiwiY2F0ZWdvcmllcyI6W3sic2l0ZSI6IndhbnRlZCIsImNhdGVnb3J5TmFtZSI6IlZSIOyXlOyngOuLiOyWtCJ9XSwic2tpbGxOYW1lcyI6W10sInRhZ0lkcyI6W10sIm1pbkNhcmVlciI6MCwibWF4Q2FyZWVyIjowfQ==")
	require.NoError(t, err)
}
