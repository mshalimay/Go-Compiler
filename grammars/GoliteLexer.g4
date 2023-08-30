lexer grammar GoliteLexer; 

TYPE: 'type'        ;

// types terms
STRUCT: 'struct'    ;
VAR:    'var'       ;
INT:    'int'       ;
BOOL:   'bool'      ;
NIL:    'nil'       ;
FUNC:  'func'       ;

// boolean literals
TRUE: 'true'        ;
FALSE: 'false'      ;

// reserved keywords
DELETE: 'delete'    ;
NEW: 'new'          ;
SCAN: 'scan'        ;
PRINTF: 'printf'    ;

// statements
IF: 'if'            ;
ELSE: 'else'        ;
FOR: 'for'          ;
RETURN: 'return'    ;

// access items in structs
DOT: '.'            ;

// surrounding terms
LBRACE: '{'         ;
RBRACE: '}'         ; 
LPAREN: '('        ; 
RPAREN: ')'        ; 

// separators
COMMA: ','          ;
SEMICOLON: ';'      ;


// special symbols
NOT: '!'    ;
FSLASH: '/'         ;
ASTERISK: '*'       ;
MINUS: '-'           ;
PLUS: '+'           ;

// operators
EQUALS: '='         ;
LT: '<'             ;
GT: '>'             ;
LTE: '<='           ;
GTE: '>='           ;
EQ: '=='            ;
NEQ: '!='           ;
AND: '&&'           ;
OR: '||'            ;

// ID: begins w/ letter from eng alphabet, followed by 0 or more alphanumeric chars
ID: [a-zA-Z] [a-zA-Z0-9]*   ;

// NUMBER: a positive integer = sequence of digits | a negative integer = '-' followed by sequence of digits
NUMBER: [0-9]+       ;

// STRING: a sequence of characters surrounded by double quotes
STRING: '"'.*?'"'        ;

// WHITESPACE: any whitespace character; ignored
WS : [ \r\t\n]+ -> skip ;

// COMMENT: a sequence of characters starting with '//' and ending with a newline character
COMMENT: '//' .*? '\n' -> skip ;
