package main

import (
	"learning/base"
	"learning/concurrency/channels"
	"learning/concurrency/issues"
	"learning/concurrency/mutex"
	"learning/concurrency/patterns/fan"
	"learning/concurrency/patterns/pipeline"
	"learning/concurrency/patterns/pool"
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

	println("\nSimple Worker")
	channels.SimpleWorker()

	println("\nMutex")
	mutex.Mutex() // Mutex para controle de concorrência

	println("\nChannelSnippet")
	channels.ChannelSnippet() // Canal padrão aprendido com multiplos canais e select com comma ok e quit channel para finalizar

	println("\nConverge Pattern")
	channels.ConvergePattern() // Padrão de convergência de canais

	println("\nWorkerPool Pattern")
	pool.WorkerPoolPattern() // Padrão de pool de workers

	println("\nPipeline Pattern")
	pipeline.PipelinePattern() // Padrão de pipeline

	println("\nFanOut/FanIn Pattern")
	fan.FanOutFanInPattern() // Padrão de fan out/fan in

}
