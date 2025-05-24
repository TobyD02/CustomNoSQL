package Interface

type BSONSerializableInterface interface {
	ToBSON() ([]byte, error)
	FromBson([]byte) error
}
