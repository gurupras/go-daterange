package daterange

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSameDate(t *testing.T) {
	require := require.New(t)

	dr := NewDateRange("19700101", "19700101")
	require.NotNil(dr)

	require.Equal(1, len(dr.AsList()))
}

func TestNegative(t *testing.T) {
	assert := assert.New(t)

	dr := NewDateRange("19700102", "19700101")
	assert.Nil(dr)
}

func TestSmallDateRange(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	var dr *DateRange

	dr = NewDateRange("19700101", "19700201")
	require.NotNil(dr)
	assert.Equal(31, len(dr.AsList()))

	dr = NewDateRange("19700101", "19700301")
	require.NotNil(dr)
	assert.Equal(59, len(dr.AsList()))

	dr = NewDateRange("19700101", "19700401")
	require.NotNil(dr)
	assert.Equal(90, len(dr.AsList()))

	dr = NewDateRange("19700101", "19700501")
	require.NotNil(dr)
	assert.Equal(120, len(dr.AsList()))

	dr = NewDateRange("19700101", "19700601")
	require.NotNil(dr)
	assert.Equal(151, len(dr.AsList()))
}

func TestLargeDateRange(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	dr := NewDateRange("19700101", "20170101")
	require.NotNil(dr)
	assert.Equal(17167, len(dr.AsList()))
}
