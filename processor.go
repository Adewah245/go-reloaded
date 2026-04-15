package main

func processor(s string) string{
	s = applycase(s)
	s = convert(s)
	s = fixquote(s)
	s = lastnword(s)
	return s
}