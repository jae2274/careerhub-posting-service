package mongocfg

import (
	"testing"

	"github.com/jae2274/Careerhub-dataProcessor/careerhub/processor/common/mongocfg"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestCheckIndex(t *testing.T) {
	indexModel := map[string]*mongo.IndexModel{
		"match_id_1_placement_1": {
			Keys: bson.D{
				{Key: "match_id", Value: 1},
				{Key: "placement", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
		"match_id_1_puuid_1": {
			Keys: bson.D{
				{Key: "match_id", Value: 1},
				{Key: "puuid", Value: 1},
			},
		},
		"info_game_datetime_1": {
			Keys: bson.D{
				{Key: "info.game_datetime", Value: 1},
			},
		},
	}

	t.Run("Valid indexes", func(t *testing.T) {
		indexes := []bson.M{
			{
				"v":    2,
				"key":  bson.M{"_id": 1},
				"name": "_id_",
			},
			{
				"v":      2,
				"key":    bson.M{"match_id": 1, "placement": 1},
				"name":   "match_id_1_placement_1",
				"unique": true,
			},
			{
				"v":    2,
				"key":  bson.M{"match_id": 1, "puuid": 1},
				"name": "match_id_1_puuid_1",
			},
			{
				"v":    2,
				"key":  bson.M{"info.game_datetime": 1},
				"name": "info_game_datetime_1",
			},
		}

		err := mongocfg.CheckIndex(indexes, indexModel)
		require.NoError(t, err)
	})

	invalidTestCases := []struct {
		name    string
		indexes []bson.M
	}{
		{
			"Invalid indexes: less index count",
			[]bson.M{
				{
					"v":    2,
					"key":  bson.M{"_id": 1},
					"name": "_id_",
				},
				{
					"v":      2,
					"key":    bson.M{"match_id": 1, "placement": 1},
					"name":   "match_id_1_placement_1",
					"unique": true,
				}, // match_id_1_puuid_1 is missing
				{
					"v":    2,
					"key":  bson.M{"info.game_datetime": 1},
					"name": "info_game_datetime_1",
				},
			},
		},
		{
			"Invalid indexes: more index count",
			[]bson.M{
				{
					"v":    2,
					"key":  bson.M{"_id": 1},
					"name": "_id_",
				},
				{
					"v":      2,
					"key":    bson.M{"match_id": 1, "placement": 1},
					"name":   "match_id_1_placement_1",
					"unique": true,
				},
				{
					"v":    2,
					"key":  bson.M{"match_id": 1, "puuid": 1},
					"name": "match_id_1_puuid_1",
				},
				{
					"v":    2,
					"key":  bson.M{"info.game_datetime": 1},
					"name": "info_game_datetime_1",
				},
				{
					"v":    2,
					"key":  bson.M{"info.game_datetime": 1},
					"name": "info_game_datetime_1",
				}, // info_game_datetime_1 is extra
			},
		},
		{
			"Invalid indexes: invalid index name",
			[]bson.M{
				{
					"v":    2,
					"key":  bson.M{"_id": 1},
					"name": "_id_",
				},
				{
					"v":      2,
					"key":    bson.M{"match_id": 1, "placement": 1},
					"name":   "match_id_and_placement_1", // match_id_and_placement_1 is invalid index name
					"unique": true,
				},
				{
					"v":    2,
					"key":  bson.M{"match_id": 1, "puuid": 1},
					"name": "match_id_1_puuid_1",
				},
				{
					"v":    2,
					"key":  bson.M{"info.game_datetime": 1},
					"name": "info_game_datetime_1",
				},
			},
		},
		{
			"Invalid indexes: invalid index key",
			[]bson.M{
				{
					"v":    2,
					"key":  bson.M{"_id": 1},
					"name": "_id_",
				},
				{
					"v":      2,
					"key":    bson.M{"match_id": 1, "level": 1}, // level is invalid index key
					"name":   "match_id_1_placement_1",
					"unique": true,
				},
				{
					"v":    2,
					"key":  bson.M{"match_id": 1, "puuid": 1},
					"name": "match_id_1_puuid_1",
				},
				{
					"v":    2,
					"key":  bson.M{"info.game_datetime": 1},
					"name": "info_game_datetime_1",
				},
			},
		},
		{
			"Invalid indexes: invalid single index key",
			[]bson.M{
				{
					"v":    2,
					"key":  bson.M{"_id": 1},
					"name": "_id_",
				},
				{
					"v":      2,
					"key":    bson.M{"match_id": 1}, // placement is missing
					"name":   "match_id_1_placement_1",
					"unique": true,
				},
				{
					"v":    2,
					"key":  bson.M{"match_id": 1, "puuid": 1},
					"name": "match_id_1_puuid_1",
				},
				{
					"v":    2,
					"key":  bson.M{"info.game_datetime": 1},
					"name": "info_game_datetime_1",
				},
			},
		},
		{
			"Invalid indexes: invalid triple index key",
			[]bson.M{
				{
					"v":    2,
					"key":  bson.M{"_id": 1},
					"name": "_id_",
				},
				{
					"v":      2,
					"key":    bson.M{"match_id": 1, "placement": 1, "level": 1}, // level is invalid index key
					"name":   "match_id_1_placement_1",
					"unique": true,
				},
				{
					"v":    2,
					"key":  bson.M{"match_id": 1, "puuid": 1},
					"name": "match_id_1_puuid_1",
				},
				{
					"v":    2,
					"key":  bson.M{"info.game_datetime": 1},
					"name": "info_game_datetime_1",
				},
			},
		},
		{
			"Invalid indexes: invalid index sorted",
			[]bson.M{
				{
					"v":    2,
					"key":  bson.M{"_id": 1},
					"name": "_id_",
				},
				{
					"v":      2,
					"key":    bson.M{"match_id": 1, "placement": -1}, // placement sorted is invalid
					"name":   "match_id_1_placement_1",
					"unique": true,
				},
				{
					"v":    2,
					"key":  bson.M{"match_id": 1, "puuid": 1},
					"name": "match_id_1_puuid_1",
				},
				{
					"v":    2,
					"key":  bson.M{"info.game_datetime": 1},
					"name": "info_game_datetime_1",
				},
			},
		},
		{
			"Invalid indexes: not unique",
			[]bson.M{
				{
					"v":    2,
					"key":  bson.M{"_id": 1},
					"name": "_id_",
				},
				{
					"v":    2,
					"key":  bson.M{"match_id": 1, "placement": 1},
					"name": "match_id_1_placement_1", // This index needs to be unique
				},
				{
					"v":    2,
					"key":  bson.M{"match_id": 1, "puuid": 1},
					"name": "match_id_1_puuid_1",
				},
				{
					"v":    2,
					"key":  bson.M{"info.game_datetime": 1},
					"name": "info_game_datetime_1",
				},
			},
		},
		{
			"Invalid indexes: unique",
			[]bson.M{
				{
					"v":    2,
					"key":  bson.M{"_id": 1},
					"name": "_id_",
				},
				{
					"v":      2,
					"key":    bson.M{"match_id": 1, "placement": 1},
					"name":   "match_id_1_placement_1",
					"unique": true,
				},
				{
					"v":      2,
					"key":    bson.M{"match_id": 1, "puuid": 1},
					"name":   "match_id_1_puuid_1",
					"unique": true, // This index needs to be not unique
				},
				{
					"v":    2,
					"key":  bson.M{"info.game_datetime": 1},
					"name": "info_game_datetime_1",
				},
			},
		},
	}

	for _, tc := range invalidTestCases {
		t.Run(tc.name, func(t *testing.T) {
			err := mongocfg.CheckIndex(tc.indexes, indexModel)
			require.Error(t, err)
		})
	}

}
