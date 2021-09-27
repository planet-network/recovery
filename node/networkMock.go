package node

type NetworkMock struct {
}

func CreateNetwork() *NetworkMock {
	return &NetworkMock{}
}
