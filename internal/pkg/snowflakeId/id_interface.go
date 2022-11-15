package snowflakeId

type ISnowflakeId interface {
	Generate() ID
}

type ID interface {
	Int64() int64
	String() string
}
