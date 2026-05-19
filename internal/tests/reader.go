package tests

type NeverEnding byte

func (b NeverEnding) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
