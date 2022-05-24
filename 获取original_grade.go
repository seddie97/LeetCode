package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "\"{\\\"gid_result\\\":\\\"{\\\\\\\"1725988080398408\\\\\\\":3,\\\\\\\"1725988680624142\\\\\\\":2.5,\\\\\\\"1725988690227200\\\\\\\":3,\\\\\\\"1725988697232384\\\\\\\":2.5,\\\\\\\"1725988719513607\\\\\\\":3,\\\\\\\"1725988727651470\\\\\\\":3,\\\\\\\"1725988734421005\\\\\\\":2.5,\\\\\\\"1725988743097359\\\\\\\":1.5,\\\\\\\"1725988758700039\\\\\\\":2,\\\\\\\"1725989916251144\\\\\\\":3.5,\\\\\\\"1725989947058183\\\\\\\":4.5,\"original_grade\":\"1.5\",\\\\\\\"7069644570971243016\\\\\\\":4,\\\\\\\"7069650327015784992\\\\\\\":1,\\\\\\\"7069655512576573982\\\\\\\":3}\\\",\\\"grade_source\\\":2,\\\"task_id\\\":7069582853180310049,\\\"timestamp\\\":1646067041000,\"verifiers\":\"fengjiawei\"}\""
	original_grade_index := strings.Index(s, "original_grade")
	fmt.Printf("original_grade_index: %v\n", original_grade_index)
	star := original_grade_index + 17
	end := star
	for end < len(s) && (s[end] >= '0' && s[end] <= '9' || s[end] == '.') {
		end++
	}
	fmt.Printf("original_grade: %s\n", s[star:end])
}
