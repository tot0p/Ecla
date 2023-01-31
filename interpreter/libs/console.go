package libs

import (
	"github.com/Eclalang/console"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/interpreter/libs/utils"
)

type Console struct {
}

var (
	functionMap = map[string]interface{}{
		"printf":       nil,
		"println":      nil,
		"print":        nil,
		"input":        nil,
		"printInColor": nil,
	}
)

func (c *Console) Call(name string, args []eclaType.Type) eclaType.Type {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := functionMap[name]; !ok {
		return nil
	}
	switch name {
	case "printf":
		// TODO: refactor this line
		console.Printf(newArgs[0].(string), newArgs[1:]...)
	case "println":
		console.Println(newArgs...)
	case "print":
		console.Print(newArgs...)
	case "input":
		return utils.GoToEclaType(console.Input(newArgs...))
	case "printInColor":
		console.PrintInColor(newArgs[0].(string), newArgs[1:]...)
		// To add later
		//case "printlnInColor":
		//	console.PrintlnInColor(newArgs[0].(string), newArgs[1:]...)
	}

	return eclaType.Null{}
}
