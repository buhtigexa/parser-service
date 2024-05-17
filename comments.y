%{
package main

import (
	"fmt"
)

type yySymType struct {
	yys  int
	word string
}

var yylval yySymType

var productInfo struct {
	Origin string
	Weight string
	Color  string
	Radius string
	Price  string
}

%}

%token ORIGIN WEIGHT COLOR RADIUS PRICE WORD NUMBER EOF

%%

input: origen
    ;

/*
peso:
    optional_words WEIGHT optional_words NUMBER EOF { }
    ;
*/

origen:
    optional_words ORIGIN pais optional_words EOF { }
    ;

optional_words:
    /* empty */
    | optional_words WORD
    {
        fmt.Printf(" Optional words=[%s] - Word=[%s] \n", $1.word, $2.word)
        $$.word +=" " + $2.word
    }
    ;

pais:
    WORD {
            fmt.Printf( "RULE : words =  %s \n",$1.word)
            $$ = $1
        }
    ;


%%

func yyError(s string) {
	fmt.Printf("Syntax error: %s\n", s)
}
