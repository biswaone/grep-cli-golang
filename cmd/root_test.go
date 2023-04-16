package cmd

import (
	"os"
	"reflect"
	"testing"
)

func TestSearchStringFile(t *testing.T) {
	tests := []struct {
		name      string
		filename  string
		searchStr string
		want      []string
		err       error
	}{
		{
			name:      "Zero matches in file",
			filename:  "sample.txt",
			searchStr: "test",
			want:      []string{},
			err:       nil,
		},
		{
			name:      "One match in file",
			filename:  "sample.txt",
			searchStr: "found",
			want:      []string{"I found the search_string in the file."},
			err:       nil,
		},
		{
			name:      "Two matches in file",
			filename:  "sample.txt",
			searchStr: "search_string",
			want:      []string{"I found the search_string in the file.", "Another line also contains the search_string"},
			err:       nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file, _ := os.Open(tt.filename)
			defer file.Close()

			got, err := grep(file, tt.searchStr)
			if err != nil {
				t.Errorf("searchString() error = %v, wantErr %v", err, tt.err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchString() got = %v, want %v", got, tt.want)
			}

		})
	}
}
