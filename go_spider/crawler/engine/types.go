package engine

type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

//定义一个空parser,占空位
func NilParser([]byte) ParseResult{
	return ParseResult{}
}