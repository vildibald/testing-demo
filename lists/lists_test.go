package lists

import (
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSorted(t *testing.T) {
	// Valid lists
	if !IsSorted([]int{1, 2, 3}) {
		t.Errorf("Expected true but got false")
	}
	if !IsSorted([]int{1, 1, 1}) {
		t.Errorf("Expected true but got false")
	}
	if !IsSorted([]int{}) {
		t.Errorf("Expected true but got false")
	}
	if !IsSorted(nil) {
		t.Errorf("Expected true but got false")
	}

	// Invalid lists
	if IsSorted([]int{1, 3, 2}) {
		t.Errorf("Expected false but got true")
	}
}

func TestIsSortedUsingAsserts(t *testing.T) {
	// Valid lists
	assert.True(t, IsSorted([]int{1, 2, 3}))
	assert.True(t, IsSorted([]int{1, 1, 1}))
	assert.True(t, IsSorted([]int{}))
	assert.True(t, IsSorted(nil))

	// Invalid lists
	assert.False(t, IsSorted([]int{1, 3, 2}))
}

func TestIsSortedUsingGomega(t *testing.T) {
	g := NewGomegaWithT(t)
	// Valid lists
	g.Expect(IsSorted([]int{1, 2, 3})).To(BeTrue())
	g.Expect(IsSorted([]int{1, 1, 1})).To(BeTrue())
	g.Expect(IsSorted([]int{})).To(BeTrue())
	g.Expect(IsSorted(nil)).To(BeTrue())

	// Invalid lists
	g.Expect(IsSorted([]int{1, 3, 2})).To(BeFalse())
}

func TestIsSortedUsingSubtests(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Valid lists", func(t *testing.T) {
		g.Expect(IsSorted([]int{1, 2, 3})).To(BeTrue())
		g.Expect(IsSorted([]int{1, 1, 1})).To(BeTrue())
		g.Expect(IsSorted([]int{})).To(BeTrue())
		g.Expect(IsSorted(nil)).To(BeTrue())
	})

	t.Run("Invalid lists", func(t *testing.T) {
		g.Expect(IsSorted([]int{1, 3, 2})).To(BeFalse())
	})
}

func TestIsSortedUsingAssertsAndTable(t *testing.T) {
	g := NewGomegaWithT(t)

	tests := []struct {
		name     string
		list     []int
		expected bool
	}{
		{
			name:     "Increasing list",
			list:     []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "Constant list",
			list:     []int{1, 1, 1},
			expected: true,
		},
		{
			name:     "Empty list",
			list:     []int{},
			expected: true,
		},
		{
			name:     "Nil list",
			list:     nil,
			expected: true,
		},
		{
			name:     "Unsorted list",
			list:     []int{1, 3, 2},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			g.Expect(IsSorted(test.list)).To(Equal(test.expected))
		})
	}
}
