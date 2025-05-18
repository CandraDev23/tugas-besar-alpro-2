package Features

import (
	"fmt"
	"strings"
)

type CareerMatch struct {
	Career     Career
	MatchScore float64 // persentase kecocokan
}

func CalculateMatch(profile *Profile, career Career) float64 {
	if profile == nil {
		return 0
	}

	matchCount := 0
	totalCriteria := len(career.Interests) + len(career.Skills)

	for _, pInterest := range profile.Interests {
		for _, cInterest := range career.Interests {
			if strings.EqualFold(pInterest, cInterest) {
				matchCount++
				break
			}
		}
	}

	for _, pSkill := range profile.Skills {
		for _, cSkill := range career.Skills {
			if strings.EqualFold(pSkill, cSkill) {
				matchCount++
				break
			}
		}
	}

	if totalCriteria == 0 {
		return 0
	}
	return float64(matchCount) / float64(totalCriteria) * 100
}

func GenerateCareerRecommendations() []CareerMatch {
	var results []CareerMatch
	for _, c := range careers {
		match := CalculateMatch(profile, c)
		if match > 0 {
			results = append(results, CareerMatch{Career: c, MatchScore: match})
		}
	}
	return results
}

func DisplayRecommendations(recommendations []CareerMatch) {
	for _, r := range recommendations {
		fmt.Printf("Karier: %s\nDeskripsi: %s\nKecocokan: %.2f%%\nGaji: %d\n\n", r.Career.Name, r.Career.Desc, r.MatchScore, r.Career.Salary)
	}
}

func SelectionSort(careers []CareerMatch, bySalary bool, ascending bool) {
	n := len(careers)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			cond := false
			if bySalary {
				cond = (ascending && careers[j].Career.Salary < careers[idx].Career.Salary) ||
					(!ascending && careers[j].Career.Salary > careers[idx].Career.Salary)
			} else {
				cond = (ascending && careers[j].MatchScore < careers[idx].MatchScore) ||
					(!ascending && careers[j].MatchScore > careers[idx].MatchScore)
			}
			if cond {
				idx = j
			}
		}
		careers[i], careers[idx] = careers[idx], careers[i]
	}
}

func InsertionSort(careers []CareerMatch, bySalary bool, ascending bool) {
	for i := 1; i < len(careers); i++ {
		key := careers[i]
		j := i - 1
		for j >= 0 {
			cond := false
			if bySalary {
				cond = (ascending && careers[j].Career.Salary > key.Career.Salary) ||
					(!ascending && careers[j].Career.Salary < key.Career.Salary)
			} else {
				cond = (ascending && careers[j].MatchScore > key.MatchScore) ||
					(!ascending && careers[j].MatchScore < key.MatchScore)
			}
			if !cond {
				break
			}
			careers[j+1] = careers[j]
			j--
		}
		careers[j+1] = key
	}
}