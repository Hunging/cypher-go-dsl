package cypher_go_dsl

import (
	"errors"
	"fmt"
)

type AliasedExpression struct {
	delegate Expression
	alias    string
	key      string
	notNil   bool
	err      error
}

func (aliased AliasedExpression) getError() error {
	return aliased.err
}

func AliasedExpressionCreate(delegate Expression, alias string) AliasedExpression {
	return AliasedExpression{
		delegate: delegate,
		alias:    alias,
		notNil:   true,
	}
}

func AliasedExpressionError(err error) AliasedExpression {
	return AliasedExpression{
		err: err,
	}
}

func (aliased AliasedExpression) isNotNil() bool {
	return aliased.notNil
}

func (aliased AliasedExpression) GetExpressionType() ExpressionType {
	return EXPRESSION
}

func (aliased AliasedExpression) As(newAlias string) AliasedExpression {
	if newAlias == "" {
		return AliasedExpressionError(errors.New("the alias may not be empty"))
	}
	return AliasedExpressionCreate(aliased.delegate, newAlias)
}

func (aliased AliasedExpression) accept(visitor *CypherRenderer) {
	aliased.key = fmt.Sprint(&aliased)
	(*visitor).enter(aliased)
	NameOrExpression(aliased.delegate).accept(visitor)
	(*visitor).leave(aliased)
}

func (aliased AliasedExpression) getKey() string {
	return aliased.key
}

func (aliased AliasedExpression) enter(renderer *CypherRenderer) {
	if _, visited := renderer.visitableToAliased[aliased.key]; visited {
		renderer.append(escapeName(aliased.alias))
	}
}

func (aliased AliasedExpression) leave(renderer *CypherRenderer) {
	if _, visited := renderer.visitableToAliased[aliased.key]; !visited {
		renderer.append(" AS ")
		renderer.append(escapeName(aliased.alias))
	}
}
