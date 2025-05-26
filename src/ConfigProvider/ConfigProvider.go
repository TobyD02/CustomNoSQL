package ConfigProvider

type Config struct {
	CollectionFilePath   string
	RecordsDirectoryPath string
	IndexesFilePath      string
}

func (c *Config) GetConfig() {
	c.CollectionFilePath = "./data/collections.bson"
	c.RecordsDirectoryPath = "./data/records/"
	c.IndexesFilePath = "./data/indexes.bson"
}
