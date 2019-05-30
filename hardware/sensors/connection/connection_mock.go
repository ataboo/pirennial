package connection

import "fmt"

type ConnectionMock struct {
	GetResponse []byte
	ThrowError  bool
	Open        bool
}

func CreateConnectionMock() *ConnectionMock {
	c := ConnectionMock{
		ThrowError: false,
		Open:       true,
	}

	return &c
}

func (c *ConnectionMock) Read(buff []byte) (n int, err error) {
	if c.ThrowError {
		return 0, fmt.Errorf("failed to read (mock)")
	}

	copy(buff, c.GetResponse)

	return len(c.GetResponse), nil
}

func (c *ConnectionMock) Write(val []byte) (n int, err error) {
	if c.ThrowError {
		return 0, fmt.Errorf("failed to write (mock)")
	}

	return len(val), nil
}

func (c *ConnectionMock) Close() error {
	if c.Open {
		c.Open = false
		return nil
	}

	return fmt.Errorf("already closed")
}
