// generated by Textmapper; DO NOT EDIT

package js

import (
	"fmt"
)

type NodeType int

type Listener func(t NodeType, offset, endoffset int)

const (
	NoType        NodeType = iota
	SyntaxProblem          // ReferenceIdent? Initializer?
	NameIdent
	ReferenceIdent
	LabelIdent
	This
	IdentExpr // ReferenceIdent
	Regexp
	Parenthesized // Expr? SyntaxProblem?
	Literal
	ArrayLiteral // list=(Expr)*
	NoElement
	SpreadElement        // Expr
	ObjectLiteral        // (PropertyDefinition)*
	ShorthandProperty    // ReferenceIdent
	Property             // (Modifier)* PropertyName value=Expr
	ObjectMethod         // (Modifier)* MethodDefinition
	SpreadProperty       // Expr
	LiteralPropertyName  // NameIdent?
	ComputedPropertyName // Expr
	Initializer          // Expr
	TemplateLiteral      // template=(NoSubstitutionTemplate | TemplateHead | TemplateMiddle | TemplateTail)+ substitution=(Expr)*
	IndexAccess          // expr=Expr index=Expr
	PropertyAccess       // expr=Expr? selector=ReferenceIdent
	TaggedTemplate       // tag=Expr literal=TemplateLiteral
	TsNonNull            // expr=Expr
	NewExpr              // expr=Expr Arguments?
	SuperExpr
	NewTarget
	CallExpr               // expr=Expr Arguments
	TsDynamicImport        // Arguments
	Arguments              // TypeArguments? list=(Expr)*
	OptionalIndexAccess    // expr=Expr index=Expr
	OptionalPropertyAccess // expr=Expr selector=ReferenceIdent
	OptionalCallExpr       // expr=Expr Arguments
	OptionalTaggedTemplate // tag=Expr literal=TemplateLiteral
	PostInc                // Expr
	PostDec                // Expr
	PreInc                 // Expr
	PreDec                 // Expr
	UnaryExpr              // Expr
	TsCastExpr             // TsType Expr
	AdditiveExpr           // left=Expr right=Expr
	ShiftExpr              // left=Expr right=Expr
	MultiplicativeExpr     // left=Expr right=Expr
	ExponentiationExpr     // left=Expr right=Expr
	RelationalExpr         // left=Expr right=Expr
	InstanceOfExpr         // left=Expr right=Expr
	TsAsExpr               // left=Expr TsType
	TsConst
	TsAsConstExpr   // left=Expr TsConst
	EqualityExpr    // left=Expr right=Expr
	BitwiseAND      // left=Expr right=Expr
	BitwiseXOR      // left=Expr right=Expr
	BitwiseOR       // left=Expr right=Expr
	LogicalAND      // left=Expr right=Expr
	LogicalOR       // left=Expr right=Expr
	CoalesceExpr    // left=Expr right=Expr
	InExpr          // left=Expr right=Expr
	ConditionalExpr // cond=Expr then=Expr else=Expr
	AssignmentExpr  // left=Expr AssignmentOperator? right=Expr
	AssignmentOperator
	CommaExpr   // left=Expr right=Expr
	Block       // (CaseClause)* (StmtListItem)*
	LexicalDecl // LetOrConst (LexicalBinding)+
	LetOrConst
	TsExclToken
	LexicalBinding     // BindingPattern? NameIdent? TsExclToken? TypeAnnotation? Initializer?
	VarStmt            // (VarDecl)+
	VarDecl            // BindingPattern? NameIdent? TsExclToken? TypeAnnotation? Initializer?
	ObjectPattern      // (PropertyPattern)* BindingRestElement?
	ArrayPattern       // list=(ElementPattern | Expr)* BindingRestElement?
	PropertyBinding    // PropertyName ElementPattern
	ElementBinding     // BindingPattern Initializer?
	SingleNameBinding  // NameIdent Initializer?
	BindingRestElement // NameIdent
	EmptyStmt
	ExprStmt    // Expr
	IfStmt      // Expr then=Stmt else=Stmt?
	DoWhileStmt // Stmt Expr
	WhileStmt   // Expr Stmt
	ForStmt     // var=Expr? ForCondition ForFinalExpr Stmt
	Var
	ForStmtWithVar   // LetOrConst? Var? (LexicalBinding)* (VarDecl)* ForCondition ForFinalExpr Stmt
	ForInStmt        // var=Expr object=Expr Stmt
	ForInStmtWithVar // LetOrConst? Var? ForBinding object=Expr Stmt
	ForOfStmt        // Await? var=Expr iterable=Expr Stmt
	ForOfStmtWithVar // Await? LetOrConst? Var? ForBinding iterable=Expr Stmt
	Await
	ForBinding   // BindingPattern? NameIdent?
	ForCondition // Expr?
	ForFinalExpr // Expr?
	ContinueStmt // LabelIdent?
	BreakStmt    // LabelIdent?
	ReturnStmt   // Expr?
	WithStmt     // Expr Stmt
	SwitchStmt   // Expr Block
	Case         // Expr (StmtListItem)*
	Default      // (StmtListItem)*
	LabelledStmt // LabelIdent Func? Stmt?
	ThrowStmt    // Expr
	TryStmt      // Block Catch? Finally?
	Catch        // BindingPattern? NameIdent? TypeAnnotation? Block
	Finally      // Block
	DebuggerStmt
	Func                      // NameIdent? TypeParameters? Parameters TypeAnnotation? Body
	FuncExpr                  // NameIdent? TypeParameters? Parameters TypeAnnotation? Body
	Body                      // (StmtListItem)*
	ArrowFunc                 // NameIdent? TypeParameters? Parameters? TypeAnnotation? Body? ConciseBody?
	ConciseBody               // Expr
	AsyncArrowFunc            // NameIdent? TypeParameters? Parameters? TypeAnnotation? Body? ConciseBody?
	Method                    // PropertyName TypeParameters? Parameters TypeAnnotation? Body
	Getter                    // (Modifier)* PropertyName TypeAnnotation? Body?
	Setter                    // (Modifier)* PropertyName Parameter Body?
	GeneratorMethod           // PropertyName TypeParameters? Parameters TypeAnnotation? Body
	Generator                 // NameIdent? TypeParameters? Parameters TypeAnnotation? Body
	GeneratorExpr             // NameIdent? TypeParameters? Parameters TypeAnnotation? Body
	Yield                     // Expr?
	AsyncMethod               // PropertyName TypeParameters? Parameters TypeAnnotation? Body
	AsyncFunc                 // NameIdent? TypeParameters? Parameters TypeAnnotation? Body
	AsyncFuncExpr             // NameIdent? TypeParameters? Parameters TypeAnnotation? Body
	AsyncGeneratorMethod      // PropertyName TypeParameters? Parameters TypeAnnotation? Body
	AsyncGeneratorDeclaration // NameIdent? TypeParameters? Parameters TypeAnnotation? Body
	AsyncGeneratorExpression  // NameIdent? TypeParameters? Parameters TypeAnnotation? Body
	AwaitExpr                 // Expr
	Class                     // (Modifier)* NameIdent? TypeParameters? Extends? TsImplementsClause? ClassBody
	ClassExpr                 // (Modifier)* NameIdent? TypeParameters? Extends? TsImplementsClause? ClassBody
	Extends                   // Expr? TypeReference?
	TsImplementsClause        // (TypeReference)+
	ClassBody                 // (ClassElement)*
	Static
	Abstract
	Override
	Readonly
	Declare
	MemberMethod      // (Modifier)* MethodDefinition
	MemberVar         // (Modifier)* PropertyName TypeAnnotation? Initializer?
	TsIndexMemberDecl // IndexSignature
	EmptyDecl
	Module     // (ModuleItem)*
	ImportDecl // TsTypeOnly? NameIdent? NameSpaceImport? NamedImports? ModuleSpec
	TsTypeOnly
	TsExport
	TsImportRequireDecl // TsExport? NameIdent
	NameSpaceImport     // NameIdent
	NamedImports        // (NamedImport)*
	ImportSpec          // ReferenceIdent? NameIdent
	ModuleSpec
	ExportDecl            // (Modifier)* TsTypeOnly? VarStmt? Decl? ExportClause? NameIdent? ModuleSpec?
	ExportDefault         // Expr? (Modifier)* Decl?
	TsExportAssignment    // Expr
	TsNamespaceExportDecl // NameIdent
	ExportClause          // (ExportElement)*
	ExportSpec            // ReferenceIdent NameIdent?
	DecoratorExpr         // (ReferenceIdent)+
	DecoratorCall         // (ReferenceIdent)+ Arguments
	JSXElement            // JSXOpeningElement? JSXSelfClosingElement? (JSXChild)* JSXClosingElement?
	JSXSelfClosingElement // JSXElementName TypeArguments? (JSXAttribute)*
	JSXOpeningElement     // JSXElementName TypeArguments? (JSXAttribute)*
	JSXClosingElement     // JSXElementName
	JSXElementName
	JSXNormalAttribute // JSXAttributeName JSXAttributeValue?
	JSXSpreadAttribute // Expr
	JSXAttributeName
	JSXLiteral
	JSXExpr // Expr?
	JSXText
	JSXSpreadExpr    // Expr?
	TsConditional    // check=TsType ext=TsType truet=TsType falset=TsType
	TypePredicate    // paramref=ReferenceIdent TsType
	AssertsType      // ReferenceIdent? This? TsType?
	TypeParameters   // (TypeParameter)+
	TypeParameter    // NameIdent TypeConstraint? TsType?
	TypeConstraint   // TsType
	TypeArguments    // (TsType)+
	UnionType        // inner=(TsType)*
	IntersectionType // inner=(TsType)*
	KeyOfType        // TsType
	UniqueType       // TsType
	ReadonlyType     // TsType
	TypeVar          // ReferenceIdent
	ThisType
	NonNullableType   // TsType
	NullableType      // TsType
	ParenthesizedType // TsType
	LiteralType       // TemplateLiteral?
	PredefinedType
	TypeReference     // TypeName TypeArguments?
	TypeName          // ref=(ReferenceIdent)+
	ObjectType        // (TypeMember)*
	ArrayType         // TsType
	IndexedAccessType // left=TsType index=TsType
	MappedType        // NameIdent inType=TsType asType=TsType? TypeAnnotation
	TupleType         // (TupleMember)*
	NamedTupleMember  // TsType
	RestType          // TsType
	FuncType          // TypeParameters? Parameters TsType
	Parameters        // (Parameter)*
	ConstructorType   // Abstract? TypeParameters? Parameters TsType
	TypeQuery         // (ReferenceIdent)+
	ImportType        // TsType (ReferenceIdent)* TypeArguments?
	PropertySignature // (Modifier)* PropertyName TypeAnnotation?
	TypeAnnotation    // TsType
	CallSignature     // TypeParameters? Parameters TypeAnnotation?
	DefaultParameter  // (Modifier)* BindingPattern? NameIdent? TypeAnnotation? Initializer?
	RestParameter     // BindingPattern? NameIdent? TypeAnnotation?
	TsThisParameter   // TypeAnnotation
	AccessibilityModifier
	ConstructSignature   // (Modifier)* TypeParameters? Parameters TypeAnnotation?
	IndexSignature       // (Modifier)* NameIdent TsType TypeAnnotation
	MethodSignature      // (Modifier)* PropertyName TypeParameters? Parameters TypeAnnotation?
	TypeAliasDecl        // NameIdent TypeParameters? TsType
	TsInterface          // NameIdent TypeParameters? TsInterfaceExtends? ObjectType
	TsInterfaceExtends   // (TypeReference)+
	TsEnum               // TsConst? NameIdent TsEnumBody
	TsEnumBody           // (TsEnumMember)*
	TsEnumMember         // PropertyName Expr?
	TsNamespace          // (NameIdent)+ TsNamespaceBody
	TsNamespaceBody      // (ModuleItem)*
	TsImportAliasDecl    // NameIdent ref=(ReferenceIdent)+
	TsAmbientVar         // LetOrConst? Var? (TsAmbientBinding)+
	TsAmbientFunc        // NameIdent TypeParameters? Parameters TypeAnnotation?
	TsAmbientClass       // (Modifier)* NameIdent TypeParameters? Extends? TsImplementsClause? ClassBody
	TsAmbientInterface   // (Modifier)* NameIdent TypeParameters? TsInterfaceExtends? ObjectType
	TsAmbientEnum        // TsConst? NameIdent TsEnumBody
	TsAmbientNamespace   // (NameIdent)+ (TsAmbientElement)*
	TsAmbientModule      // (NameIdent)* (ModuleItem)*
	TsAmbientGlobal      // (ModuleItem)*
	TsAmbientTypeAlias   // TypeAliasDecl
	TsAmbientBinding     // NameIdent TypeAnnotation? Initializer?
	TsAmbientImportAlias // TsImportAliasDecl
	TsAmbientExportDecl  // ExportClause
	InsertedSemicolon
	MultiLineComment
	SingleLineComment
	InvalidToken
	NoSubstitutionTemplate
	TemplateHead
	TemplateMiddle
	TemplateTail
	NodeTypeMax
)

var nodeTypeStr = [...]string{
	"NONE",
	"SyntaxProblem",
	"NameIdent",
	"ReferenceIdent",
	"LabelIdent",
	"This",
	"IdentExpr",
	"Regexp",
	"Parenthesized",
	"Literal",
	"ArrayLiteral",
	"NoElement",
	"SpreadElement",
	"ObjectLiteral",
	"ShorthandProperty",
	"Property",
	"ObjectMethod",
	"SpreadProperty",
	"LiteralPropertyName",
	"ComputedPropertyName",
	"Initializer",
	"TemplateLiteral",
	"IndexAccess",
	"PropertyAccess",
	"TaggedTemplate",
	"TsNonNull",
	"NewExpr",
	"SuperExpr",
	"NewTarget",
	"CallExpr",
	"TsDynamicImport",
	"Arguments",
	"OptionalIndexAccess",
	"OptionalPropertyAccess",
	"OptionalCallExpr",
	"OptionalTaggedTemplate",
	"PostInc",
	"PostDec",
	"PreInc",
	"PreDec",
	"UnaryExpr",
	"TsCastExpr",
	"AdditiveExpr",
	"ShiftExpr",
	"MultiplicativeExpr",
	"ExponentiationExpr",
	"RelationalExpr",
	"InstanceOfExpr",
	"TsAsExpr",
	"TsConst",
	"TsAsConstExpr",
	"EqualityExpr",
	"BitwiseAND",
	"BitwiseXOR",
	"BitwiseOR",
	"LogicalAND",
	"LogicalOR",
	"CoalesceExpr",
	"InExpr",
	"ConditionalExpr",
	"AssignmentExpr",
	"AssignmentOperator",
	"CommaExpr",
	"Block",
	"LexicalDecl",
	"LetOrConst",
	"TsExclToken",
	"LexicalBinding",
	"VarStmt",
	"VarDecl",
	"ObjectPattern",
	"ArrayPattern",
	"PropertyBinding",
	"ElementBinding",
	"SingleNameBinding",
	"BindingRestElement",
	"EmptyStmt",
	"ExprStmt",
	"IfStmt",
	"DoWhileStmt",
	"WhileStmt",
	"ForStmt",
	"Var",
	"ForStmtWithVar",
	"ForInStmt",
	"ForInStmtWithVar",
	"ForOfStmt",
	"ForOfStmtWithVar",
	"Await",
	"ForBinding",
	"ForCondition",
	"ForFinalExpr",
	"ContinueStmt",
	"BreakStmt",
	"ReturnStmt",
	"WithStmt",
	"SwitchStmt",
	"Case",
	"Default",
	"LabelledStmt",
	"ThrowStmt",
	"TryStmt",
	"Catch",
	"Finally",
	"DebuggerStmt",
	"Func",
	"FuncExpr",
	"Body",
	"ArrowFunc",
	"ConciseBody",
	"AsyncArrowFunc",
	"Method",
	"Getter",
	"Setter",
	"GeneratorMethod",
	"Generator",
	"GeneratorExpr",
	"Yield",
	"AsyncMethod",
	"AsyncFunc",
	"AsyncFuncExpr",
	"AsyncGeneratorMethod",
	"AsyncGeneratorDeclaration",
	"AsyncGeneratorExpression",
	"AwaitExpr",
	"Class",
	"ClassExpr",
	"Extends",
	"TsImplementsClause",
	"ClassBody",
	"Static",
	"Abstract",
	"Override",
	"Readonly",
	"Declare",
	"MemberMethod",
	"MemberVar",
	"TsIndexMemberDecl",
	"EmptyDecl",
	"Module",
	"ImportDecl",
	"TsTypeOnly",
	"TsExport",
	"TsImportRequireDecl",
	"NameSpaceImport",
	"NamedImports",
	"ImportSpec",
	"ModuleSpec",
	"ExportDecl",
	"ExportDefault",
	"TsExportAssignment",
	"TsNamespaceExportDecl",
	"ExportClause",
	"ExportSpec",
	"DecoratorExpr",
	"DecoratorCall",
	"JSXElement",
	"JSXSelfClosingElement",
	"JSXOpeningElement",
	"JSXClosingElement",
	"JSXElementName",
	"JSXNormalAttribute",
	"JSXSpreadAttribute",
	"JSXAttributeName",
	"JSXLiteral",
	"JSXExpr",
	"JSXText",
	"JSXSpreadExpr",
	"TsConditional",
	"TypePredicate",
	"AssertsType",
	"TypeParameters",
	"TypeParameter",
	"TypeConstraint",
	"TypeArguments",
	"UnionType",
	"IntersectionType",
	"KeyOfType",
	"UniqueType",
	"ReadonlyType",
	"TypeVar",
	"ThisType",
	"NonNullableType",
	"NullableType",
	"ParenthesizedType",
	"LiteralType",
	"PredefinedType",
	"TypeReference",
	"TypeName",
	"ObjectType",
	"ArrayType",
	"IndexedAccessType",
	"MappedType",
	"TupleType",
	"NamedTupleMember",
	"RestType",
	"FuncType",
	"Parameters",
	"ConstructorType",
	"TypeQuery",
	"ImportType",
	"PropertySignature",
	"TypeAnnotation",
	"CallSignature",
	"DefaultParameter",
	"RestParameter",
	"TsThisParameter",
	"AccessibilityModifier",
	"ConstructSignature",
	"IndexSignature",
	"MethodSignature",
	"TypeAliasDecl",
	"TsInterface",
	"TsInterfaceExtends",
	"TsEnum",
	"TsEnumBody",
	"TsEnumMember",
	"TsNamespace",
	"TsNamespaceBody",
	"TsImportAliasDecl",
	"TsAmbientVar",
	"TsAmbientFunc",
	"TsAmbientClass",
	"TsAmbientInterface",
	"TsAmbientEnum",
	"TsAmbientNamespace",
	"TsAmbientModule",
	"TsAmbientGlobal",
	"TsAmbientTypeAlias",
	"TsAmbientBinding",
	"TsAmbientImportAlias",
	"TsAmbientExportDecl",
	"InsertedSemicolon",
	"MultiLineComment",
	"SingleLineComment",
	"InvalidToken",
	"NoSubstitutionTemplate",
	"TemplateHead",
	"TemplateMiddle",
	"TemplateTail",
}

func (t NodeType) String() string {
	if t >= 0 && int(t) < len(nodeTypeStr) {
		return nodeTypeStr[t]
	}
	return fmt.Sprintf("node(%d)", t)
}

var BindingPattern = []NodeType{
	ArrayPattern,
	ObjectPattern,
}

var CaseClause = []NodeType{
	Case,
	Default,
}

var ClassElement = []NodeType{
	EmptyDecl,
	MemberMethod,
	MemberVar,
	TsIndexMemberDecl,
}

var Decl = []NodeType{
	AsyncFunc,
	AsyncGeneratorDeclaration,
	Class,
	Func,
	Generator,
	LexicalDecl,
	TsAmbientClass,
	TsAmbientEnum,
	TsAmbientExportDecl,
	TsAmbientFunc,
	TsAmbientGlobal,
	TsAmbientImportAlias,
	TsAmbientInterface,
	TsAmbientModule,
	TsAmbientNamespace,
	TsAmbientTypeAlias,
	TsAmbientVar,
	TsEnum,
	TsImportAliasDecl,
	TsInterface,
	TsNamespace,
	TypeAliasDecl,
}

var Decorator = []NodeType{
	DecoratorCall,
	DecoratorExpr,
}

var ElementPattern = []NodeType{
	ElementBinding,
	NoElement,
	SingleNameBinding,
	SyntaxProblem,
}

var ExportElement = []NodeType{
	ExportSpec,
	SyntaxProblem,
}

var Expr = []NodeType{
	AdditiveExpr,
	ArrayLiteral,
	ArrowFunc,
	AssignmentExpr,
	AsyncArrowFunc,
	AsyncFuncExpr,
	AsyncGeneratorExpression,
	AwaitExpr,
	BitwiseAND,
	BitwiseOR,
	BitwiseXOR,
	CallExpr,
	ClassExpr,
	CoalesceExpr,
	CommaExpr,
	ConditionalExpr,
	EqualityExpr,
	ExponentiationExpr,
	FuncExpr,
	GeneratorExpr,
	IdentExpr,
	InExpr,
	IndexAccess,
	InstanceOfExpr,
	JSXElement,
	Literal,
	LogicalAND,
	LogicalOR,
	MultiplicativeExpr,
	NewExpr,
	NewTarget,
	NoElement,
	ObjectLiteral,
	OptionalCallExpr,
	OptionalIndexAccess,
	OptionalPropertyAccess,
	OptionalTaggedTemplate,
	Parenthesized,
	PostDec,
	PostInc,
	PreDec,
	PreInc,
	PropertyAccess,
	Regexp,
	RelationalExpr,
	ShiftExpr,
	SpreadElement,
	SuperExpr,
	TaggedTemplate,
	TemplateLiteral,
	This,
	TsAsConstExpr,
	TsAsExpr,
	TsCastExpr,
	TsDynamicImport,
	TsNonNull,
	UnaryExpr,
	Yield,
}

var IterationStmt = []NodeType{
	DoWhileStmt,
	ForInStmt,
	ForInStmtWithVar,
	ForOfStmt,
	ForOfStmtWithVar,
	ForStmt,
	ForStmtWithVar,
	WhileStmt,
}

var JSXAttribute = []NodeType{
	JSXNormalAttribute,
	JSXSpreadAttribute,
}

var JSXAttributeValue = []NodeType{
	JSXElement,
	JSXExpr,
	JSXLiteral,
}

var JSXChild = []NodeType{
	JSXElement,
	JSXExpr,
	JSXSpreadExpr,
	JSXText,
}

var MethodDefinition = []NodeType{
	AsyncGeneratorMethod,
	AsyncMethod,
	GeneratorMethod,
	Getter,
	Method,
	Setter,
}

var Modifier = []NodeType{
	Abstract,
	AccessibilityModifier,
	Declare,
	DecoratorCall,
	DecoratorExpr,
	Override,
	Readonly,
	Static,
}

var ModuleItem = []NodeType{
	AsyncFunc,
	AsyncGeneratorDeclaration,
	Block,
	BreakStmt,
	Class,
	ContinueStmt,
	DebuggerStmt,
	DoWhileStmt,
	EmptyStmt,
	ExportDecl,
	ExportDefault,
	ExprStmt,
	ForInStmt,
	ForInStmtWithVar,
	ForOfStmt,
	ForOfStmtWithVar,
	ForStmt,
	ForStmtWithVar,
	Func,
	Generator,
	IfStmt,
	ImportDecl,
	LabelledStmt,
	LexicalDecl,
	ReturnStmt,
	SwitchStmt,
	SyntaxProblem,
	ThrowStmt,
	TryStmt,
	TsAmbientClass,
	TsAmbientEnum,
	TsAmbientExportDecl,
	TsAmbientFunc,
	TsAmbientGlobal,
	TsAmbientImportAlias,
	TsAmbientInterface,
	TsAmbientModule,
	TsAmbientNamespace,
	TsAmbientTypeAlias,
	TsAmbientVar,
	TsEnum,
	TsExportAssignment,
	TsImportAliasDecl,
	TsImportRequireDecl,
	TsInterface,
	TsNamespace,
	TsNamespaceExportDecl,
	TypeAliasDecl,
	VarStmt,
	WhileStmt,
	WithStmt,
}

var NamedImport = []NodeType{
	ImportSpec,
	SyntaxProblem,
}

var Parameter = []NodeType{
	DefaultParameter,
	RestParameter,
	SyntaxProblem,
	TsThisParameter,
}

var PropertyDefinition = []NodeType{
	ObjectMethod,
	Property,
	ShorthandProperty,
	SpreadProperty,
	SyntaxProblem,
}

var PropertyName = []NodeType{
	ComputedPropertyName,
	LiteralPropertyName,
}

var PropertyPattern = []NodeType{
	PropertyBinding,
	SingleNameBinding,
	SyntaxProblem,
}

var Stmt = []NodeType{
	Block,
	BreakStmt,
	ContinueStmt,
	DebuggerStmt,
	DoWhileStmt,
	EmptyStmt,
	ExprStmt,
	ForInStmt,
	ForInStmtWithVar,
	ForOfStmt,
	ForOfStmtWithVar,
	ForStmt,
	ForStmtWithVar,
	IfStmt,
	LabelledStmt,
	ReturnStmt,
	SwitchStmt,
	ThrowStmt,
	TryStmt,
	VarStmt,
	WhileStmt,
	WithStmt,
}

var StmtListItem = []NodeType{
	AsyncFunc,
	AsyncGeneratorDeclaration,
	Block,
	BreakStmt,
	Class,
	ContinueStmt,
	DebuggerStmt,
	DoWhileStmt,
	EmptyStmt,
	ExprStmt,
	ForInStmt,
	ForInStmtWithVar,
	ForOfStmt,
	ForOfStmtWithVar,
	ForStmt,
	ForStmtWithVar,
	Func,
	Generator,
	IfStmt,
	LabelledStmt,
	LexicalDecl,
	ReturnStmt,
	SwitchStmt,
	SyntaxProblem,
	ThrowStmt,
	TryStmt,
	TsAmbientClass,
	TsAmbientEnum,
	TsAmbientExportDecl,
	TsAmbientFunc,
	TsAmbientGlobal,
	TsAmbientImportAlias,
	TsAmbientInterface,
	TsAmbientModule,
	TsAmbientNamespace,
	TsAmbientTypeAlias,
	TsAmbientVar,
	TsEnum,
	TsImportAliasDecl,
	TsInterface,
	TsNamespace,
	TypeAliasDecl,
	VarStmt,
	WhileStmt,
	WithStmt,
}

var TokenSet = []NodeType{
	NoSubstitutionTemplate,
	TemplateHead,
	TemplateMiddle,
	TemplateTail,
}

var TsAmbientElement = []NodeType{
	TsAmbientClass,
	TsAmbientEnum,
	TsAmbientExportDecl,
	TsAmbientFunc,
	TsAmbientGlobal,
	TsAmbientImportAlias,
	TsAmbientInterface,
	TsAmbientModule,
	TsAmbientNamespace,
	TsAmbientTypeAlias,
	TsAmbientVar,
}

var TsType = []NodeType{
	ArrayType,
	AssertsType,
	ConstructorType,
	FuncType,
	ImportType,
	IndexedAccessType,
	IntersectionType,
	KeyOfType,
	LiteralType,
	MappedType,
	NonNullableType,
	NullableType,
	ObjectType,
	ParenthesizedType,
	PredefinedType,
	ReadonlyType,
	RestType,
	ThisType,
	TsConditional,
	TupleType,
	TypePredicate,
	TypeQuery,
	TypeReference,
	TypeVar,
	UnionType,
	UniqueType,
}

var TupleMember = []NodeType{
	ArrayType,
	AssertsType,
	ConstructorType,
	FuncType,
	ImportType,
	IndexedAccessType,
	IntersectionType,
	KeyOfType,
	LiteralType,
	MappedType,
	NamedTupleMember,
	NonNullableType,
	NullableType,
	ObjectType,
	ParenthesizedType,
	PredefinedType,
	ReadonlyType,
	RestType,
	ThisType,
	TsConditional,
	TupleType,
	TypePredicate,
	TypeQuery,
	TypeReference,
	TypeVar,
	UnionType,
	UniqueType,
}

var TypeMember = []NodeType{
	CallSignature,
	ConstructSignature,
	Getter,
	IndexSignature,
	MethodSignature,
	PropertySignature,
	Setter,
}
