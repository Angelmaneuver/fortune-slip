package lottery

import (
	"errors"

	"github.com/Angelmaneuver/fortune-slip/internal/lottery/lot"
)

type Lottery struct {
	lot lot.Lot
}

func New(path string) (*Lottery, error) {
	lot, err := lot.New(path)
	if err != nil {
		return nil, err
	}

	if lot.Size() == 0 {
		return nil, errors.New("there is no fortune")
	}

	return &Lottery{
		lot: *lot,
	}, nil
}

func (l Lottery) Draw() string {
	return l.lot.Picking()
}
