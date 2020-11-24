package cypher_go_dsl

type MatchBuilder struct {
	patternList []PatternElement
	condition   ConditionBuilder
	optional    bool
}

func (builder MatchBuilder) buildMatch() Match  {
	pattern := Pattern{PatternElements: builder.patternList}
	conditionBuilder := builder.condition
	var optionalWhere *Where = nil
	if conditionBuilder.condition != nil {
		builtCondition := *conditionBuilder.BuildCondition()
		optionalWhere = &Where{condition: builtCondition}
	}
	return Match{
		optional: builder.optional,
		pattern: pattern,
		optionalWhere: optionalWhere,
	}
}
