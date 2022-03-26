package adder_test

import (
	"bufio"
	"github.com/google/go-cmp/cmp"
	"learning-go/ch13/adder"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	f, err := os.Open("testdata/input.txt")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Cleanup(func() {
		f.Close()
	})

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		inputs := strings.Split(scanner.Text(), " ")

		t.Run("", func(t *testing.T) {
			input1, err := strconv.Atoi(inputs[0])
			if err != nil {
				t.Fatal(err.Error())
			}

			input2, err := strconv.Atoi(inputs[1])
			if err != nil {
				t.Fatal(err.Error())
			}

			expected, err := strconv.Atoi(inputs[2])
			if err != nil {
				t.Fatal(err.Error())
			}

			actual := adder.Add(input1, input2)

			if diff := cmp.Diff(expected, actual); diff != "" {
				t.Error(diff)
			}
		})
	}
}
