package lot

import (
	"math/rand"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/Angelmaneuver/fortune-slip/internal/lottery/lot/parameter"
	"github.com/gabriel-vasile/mimetype"
)

type Index struct {
	selected   []int
	unselected []int
}

func (i *Index) move(index int) int {
	value := i.unselected[index]

	i.selected = append(i.selected, value)
	i.unselected = slices.Delete(i.unselected, index, index+1)

	return value
}

type Lot struct {
	parameter parameter.Parameter
	contents  []string
	index     *Index
}

func New(path string) (*Lot, error) {
	var lot = Lot{
		parameter: *parameter.New(path),
		index:     &Index{},
	}

	var patterns = lot.parameter.Patterns()

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		mtype, err := mimetype.DetectFile(path)
		if err != nil {
			return nil
		}

		for _, pattern := range patterns {
			if strings.HasPrefix(mtype.String(), pattern) {
				lot.contents = append(lot.contents, path)
				continue
			}
		}

		return nil
	})

	lot.index.selected = make([]int, 0, len(lot.contents))
	lot.index.unselected = rand.Perm(len(lot.contents))

	return &lot, err
}

func (l Lot) Parameter() parameter.Parameter {
	return l.parameter
}

func (l Lot) Contents(index int) string {
	return l.contents[index]
}

func (l Lot) Size() int {
	return len(l.contents)
}

func (l Lot) Picking() string {
	if len(l.index.unselected) == 0 {
		l.index.unselected = l.index.selected
		l.index.selected = make([]int, 0, len(l.contents))
	}

	var index int

	if len(l.index.unselected) == 1 {
		index = 0
	} else {
		index = rand.Intn(len(l.index.unselected) - 1)
	}

	value := l.index.move(index)

	return l.Contents(value)
}
