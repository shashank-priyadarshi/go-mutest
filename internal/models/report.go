package models

import "math/big"

// ReportFileName File name for json report
var ReportFileName string = "report.json"

// Report Structure for mutation report
type Report struct {
	Stats     Stats    `json:"stats"`
	Escaped   []Mutant `json:"escaped"`
	Timeouted []Mutant `json:"timeouted"`
	Killed    []Mutant `json:"killed"`
	Errored   []Mutant `json:"errored"`
}

// Stats There is stats for mutations
type Stats struct {
	TotalMutantsCount    int64      `json:"totalMutantsCount"`
	KilledCount          int64      `json:"killedCount"`
	NotCoveredCount      int64      `json:"notCoveredCount"`
	EscapedCount         int64      `json:"escapedCount"`
	ErrorCount           int64      `json:"errorCount"`
	SkippedCount         int64      `json:"skippedCount"`
	TimeOutCount         int64      `json:"timeOutCount"`
	Msi                  *big.Float `json:"msi"`
	MutationCodeCoverage int64      `json:"mutationCodeCoverage"`
	CoveredCodeMsi       float64    `json:"coveredCodeMsi"`
	DuplicatedCount      int64      `json:"-"`
}

// Mutant report by mutant for one mutation on one file
type Mutant struct {
	Mutator       Mutator `json:"mutator"`
	Diff          string  `json:"diff"`
	ProcessOutput string  `json:"processOutput,omitempty"`
}

// Mutator mutator and changes in file
type Mutator struct {
	MutatorName        string `json:"mutatorName"`
	OriginalSourceCode string `json:"originalSourceCode"`
	MutatedSourceCode  string `json:"mutatedSourceCode"`
	OriginalFilePath   string `json:"originalFilePath"`
	OriginalStartLine  int64  `json:"originalStartLine"`
}

// Calculate calculation for final report
func (report *Report) Calculate() {
	report.Stats.Msi = report.MsiScore()
	report.Stats.TotalMutantsCount = report.TotalCount()
}

// MsiScore msi score calculation
func (report *Report) MsiScore() *big.Float {
	total := report.TotalCount()

	msi := big.NewFloat(0)

	if total == 0 {
		return msi
	}

	killedCount := big.NewFloat(float64(report.Stats.KilledCount))
	errorCount := big.NewFloat(float64(report.Stats.ErrorCount))
	skippedCount := big.NewFloat(float64(report.Stats.SkippedCount))

	subTotalSuccess := big.NewFloat(0).Add(killedCount, errorCount)
	totalSuccess := big.NewFloat(0).Add(subTotalSuccess, skippedCount)

	return msi.Quo(totalSuccess, big.NewFloat(float64(total)))
}

// TotalCount total mutations count
func (report *Report) TotalCount() int64 {
	return report.Stats.KilledCount + report.Stats.EscapedCount + report.Stats.ErrorCount + report.Stats.SkippedCount
}
