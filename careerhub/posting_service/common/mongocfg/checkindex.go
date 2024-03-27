package mongocfg

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBModel interface {
	Collection() string
	IndexModels() map[string]*mongo.IndexModel
}

type InvalidIndexError struct {
	Collection string
	error
}

func (e InvalidIndexError) Error() string {
	return fmt.Sprintf("in collection %s: %s", e.Collection, e.error.Error())
}

func CheckIndex(indexes []bson.M, indexModels map[string]*mongo.IndexModel) error {
	if len(indexes)-1 != len(indexModels) { // -1 because of _id_
		return fmt.Errorf("index length is not equal. indexes: " + fmt.Sprint(indexes) + ", indexModels: " + fmt.Sprint(indexModels))
	}

	err := checkIndexes(indexes, indexModels)

	if err != nil {
		return err
	}

	return nil
}

func CheckIndexViaCollection(col *mongo.Collection, model MongoDBModel) error {
	cursor, err := col.Indexes().List(context.TODO())
	if err != nil {
		return InvalidIndexError{Collection: model.Collection(), error: err}
	}

	indexes, err := indexesFromCursor(cursor)
	if err != nil {
		return InvalidIndexError{Collection: model.Collection(), error: err}
	}

	err = CheckIndex(indexes, model.IndexModels())
	if err != nil {
		return InvalidIndexError{Collection: model.Collection(), error: err}
	}

	return nil
}

func indexesFromCursor(cursor *mongo.Cursor) ([]bson.M, error) {
	var indexes []bson.M

	for cursor.Next(context.TODO()) {
		var index bson.M
		err := cursor.Decode(&index)
		if err != nil {
			return nil, err
		}

		indexes = append(indexes, index)
	}

	return indexes, nil
}

func checkIndexes(indexes []bson.M, indexModels map[string]*mongo.IndexModel) error {
	for _, indexSpec := range indexes {
		indexName, ok := indexSpec["name"].(string)
		if !ok {
			return fmt.Errorf("name is not string")
		}

		if indexName == "_id_" {
			continue
		}

		indexModel, ok := indexModels[indexName]
		if !ok {
			return fmt.Errorf("index name is not exist. indexName:%s", indexName)
		}

		err := isEqualIndex(indexName, indexSpec, indexModel)

		if err != nil {
			return err
		}
	}

	return nil
}

func isEqualIndex(indexName string, indexSpec bson.M, indexModel *mongo.IndexModel) error {
	// fmt.Println(indexSpec)
	// fmt.Println(indexModel)
	modelKey, ok := indexModel.Keys.(bson.D)
	if !ok {
		return fmt.Errorf("%s. keys is not bson.D", indexName)
	}

	specKey, ok := indexSpec["key"].(bson.M)

	if !ok {
		return fmt.Errorf("%s. key is not bson.M", indexName)
	}

	if len(modelKey) != len(specKey) {
		return fmt.Errorf("%s. key length is not equal. len(modelKey): %d, len(specKey): %d", indexName, len(modelKey), len(specKey))
	}

	for _, modelKeyElem := range modelKey {
		var specValue int

		switch v := specKey[modelKeyElem.Key].(type) {
		case int32:
			specValue = int(v)
		case int:
			specValue = v
		default:
			return fmt.Errorf("unsupported type %T for specKey[modelKeyElem.Key]", v)
		}

		modelValue := modelKeyElem.Value.(int)

		if specValue != modelValue {
			return fmt.Errorf("%s. key is not equal. specKey[modelKeyElem.Key]: %v, modelKeyElem.Value: %v", indexName, specValue, modelValue)
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
			return fmt.Errorf("%s. unique is not bool", indexName)
		}
	}

	if existedModelUnique != existedSpecUnique {
		return fmt.Errorf("%s. unique is not equal. existedModelUnique:%v, existedSpecUnique:%v", indexName, existedModelUnique, existedSpecUnique)
	}

	if modelUnique != specUnique {
		return fmt.Errorf("%s. unique is not equal. modelUnique:%v, specUnique:%v", indexName, modelUnique, specUnique)
	}

	return nil
}
