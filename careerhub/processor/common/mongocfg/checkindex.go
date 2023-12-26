package mongocfg

import (
	"context"

	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckIndex(indexes []bson.M, indexModels map[string]*mongo.IndexModel) error {
	if len(indexes)-1 != len(indexModels) { // -1 because of _id_
		return terr.New("invalid index")
	}

	isValidIndex, err := checkIndexes(indexes, indexModels)

	if err != nil {
		return err
	}

	if !isValidIndex {
		return terr.New("invalid index")
	}

	return nil
}

func IndexesFromCursor(cursor *mongo.Cursor) ([]bson.M, error) {
	var indexes []bson.M

	for cursor.Next(context.TODO()) {
		var index bson.M
		err := cursor.Decode(&index)
		if err != nil {
			return nil, terr.Wrap(err)
		}

		indexes = append(indexes, index)
	}

	return indexes, nil
}

func checkIndexes(indexes []bson.M, indexModels map[string]*mongo.IndexModel) (bool, error) {
	for _, indexSpec := range indexes {
		indexName, ok := indexSpec["name"].(string)
		if !ok {
			return false, terr.New("invalid index")
		}

		if indexName == "_id_" {
			continue
		}

		indexModel, ok := indexModels[indexName]
		if !ok {
			return false, nil
		}

		isEqual, err := isEqualIndex(indexSpec, indexModel)
		if !isEqual {
			return false, err
		}
	}

	return true, nil
}

func isEqualIndex(indexSpec bson.M, indexModel *mongo.IndexModel) (bool, error) {
	// fmt.Println(indexSpec)
	// fmt.Println(indexModel)
	modelKey, ok := indexModel.Keys.(bson.D)
	if !ok {
		return false, terr.New("invalid index")
	}

	specKey, ok := indexSpec["key"].(bson.M)

	if !ok {
		return false, terr.New("invalid index")
	}

	if len(modelKey) != len(specKey) {
		return false, nil
	}

	for _, modelKeyElem := range modelKey {
		if specKey[modelKeyElem.Key] != modelKeyElem.Value {
			return false, nil
		}
	}

	modelUnique, existedModelUnique := false, false
	if indexModel.Options != nil && indexModel.Options.Unique != nil {
		existedModelUnique = *indexModel.Options.Unique
		modelUnique = *indexModel.Options.Unique
	}

	specUnique, existedSpecUnique := false, false
	uniqueM, existedSpecUnique := indexSpec["unique"]

	if existedSpecUnique {
		specUnique, ok = uniqueM.(bool)
		if !ok {
			return false, terr.New("invalid index")
		}
	}

	if existedModelUnique != existedSpecUnique {
		return false, nil
	}

	if modelUnique != specUnique {
		return false, nil
	}

	return true, nil
}
