package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func ExperimentalMin[T constraints.Ordered](first T, last T) T {
	if first < last {
		return first
	} else {
		return last
	}
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, 100.0, ExperimentalMin(100.0, 200.0))
}

func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"Name": "Arif",
	}

	last := map[string]string{
		"Name": "Arif",
	}

	assert.True(t, maps.Equal(first, last))
}

func TestExperimentalSlice(t *testing.T) {
	first := []string{"Arif"}
	last := []string{"Arif"}

	assert.True(t, slices.Equal(first, last))
}
