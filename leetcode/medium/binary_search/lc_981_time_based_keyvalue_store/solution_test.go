package lc_981_time_based_keyvalue_store

import (
	"testing"
)

func TestTimeMap(t *testing.T) {
	t.Run("TestSetAndGet", testSetAndGet)
	t.Run("TestGetWithNonexistentKey", testGetWithNonexistentKey)
	t.Run("TestGetWithEmptyValue", testGetWithEmptyValue)
	t.Run("TestGetWithNegativeTimestamp", testGetWithNegativeTimestamp)
	t.Run("TestGetWithOutOfRangeTimestamp", testGetWithOutOfRangeTimestamp)
}

func testSetAndGet(t *testing.T) {
	tm := Constructor()

	type testCase struct {
		key       string
		value     string
		timestamp int
	}

	tests := []testCase{
		{"key1", "value1", 1},
		{"key1", "value2", 2},
		{"key1", "value3", 3},
		{"key2", "valueA", 5},
		{"key2", "valueB", 6},
		{"empty", "", 10},
		{"empty", "", 11},
		{"empty", "nonempty", 12},
	}

	for _, tc := range tests {
		t.Run("", func(t *testing.T) {
			tm.Set(tc.key, tc.value, tc.timestamp)
			result := tm.Get(tc.key, tc.timestamp)
			if result != tc.value {
				t.Errorf("Get(%q, %d) returned %q, expected %q", tc.key, tc.timestamp, result, tc.value)
			}
		})
	}
}

func testGetWithNonexistentKey(t *testing.T) {
	tm := Constructor()
	result := tm.Get("nonexistent", 1)
	if result != "" {
		t.Errorf("Get with nonexistent key returned %q, expected \"\"", result)
	}
}

func testGetWithEmptyValue(t *testing.T) {
	tm := Constructor()
	tm.Set("empty", "", 10)
	result := tm.Get("empty", 11)
	if result != "" {
		t.Errorf("Get with empty value returned %q, expected \"\"", result)
	}

	tm.Set("empty", "nonempty", 12)
	result = tm.Get("empty", 13)
	if result != "nonempty" {
		t.Errorf("Get with empty value returned %q, expected \"nonempty\"", result)
	}
}

func testGetWithNegativeTimestamp(t *testing.T) {
	tm := Constructor()
	result := tm.Get("key1", -1)
	if result != "" {
		t.Errorf("Get with negative timestamp returned %q, expected \"\"", result)
	}
}

func testGetWithOutOfRangeTimestamp(t *testing.T) {
	tm := Constructor()
	tm.Set("key1", "value1", 1)
	tm.Set("key1", "value2", 2)
	tm.Set("key1", "value3", 3)
	result := tm.Get("key1", 100)
	if result != "value3" {
		t.Errorf("Get with out-of-range timestamp returned %q, expected \"value3\"", result)
	}
}
