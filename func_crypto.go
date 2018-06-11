package neovm

func opHash(e *ExecutionEngine) (VMState, error) {
	x := PopByteArray(e)
	PushData(e, Hash(x, e))
	return NONE, nil
}
