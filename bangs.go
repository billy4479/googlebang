package main

//Bang struct, must escape %
type Bang struct {
	searchString string
	acronym      string
}

var (
	bangs = []Bang{
		{"https://github.com/search?utf8=%%E2%%9C%%93&q=%s", "gh"},
		{"https://duckduckgo.com/?q=%s", "duck"},
		{"https://www.amazon.it/s?k=%s&__mk_it_IT=%%C3%%85M%%C3%%85%%C5%%BD%%C3%%95%%C3%%91&ref=nb_sb_noss", "a"},
		{"https://www.youtube.com/results?search_query=%s", "y"},
		{"https://it.wikipedia.org/w/index.php?search=%s", "w"},
		{"https://www.ebay.it/sch/i.html?&_nkw=%s", "ebay"},
	}
)
