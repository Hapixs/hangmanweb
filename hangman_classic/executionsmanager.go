package hangman_classic

var Executions = []GameExecution{}

func InitGameExecutions() {

	Executions = append(Executions, executionLookForAutoSave)
	Executions = append(Executions, executionDisplayBody)
	Executions = append(Executions, executionCheckForRemainingTries)
	Executions = append(Executions, executionWaitForInput)
	Executions = append(Executions, executionCheckForWord)
	Executions = append(Executions, executionCheckForVowel)
	Executions = append(Executions, executionCheckLetterIsUsed)
	Executions = append(Executions, executionCheckForLetterOccurence)
	Executions = append(Executions, executionCheckForWordDiscover)
	Executions = append(Executions, executionAddToUsedLetter)

}

func AddGameExecution(exec GameExecution) {
	Executions = append(Executions, exec)
}

func ReplaceExecution(exec GameExecution, targetName string) {
	for i, e := range Executions {
		if e.Name == targetName {
			Executions[i] = exec
			return
		}
	}
	println("Unable to find " + targetName + " execution in the list.")
}

func AddAfterExecution(exec GameExecution, targetName string) {
	for i, e := range Executions {
		if e.Name == targetName {
			t := Executions[i:]
			Executions = append(append(Executions[:i], exec), t...)
			return
		}
	}
	println("Unable to find " + targetName + " execution in the list.")
}

func AddBeforeExecution(exec GameExecution, targetName string) {
	for i, e := range Executions {
		if e.Name == targetName {
			t := Executions[i-1:]
			Executions = append(append(Executions[:i-1], exec), t...)
			return
		}
	}
	println("Unable to find " + targetName + " execution in the list.")
}
