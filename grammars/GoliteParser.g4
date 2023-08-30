parser grammar GoliteParser; 

options {
    tokenVocab = GoliteLexer;  
}

//=============================================================================
// Golite grammar
// ================================================================================
/*
// {} repetition. Zero or more occurrences of the symbol(s) defined in the braces.
// [] option. The symbol(s) defined in subscript brackets are not required to appear.
// () grouping. Any terminal symbol from within this group is valid.

Program = Types Declarations Functions 'eof'                                                       ;
Types = {TypeDeclaration}                                                                          ;
TypeDeclaration = 'type' 'id' 'struct' '{' Fields '}' ';'                                          ;
Fields = Decl ';' {Decl ';'}                                                                       ;
Decl = 'id' Type                                                                                   ;
Type = 'int' | 'bool' | '*' 'id'
                                                                                                   ;
Declarations = {Declaration}                                                                       ;
Declaration = 'var' Ids Type ';'                                                                   ;
Ids = 'id' {',' 'id'}
                                                                                                   ;
Functions = {Function}                                                                             ;
Function = 'func' 'id' Parameters [ReturnType] '{' Declarations Statements '}'                     ;
Parameters = '(' [ Decl {',' Decl}] ')'                                                            ;
ReturnType = Type
                                                                                                   ;
Statements = {Statement}                                                                           ;
Statement = Block | Assignment | Print | Delete | Conditional | Loop | Return  | Invocation        ;
Block = '{' Statements '}'                                                                         ;
Delete = 'delete' Expression ';'                                                                   ;
Assignment = LValue '=' ( Expression | 'scan' ) ';'                                                ;
Print = 'printf' '(' 'string' { ',' Expression} ')'  ';'                                           ;
Conditional = 'if' '(' Expression ')' Block ['else' Block]                                         ;
Loop = 'for' '(' Expression ')' Block                                                              ;
Return = 'return' [Expression] ';'                                                                 ;
Invocation = 'id' Arguments ';'                                                                    ;
Arguments = '(' [Expression {',' Expression}] ')'                                                  ;
LValue = 'id' {'.' id}
                                                                             ;
Expression = BoolTerm {'||' BoolTerm}                                                              ;
BoolTerm = EqualTerm {'&&' EqualTerm}                                                              ;
EqualTerm =  RelationTerm {('=='| '!=') RelationTerm}                                              ;
RelationTerm = SimpleTerm {('>'| '<' | '<=' | '>=') SimpleTerm}                                    ;
SimpleTerm = Term {('+'| '-') Term}                                                                ;
Term = UnaryTerm {('*'| '/') UnaryTerm}                                                            ;
UnaryTerm = '!' SelectorTerm | '-' SelectorTerm | SelectorTerm                                     ;
SelectorTerm = Factor {'.' 'id'}                                                                   ;
Factor = '(' Expression ')' | 'id' [Arguments] | 'number' | 'new' 'id' | 'true' | 'false' | 'nil'  ;
 */


program: types declarations functions EOF;

//NOTE no exit function
types: typeDeclaration*;

typeDeclaration: TYPE ID STRUCT LBRACE fields RBRACE SEMICOLON;

// fields: decl SEMICOLON (decl SEMICOLON)*; // eg: a int ; b bool ; c *d;
fields: decl SEMICOLON fieldsPrime*;  
fieldsPrime: decl SEMICOLON;

//NOTE no exit function
decl: ID type; // eg: a int;

//NOTE: no exit function
type: INT | BOOL | ASTERISK ID;

//NOTE no exit function
declarations: declaration*;

declaration: VAR ids type SEMICOLON;

//NOTE: no exit function
// ids: ID (COMMA ID)*; 
ids: ID idsPrime*; 
idsPrime: COMMA ID;

//NOTE: no exit function
functions: function*;

function: FUNC ID parameters (returnType)? LBRACE declarations statements RBRACE;

// parameters: LPAREN (decl (COMMA decl)*)? RPAREN;
parameters: LPAREN (decl parametersPrime*)? RPAREN;    
parametersPrime: COMMA decl;

//NOTE no exit function
returnType: type;

//NOTE no exit function
statements: statement*;

statement: block | assignment | print | delete | conditional | loop | ret | invocation;

block: LBRACE statements RBRACE;

delete: DELETE expression SEMICOLON;

// a = 1; a = scan; a.b = x.y.w.z * 3
assignment: lvalue EQUALS (expression | SCAN) SEMICOLON;

//print: PRINTF LPAREN STRING (COMMA expression)* RPAREN SEMICOLON;
print: PRINTF LPAREN STRING printPrime* RPAREN SEMICOLON; 
printPrime: COMMA expression;

conditional: IF LPAREN expression RPAREN block (ELSE block)? ;

loop: FOR LPAREN expression RPAREN block ;

ret: RETURN (expression)? SEMICOLON;

invocation: ID arguments SEMICOLON ;

//arguments: LPAREN (expression (COMMA expression)*)? RPAREN;
arguments: LPAREN (expression argumentsPrime*)? RPAREN; 
argumentsPrime: COMMA expression;

// lvalue: ID (DOT ID)*;
lvalue: ID lvaluePrime*;                        
lvaluePrime: DOT ID;

// expression: boolTerm (OR boolTerm)*;
expression: boolTerm expressionPrime*;          
expressionPrime: OR boolTerm;

// boolTerm: equalTerm (AND equalTerm)*;
boolTerm: equalTerm booltermPrime*;             
booltermPrime: AND equalTerm;

// equalTerm: relationTerm ((EQ | NEQ) relationTerm)*;

// Suppose var a *s ; var b *s
// a == b; a.b.c == x.y; 
equalTerm: relationTerm equalTermPrime*;             
equalTermPrime: op=(EQ | NEQ) relt=relationTerm;

// relationTerm: simpleTerm ((GT | LT | LTE | GTE) simpleTerm)*;
relationTerm: simpleTerm relationTermPrime*;    
relationTermPrime: op=(GT | LT | LTE | GTE) smpt=simpleTerm;

// simpleTerm: term ((PLUS | MINUS) term)*;
simpleTerm: term simpleTermPrime*;              
simpleTermPrime: op=(PLUS | MINUS) term;
 
// term: unaryTerm ((ASTERISK | FSLASH) unaryTerm)*;
term: unaryTerm termPrime*;                      
termPrime : op=(ASTERISK | FSLASH) unat=unaryTerm;

unaryTerm: NOT selectorTerm | MINUS selectorTerm | selectorTerm;

//NOTE: no exit function
// selectorTerm: factor (DOT ID)*;
selectorTerm: factor selectorTermPrime*;  
selectorTermPrime: DOT ID;

// factor: LPAREN expression RPAREN | ID (arguments)? | NUMBER | NEW ID | TRUE | FALSE | NIL;
factor: factorPrime1 | factorPrime2 | NUMBER | NEW ID | TRUE | FALSE | NIL;
factorPrime1: LPAREN expression RPAREN;
factorPrime2: ID (arguments)?;




// memo notes:

// Invocation != factorPrime2 = ID(arguments)
// invocation is used to something like `doSomething()`
// factorPrime2 is used to something like `x = doSomething() + 3;`