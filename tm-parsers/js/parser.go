// generated by Textmapper; DO NOT EDIT

package js

import (
	"context"
	"fmt"
)

// ErrorHandler is called every time a parser is unable to process some part of the input.
// This handler can return false to abort the parser.
type ErrorHandler func(err SyntaxError) bool

// StopOnFirstError is an error handler that forces the parser to stop on and return the first
// error.
func StopOnFirstError(_ SyntaxError) bool { return false }

type SyntaxError struct {
	Line      int
	Offset    int
	Endoffset int
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf("syntax error at line %v", e.Line)
}

func (p *Parser) Parse(ctx context.Context, lexer *Lexer) error {
	return p.parse(ctx, 6, 6362, lexer)
}

func lookaheadRule(ctx context.Context, lexer *Lexer, next, rule int32, s *session) (sym int32, err error) {
	switch rule {
	case 4212:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 0, 6355, s); ok {
			sym = 699 /* lookahead_StartOfArrowFunction */
		} else {
			sym = 167 /* lookahead_notStartOfArrowFunction */
		}
		return
	case 4213:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 1, 6356, s); ok {
			sym = 344 /* lookahead_StartOfParametrizedCall */
		} else {
			sym = 300 /* lookahead_notStartOfParametrizedCall */
		}
		return
	case 4214:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 3, 6358, s); ok {
			sym = 832 /* lookahead_StartOfIs */
		} else {
			sym = 837 /* lookahead_notStartOfIs */
		}
		return
	case 4215:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 5, 6360, s); ok {
			sym = 870 /* lookahead_StartOfMappedType */
		} else {
			sym = 860 /* lookahead_notStartOfMappedType */
		}
		return
	case 4216:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 4, 6359, s); ok {
			sym = 879 /* lookahead_StartOfFunctionType */
		} else {
			sym = 853 /* lookahead_notStartOfFunctionType */
		}
		return
	case 4217:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 2, 6357, s); ok {
			sym = 755 /* lookahead_StartOfExtendsTypeRef */
		} else {
			sym = 754 /* lookahead_notStartOfExtendsTypeRef */
		}
		return
	}
	return 0, nil
}

func AtStartOfArrowFunction(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 0, 6355, s)
}

func AtStartOfParametrizedCall(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 1, 6356, s)
}

func AtStartOfExtendsTypeRef(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 2, 6357, s)
}

func AtStartOfIs(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 3, 6358, s)
}

func AtStartOfFunctionType(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 4, 6359, s)
}

func AtStartOfMappedType(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	return lookahead(ctx, lexer, next, 5, 6360, s)
}

func lookahead(ctx context.Context, l *Lexer, next int32, start, end int16, s *session) (bool, error) {
	var lexer Lexer
	lexer.source = l.source
	lexer.ch = l.ch
	lexer.offset = l.offset
	lexer.tokenOffset = l.tokenOffset
	lexer.line = l.line
	lexer.tokenLine = l.tokenLine
	lexer.scanOffset = l.scanOffset
	lexer.State = l.State
	lexer.Dialect = l.Dialect
	lexer.token = l.token
	// Note: Stack is intentionally omitted.

	// Use memoization for recursive lookaheads.
	if next == noToken {
		next = lookaheadNext(&lexer, end, nil /*empty stack*/)
	}
	key := uint64(l.tokenOffset) + uint64(end)<<40
	if ret, ok := s.cache[key]; ok {
		return ret, nil
	}

	var allocated [64]stackEntry
	state := start
	stack := append(allocated[:0], stackEntry{state: state})

	for state != end {
		action := tmAction[state]
		if action < -2 {
			// Lookahead is needed.
			if next == noToken {
				next = lookaheadNext(&lexer, end, stack)
			}
			action = lalr(action, next)
		}

		if action >= 0 {
			// Reduce.
			rule := action
			ln := int(tmRuleLen[rule])

			var entry stackEntry
			entry.sym.symbol = tmRuleSymbol[rule]
			stack = stack[:len(stack)-ln]
			sym, err := lookaheadRule(ctx, &lexer, next, rule, s)
			if err != nil {
				return false, err
			}
			if sym != 0 {
				entry.sym.symbol = sym
			}
			state = gotoState(stack[len(stack)-1].state, entry.sym.symbol)
			entry.state = state
			stack = append(stack, entry)

		} else if action == -1 {
			if s.shiftCounter++; s.shiftCounter&0x1ff == 0 {
				// Note: checking for context cancellation is expensive so we do it from time to time.
				select {
				case <-ctx.Done():
					return false, ctx.Err()
				default:
				}
			}

			// Shift.
			if next == noToken {
				next = lookaheadNext(&lexer, end, stack)
			}
			state = gotoState(state, next)
			stack = append(stack, stackEntry{
				sym:   symbol{symbol: next},
				state: state,
			})
			if state != -1 && next != eoiToken {
				next = noToken
			}
		}

		if action == -2 || state == -1 {
			break
		}
	}

	s.cache[key] = state == end
	return state == end, nil
}

func lalr(action, next int32) int32 {
	a := -action - 3
	for ; tmLalr[a] >= 0; a += 2 {
		if tmLalr[a] == next {
			break
		}
	}
	return tmLalr[a+1]
}

func gotoState(state int16, symbol int32) int16 {
	min := tmGoto[symbol]
	max := tmGoto[symbol+1]

	if max-min < 32 {
		for i := min; i < max; i += 2 {
			if tmFromTo[i] == state {
				return tmFromTo[i+1]
			}
		}
	} else {
		for min < max {
			e := (min + max) >> 1 &^ int32(1)
			i := tmFromTo[e]
			if i == state {
				return tmFromTo[e+1]
			} else if i < state {
				min = e + 2
			} else {
				max = e
			}
		}
	}
	return -1
}

func (p *Parser) applyRule(ctx context.Context, rule int32, lhs *stackEntry, rhs []stackEntry, lexer *Lexer, s *session) (err error) {
	switch rule {
	case 1144: // LiteralPropertyName : PrivateIdentifier
		p.listener(BindingIdentifier, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 1148: // LiteralPropertyName_WithoutNew : PrivateIdentifier
		p.listener(BindingIdentifier, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 3209: // IterationStatement : 'for' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In ')' Statement
		p.listener(IdentifierReference, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3223: // IterationStatement_Await : 'for' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_Await_In ')' Statement_Await
		p.listener(IdentifierReference, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3237: // IterationStatement_Yield : 'for' '(' 'async' lookahead_notStartOfArrowFunction 'of' AssignmentExpression_In_Yield ')' Statement_Yield
		p.listener(IdentifierReference, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 3723: // TypePredicate : 'asserts' lookahead_StartOfIs 'is' Type1
		p.listener(IdentifierReference, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 3725: // TypePredicate1 : 'asserts' lookahead_StartOfIs 'is' Type1
		p.listener(IdentifierReference, rhs[0].sym.offset, rhs[0].sym.endoffset)
	case 4212:
		var ok bool
		if ok, err = AtStartOfArrowFunction(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 699 /* lookahead_StartOfArrowFunction */
		} else {
			lhs.sym.symbol = 167 /* lookahead_notStartOfArrowFunction */
		}
		return
	case 4213:
		var ok bool
		if ok, err = AtStartOfParametrizedCall(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 344 /* lookahead_StartOfParametrizedCall */
		} else {
			lhs.sym.symbol = 300 /* lookahead_notStartOfParametrizedCall */
		}
		return
	case 4214:
		var ok bool
		if ok, err = AtStartOfIs(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 832 /* lookahead_StartOfIs */
		} else {
			lhs.sym.symbol = 837 /* lookahead_notStartOfIs */
		}
		return
	case 4215:
		var ok bool
		if ok, err = AtStartOfMappedType(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 870 /* lookahead_StartOfMappedType */
		} else {
			lhs.sym.symbol = 860 /* lookahead_notStartOfMappedType */
		}
		return
	case 4216:
		var ok bool
		if ok, err = AtStartOfFunctionType(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 879 /* lookahead_StartOfFunctionType */
		} else {
			lhs.sym.symbol = 853 /* lookahead_notStartOfFunctionType */
		}
		return
	case 4217:
		var ok bool
		if ok, err = AtStartOfExtendsTypeRef(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 755 /* lookahead_StartOfExtendsTypeRef */
		} else {
			lhs.sym.symbol = 754 /* lookahead_notStartOfExtendsTypeRef */
		}
		return
	}
	nt := ruleNodeType[rule]
	if nt != 0 {
		p.listener(nt, lhs.sym.offset, lhs.sym.endoffset)
	}
	return
}

const errSymbol = 2

func (p *Parser) skipBrokenCode(lexer *Lexer, stack []stackEntry, canRecover func(symbol int32) bool) int {
	var e int
	for p.next.symbol != eoiToken && !canRecover(p.next.symbol) {
		if debugSyntax {
			fmt.Printf("skipped while recovering: %v (%s)\n", Symbol(p.next.symbol), lexer.Text())
		}
		switch Token(p.next.symbol) {
		case NOSUBSTITUTIONTEMPLATE:
			p.listener(NoSubstitutionTemplate, p.next.offset, p.next.endoffset)
		case TEMPLATEHEAD:
			p.listener(TemplateHead, p.next.offset, p.next.endoffset)
		case TEMPLATEMIDDLE:
			p.listener(TemplateMiddle, p.next.offset, p.next.endoffset)
		case TEMPLATETAIL:
			p.listener(TemplateTail, p.next.offset, p.next.endoffset)
		}
		e = p.next.endoffset
		p.fetchNext(lexer, stack, nil)
	}
	return e
}

// willShift checks if "symbol" is going to be shifted in the given state.
// This function does not support empty productions and returns false if they occur before "symbol".
func (p *Parser) willShift(stackPos int, state int16, symbol int32, stack []stackEntry) bool {
	if state == -1 {
		return false
	}

	for state != p.endState {
		action := tmAction[state]
		if action < -2 {
			action = lalr(action, symbol)
		}

		if action >= 0 {
			// Reduce.
			rule := action
			ln := int(tmRuleLen[rule])
			if ln == 0 {
				// we do not support empty productions
				return false
			}
			stackPos -= ln - 1
			state = gotoState(stack[stackPos-1].state, tmRuleSymbol[rule])
		} else {
			return action == -1 && gotoState(state, symbol) >= 0
		}
	}
	return symbol == eoiToken
}

func (p *Parser) reportIgnoredToken(tok symbol) {
	var t NodeType
	switch Token(tok.symbol) {
	case MULTILINECOMMENT:
		t = MultiLineComment
	case SINGLELINECOMMENT:
		t = SingleLineComment
	case INVALID_TOKEN:
		t = InvalidToken
	default:
		return
	}
	if debugSyntax {
		fmt.Printf("ignored: %v as %v\n", Token(tok.symbol), t)
	}
	p.listener(t, tok.offset, tok.endoffset)
}
