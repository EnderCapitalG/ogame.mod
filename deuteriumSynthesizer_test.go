package ogame

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeuteriumSynthesizerPrice(t *testing.T) {
	ds := newDeuteriumSynthesizer()

	price := ds.GetPrice(1)
	assert.Equal(t, Resources{Metal: 225, Crystal: 75}, price)

	price = ds.GetPrice(2)
	assert.Equal(t, Resources{Metal: 337, Crystal: 112}, price)

	price = ds.GetPrice(3)
	assert.Equal(t, Resources{Metal: 506, Crystal: 168}, price)

	price = ds.GetPrice(4)
	assert.Equal(t, Resources{Metal: 759, Crystal: 253}, price)

	price = ds.GetPrice(5)
	assert.Equal(t, Resources{Metal: 1139, Crystal: 379}, price)

	price = ds.GetPrice(11)
	assert.Equal(t, Resources{Metal: 12974, Crystal: 4324}, price)
}

func TestDeuteriumSynthesizerConstructionTime(t *testing.T) {
	ds := newDeuteriumSynthesizer()
	assert.Equal(t, 1845, ds.ConstructionTime(9, 6, Facilities{}))
}
