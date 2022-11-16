package hangman_classic

func (g *HangmanGame) InitGameExecutions() {
	g.executions = append(g.executions, executionLookForAutoSave)
	g.executions = append(g.executions, executionDisplayBody)
	g.executions = append(g.executions, executionCheckForRemainingTries)
	g.executions = append(g.executions, executionWaitForInput)
	g.executions = append(g.executions, executionCheckForWord)
	g.executions = append(g.executions, executionCheckForVowel)
	g.executions = append(g.executions, executionCheckLetterIsUsed)
	g.executions = append(g.executions, executionCheckForLetterOccurence)
	g.executions = append(g.executions, executionCheckForWordDiscover)
	g.executions = append(g.executions, executionAddToUsedLetter)
}

func (g *HangmanGame) AddGameExecution(exec GameExecution) {
	g.executions = append(g.executions, exec)
}

func (g *HangmanGame) ReplaceExecution(exec GameExecution, targetName string) {
	for i, e := range g.executions {
		if e.Name == targetName {
			g.executions[i] = exec
			return
		}
	}
	println("Unable to find " + targetName + " execution in the list.")
}

func (g *HangmanGame) AddAfterExecution(exec GameExecution, targetName string) {
	for i, e := range g.executions {
		if e.Name == targetName {
			t := g.executions[i:]
			g.executions = append(append(g.executions[:i], exec), t...)
			return
		}
	}
	println("Unable to find " + targetName + " execution in the list.")
}

func (g *HangmanGame) AddBeforeExecution(exec GameExecution, targetName string) {
	for i, e := range g.executions {
		if e.Name == targetName {
			t := g.executions[i-1:]
			g.executions = append(append(g.executions[:i-1], exec), t...)
			return
		}
	}
	println("Unable to find " + targetName + " execution in the list.")
}
