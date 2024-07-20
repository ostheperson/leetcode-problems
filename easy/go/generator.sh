#!/bin/bash

# Check for function name argument
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <function_name>"
    exit 1
fi

FUNC_NAME=$1
PROBLEMS_DIR="./problems"
TESTS_DIR="./tests"

# Convert function name to snake_case
snake_case_name=$(echo "$FUNC_NAME" | sed 's/\([a-z0-9]\)\([A-Z]\)/\1_\L\2/g' | sed 's/[A-Z]/\L&/g')

# Create directories if they do not exist
mkdir -p $PROBLEMS_DIR
mkdir -p $TESTS_DIR

# Create the problem file with the function template
cat <<EOF > $PROBLEMS_DIR/${snake_case_name}.go
package problems

// ${FUNC_NAME} takes parameters and performs its operation.
// Add your implementation here.
func ${FUNC_NAME}() []int {
    // TODO: Implement ${FUNC_NAME}
    return nil
}
EOF

# Create the test file with the provided template
cat <<EOF > $TESTS_DIR/${snake_case_name}_test.go
package tests

import (
	"reflect"
	"testing"

	"leetcode-problems/problems" // Update this path to match your module structure
)

func Test${FUNC_NAME}(t *testing.T) {
	tests := []struct {
		expected []int
	}{
		// Add your test cases here
		{
			expected: []int{}, // Example expected output
		},
	}

	for _, test := range tests {
		result := problems.${FUNC_NAME}()
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For expected %v, but got %v", test.expected, result)
		}
	}
}
EOF

echo "Generated files for function '${FUNC_NAME}' with filenames in snake_case in the '${PROBLEMS_DIR}' and '${TESTS_DIR}' directories."

