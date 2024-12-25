package asciitable

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
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

type gridData struct {
	Name string `asciitable:"param_name"`
	ROM1 string `asciitable:"rom1"`
	ROM2 string `asciitable:"rom2"`
	ROM3 string `asciitable:"rom3"`
}

func TestDecodeGridData(t *testing.T) {
	bytes, err := os.ReadFile("decode_grid.txt")
	require.NoError(t, err)

	table := string(bytes)
	headers, results, err := Unmarshal(table, gridData{})
	require.NoError(t, err)

	require.Equal(t, []string{"param_name", "rom1", "rom2", "rom3"}, headers)

	expected := []gridData{
		{Name: "chip_tech", ROM1: "CB", ROM2: "EB", ROM3: "PB"},
		{Name: "default_v", ROM1: "14000", ROM2: "14000", ROM3: "14000"},
		{Name: "default_clk", ROM1: "490", ROM2: "490", ROM3: "490"},
		{Name: "warmup_temp", ROM1: "41", ROM2: "41", ROM3: "45"},
		{Name: "cooldown_temp", ROM1: "56", ROM2: "59", ROM3: "62"},
	}
	require.Equal(t, expected, results)
}
