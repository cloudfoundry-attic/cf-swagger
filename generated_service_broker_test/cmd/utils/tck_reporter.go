package utils

import (
	"fmt"
	"strings"

	"github.com/onsi/ginkgo/config"
	"github.com/onsi/ginkgo/types"
)

type tckReporter struct {
	apiTestCountMap   map[string]int
	apiTestPassedMap  map[string]int
	apiTestFailedMap  map[string]int
	failMessages      []string
	totalTestCount    int
	totalTestPassed   int
	failsWithoutAppId bool
	succeedsWithAppId bool
}

func NewTckReporter() *tckReporter {
	return &tckReporter{
		apiTestCountMap:   map[string]int{},
		apiTestPassedMap:  map[string]int{},
		apiTestFailedMap:  map[string]int{},
		failMessages:      []string{},
		totalTestCount:    0,
		totalTestPassed:   0,
		failsWithoutAppId: false,
		succeedsWithAppId: false,
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

	if specSummary.HasFailureState() {
		if strings.EqualFold(specSummary.ComponentTexts[1], "#serviceBind") && strings.EqualFold(specSummary.ComponentTexts[2], "when serviceBind succeed") && strings.EqualFold(specSummary.ComponentTexts[3], "serviceBind with serviceBindData without app_id") {
			tck.failsWithoutAppId = true
			tck.failMessages = append(tck.failMessages, specSummary.ComponentTexts[1]+" "+specSummary.ComponentTexts[2]+" "+specSummary.ComponentTexts[3])
		}
	} else {
		if strings.EqualFold(specSummary.ComponentTexts[1], "#serviceBind") && strings.EqualFold(specSummary.ComponentTexts[2], "when serviceBind succeed") && strings.EqualFold(specSummary.ComponentTexts[3], "serviceBind with serviceBindData with app_id") {
			tck.succeedsWithAppId = true
		}
	}
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

	fmt.Printf("summary %#v\n", summary)
	for _, v := range tck.apiTestCountMap {
		tck.totalTestCount += v
	}

	for _, v := range tck.apiTestPassedMap {
		tck.totalTestPassed += v
	}

	fmt.Println()
	fmt.Printf("Tck Compliance v2.6: %3.2f % \n", tck.TckCompliance()*100)
	if tck.succeedsWithAppId && tck.failsWithoutAppId {
		tck.totalTestCount -= 1
	}
	if (tck.failMessages!=nil){
	fmt.Println("Cause(s) of the failure")
	for _, v := range tck.failMessages {
		fmt.Println(v)
	}
	}
	fmt.Println()
	fmt.Printf("Tck Compliance v2.5: %3.2f % \n", tck.TckCompliance()*100)
}

func (tck *tckReporter) TckCompliance() float64 {
	return (float64)(tck.totalTestPassed) / (float64)(tck.totalTestCount)
}
