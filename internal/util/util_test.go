package util

import (
	"bufio"
	"errors"
	"strings"
	"testing"
	"time"
)

func newScanner(s string) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(s))
}

func timeParse(s string) time.Time {
	pasedTime, _ := time.Parse("2006-01-02", s)
	return pasedTime
}

// func errDateImitation(s string) error {
// 	_, err := time.Parse("2006-01-02", s)
// 	return err
// }

func equalMaps(m1, m2 map[string][]UserCard) bool {
	if len(m1) != len(m2) {
		return false
	}

	for key, cards1 := range m1 {
		cards2, exists := m2[key]
		if !exists {
			return false
		}
		if len(cards1) != len(cards2) {
			return false
		}
		for i := range cards1 {
			if cards1[i].Specialization != cards2[i].Specialization {
				return false
			}
			if !cards1[i].Date.Equal(cards2[i].Date) {
				return false
			}
		}
	}

	return true
}

func TestSave(t *testing.T) {
	tests := []struct {
		name           string
		inc            *bufio.Scanner
		want           map[string][]UserCard
		wantParseError bool
		err            error
	}{
		{
			name: "TestSave - test #1",
			inc:  newScanner("Иванов Иван Иванович\nТерапевт\n2024-01-01\n"),
			want: map[string][]UserCard{
				"Иванов Иван Иванович": {
					{
						Specialization: "терапевт",
						Date:           timeParse("2024-01-01"),
					},
				},
			},
			wantParseError: false,
			err:            nil,
		},
		{
			name: "TestSave - test #2",
			inc:  newScanner("   ивАнОв иВан   иванОвич\n  тераПевт  \n  2024-01-01   \n"),
			want: map[string][]UserCard{
				"Иванов Иван Иванович": {
					{
						Specialization: "терапевт",
						Date:           timeParse("2024-01-01"),
					},
				},
			},
			wantParseError: false,
			err:            nil,
		},
		{
			name:           "TestSave - test #3",
			inc:            newScanner("Иванов Иван\nТерапевт\n2024-01-01\n"),
			want:           map[string][]UserCard{},
			wantParseError: false,
			err:            ErrorInsData,
		},
		{
			name:           "TestSave - test #4",
			inc:            newScanner("\nТерапевт\n2024-01-01\n"),
			want:           map[string][]UserCard{},
			wantParseError: false,
			err:            ErrEmptyLine,
		},
		{
			name:           "TestSave - test #5",
			inc:            newScanner("Иванов Иван Иванович\n\n2024-01-01\n"),
			want:           map[string][]UserCard{},
			wantParseError: false,
			err:            ErrEmptyLine,
		},
		{
			name:           "TestSave - test #6",
			inc:            newScanner("Иванов Иван Иванович\nТерапевт\n\n"),
			want:           map[string][]UserCard{},
			wantParseError: true,
			err:            nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base = map[string][]UserCard{}
			got, err := Save(tt.inc)

			if tt.wantParseError {
				var parseErr *time.ParseError
				if !errors.As(err, &parseErr) {
					t.Fatalf("expected *time.ParseError, got %T: %v", err, err)
				}
				return
			}

			if !errors.Is(err, tt.err) && err != tt.err {
				t.Fatalf("expected error %v, got %v", tt.err, err)
			}

			if err == nil && !equalMaps(got, tt.want) {
				t.Errorf("Save() = %v, want %v", got, tt.want)
			}
		})
	}
}
