// Generated from /home/mshalimay/Compilers/proj-mashalimay/grammars/GoliteParser.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class GoliteParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		TYPE=1, STRUCT=2, VAR=3, INT=4, BOOL=5, NIL=6, FUNC=7, TRUE=8, FALSE=9, 
		DELETE=10, NEW=11, SCAN=12, PRINTF=13, IF=14, ELSE=15, FOR=16, RETURN=17, 
		DOT=18, LBRACE=19, RBRACE=20, LPAREN=21, RPAREN=22, COMMA=23, SEMICOLON=24, 
		NOT=25, FSLASH=26, ASTERISK=27, MINUS=28, PLUS=29, EQUALS=30, LT=31, GT=32, 
		LTE=33, GTE=34, EQ=35, NEQ=36, AND=37, OR=38, ID=39, NUMBER=40, STRING=41, 
		WS=42, COMMENT=43;
	public static final int
		RULE_program = 0, RULE_types = 1, RULE_typeDeclaration = 2, RULE_fields = 3, 
		RULE_fieldsPrime = 4, RULE_decl = 5, RULE_type = 6, RULE_declarations = 7, 
		RULE_declaration = 8, RULE_ids = 9, RULE_idsPrime = 10, RULE_functions = 11, 
		RULE_function = 12, RULE_parameters = 13, RULE_parametersPrime = 14, RULE_returnType = 15, 
		RULE_statements = 16, RULE_statement = 17, RULE_block = 18, RULE_delete = 19, 
		RULE_assignment = 20, RULE_print = 21, RULE_printPrime = 22, RULE_conditional = 23, 
		RULE_loop = 24, RULE_ret = 25, RULE_invocation = 26, RULE_arguments = 27, 
		RULE_argumentsPrime = 28, RULE_lvalue = 29, RULE_lvaluePrime = 30, RULE_expression = 31, 
		RULE_expressionPrime = 32, RULE_boolTerm = 33, RULE_booltermPrime = 34, 
		RULE_equalTerm = 35, RULE_equalTermPrime = 36, RULE_relationTerm = 37, 
		RULE_relationTermPrime = 38, RULE_simpleTerm = 39, RULE_simpleTermPrime = 40, 
		RULE_term = 41, RULE_termPrime = 42, RULE_unaryTerm = 43, RULE_selectorTerm = 44, 
		RULE_selectorTermPrime = 45, RULE_factor = 46, RULE_factorPrime1 = 47, 
		RULE_factorPrime2 = 48;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "types", "typeDeclaration", "fields", "fieldsPrime", "decl", 
			"type", "declarations", "declaration", "ids", "idsPrime", "functions", 
			"function", "parameters", "parametersPrime", "returnType", "statements", 
			"statement", "block", "delete", "assignment", "print", "printPrime", 
			"conditional", "loop", "ret", "invocation", "arguments", "argumentsPrime", 
			"lvalue", "lvaluePrime", "expression", "expressionPrime", "boolTerm", 
			"booltermPrime", "equalTerm", "equalTermPrime", "relationTerm", "relationTermPrime", 
			"simpleTerm", "simpleTermPrime", "term", "termPrime", "unaryTerm", "selectorTerm", 
			"selectorTermPrime", "factor", "factorPrime1", "factorPrime2"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'type'", "'struct'", "'var'", "'int'", "'bool'", "'nil'", "'func'", 
			"'true'", "'false'", "'delete'", "'new'", "'scan'", "'printf'", "'if'", 
			"'else'", "'for'", "'return'", "'.'", "'{'", "'}'", "'('", "')'", "','", 
			"';'", "'!'", "'/'", "'*'", "'-'", "'+'", "'='", "'<'", "'>'", "'<='", 
			"'>='", "'=='", "'!='", "'&&'", "'||'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "TYPE", "STRUCT", "VAR", "INT", "BOOL", "NIL", "FUNC", "TRUE", 
			"FALSE", "DELETE", "NEW", "SCAN", "PRINTF", "IF", "ELSE", "FOR", "RETURN", 
			"DOT", "LBRACE", "RBRACE", "LPAREN", "RPAREN", "COMMA", "SEMICOLON", 
			"NOT", "FSLASH", "ASTERISK", "MINUS", "PLUS", "EQUALS", "LT", "GT", "LTE", 
			"GTE", "EQ", "NEQ", "AND", "OR", "ID", "NUMBER", "STRING", "WS", "COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "GoliteParser.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public GoliteParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ProgramContext extends ParserRuleContext {
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public DeclarationsContext declarations() {
			return getRuleContext(DeclarationsContext.class,0);
		}
		public FunctionsContext functions() {
			return getRuleContext(FunctionsContext.class,0);
		}
		public TerminalNode EOF() { return getToken(GoliteParser.EOF, 0); }
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(98);
			types();
			setState(99);
			declarations();
			setState(100);
			functions();
			setState(101);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypesContext extends ParserRuleContext {
		public List<TypeDeclarationContext> typeDeclaration() {
			return getRuleContexts(TypeDeclarationContext.class);
		}
		public TypeDeclarationContext typeDeclaration(int i) {
			return getRuleContext(TypeDeclarationContext.class,i);
		}
		public TypesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_types; }
	}

	public final TypesContext types() throws RecognitionException {
		TypesContext _localctx = new TypesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_types);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(106);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==TYPE) {
				{
				{
				setState(103);
				typeDeclaration();
				}
				}
				setState(108);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeDeclarationContext extends ParserRuleContext {
		public TerminalNode TYPE() { return getToken(GoliteParser.TYPE, 0); }
		public TerminalNode ID() { return getToken(GoliteParser.ID, 0); }
		public TerminalNode STRUCT() { return getToken(GoliteParser.STRUCT, 0); }
		public TerminalNode LBRACE() { return getToken(GoliteParser.LBRACE, 0); }
		public FieldsContext fields() {
			return getRuleContext(FieldsContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(GoliteParser.RBRACE, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public TypeDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_typeDeclaration; }
	}

	public final TypeDeclarationContext typeDeclaration() throws RecognitionException {
		TypeDeclarationContext _localctx = new TypeDeclarationContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_typeDeclaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(109);
			match(TYPE);
			setState(110);
			match(ID);
			setState(111);
			match(STRUCT);
			setState(112);
			match(LBRACE);
			setState(113);
			fields();
			setState(114);
			match(RBRACE);
			setState(115);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FieldsContext extends ParserRuleContext {
		public DeclContext decl() {
			return getRuleContext(DeclContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public List<FieldsPrimeContext> fieldsPrime() {
			return getRuleContexts(FieldsPrimeContext.class);
		}
		public FieldsPrimeContext fieldsPrime(int i) {
			return getRuleContext(FieldsPrimeContext.class,i);
		}
		public FieldsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fields; }
	}

	public final FieldsContext fields() throws RecognitionException {
		FieldsContext _localctx = new FieldsContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_fields);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(117);
			decl();
			setState(118);
			match(SEMICOLON);
			setState(122);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ID) {
				{
				{
				setState(119);
				fieldsPrime();
				}
				}
				setState(124);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FieldsPrimeContext extends ParserRuleContext {
		public DeclContext decl() {
			return getRuleContext(DeclContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public FieldsPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fieldsPrime; }
	}

	public final FieldsPrimeContext fieldsPrime() throws RecognitionException {
		FieldsPrimeContext _localctx = new FieldsPrimeContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_fieldsPrime);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(125);
			decl();
			setState(126);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DeclContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GoliteParser.ID, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public DeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decl; }
	}

	public final DeclContext decl() throws RecognitionException {
		DeclContext _localctx = new DeclContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_decl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(128);
			match(ID);
			setState(129);
			type();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypeContext extends ParserRuleContext {
		public TerminalNode INT() { return getToken(GoliteParser.INT, 0); }
		public TerminalNode BOOL() { return getToken(GoliteParser.BOOL, 0); }
		public TerminalNode ASTERISK() { return getToken(GoliteParser.ASTERISK, 0); }
		public TerminalNode ID() { return getToken(GoliteParser.ID, 0); }
		public TypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type; }
	}

	public final TypeContext type() throws RecognitionException {
		TypeContext _localctx = new TypeContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_type);
		try {
			setState(135);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(131);
				match(INT);
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 2);
				{
				setState(132);
				match(BOOL);
				}
				break;
			case ASTERISK:
				enterOuterAlt(_localctx, 3);
				{
				setState(133);
				match(ASTERISK);
				setState(134);
				match(ID);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DeclarationsContext extends ParserRuleContext {
		public List<DeclarationContext> declaration() {
			return getRuleContexts(DeclarationContext.class);
		}
		public DeclarationContext declaration(int i) {
			return getRuleContext(DeclarationContext.class,i);
		}
		public DeclarationsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarations; }
	}

	public final DeclarationsContext declarations() throws RecognitionException {
		DeclarationsContext _localctx = new DeclarationsContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_declarations);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(140);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR) {
				{
				{
				setState(137);
				declaration();
				}
				}
				setState(142);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DeclarationContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(GoliteParser.VAR, 0); }
		public IdsContext ids() {
			return getRuleContext(IdsContext.class,0);
		}
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public DeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaration; }
	}

	public final DeclarationContext declaration() throws RecognitionException {
		DeclarationContext _localctx = new DeclarationContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_declaration);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(143);
			match(VAR);
			setState(144);
			ids();
			setState(145);
			type();
			setState(146);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IdsContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GoliteParser.ID, 0); }
		public List<IdsPrimeContext> idsPrime() {
			return getRuleContexts(IdsPrimeContext.class);
		}
		public IdsPrimeContext idsPrime(int i) {
			return getRuleContext(IdsPrimeContext.class,i);
		}
		public IdsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ids; }
	}

	public final IdsContext ids() throws RecognitionException {
		IdsContext _localctx = new IdsContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_ids);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(148);
			match(ID);
			setState(152);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(149);
				idsPrime();
				}
				}
				setState(154);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IdsPrimeContext extends ParserRuleContext {
		public TerminalNode COMMA() { return getToken(GoliteParser.COMMA, 0); }
		public TerminalNode ID() { return getToken(GoliteParser.ID, 0); }
		public IdsPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_idsPrime; }
	}

	public final IdsPrimeContext idsPrime() throws RecognitionException {
		IdsPrimeContext _localctx = new IdsPrimeContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_idsPrime);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(155);
			match(COMMA);
			setState(156);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionsContext extends ParserRuleContext {
		public List<FunctionContext> function() {
			return getRuleContexts(FunctionContext.class);
		}
		public FunctionContext function(int i) {
			return getRuleContext(FunctionContext.class,i);
		}
		public FunctionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functions; }
	}

	public final FunctionsContext functions() throws RecognitionException {
		FunctionsContext _localctx = new FunctionsContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_functions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(161);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==FUNC) {
				{
				{
				setState(158);
				function();
				}
				}
				setState(163);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionContext extends ParserRuleContext {
		public TerminalNode FUNC() { return getToken(GoliteParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(GoliteParser.ID, 0); }
		public ParametersContext parameters() {
			return getRuleContext(ParametersContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(GoliteParser.LBRACE, 0); }
		public DeclarationsContext declarations() {
			return getRuleContext(DeclarationsContext.class,0);
		}
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(GoliteParser.RBRACE, 0); }
		public ReturnTypeContext returnType() {
			return getRuleContext(ReturnTypeContext.class,0);
		}
		public FunctionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_function; }
	}

	public final FunctionContext function() throws RecognitionException {
		FunctionContext _localctx = new FunctionContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_function);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(164);
			match(FUNC);
			setState(165);
			match(ID);
			setState(166);
			parameters();
			setState(168);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << INT) | (1L << BOOL) | (1L << ASTERISK))) != 0)) {
				{
				setState(167);
				returnType();
				}
			}

			setState(170);
			match(LBRACE);
			setState(171);
			declarations();
			setState(172);
			statements();
			setState(173);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParametersContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public DeclContext decl() {
			return getRuleContext(DeclContext.class,0);
		}
		public List<ParametersPrimeContext> parametersPrime() {
			return getRuleContexts(ParametersPrimeContext.class);
		}
		public ParametersPrimeContext parametersPrime(int i) {
			return getRuleContext(ParametersPrimeContext.class,i);
		}
		public ParametersContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parameters; }
	}

	public final ParametersContext parameters() throws RecognitionException {
		ParametersContext _localctx = new ParametersContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_parameters);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(175);
			match(LPAREN);
			setState(183);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(176);
				decl();
				setState(180);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(177);
					parametersPrime();
					}
					}
					setState(182);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(185);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ParametersPrimeContext extends ParserRuleContext {
		public TerminalNode COMMA() { return getToken(GoliteParser.COMMA, 0); }
		public DeclContext decl() {
			return getRuleContext(DeclContext.class,0);
		}
		public ParametersPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametersPrime; }
	}

	public final ParametersPrimeContext parametersPrime() throws RecognitionException {
		ParametersPrimeContext _localctx = new ParametersPrimeContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_parametersPrime);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(187);
			match(COMMA);
			setState(188);
			decl();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ReturnTypeContext extends ParserRuleContext {
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public ReturnTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnType; }
	}

	public final ReturnTypeContext returnType() throws RecognitionException {
		ReturnTypeContext _localctx = new ReturnTypeContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_returnType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(190);
			type();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementsContext extends ParserRuleContext {
		public List<StatementContext> statement() {
			return getRuleContexts(StatementContext.class);
		}
		public StatementContext statement(int i) {
			return getRuleContext(StatementContext.class,i);
		}
		public StatementsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statements; }
	}

	public final StatementsContext statements() throws RecognitionException {
		StatementsContext _localctx = new StatementsContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_statements);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(195);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << DELETE) | (1L << PRINTF) | (1L << IF) | (1L << FOR) | (1L << RETURN) | (1L << LBRACE) | (1L << ID))) != 0)) {
				{
				{
				setState(192);
				statement();
				}
				}
				setState(197);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StatementContext extends ParserRuleContext {
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public AssignmentContext assignment() {
			return getRuleContext(AssignmentContext.class,0);
		}
		public PrintContext print() {
			return getRuleContext(PrintContext.class,0);
		}
		public DeleteContext delete() {
			return getRuleContext(DeleteContext.class,0);
		}
		public ConditionalContext conditional() {
			return getRuleContext(ConditionalContext.class,0);
		}
		public LoopContext loop() {
			return getRuleContext(LoopContext.class,0);
		}
		public RetContext ret() {
			return getRuleContext(RetContext.class,0);
		}
		public InvocationContext invocation() {
			return getRuleContext(InvocationContext.class,0);
		}
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_statement);
		try {
			setState(206);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(198);
				block();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(199);
				assignment();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(200);
				print();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(201);
				delete();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(202);
				conditional();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(203);
				loop();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(204);
				ret();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(205);
				invocation();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BlockContext extends ParserRuleContext {
		public TerminalNode LBRACE() { return getToken(GoliteParser.LBRACE, 0); }
		public StatementsContext statements() {
			return getRuleContext(StatementsContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(GoliteParser.RBRACE, 0); }
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_block);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(208);
			match(LBRACE);
			setState(209);
			statements();
			setState(210);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DeleteContext extends ParserRuleContext {
		public TerminalNode DELETE() { return getToken(GoliteParser.DELETE, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public DeleteContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_delete; }
	}

	public final DeleteContext delete() throws RecognitionException {
		DeleteContext _localctx = new DeleteContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_delete);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(212);
			match(DELETE);
			setState(213);
			expression();
			setState(214);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AssignmentContext extends ParserRuleContext {
		public LvalueContext lvalue() {
			return getRuleContext(LvalueContext.class,0);
		}
		public TerminalNode EQUALS() { return getToken(GoliteParser.EQUALS, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode SCAN() { return getToken(GoliteParser.SCAN, 0); }
		public AssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment; }
	}

	public final AssignmentContext assignment() throws RecognitionException {
		AssignmentContext _localctx = new AssignmentContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(216);
			lvalue();
			setState(217);
			match(EQUALS);
			setState(220);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NIL:
			case TRUE:
			case FALSE:
			case NEW:
			case LPAREN:
			case NOT:
			case MINUS:
			case ID:
			case NUMBER:
				{
				setState(218);
				expression();
				}
				break;
			case SCAN:
				{
				setState(219);
				match(SCAN);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(222);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PrintContext extends ParserRuleContext {
		public TerminalNode PRINTF() { return getToken(GoliteParser.PRINTF, 0); }
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public TerminalNode STRING() { return getToken(GoliteParser.STRING, 0); }
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public List<PrintPrimeContext> printPrime() {
			return getRuleContexts(PrintPrimeContext.class);
		}
		public PrintPrimeContext printPrime(int i) {
			return getRuleContext(PrintPrimeContext.class,i);
		}
		public PrintContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_print; }
	}

	public final PrintContext print() throws RecognitionException {
		PrintContext _localctx = new PrintContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_print);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(224);
			match(PRINTF);
			setState(225);
			match(LPAREN);
			setState(226);
			match(STRING);
			setState(230);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(227);
				printPrime();
				}
				}
				setState(232);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(233);
			match(RPAREN);
			setState(234);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PrintPrimeContext extends ParserRuleContext {
		public TerminalNode COMMA() { return getToken(GoliteParser.COMMA, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public PrintPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printPrime; }
	}

	public final PrintPrimeContext printPrime() throws RecognitionException {
		PrintPrimeContext _localctx = new PrintPrimeContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_printPrime);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(236);
			match(COMMA);
			setState(237);
			expression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConditionalContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(GoliteParser.IF, 0); }
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(GoliteParser.ELSE, 0); }
		public ConditionalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditional; }
	}

	public final ConditionalContext conditional() throws RecognitionException {
		ConditionalContext _localctx = new ConditionalContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_conditional);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(239);
			match(IF);
			setState(240);
			match(LPAREN);
			setState(241);
			expression();
			setState(242);
			match(RPAREN);
			setState(243);
			block();
			setState(246);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(244);
				match(ELSE);
				setState(245);
				block();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LoopContext extends ParserRuleContext {
		public TerminalNode FOR() { return getToken(GoliteParser.FOR, 0); }
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public LoopContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_loop; }
	}

	public final LoopContext loop() throws RecognitionException {
		LoopContext _localctx = new LoopContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_loop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(248);
			match(FOR);
			setState(249);
			match(LPAREN);
			setState(250);
			expression();
			setState(251);
			match(RPAREN);
			setState(252);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RetContext extends ParserRuleContext {
		public TerminalNode RETURN() { return getToken(GoliteParser.RETURN, 0); }
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public RetContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ret; }
	}

	public final RetContext ret() throws RecognitionException {
		RetContext _localctx = new RetContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_ret);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(254);
			match(RETURN);
			setState(256);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NIL) | (1L << TRUE) | (1L << FALSE) | (1L << NEW) | (1L << LPAREN) | (1L << NOT) | (1L << MINUS) | (1L << ID) | (1L << NUMBER))) != 0)) {
				{
				setState(255);
				expression();
				}
			}

			setState(258);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InvocationContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GoliteParser.ID, 0); }
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GoliteParser.SEMICOLON, 0); }
		public InvocationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_invocation; }
	}

	public final InvocationContext invocation() throws RecognitionException {
		InvocationContext _localctx = new InvocationContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_invocation);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(260);
			match(ID);
			setState(261);
			arguments();
			setState(262);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArgumentsContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public List<ArgumentsPrimeContext> argumentsPrime() {
			return getRuleContexts(ArgumentsPrimeContext.class);
		}
		public ArgumentsPrimeContext argumentsPrime(int i) {
			return getRuleContext(ArgumentsPrimeContext.class,i);
		}
		public ArgumentsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arguments; }
	}

	public final ArgumentsContext arguments() throws RecognitionException {
		ArgumentsContext _localctx = new ArgumentsContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_arguments);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(264);
			match(LPAREN);
			setState(272);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << NIL) | (1L << TRUE) | (1L << FALSE) | (1L << NEW) | (1L << LPAREN) | (1L << NOT) | (1L << MINUS) | (1L << ID) | (1L << NUMBER))) != 0)) {
				{
				setState(265);
				expression();
				setState(269);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(266);
					argumentsPrime();
					}
					}
					setState(271);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(274);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArgumentsPrimeContext extends ParserRuleContext {
		public TerminalNode COMMA() { return getToken(GoliteParser.COMMA, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ArgumentsPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_argumentsPrime; }
	}

	public final ArgumentsPrimeContext argumentsPrime() throws RecognitionException {
		ArgumentsPrimeContext _localctx = new ArgumentsPrimeContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_argumentsPrime);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(276);
			match(COMMA);
			setState(277);
			expression();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LvalueContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GoliteParser.ID, 0); }
		public List<LvaluePrimeContext> lvaluePrime() {
			return getRuleContexts(LvaluePrimeContext.class);
		}
		public LvaluePrimeContext lvaluePrime(int i) {
			return getRuleContext(LvaluePrimeContext.class,i);
		}
		public LvalueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lvalue; }
	}

	public final LvalueContext lvalue() throws RecognitionException {
		LvalueContext _localctx = new LvalueContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_lvalue);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(279);
			match(ID);
			setState(283);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(280);
				lvaluePrime();
				}
				}
				setState(285);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LvaluePrimeContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(GoliteParser.DOT, 0); }
		public TerminalNode ID() { return getToken(GoliteParser.ID, 0); }
		public LvaluePrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lvaluePrime; }
	}

	public final LvaluePrimeContext lvaluePrime() throws RecognitionException {
		LvaluePrimeContext _localctx = new LvaluePrimeContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_lvaluePrime);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(286);
			match(DOT);
			setState(287);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public BoolTermContext boolTerm() {
			return getRuleContext(BoolTermContext.class,0);
		}
		public List<ExpressionPrimeContext> expressionPrime() {
			return getRuleContexts(ExpressionPrimeContext.class);
		}
		public ExpressionPrimeContext expressionPrime(int i) {
			return getRuleContext(ExpressionPrimeContext.class,i);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(289);
			boolTerm();
			setState(293);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==OR) {
				{
				{
				setState(290);
				expressionPrime();
				}
				}
				setState(295);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionPrimeContext extends ParserRuleContext {
		public TerminalNode OR() { return getToken(GoliteParser.OR, 0); }
		public BoolTermContext boolTerm() {
			return getRuleContext(BoolTermContext.class,0);
		}
		public ExpressionPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionPrime; }
	}

	public final ExpressionPrimeContext expressionPrime() throws RecognitionException {
		ExpressionPrimeContext _localctx = new ExpressionPrimeContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_expressionPrime);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(296);
			match(OR);
			setState(297);
			boolTerm();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BoolTermContext extends ParserRuleContext {
		public EqualTermContext equalTerm() {
			return getRuleContext(EqualTermContext.class,0);
		}
		public List<BooltermPrimeContext> booltermPrime() {
			return getRuleContexts(BooltermPrimeContext.class);
		}
		public BooltermPrimeContext booltermPrime(int i) {
			return getRuleContext(BooltermPrimeContext.class,i);
		}
		public BoolTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_boolTerm; }
	}

	public final BoolTermContext boolTerm() throws RecognitionException {
		BoolTermContext _localctx = new BoolTermContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_boolTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(299);
			equalTerm();
			setState(303);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==AND) {
				{
				{
				setState(300);
				booltermPrime();
				}
				}
				setState(305);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BooltermPrimeContext extends ParserRuleContext {
		public TerminalNode AND() { return getToken(GoliteParser.AND, 0); }
		public EqualTermContext equalTerm() {
			return getRuleContext(EqualTermContext.class,0);
		}
		public BooltermPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_booltermPrime; }
	}

	public final BooltermPrimeContext booltermPrime() throws RecognitionException {
		BooltermPrimeContext _localctx = new BooltermPrimeContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_booltermPrime);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(306);
			match(AND);
			setState(307);
			equalTerm();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class EqualTermContext extends ParserRuleContext {
		public RelationTermContext relationTerm() {
			return getRuleContext(RelationTermContext.class,0);
		}
		public List<EqualTermPrimeContext> equalTermPrime() {
			return getRuleContexts(EqualTermPrimeContext.class);
		}
		public EqualTermPrimeContext equalTermPrime(int i) {
			return getRuleContext(EqualTermPrimeContext.class,i);
		}
		public EqualTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_equalTerm; }
	}

	public final EqualTermContext equalTerm() throws RecognitionException {
		EqualTermContext _localctx = new EqualTermContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_equalTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(309);
			relationTerm();
			setState(313);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==EQ || _la==NEQ) {
				{
				{
				setState(310);
				equalTermPrime();
				}
				}
				setState(315);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class EqualTermPrimeContext extends ParserRuleContext {
		public Token op;
		public RelationTermContext relt;
		public RelationTermContext relationTerm() {
			return getRuleContext(RelationTermContext.class,0);
		}
		public TerminalNode EQ() { return getToken(GoliteParser.EQ, 0); }
		public TerminalNode NEQ() { return getToken(GoliteParser.NEQ, 0); }
		public EqualTermPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_equalTermPrime; }
	}

	public final EqualTermPrimeContext equalTermPrime() throws RecognitionException {
		EqualTermPrimeContext _localctx = new EqualTermPrimeContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_equalTermPrime);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(316);
			((EqualTermPrimeContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==EQ || _la==NEQ) ) {
				((EqualTermPrimeContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(317);
			((EqualTermPrimeContext)_localctx).relt = relationTerm();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RelationTermContext extends ParserRuleContext {
		public SimpleTermContext simpleTerm() {
			return getRuleContext(SimpleTermContext.class,0);
		}
		public List<RelationTermPrimeContext> relationTermPrime() {
			return getRuleContexts(RelationTermPrimeContext.class);
		}
		public RelationTermPrimeContext relationTermPrime(int i) {
			return getRuleContext(RelationTermPrimeContext.class,i);
		}
		public RelationTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationTerm; }
	}

	public final RelationTermContext relationTerm() throws RecognitionException {
		RelationTermContext _localctx = new RelationTermContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_relationTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(319);
			simpleTerm();
			setState(323);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LT) | (1L << GT) | (1L << LTE) | (1L << GTE))) != 0)) {
				{
				{
				setState(320);
				relationTermPrime();
				}
				}
				setState(325);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RelationTermPrimeContext extends ParserRuleContext {
		public Token op;
		public SimpleTermContext smpt;
		public SimpleTermContext simpleTerm() {
			return getRuleContext(SimpleTermContext.class,0);
		}
		public TerminalNode GT() { return getToken(GoliteParser.GT, 0); }
		public TerminalNode LT() { return getToken(GoliteParser.LT, 0); }
		public TerminalNode LTE() { return getToken(GoliteParser.LTE, 0); }
		public TerminalNode GTE() { return getToken(GoliteParser.GTE, 0); }
		public RelationTermPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_relationTermPrime; }
	}

	public final RelationTermPrimeContext relationTermPrime() throws RecognitionException {
		RelationTermPrimeContext _localctx = new RelationTermPrimeContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_relationTermPrime);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(326);
			((RelationTermPrimeContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LT) | (1L << GT) | (1L << LTE) | (1L << GTE))) != 0)) ) {
				((RelationTermPrimeContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(327);
			((RelationTermPrimeContext)_localctx).smpt = simpleTerm();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SimpleTermContext extends ParserRuleContext {
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public List<SimpleTermPrimeContext> simpleTermPrime() {
			return getRuleContexts(SimpleTermPrimeContext.class);
		}
		public SimpleTermPrimeContext simpleTermPrime(int i) {
			return getRuleContext(SimpleTermPrimeContext.class,i);
		}
		public SimpleTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_simpleTerm; }
	}

	public final SimpleTermContext simpleTerm() throws RecognitionException {
		SimpleTermContext _localctx = new SimpleTermContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_simpleTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(329);
			term();
			setState(333);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==MINUS || _la==PLUS) {
				{
				{
				setState(330);
				simpleTermPrime();
				}
				}
				setState(335);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SimpleTermPrimeContext extends ParserRuleContext {
		public Token op;
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public TerminalNode PLUS() { return getToken(GoliteParser.PLUS, 0); }
		public TerminalNode MINUS() { return getToken(GoliteParser.MINUS, 0); }
		public SimpleTermPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_simpleTermPrime; }
	}

	public final SimpleTermPrimeContext simpleTermPrime() throws RecognitionException {
		SimpleTermPrimeContext _localctx = new SimpleTermPrimeContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_simpleTermPrime);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(336);
			((SimpleTermPrimeContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==MINUS || _la==PLUS) ) {
				((SimpleTermPrimeContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(337);
			term();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TermContext extends ParserRuleContext {
		public UnaryTermContext unaryTerm() {
			return getRuleContext(UnaryTermContext.class,0);
		}
		public List<TermPrimeContext> termPrime() {
			return getRuleContexts(TermPrimeContext.class);
		}
		public TermPrimeContext termPrime(int i) {
			return getRuleContext(TermPrimeContext.class,i);
		}
		public TermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_term; }
	}

	public final TermContext term() throws RecognitionException {
		TermContext _localctx = new TermContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_term);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(339);
			unaryTerm();
			setState(343);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==FSLASH || _la==ASTERISK) {
				{
				{
				setState(340);
				termPrime();
				}
				}
				setState(345);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TermPrimeContext extends ParserRuleContext {
		public Token op;
		public UnaryTermContext unat;
		public UnaryTermContext unaryTerm() {
			return getRuleContext(UnaryTermContext.class,0);
		}
		public TerminalNode ASTERISK() { return getToken(GoliteParser.ASTERISK, 0); }
		public TerminalNode FSLASH() { return getToken(GoliteParser.FSLASH, 0); }
		public TermPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_termPrime; }
	}

	public final TermPrimeContext termPrime() throws RecognitionException {
		TermPrimeContext _localctx = new TermPrimeContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_termPrime);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(346);
			((TermPrimeContext)_localctx).op = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==FSLASH || _la==ASTERISK) ) {
				((TermPrimeContext)_localctx).op = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(347);
			((TermPrimeContext)_localctx).unat = unaryTerm();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class UnaryTermContext extends ParserRuleContext {
		public TerminalNode NOT() { return getToken(GoliteParser.NOT, 0); }
		public SelectorTermContext selectorTerm() {
			return getRuleContext(SelectorTermContext.class,0);
		}
		public TerminalNode MINUS() { return getToken(GoliteParser.MINUS, 0); }
		public UnaryTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_unaryTerm; }
	}

	public final UnaryTermContext unaryTerm() throws RecognitionException {
		UnaryTermContext _localctx = new UnaryTermContext(_ctx, getState());
		enterRule(_localctx, 86, RULE_unaryTerm);
		try {
			setState(354);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NOT:
				enterOuterAlt(_localctx, 1);
				{
				setState(349);
				match(NOT);
				setState(350);
				selectorTerm();
				}
				break;
			case MINUS:
				enterOuterAlt(_localctx, 2);
				{
				setState(351);
				match(MINUS);
				setState(352);
				selectorTerm();
				}
				break;
			case NIL:
			case TRUE:
			case FALSE:
			case NEW:
			case LPAREN:
			case ID:
			case NUMBER:
				enterOuterAlt(_localctx, 3);
				{
				setState(353);
				selectorTerm();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SelectorTermContext extends ParserRuleContext {
		public FactorContext factor() {
			return getRuleContext(FactorContext.class,0);
		}
		public List<SelectorTermPrimeContext> selectorTermPrime() {
			return getRuleContexts(SelectorTermPrimeContext.class);
		}
		public SelectorTermPrimeContext selectorTermPrime(int i) {
			return getRuleContext(SelectorTermPrimeContext.class,i);
		}
		public SelectorTermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selectorTerm; }
	}

	public final SelectorTermContext selectorTerm() throws RecognitionException {
		SelectorTermContext _localctx = new SelectorTermContext(_ctx, getState());
		enterRule(_localctx, 88, RULE_selectorTerm);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(356);
			factor();
			setState(360);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOT) {
				{
				{
				setState(357);
				selectorTermPrime();
				}
				}
				setState(362);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SelectorTermPrimeContext extends ParserRuleContext {
		public TerminalNode DOT() { return getToken(GoliteParser.DOT, 0); }
		public TerminalNode ID() { return getToken(GoliteParser.ID, 0); }
		public SelectorTermPrimeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selectorTermPrime; }
	}

	public final SelectorTermPrimeContext selectorTermPrime() throws RecognitionException {
		SelectorTermPrimeContext _localctx = new SelectorTermPrimeContext(_ctx, getState());
		enterRule(_localctx, 90, RULE_selectorTermPrime);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(363);
			match(DOT);
			setState(364);
			match(ID);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FactorContext extends ParserRuleContext {
		public FactorPrime1Context factorPrime1() {
			return getRuleContext(FactorPrime1Context.class,0);
		}
		public FactorPrime2Context factorPrime2() {
			return getRuleContext(FactorPrime2Context.class,0);
		}
		public TerminalNode NUMBER() { return getToken(GoliteParser.NUMBER, 0); }
		public TerminalNode NEW() { return getToken(GoliteParser.NEW, 0); }
		public TerminalNode ID() { return getToken(GoliteParser.ID, 0); }
		public TerminalNode TRUE() { return getToken(GoliteParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(GoliteParser.FALSE, 0); }
		public TerminalNode NIL() { return getToken(GoliteParser.NIL, 0); }
		public FactorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_factor; }
	}

	public final FactorContext factor() throws RecognitionException {
		FactorContext _localctx = new FactorContext(_ctx, getState());
		enterRule(_localctx, 92, RULE_factor);
		try {
			setState(374);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LPAREN:
				enterOuterAlt(_localctx, 1);
				{
				setState(366);
				factorPrime1();
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 2);
				{
				setState(367);
				factorPrime2();
				}
				break;
			case NUMBER:
				enterOuterAlt(_localctx, 3);
				{
				setState(368);
				match(NUMBER);
				}
				break;
			case NEW:
				enterOuterAlt(_localctx, 4);
				{
				setState(369);
				match(NEW);
				setState(370);
				match(ID);
				}
				break;
			case TRUE:
				enterOuterAlt(_localctx, 5);
				{
				setState(371);
				match(TRUE);
				}
				break;
			case FALSE:
				enterOuterAlt(_localctx, 6);
				{
				setState(372);
				match(FALSE);
				}
				break;
			case NIL:
				enterOuterAlt(_localctx, 7);
				{
				setState(373);
				match(NIL);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FactorPrime1Context extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(GoliteParser.LPAREN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(GoliteParser.RPAREN, 0); }
		public FactorPrime1Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_factorPrime1; }
	}

	public final FactorPrime1Context factorPrime1() throws RecognitionException {
		FactorPrime1Context _localctx = new FactorPrime1Context(_ctx, getState());
		enterRule(_localctx, 94, RULE_factorPrime1);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(376);
			match(LPAREN);
			setState(377);
			expression();
			setState(378);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FactorPrime2Context extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GoliteParser.ID, 0); }
		public ArgumentsContext arguments() {
			return getRuleContext(ArgumentsContext.class,0);
		}
		public FactorPrime2Context(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_factorPrime2; }
	}

	public final FactorPrime2Context factorPrime2() throws RecognitionException {
		FactorPrime2Context _localctx = new FactorPrime2Context(_ctx, getState());
		enterRule(_localctx, 96, RULE_factorPrime2);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(380);
			match(ID);
			setState(382);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LPAREN) {
				{
				setState(381);
				arguments();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3-\u0183\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\3\2\3\2\3\2\3\2\3"+
		"\2\3\3\7\3k\n\3\f\3\16\3n\13\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5"+
		"\3\5\7\5{\n\5\f\5\16\5~\13\5\3\6\3\6\3\6\3\7\3\7\3\7\3\b\3\b\3\b\3\b\5"+
		"\b\u008a\n\b\3\t\7\t\u008d\n\t\f\t\16\t\u0090\13\t\3\n\3\n\3\n\3\n\3\n"+
		"\3\13\3\13\7\13\u0099\n\13\f\13\16\13\u009c\13\13\3\f\3\f\3\f\3\r\7\r"+
		"\u00a2\n\r\f\r\16\r\u00a5\13\r\3\16\3\16\3\16\3\16\5\16\u00ab\n\16\3\16"+
		"\3\16\3\16\3\16\3\16\3\17\3\17\3\17\7\17\u00b5\n\17\f\17\16\17\u00b8\13"+
		"\17\5\17\u00ba\n\17\3\17\3\17\3\20\3\20\3\20\3\21\3\21\3\22\7\22\u00c4"+
		"\n\22\f\22\16\22\u00c7\13\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\5"+
		"\23\u00d1\n\23\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\26\3\26\3\26"+
		"\3\26\5\26\u00df\n\26\3\26\3\26\3\27\3\27\3\27\3\27\7\27\u00e7\n\27\f"+
		"\27\16\27\u00ea\13\27\3\27\3\27\3\27\3\30\3\30\3\30\3\31\3\31\3\31\3\31"+
		"\3\31\3\31\3\31\5\31\u00f9\n\31\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33"+
		"\5\33\u0103\n\33\3\33\3\33\3\34\3\34\3\34\3\34\3\35\3\35\3\35\7\35\u010e"+
		"\n\35\f\35\16\35\u0111\13\35\5\35\u0113\n\35\3\35\3\35\3\36\3\36\3\36"+
		"\3\37\3\37\7\37\u011c\n\37\f\37\16\37\u011f\13\37\3 \3 \3 \3!\3!\7!\u0126"+
		"\n!\f!\16!\u0129\13!\3\"\3\"\3\"\3#\3#\7#\u0130\n#\f#\16#\u0133\13#\3"+
		"$\3$\3$\3%\3%\7%\u013a\n%\f%\16%\u013d\13%\3&\3&\3&\3\'\3\'\7\'\u0144"+
		"\n\'\f\'\16\'\u0147\13\'\3(\3(\3(\3)\3)\7)\u014e\n)\f)\16)\u0151\13)\3"+
		"*\3*\3*\3+\3+\7+\u0158\n+\f+\16+\u015b\13+\3,\3,\3,\3-\3-\3-\3-\3-\5-"+
		"\u0165\n-\3.\3.\7.\u0169\n.\f.\16.\u016c\13.\3/\3/\3/\3\60\3\60\3\60\3"+
		"\60\3\60\3\60\3\60\3\60\5\60\u0179\n\60\3\61\3\61\3\61\3\61\3\62\3\62"+
		"\5\62\u0181\n\62\3\62\2\2\63\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \""+
		"$&(*,.\60\62\64\668:<>@BDFHJLNPRTVXZ\\^`b\2\6\3\2%&\3\2!$\3\2\36\37\3"+
		"\2\34\35\2\u017a\2d\3\2\2\2\4l\3\2\2\2\6o\3\2\2\2\bw\3\2\2\2\n\177\3\2"+
		"\2\2\f\u0082\3\2\2\2\16\u0089\3\2\2\2\20\u008e\3\2\2\2\22\u0091\3\2\2"+
		"\2\24\u0096\3\2\2\2\26\u009d\3\2\2\2\30\u00a3\3\2\2\2\32\u00a6\3\2\2\2"+
		"\34\u00b1\3\2\2\2\36\u00bd\3\2\2\2 \u00c0\3\2\2\2\"\u00c5\3\2\2\2$\u00d0"+
		"\3\2\2\2&\u00d2\3\2\2\2(\u00d6\3\2\2\2*\u00da\3\2\2\2,\u00e2\3\2\2\2."+
		"\u00ee\3\2\2\2\60\u00f1\3\2\2\2\62\u00fa\3\2\2\2\64\u0100\3\2\2\2\66\u0106"+
		"\3\2\2\28\u010a\3\2\2\2:\u0116\3\2\2\2<\u0119\3\2\2\2>\u0120\3\2\2\2@"+
		"\u0123\3\2\2\2B\u012a\3\2\2\2D\u012d\3\2\2\2F\u0134\3\2\2\2H\u0137\3\2"+
		"\2\2J\u013e\3\2\2\2L\u0141\3\2\2\2N\u0148\3\2\2\2P\u014b\3\2\2\2R\u0152"+
		"\3\2\2\2T\u0155\3\2\2\2V\u015c\3\2\2\2X\u0164\3\2\2\2Z\u0166\3\2\2\2\\"+
		"\u016d\3\2\2\2^\u0178\3\2\2\2`\u017a\3\2\2\2b\u017e\3\2\2\2de\5\4\3\2"+
		"ef\5\20\t\2fg\5\30\r\2gh\7\2\2\3h\3\3\2\2\2ik\5\6\4\2ji\3\2\2\2kn\3\2"+
		"\2\2lj\3\2\2\2lm\3\2\2\2m\5\3\2\2\2nl\3\2\2\2op\7\3\2\2pq\7)\2\2qr\7\4"+
		"\2\2rs\7\25\2\2st\5\b\5\2tu\7\26\2\2uv\7\32\2\2v\7\3\2\2\2wx\5\f\7\2x"+
		"|\7\32\2\2y{\5\n\6\2zy\3\2\2\2{~\3\2\2\2|z\3\2\2\2|}\3\2\2\2}\t\3\2\2"+
		"\2~|\3\2\2\2\177\u0080\5\f\7\2\u0080\u0081\7\32\2\2\u0081\13\3\2\2\2\u0082"+
		"\u0083\7)\2\2\u0083\u0084\5\16\b\2\u0084\r\3\2\2\2\u0085\u008a\7\6\2\2"+
		"\u0086\u008a\7\7\2\2\u0087\u0088\7\35\2\2\u0088\u008a\7)\2\2\u0089\u0085"+
		"\3\2\2\2\u0089\u0086\3\2\2\2\u0089\u0087\3\2\2\2\u008a\17\3\2\2\2\u008b"+
		"\u008d\5\22\n\2\u008c\u008b\3\2\2\2\u008d\u0090\3\2\2\2\u008e\u008c\3"+
		"\2\2\2\u008e\u008f\3\2\2\2\u008f\21\3\2\2\2\u0090\u008e\3\2\2\2\u0091"+
		"\u0092\7\5\2\2\u0092\u0093\5\24\13\2\u0093\u0094\5\16\b\2\u0094\u0095"+
		"\7\32\2\2\u0095\23\3\2\2\2\u0096\u009a\7)\2\2\u0097\u0099\5\26\f\2\u0098"+
		"\u0097\3\2\2\2\u0099\u009c\3\2\2\2\u009a\u0098\3\2\2\2\u009a\u009b\3\2"+
		"\2\2\u009b\25\3\2\2\2\u009c\u009a\3\2\2\2\u009d\u009e\7\31\2\2\u009e\u009f"+
		"\7)\2\2\u009f\27\3\2\2\2\u00a0\u00a2\5\32\16\2\u00a1\u00a0\3\2\2\2\u00a2"+
		"\u00a5\3\2\2\2\u00a3\u00a1\3\2\2\2\u00a3\u00a4\3\2\2\2\u00a4\31\3\2\2"+
		"\2\u00a5\u00a3\3\2\2\2\u00a6\u00a7\7\t\2\2\u00a7\u00a8\7)\2\2\u00a8\u00aa"+
		"\5\34\17\2\u00a9\u00ab\5 \21\2\u00aa\u00a9\3\2\2\2\u00aa\u00ab\3\2\2\2"+
		"\u00ab\u00ac\3\2\2\2\u00ac\u00ad\7\25\2\2\u00ad\u00ae\5\20\t\2\u00ae\u00af"+
		"\5\"\22\2\u00af\u00b0\7\26\2\2\u00b0\33\3\2\2\2\u00b1\u00b9\7\27\2\2\u00b2"+
		"\u00b6\5\f\7\2\u00b3\u00b5\5\36\20\2\u00b4\u00b3\3\2\2\2\u00b5\u00b8\3"+
		"\2\2\2\u00b6\u00b4\3\2\2\2\u00b6\u00b7\3\2\2\2\u00b7\u00ba\3\2\2\2\u00b8"+
		"\u00b6\3\2\2\2\u00b9\u00b2\3\2\2\2\u00b9\u00ba\3\2\2\2\u00ba\u00bb\3\2"+
		"\2\2\u00bb\u00bc\7\30\2\2\u00bc\35\3\2\2\2\u00bd\u00be\7\31\2\2\u00be"+
		"\u00bf\5\f\7\2\u00bf\37\3\2\2\2\u00c0\u00c1\5\16\b\2\u00c1!\3\2\2\2\u00c2"+
		"\u00c4\5$\23\2\u00c3\u00c2\3\2\2\2\u00c4\u00c7\3\2\2\2\u00c5\u00c3\3\2"+
		"\2\2\u00c5\u00c6\3\2\2\2\u00c6#\3\2\2\2\u00c7\u00c5\3\2\2\2\u00c8\u00d1"+
		"\5&\24\2\u00c9\u00d1\5*\26\2\u00ca\u00d1\5,\27\2\u00cb\u00d1\5(\25\2\u00cc"+
		"\u00d1\5\60\31\2\u00cd\u00d1\5\62\32\2\u00ce\u00d1\5\64\33\2\u00cf\u00d1"+
		"\5\66\34\2\u00d0\u00c8\3\2\2\2\u00d0\u00c9\3\2\2\2\u00d0\u00ca\3\2\2\2"+
		"\u00d0\u00cb\3\2\2\2\u00d0\u00cc\3\2\2\2\u00d0\u00cd\3\2\2\2\u00d0\u00ce"+
		"\3\2\2\2\u00d0\u00cf\3\2\2\2\u00d1%\3\2\2\2\u00d2\u00d3\7\25\2\2\u00d3"+
		"\u00d4\5\"\22\2\u00d4\u00d5\7\26\2\2\u00d5\'\3\2\2\2\u00d6\u00d7\7\f\2"+
		"\2\u00d7\u00d8\5@!\2\u00d8\u00d9\7\32\2\2\u00d9)\3\2\2\2\u00da\u00db\5"+
		"<\37\2\u00db\u00de\7 \2\2\u00dc\u00df\5@!\2\u00dd\u00df\7\16\2\2\u00de"+
		"\u00dc\3\2\2\2\u00de\u00dd\3\2\2\2\u00df\u00e0\3\2\2\2\u00e0\u00e1\7\32"+
		"\2\2\u00e1+\3\2\2\2\u00e2\u00e3\7\17\2\2\u00e3\u00e4\7\27\2\2\u00e4\u00e8"+
		"\7+\2\2\u00e5\u00e7\5.\30\2\u00e6\u00e5\3\2\2\2\u00e7\u00ea\3\2\2\2\u00e8"+
		"\u00e6\3\2\2\2\u00e8\u00e9\3\2\2\2\u00e9\u00eb\3\2\2\2\u00ea\u00e8\3\2"+
		"\2\2\u00eb\u00ec\7\30\2\2\u00ec\u00ed\7\32\2\2\u00ed-\3\2\2\2\u00ee\u00ef"+
		"\7\31\2\2\u00ef\u00f0\5@!\2\u00f0/\3\2\2\2\u00f1\u00f2\7\20\2\2\u00f2"+
		"\u00f3\7\27\2\2\u00f3\u00f4\5@!\2\u00f4\u00f5\7\30\2\2\u00f5\u00f8\5&"+
		"\24\2\u00f6\u00f7\7\21\2\2\u00f7\u00f9\5&\24\2\u00f8\u00f6\3\2\2\2\u00f8"+
		"\u00f9\3\2\2\2\u00f9\61\3\2\2\2\u00fa\u00fb\7\22\2\2\u00fb\u00fc\7\27"+
		"\2\2\u00fc\u00fd\5@!\2\u00fd\u00fe\7\30\2\2\u00fe\u00ff\5&\24\2\u00ff"+
		"\63\3\2\2\2\u0100\u0102\7\23\2\2\u0101\u0103\5@!\2\u0102\u0101\3\2\2\2"+
		"\u0102\u0103\3\2\2\2\u0103\u0104\3\2\2\2\u0104\u0105\7\32\2\2\u0105\65"+
		"\3\2\2\2\u0106\u0107\7)\2\2\u0107\u0108\58\35\2\u0108\u0109\7\32\2\2\u0109"+
		"\67\3\2\2\2\u010a\u0112\7\27\2\2\u010b\u010f\5@!\2\u010c\u010e\5:\36\2"+
		"\u010d\u010c\3\2\2\2\u010e\u0111\3\2\2\2\u010f\u010d\3\2\2\2\u010f\u0110"+
		"\3\2\2\2\u0110\u0113\3\2\2\2\u0111\u010f\3\2\2\2\u0112\u010b\3\2\2\2\u0112"+
		"\u0113\3\2\2\2\u0113\u0114\3\2\2\2\u0114\u0115\7\30\2\2\u01159\3\2\2\2"+
		"\u0116\u0117\7\31\2\2\u0117\u0118\5@!\2\u0118;\3\2\2\2\u0119\u011d\7)"+
		"\2\2\u011a\u011c\5> \2\u011b\u011a\3\2\2\2\u011c\u011f\3\2\2\2\u011d\u011b"+
		"\3\2\2\2\u011d\u011e\3\2\2\2\u011e=\3\2\2\2\u011f\u011d\3\2\2\2\u0120"+
		"\u0121\7\24\2\2\u0121\u0122\7)\2\2\u0122?\3\2\2\2\u0123\u0127\5D#\2\u0124"+
		"\u0126\5B\"\2\u0125\u0124\3\2\2\2\u0126\u0129\3\2\2\2\u0127\u0125\3\2"+
		"\2\2\u0127\u0128\3\2\2\2\u0128A\3\2\2\2\u0129\u0127\3\2\2\2\u012a\u012b"+
		"\7(\2\2\u012b\u012c\5D#\2\u012cC\3\2\2\2\u012d\u0131\5H%\2\u012e\u0130"+
		"\5F$\2\u012f\u012e\3\2\2\2\u0130\u0133\3\2\2\2\u0131\u012f\3\2\2\2\u0131"+
		"\u0132\3\2\2\2\u0132E\3\2\2\2\u0133\u0131\3\2\2\2\u0134\u0135\7\'\2\2"+
		"\u0135\u0136\5H%\2\u0136G\3\2\2\2\u0137\u013b\5L\'\2\u0138\u013a\5J&\2"+
		"\u0139\u0138\3\2\2\2\u013a\u013d\3\2\2\2\u013b\u0139\3\2\2\2\u013b\u013c"+
		"\3\2\2\2\u013cI\3\2\2\2\u013d\u013b\3\2\2\2\u013e\u013f\t\2\2\2\u013f"+
		"\u0140\5L\'\2\u0140K\3\2\2\2\u0141\u0145\5P)\2\u0142\u0144\5N(\2\u0143"+
		"\u0142\3\2\2\2\u0144\u0147\3\2\2\2\u0145\u0143\3\2\2\2\u0145\u0146\3\2"+
		"\2\2\u0146M\3\2\2\2\u0147\u0145\3\2\2\2\u0148\u0149\t\3\2\2\u0149\u014a"+
		"\5P)\2\u014aO\3\2\2\2\u014b\u014f\5T+\2\u014c\u014e\5R*\2\u014d\u014c"+
		"\3\2\2\2\u014e\u0151\3\2\2\2\u014f\u014d\3\2\2\2\u014f\u0150\3\2\2\2\u0150"+
		"Q\3\2\2\2\u0151\u014f\3\2\2\2\u0152\u0153\t\4\2\2\u0153\u0154\5T+\2\u0154"+
		"S\3\2\2\2\u0155\u0159\5X-\2\u0156\u0158\5V,\2\u0157\u0156\3\2\2\2\u0158"+
		"\u015b\3\2\2\2\u0159\u0157\3\2\2\2\u0159\u015a\3\2\2\2\u015aU\3\2\2\2"+
		"\u015b\u0159\3\2\2\2\u015c\u015d\t\5\2\2\u015d\u015e\5X-\2\u015eW\3\2"+
		"\2\2\u015f\u0160\7\33\2\2\u0160\u0165\5Z.\2\u0161\u0162\7\36\2\2\u0162"+
		"\u0165\5Z.\2\u0163\u0165\5Z.\2\u0164\u015f\3\2\2\2\u0164\u0161\3\2\2\2"+
		"\u0164\u0163\3\2\2\2\u0165Y\3\2\2\2\u0166\u016a\5^\60\2\u0167\u0169\5"+
		"\\/\2\u0168\u0167\3\2\2\2\u0169\u016c\3\2\2\2\u016a\u0168\3\2\2\2\u016a"+
		"\u016b\3\2\2\2\u016b[\3\2\2\2\u016c\u016a\3\2\2\2\u016d\u016e\7\24\2\2"+
		"\u016e\u016f\7)\2\2\u016f]\3\2\2\2\u0170\u0179\5`\61\2\u0171\u0179\5b"+
		"\62\2\u0172\u0179\7*\2\2\u0173\u0174\7\r\2\2\u0174\u0179\7)\2\2\u0175"+
		"\u0179\7\n\2\2\u0176\u0179\7\13\2\2\u0177\u0179\7\b\2\2\u0178\u0170\3"+
		"\2\2\2\u0178\u0171\3\2\2\2\u0178\u0172\3\2\2\2\u0178\u0173\3\2\2\2\u0178"+
		"\u0175\3\2\2\2\u0178\u0176\3\2\2\2\u0178\u0177\3\2\2\2\u0179_\3\2\2\2"+
		"\u017a\u017b\7\27\2\2\u017b\u017c\5@!\2\u017c\u017d\7\30\2\2\u017da\3"+
		"\2\2\2\u017e\u0180\7)\2\2\u017f\u0181\58\35\2\u0180\u017f\3\2\2\2\u0180"+
		"\u0181\3\2\2\2\u0181c\3\2\2\2\36l|\u0089\u008e\u009a\u00a3\u00aa\u00b6"+
		"\u00b9\u00c5\u00d0\u00de\u00e8\u00f8\u0102\u010f\u0112\u011d\u0127\u0131"+
		"\u013b\u0145\u014f\u0159\u0164\u016a\u0178\u0180";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}