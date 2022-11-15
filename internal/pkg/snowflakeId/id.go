package snowflakeId

type IdType int

const (
	IdTypeTwitterSF = iota
)

func New(workerId int64) (ISnowflakeId, error) {
	return NewGenerater(IdTypeTwitterSF, workerId)
}

func NewGenerater(idType IdType, workerId int64) (ISnowflakeId, error) {
	switch idType {
	case IdTypeTwitterSF:
		return NewTwitter(workerId)
	default:
		return NewTwitter(workerId)
	}
}
