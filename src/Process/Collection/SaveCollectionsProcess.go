package Collection

import (
	"CustomNoSQL/src/ConfigProvider"
	"CustomNoSQL/src/Model/Entity"
	"os"
	"path/filepath"
)

func SaveCollectionsProcess(collections Entity.AllCollectionsModel) error {
	var config ConfigProvider.Config
	config.GetConfig()

	dir := filepath.Dir(config.CollectionFilePath)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	data, err := collections.ToBson()

	if err != nil {
		return err
	}

	err = os.WriteFile(config.CollectionFilePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
