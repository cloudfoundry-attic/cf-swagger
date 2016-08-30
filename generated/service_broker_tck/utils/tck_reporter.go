package utils

import (
	"fmt"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

type tckReporter struct {
	apiTestCountMap  map[string]int
	apiTestPassedMap map[string]int
	totalTestCount   int
	totalTestPassed  int
}

func NewTckReporter() *tckReporter {
	return &tckReporter{
		apiTestCountMap:  map[string]int{},
		apiTestPassedMap: map[string]int{},
		totalTestCount:   0,
		totalTestPassed:  0,
	}
}

func (tck *tckReporter) SpecSuiteWillBegin(config config.GinkgoConfigType, summary *types.SuiteSummary) {
}

func (tck *tckReporter) BeforeSuiteDidRun(setupSummary *types.SetupSummary) {
}

func (tck *tckReporter) SpecWillRun(specSummary *types.SpecSummary) {
	_, ok := tck.apiTestCountMap[specSummary.ComponentTexts[1]]
	if !ok {
		tck.apiTestCountMap[specSummary.ComponentTexts[1]] = 1
	} else {
		tck.apiTestCountMap[specSummary.ComponentTexts[1]] += 1
	}
}

func (tck *tckReporter) SpecDidComplete(specSummary *types.SpecSummary) {
	_, ok := tck.apiTestPassedMap[specSummary.ComponentTexts[1]]
	if !ok {
		if specSummary.Failure == (types.SpecFailure{}) {
			tck.apiTestPassedMap[specSummary.ComponentTexts[1]] = 1
		}
	} else {
		if specSummary.Failure == (types.SpecFailure{}) {
			tck.apiTestPassedMap[specSummary.ComponentTexts[1]] += 1
		}
	}
}

func (tck *tckReporter) AfterSuiteDidRun(setupSummary *types.SetupSummary) {
}

func (tck *tckReporter) SpecSuiteDidEnd(summary *types.SuiteSummary) {
	for _, v := range tck.apiTestCountMap {
		tck.totalTestCount += v
	}

	for _, v := range tck.apiTestPassedMap {
		tck.totalTestPassed += v
	}

	fmt.Println()
	fmt.Printf("TckCompliance: %f\n", tck.TckCompliance()*100)
}

func (tck *tckReporter) TckCompliance() float64 {
	return (float64)(tck.totalTestPassed) / (float64)(tck.totalTestCount)
}
