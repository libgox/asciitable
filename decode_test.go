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
