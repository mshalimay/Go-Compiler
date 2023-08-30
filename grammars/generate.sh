#!/bin/bash


alias antlr4='java -Xmx500M -cp "./antlr-4.13.0-complete.jar:$CLASSPATH" org.antlr.v4.Tool'

antlr4 -Dlanguage=Go -no-visitor -package lexer -o ../lexer GoliteLexer.g4
antlr4 -Dlanguage=Go -no-visitor -package parser -lib ../lexer -o ../parser GoliteParser.g4