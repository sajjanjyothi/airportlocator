package flightpath

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetSourceDestinationAirports(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "valid-case-1",
			input: `[["SFO", "EWR"]]`,
			want:  []string{"SFO", "EWR"},
		},
		{
			name:  "valid-case-2",
			input: `[["ATL", "EWR"], ["SFO", "ATL"]]`,
			want:  []string{"SFO", "EWR"},
		},
		{
			name:  "valid-case-3",
			input: `[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]`,
			want:  []string{"SFO", "EWR"},
		},
		{
			name:  "valid-case-4-loop-test",
			input: `[["IND", "EWR"], ["EWR", "JP"], ["IND", "IND"], ["JP", "GER"]]`,
			want:  []string{"IND", "GER"},
		},
	}

	flightDetector := NewAirportDetector()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var flightPath [][]string
			err := json.Unmarshal([]byte(tt.input), &flightPath)
			assert.NoError(t, err)

			paths, err := flightDetector.GetSourceDestinationAirports(flightPath)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, paths, tt.name)
		})
	}
}
