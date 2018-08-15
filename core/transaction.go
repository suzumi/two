package core

type (
	Input struct {
		Amount uint64
	}

	Output struct {
		Amount uint64
	}

	Transaction struct {
		Inputs  []*Input
		Outputs []*Output
		Amount  uint64
	}
)
