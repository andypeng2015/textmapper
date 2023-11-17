// generated by Textmapper; DO NOT EDIT

package test

import (
	"fmt"

	"github.com/inspirer/textmapper/parsers/test/token"
)

var tmNonterminals = [...]string{
	"Declaration_list",
	"Test",
	"Declaration",
	"lookahead_FooLookahead",
	"lookahead_notFooLookahead",
	"setof_not_EOI_or_DOT_or_RBRACE",
	"setof_not_EOI_or_DOT_or_RBRACE_optlist",
	"FooLookahead",
	"setof_foo_la",
	"setof_foo_la_list",
	"empty1",
	"foo_la",
	"foo_nonterm",
	"foo_nonterm_A",
	"QualifiedName",
	"Decl1",
	"Decl2",
	"If",
	"expr",
	"customPlus",
	"primaryExpr",
	"primaryExpr_WithoutAs",
	"QualifiedNameopt",
}

func symbolName(sym int32) string {
	if sym == noToken {
		return "<no-token>"
	}
	if sym < int32(token.NumTokens) {
		return token.Token(sym).String()
	}
	if i := int(sym) - int(token.NumTokens); i < len(tmNonterminals) {
		return tmNonterminals[i]
	}
	return fmt.Sprintf("nonterminal(%d)", sym)
}

var tmAction = []int32{
	-1, -1, -1, -1, -3, 11, -1, -1, -27, -51, -1, -1, -55, 1, 3, 4, 82, -1, -1,
	17, 62, -75, -1, -81, -1, -1, -1, 10, -1, -1, 0, -1, 12, -1, -1, -1, -1, 78,
	-1, -105, 21, -1, -1, -1, 8, -1, -1, 9, 64, 65, 66, 67, 68, 69, 71, -1, 24,
	25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 14, 39, 40, 41, 42,
	43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	-1, 15, 16, 80, -1, -1, -1, -129, -1, -1, 86, 87, -1, 81, -137, 6, -1, 7, 63,
	70, -161, 79, -1, -1, 19, -1, 72, -1, -1, -1, 5, -167, -173, -1, 18, 85,
	-179, 84, -1, -185, -1, 20, -2, -1, -1, -2, -2,
}

var tmLalr = []int32{
	19, -1, 0, 13, 6, 13, 7, 13, 8, 13, 9, 13, 10, 13, 11, 13, 13, 13, 15, 13,
	16, 13, -1, -2, 24, -1, 0, 81, 6, 81, 7, 81, 8, 81, 9, 81, 10, 81, 11, 81,
	13, 81, 15, 81, 16, 81, -1, -2, 17, 94, -1, -2, 6, -1, 7, -1, 8, -1, 9, -1,
	10, -1, 11, -1, 13, -1, 15, -1, 0, 2, -1, -2, 6, -1, 18, 72, -1, -2, 4, -1,
	0, 93, 6, 93, 7, 93, 8, 93, 9, 93, 10, 93, 11, 93, 13, 93, 15, 93, 16, 93,
	-1, -2, 21, -1, 0, 92, 6, 92, 7, 92, 8, 92, 9, 92, 10, 92, 11, 92, 13, 92,
	15, 92, 16, 92, -1, -2, 18, 90, 27, 90, 12, 91, -1, -2, 14, -1, 0, 83, 6, 83,
	7, 83, 8, 83, 9, 83, 10, 83, 11, 83, 13, 83, 15, 83, 16, 83, -1, -2, 27, -1,
	18, 75, -1, -2, 27, -1, 18, 76, -1, -2, 27, -1, 18, 77, -1, -2, 27, -1, 18,
	89, -1, -2, 27, 88, 18, 88, -1, -2,
}

var tmGoto = []int32{
	0, 4, 6, 8, 10, 18, 20, 68, 86, 104, 124, 146, 164, 172, 194, 198, 218, 232,
	246, 264, 268, 272, 286, 288, 290, 294, 300, 302, 324, 344, 346, 354, 356,
	358, 360, 362, 364, 366, 368, 370, 372, 380, 382, 398, 400, 402, 404, 406,
	408, 412, 414, 418, 418, 420, 422, 426, 444, 464, 484, 498, 512, 530, 548,
	550,
}

var tmFromTo = []int16{
	138, 140, 139, 141, 33, 56, 33, 57, 33, 58, 22, 37, 23, 37, 33, 59, 99, 116,
	33, 60, 1, 4, 3, 17, 6, 19, 11, 4, 12, 4, 21, 34, 28, 4, 29, 4, 31, 48, 33,
	61, 41, 100, 42, 102, 45, 4, 46, 4, 55, 48, 95, 102, 103, 102, 111, 4, 117,
	102, 118, 102, 122, 102, 123, 102, 128, 102, 133, 102, 1, 5, 11, 5, 12, 5,
	28, 5, 29, 5, 33, 62, 45, 5, 46, 5, 111, 5, 1, 6, 11, 6, 12, 6, 28, 6, 29, 6,
	33, 63, 45, 6, 46, 6, 111, 6, 1, 7, 2, 7, 11, 7, 12, 7, 28, 7, 29, 7, 33, 64,
	45, 7, 46, 7, 111, 7, 1, 8, 11, 8, 12, 8, 28, 8, 29, 8, 33, 65, 43, 108, 45,
	8, 46, 8, 111, 8, 124, 108, 1, 9, 11, 9, 12, 9, 28, 9, 29, 9, 33, 66, 45, 9,
	46, 9, 111, 9, 31, 49, 33, 67, 55, 49, 107, 123, 1, 10, 11, 10, 12, 10, 28,
	10, 29, 10, 33, 68, 43, 10, 45, 10, 46, 10, 111, 10, 124, 10, 33, 69, 109,
	124, 1, 11, 6, 20, 11, 11, 12, 11, 28, 11, 29, 11, 33, 70, 45, 11, 46, 11,
	111, 11, 11, 27, 28, 44, 29, 47, 33, 71, 45, 110, 46, 112, 111, 125, 0, 3, 6,
	21, 7, 22, 10, 26, 24, 41, 25, 42, 33, 72, 26, 43, 33, 73, 35, 96, 36, 97,
	38, 98, 55, 113, 101, 119, 104, 121, 135, 136, 4, 18, 33, 74, 18, 32, 33, 75,
	17, 31, 31, 50, 34, 95, 38, 99, 39, 99, 55, 50, 100, 117, 33, 76, 33, 77, 8,
	23, 33, 78, 11, 28, 28, 45, 33, 79, 33, 80, 31, 51, 33, 81, 55, 51, 104, 122,
	115, 122, 120, 128, 126, 133, 127, 122, 131, 122, 134, 122, 135, 122, 31, 52,
	33, 82, 42, 103, 55, 52, 95, 103, 117, 103, 118, 103, 123, 103, 128, 103,
	133, 103, 33, 83, 31, 53, 33, 84, 55, 53, 100, 118, 33, 85, 33, 86, 33, 87,
	33, 88, 33, 89, 33, 90, 33, 91, 33, 92, 33, 93, 1, 12, 11, 29, 28, 46, 45,
	111, 1, 138, 1, 13, 11, 13, 12, 30, 28, 13, 29, 30, 45, 13, 46, 30, 111, 30,
	9, 24, 9, 25, 33, 94, 20, 33, 0, 137, 31, 54, 55, 114, 31, 55, 21, 35, 121,
	129, 21, 36, 41, 101, 22, 38, 23, 39, 1, 14, 2, 139, 11, 14, 12, 14, 28, 14,
	29, 14, 45, 14, 46, 14, 111, 14, 1, 15, 11, 15, 12, 15, 28, 15, 29, 15, 43,
	109, 45, 15, 46, 15, 111, 15, 124, 132, 1, 16, 11, 16, 12, 16, 28, 16, 29,
	16, 43, 16, 45, 16, 46, 16, 111, 16, 124, 16, 42, 104, 95, 115, 117, 126,
	118, 127, 123, 131, 128, 134, 133, 135, 42, 105, 95, 105, 117, 105, 118, 105,
	123, 105, 128, 105, 133, 105, 42, 106, 95, 106, 103, 120, 117, 106, 118, 106,
	122, 130, 123, 106, 128, 106, 133, 106, 42, 107, 95, 107, 103, 107, 117, 107,
	118, 107, 122, 107, 123, 107, 128, 107, 133, 107, 23, 40,
}

var tmRuleLen = []int8{
	2, 1, 1, 1, 1, 5, 4, 4, 3, 3, 2, 1, 3, 1, 4, 4, 4, 2, 6, 5, 9, 3, 0, 0, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 0, 5, 1, 1, 1, 1, 1, 1, 2, 1, 0, 3, 3, 3, 3, 3,
	1, 3, 4, 1, 1, 4, 6, 3, 1, 1, 4, 3, 1, 1, 1, 0, 0,
}

var tmRuleSymbol = []int32{
	40, 40, 41, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42,
	42, 42, 42, 43, 44, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45,
	45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45, 45,
	45, 45, 45, 45, 46, 46, 47, 48, 48, 48, 48, 48, 48, 49, 49, 50, 51, 51, 52,
	53, 53, 54, 54, 55, 56, 56, 57, 57, 58, 58, 58, 59, 60, 60, 61, 62, 62, 44,
}

var tmRuleType = [...]uint32{
	0,                  // Declaration_list : Declaration_list Declaration
	0,                  // Declaration_list : Declaration
	uint32(Test),       // Test : Declaration_list
	0,                  // Declaration : Decl1
	0,                  // Declaration : Decl2
	uint32(Block),      // Declaration : '{' '-' '-' Declaration_list '}'
	uint32(Block),      // Declaration : '{' '-' '-' '}'
	uint32(Block),      // Declaration : '{' '-' Declaration_list '}'
	uint32(Block),      // Declaration : '{' '-' '}'
	uint32(Block),      // Declaration : '{' Declaration_list '}'
	uint32(Block),      // Declaration : '{' '}'
	uint32(LastInt),    // Declaration : lastInt
	uint32(Int),        // Declaration : IntegerConstant '[' ']'
	uint32(Int),        // Declaration : IntegerConstant
	uint32(TestClause), // Declaration : 'test' '{' setof_not_EOI_or_DOT_or_RBRACE_optlist '}'
	0,                  // Declaration : 'test' '(' empty1 ')'
	0,                  // Declaration : 'test' '(' foo_nonterm ')'
	uint32(TestIntClause) + uint32(InTest|InFoo)<<16, // Declaration : 'test' IntegerConstant
	uint32(EvalEmpty1),  // Declaration : 'eval' lookahead_notFooLookahead '(' expr ')' empty1
	uint32(EvalFoo),     // Declaration : 'eval' lookahead_FooLookahead '(' foo_nonterm_A ')'
	uint32(EvalFoo2),    // Declaration : 'eval' lookahead_FooLookahead '(' IntegerConstant '.' expr '+' .greedy expr ')'
	uint32(DeclOptQual), // Declaration : 'decl2' ':' QualifiedNameopt
	0,                   // lookahead_FooLookahead :
	0,                   // lookahead_notFooLookahead :
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : invalid_token
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : WhiteSpace
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : SingleLineComment
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : Identifier
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : Identifier2
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : IntegerConstant
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : lastInt
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : 'test'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : 'decl1'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : 'decl2'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : 'eval'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : 'as'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : 'if'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : 'else'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : '{'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : '('
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : ')'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : '['
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : ']'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : '...'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : ','
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : ':'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : '-'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : '->'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : '+'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : '\\'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : '_'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : 'foo_'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : 'f_a'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : multiline
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : dquote
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : '\''
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : SharpAtID
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : 'Zfoo'
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : backtrackingToken
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : error
	0,                   // setof_not_EOI_or_DOT_or_RBRACE : MultiLineComment
	0,                   // setof_not_EOI_or_DOT_or_RBRACE_optlist : setof_not_EOI_or_DOT_or_RBRACE_optlist setof_not_EOI_or_DOT_or_RBRACE
	0,                   // setof_not_EOI_or_DOT_or_RBRACE_optlist :
	0,                   // FooLookahead : '(' IntegerConstant '.' setof_foo_la_list ')'
	0,                   // setof_foo_la : IntegerConstant
	0,                   // setof_foo_la : 'as'
	0,                   // setof_foo_la : '.'
	0,                   // setof_foo_la : '+'
	0,                   // setof_foo_la : '\\'
	0,                   // setof_foo_la : 'foo_'
	0,                   // setof_foo_la_list : setof_foo_la_list setof_foo_la
	0,                   // setof_foo_la_list : setof_foo_la
	0,                   // empty1 :
	0,                   // foo_la : IntegerConstant '.' expr
	0,                   // foo_la : IntegerConstant 'foo_' expr
	0,                   // foo_nonterm : IntegerConstant '.' expr
	0,                   // foo_nonterm_A : IntegerConstant '.' expr
	0,                   // foo_nonterm_A : IntegerConstant 'foo_' expr
	0,                   // QualifiedName : Identifier
	0,                   // QualifiedName : QualifiedName '.' Identifier
	uint32(Decl1),       // Decl1 : 'decl1' '(' QualifiedName ')'
	uint32(Decl2),       // Decl2 : 'decl2'
	0,                   // Decl2 : If
	uint32(If),          // If : 'if' '(' ')' Decl2
	uint32(If),          // If : 'if' '(' ')' Decl2 'else' Decl2
	uint32(PlusExpr),    // expr : expr '+' primaryExpr
	0,                   // expr : customPlus
	0,                   // expr : primaryExpr
	0,                   // customPlus : '\\' primaryExpr '+' expr
	uint32(AsExpr),      // primaryExpr : primaryExpr_WithoutAs 'as' expr
	uint32(IntExpr),     // primaryExpr : IntegerConstant
	uint32(IntExpr),     // primaryExpr_WithoutAs : IntegerConstant
	0,                   // QualifiedNameopt : QualifiedName
	0,                   // QualifiedNameopt :
}

// set(follow error) =
var afterErr = []token.Token{}
