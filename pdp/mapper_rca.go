package pdp

import (
	"fmt"

	"github.com/infobloxopen/go-trees/strtree"
)

type mapperRCA struct {
	argument  Expression
	rules     *strtree.Tree
	def       *Rule
	err       *Rule
	algorithm RuleCombiningAlg
}

// MapperRCAParams gathers all parameters of mapper rule combining algorithm.
type MapperRCAParams struct {
	// Argument represent expression which value is used to get nested rule
	// (or list of them).
	Argument Expression

	// DefOk indicates if Def contains valid value.
	DefOk bool
	// Def contains id of default rule (the default rule is used when Argument
	// expression evaluates to a value which doesn't match to any id).
	// This value is used only if DefOk is true.
	Def string

	// ErrOk indicateis if Err contains valid value.
	ErrOk bool
	// Err ontains id of rule to use in case of error (when  Argument can't be
	// evaluated).
	Err string

	// Algorithm is additional rule combining algorithm which is used when
	// argument can return several ids.
	Algorithm RuleCombiningAlg
}

func getSetOfIDs(v AttributeValue) ([]string, error) {
	switch v.t {
	case TypeString:
		ID, err := v.str()
		if err != nil {
			panic(err)
		}

		return []string{ID}, nil

	case TypeSetOfStrings:
		setIDs, err := v.setOfStrings()
		if err != nil {
			panic(err)
		}

		return sortSetOfStrings(setIDs), nil

	case TypeListOfStrings:
		listIDs, err := v.listOfStrings()
		if err != nil {
			panic(err)
		}

		return listIDs, nil
	}

	return nil, newMapperArgumentTypeError(v.t)
}

func collectSubRules(IDs []string, m *strtree.Tree) []*Rule {
	rules := []*Rule{}
	for _, ID := range IDs {
		rule, ok := m.Get(ID)
		if ok {
			rules = append(rules, rule.(*Rule))
		}
	}

	return rules
}

func makeMapperRCA(rules []*Rule, params interface{}) RuleCombiningAlg {
	mapperParams, ok := params.(MapperRCAParams)
	if !ok {
		panic(fmt.Errorf("Mapper rule combining algorithm maker expected MapperRCAParams structure as params "+
			"but got %T", params))
	}

	var (
		m   *strtree.Tree
		def *Rule
		err *Rule
	)

	if rules != nil {
		m = strtree.NewTree()
		count := 0
		for _, r := range rules {
			if !r.hidden {
				m.InplaceInsert(r.id, r)
				count++
			}
		}

		if count > 0 {
			if mapperParams.DefOk {
				if v, ok := m.Get(mapperParams.Def); ok {
					def = v.(*Rule)
				}
			}

			if mapperParams.ErrOk {
				if v, ok := m.Get(mapperParams.Err); ok {
					err = v.(*Rule)
				}
			}
		} else {
			m = nil
		}
	}

	return mapperRCA{
		argument:  mapperParams.Argument,
		rules:     m,
		def:       def,
		err:       err,
		algorithm: mapperParams.Algorithm}
}

func (a mapperRCA) describe() string {
	return "mapper"
}

func (a mapperRCA) calculateErrorRule(ctx *Context, err error) Response {
	if a.err != nil {
		return a.err.calculate(ctx)
	}

	return Response{EffectIndeterminate, bindError(err, a.describe()), nil}
}

func (a mapperRCA) getRulesMap(rules []*Rule) *strtree.Tree {
	if a.rules != nil {
		return a.rules
	}

	m := strtree.NewTree()
	count := 0
	for _, rule := range rules {
		if !rule.hidden {
			m.InplaceInsert(rule.id, rule)
			count++
		}
	}

	if count > 0 {
		return m
	}

	return nil
}

func (a mapperRCA) add(ID string, child, old *Rule) RuleCombiningAlg {
	def := a.def
	if old != nil && old == def {
		def = child
	}

	err := a.err
	if old != nil && old == err {
		err = child
	}

	return mapperRCA{
		argument:  a.argument,
		rules:     a.rules.Insert(ID, child),
		def:       def,
		err:       err,
		algorithm: a.algorithm}
}

func (a mapperRCA) del(ID string, old *Rule) RuleCombiningAlg {
	def := a.def
	if old != nil && old == def {
		def = nil
	}

	err := a.err
	if old != nil && old == err {
		err = nil
	}

	rules := a.rules
	if rules != nil {
		rules, _ = a.rules.Delete(ID)
		if rules.IsEmpty() {
			rules = nil
		}
	}

	return mapperRCA{
		argument:  a.argument,
		rules:     rules,
		def:       def,
		err:       err,
		algorithm: a.algorithm}
}

func (a mapperRCA) execute(rules []*Rule, ctx *Context) Response {
	v, err := a.argument.Calculate(ctx)
	if err != nil {
		switch err.(type) {
		case *missingValueError:
			if a.def != nil {
				return a.def.calculate(ctx)
			}
		}

		return a.calculateErrorRule(ctx, err)
	}

	if a.algorithm != nil {
		IDs, err := getSetOfIDs(v)
		if err != nil {
			return a.calculateErrorRule(ctx, err)
		}

		r := a.algorithm.execute(collectSubRules(IDs, a.getRulesMap(rules)), ctx)
		if r.Effect == EffectNotApplicable && a.def != nil {
			return a.def.calculate(ctx)
		}

		return r
	}

	ID, err := v.str()
	if err != nil {
		return a.calculateErrorRule(ctx, err)
	}

	if a.rules != nil {
		rule, ok := a.rules.Get(ID)
		if ok {
			return rule.(*Rule).calculate(ctx)
		}
	} else {
		for _, rule := range rules {
			if rule.id == ID {
				return rule.calculate(ctx)
			}
		}
	}

	if a.def != nil {
		return a.def.calculate(ctx)
	}

	return Response{EffectNotApplicable, nil, nil}
}
