package cypher_go_dsl

type ExposesPropertiesNode interface {
	WithRawProperties(keysAndValues ...interface{}) Node

	WithProperties(newProperties MapExpression) Node
}

type ExposesPropertiesRelation interface {
	WithRawProperties(keysAndValues ...interface{}) Relationship

	WithProperties(newProperties MapExpression) Relationship
}
