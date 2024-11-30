package asciitable

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

type nameAgeScore struct {
	Name  string `asciitable:"Name"`
	Age   int    `asciitable:"Age"`
	Score int    `asciitable:"Score"`
}

func TestDecodeNameAgeScore(t *testing.T) {
	bytes, err := os.ReadFile("decode_name_age_score.txt")
	require.NoError(t, err)

	table := string(bytes)
	headers, results, err := Unmarshal(table, nameAgeScore{})
	require.NoError(t, err)

	require.Equal(t, []string{"Name", "Age", "Score"}, headers)

	expected := []nameAgeScore{
		{Name: "Alice", Age: 24, Score: 89},
		{Name: "Bob", Age: 19, Score: 72},
		{Name: "Charlie", Age: 22, Score: 95},
	}
	require.Equal(t, expected, results)
}

type normalData struct {
	Bool    bool    `asciitable:"Bool"`
	Float32 float32 `asciitable:"Float32"`
	Float64 float64 `asciitable:"Float64"`
}

func TestDecodeNormalData(t *testing.T) {
	bytes, err := os.ReadFile("decode_normal.txt")
	require.NoError(t, err)

	table := string(bytes)
	headers, results, err := Unmarshal(table, normalData{})
	require.NoError(t, err)

	require.Equal(t, []string{"Bool", "Float32", "Float64"}, headers)

	expected := []normalData{
		{Bool: true, Float32: 24.5, Float64: 89.123},
		{Bool: true, Float32: 19.3, Float64: 72.001},
		{Bool: false, Float32: 19.3, Float64: 72.001},
		{Bool: false, Float32: 22.7, Float64: 95.5},
	}
	require.Equal(t, expected, results)
}
