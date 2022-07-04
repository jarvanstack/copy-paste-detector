package matcher

import "path/filepath"

type GlobMatcher struct {
	contains map[string]struct{}
}

func NewGlobMatcher(containRules []string) Marcher {
	g := &GlobMatcher{
		contains: make(map[string]struct{}),
	}
	for _, rule := range containRules {
		matches, err := filepath.Glob(rule)
		if err != nil {
			panic(err)
		}
		for _, match := range matches {
			g.contains[match] = struct{}{}
		}
	}
	return g
}

func (g *GlobMatcher) Match(str string) bool {
	_, ok := g.contains[str]
	return ok
}
