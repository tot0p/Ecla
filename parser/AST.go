package parser

// The AST struct contains all the information needed for the interpreter to run
type AST struct {
	Operations []Node
}

type File struct {
	ParseTree *AST
	Imports   []string
	// TODO: use a map instead of a slice for better performance
	Dependencies []string
	VariableDecl []string
	FunctionDecl []string
	Trace        string
}

func (f *File) DepChecker() (bool, []string) {
	Unresolved := []string{}
	for _, value := range f.Dependencies {
		if !contains(value, f.Imports) {
			Unresolved = append(Unresolved, value)
		}
	}
	if len(Unresolved) > 0 {
		return false, Unresolved
	}
	return true, Unresolved
}

func (f *File) AddDependency(dep string) {
	if !contains(dep, f.Dependencies) {
		f.Dependencies = append(f.Dependencies, dep)
	}
}

func contains(needle string, haystack []string) bool {
	for _, matchValue := range haystack {
		if matchValue == needle {
			return true
		}
	}
	return false
}
