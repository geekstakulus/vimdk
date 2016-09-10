package vimlparser

// ref: go/vimlparser.go
const (
	NODE_TOPLEVEL      = 1
	NODE_COMMENT       = 2
	NODE_EXCMD         = 3
	NODE_FUNCTION      = 4
	NODE_ENDFUNCTION   = 5
	NODE_DELFUNCTION   = 6
	NODE_RETURN        = 7
	NODE_EXCALL        = 8
	NODE_LET           = 9
	NODE_UNLET         = 10
	NODE_LOCKVAR       = 11
	NODE_UNLOCKVAR     = 12
	NODE_IF            = 13
	NODE_ELSEIF        = 14
	NODE_ELSE          = 15
	NODE_ENDIF         = 16
	NODE_WHILE         = 17
	NODE_ENDWHILE      = 18
	NODE_FOR           = 19
	NODE_ENDFOR        = 20
	NODE_CONTINUE      = 21
	NODE_BREAK         = 22
	NODE_TRY           = 23
	NODE_CATCH         = 24
	NODE_FINALLY       = 25
	NODE_ENDTRY        = 26
	NODE_THROW         = 27
	NODE_ECHO          = 28
	NODE_ECHON         = 29
	NODE_ECHOHL        = 30
	NODE_ECHOMSG       = 31
	NODE_ECHOERR       = 32
	NODE_EXECUTE       = 33
	NODE_TERNARY       = 34
	NODE_OR            = 35
	NODE_AND           = 36
	NODE_EQUAL         = 37
	NODE_EQUALCI       = 38
	NODE_EQUALCS       = 39
	NODE_NEQUAL        = 40
	NODE_NEQUALCI      = 41
	NODE_NEQUALCS      = 42
	NODE_GREATER       = 43
	NODE_GREATERCI     = 44
	NODE_GREATERCS     = 45
	NODE_GEQUAL        = 46
	NODE_GEQUALCI      = 47
	NODE_GEQUALCS      = 48
	NODE_SMALLER       = 49
	NODE_SMALLERCI     = 50
	NODE_SMALLERCS     = 51
	NODE_SEQUAL        = 52
	NODE_SEQUALCI      = 53
	NODE_SEQUALCS      = 54
	NODE_MATCH         = 55
	NODE_MATCHCI       = 56
	NODE_MATCHCS       = 57
	NODE_NOMATCH       = 58
	NODE_NOMATCHCI     = 59
	NODE_NOMATCHCS     = 60
	NODE_IS            = 61
	NODE_ISCI          = 62
	NODE_ISCS          = 63
	NODE_ISNOT         = 64
	NODE_ISNOTCI       = 65
	NODE_ISNOTCS       = 66
	NODE_ADD           = 67
	NODE_SUBTRACT      = 68
	NODE_CONCAT        = 69
	NODE_MULTIPLY      = 70
	NODE_DIVIDE        = 71
	NODE_REMAINDER     = 72
	NODE_NOT           = 73
	NODE_MINUS         = 74
	NODE_PLUS          = 75
	NODE_SUBSCRIPT     = 76
	NODE_SLICE         = 77
	NODE_CALL          = 78
	NODE_DOT           = 79
	NODE_NUMBER        = 80
	NODE_STRING        = 81
	NODE_LIST          = 82
	NODE_DICT          = 83
	NODE_OPTION        = 85
	NODE_IDENTIFIER    = 86
	NODE_CURLYNAME     = 87
	NODE_ENV           = 88
	NODE_REG           = 89
	NODE_CURLYNAMEPART = 90
	NODE_CURLYNAMEEXPR = 91
)
