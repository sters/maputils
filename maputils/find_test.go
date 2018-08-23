package maputils_test

import (
	"encoding/json"
	"testing"

	"github.com/sters/maputils/maputils"
)

func TestRecursiveFindKey(t *testing.T) {
	testcases := map[string]struct {
		target map[string]interface{}
		key    string
		expect []interface{}
	}{
		"flat map": {
			target: map[string]interface{}{
				"a": "hoge",
				"b": "huga",
				"c": "moge",
			},
			key: "b",
			expect: []interface{}{
				"huga",
			},
		},
		"2 level": {
			target: map[string]interface{}{
				"a": "hoge",
				"b": "huga",
				"c": map[string]interface{}{
					"a": "moge",
					"b": "muga",
				},
			},
			key: "b",
			expect: []interface{}{
				"huga",
				"muga",
			},
		},
		"more level": {
			target: map[string]interface{}{
				"a": "hoge",
				"b": "huga",
				"c": map[string]interface{}{
					"a": map[string]interface{}{
						"a": map[string]interface{}{
							"a": map[string]interface{}{
								"a": map[string]interface{}{
									"a": map[string]interface{}{
										"a": map[string]interface{}{
											"a": map[string]interface{}{
												"a": map[string]interface{}{
													"b": map[string]interface{}{
														"b": "moge",
													},
												},
											},
										},
									},
								},
							},
						},
					},
					"b": t,
					"c": map[string]interface{}{
						"b": []int{1, 1, 2, 3, 5},
					},
				},
			},
			key: "b",
			expect: []interface{}{
				"huga",
				t,
				[]int{1, 1, 2, 3, 5},
				map[string]interface{}{
					"b": "moge",
				},
				"moge",
			},
		},
	}

	for testname, testcase := range testcases {
		t.Run(testname, func(t *testing.T) {
			result := maputils.RecursiveFindKey(testcase.key, testcase.target)

			resultJson, err := json.Marshal(result)
			if err != nil {
				t.Fatalf("failed result to json: %s", err)
			}

			expectJson, err := json.Marshal(testcase.expect)
			if err != nil {
				t.Fatalf("failed expect to json: %s", err)
			}

			if string(resultJson) != string(expectJson) {
				t.Errorf("expect %s, got %s", string(expectJson), string(resultJson))
			}
		})
	}
}
