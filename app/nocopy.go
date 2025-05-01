package app

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}
