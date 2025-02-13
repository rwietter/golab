package main

import (
	"learning/base"
	"learning/concurrency/channels"
	"learning/concurrency/issues"
)

func main() {
	// Capacidade de Slices
	println("\nSlicesCap")
	base.SlicesCap()
	println("\nMaps")
	base.Maps()
	println("\nVariadicFunctions")
	base.VariadicFunctions()
	println("\nMethods")
	base.Methods()
	println("\nCreateSnippetInterface")
	base.CreateSnippetInterface()
	println("\nMutableMethod")
	base.MutableMethod()
	println("\nGenerics")
	base.Generics()
	println("\nJSONOperations")
	base.JSONOperations()
	println("\nDeadlock")
	// issues.Deadlock()
	println("\nCorridorLivelock")
	issues.CorridorLivelock()
	println("\nStarvation")
	issues.Starvation()
	println("\nChanCounter")
	channels.ChanCounter()
	println("\nWorkerPool")
	channels.WorkerPool()
	println("\nChannelSnippet")
	channels.ChannelSnippet() // Canal padrão aprendido com multiplos canais e select com comma ok e quit channel para finalizar
	println("\nConvergePattern")
	channels.ConvergePattern() // Padrão de convergência de canais
}
