package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/JekaMas/go-mutesting/internal/models"
)

func TestMainSimple(t *testing.T) {
	testMain(
		t,
		"../../example",
		[]string{"--debug", "--exec-timeout", "1"},
		returnOk,
		"The mutation score is 0.564516 (35 passed, 27 failed, 8 duplicated, 0 skipped, total is 62)",
	)
}

func TestMainRecursive(t *testing.T) {
	testMain(
		t,
		"../../example",
		[]string{"--debug", "--exec-timeout", "1", "./..."},
		returnOk,
		"The mutation score is 0.590909 (39 passed, 27 failed, 8 duplicated, 0 skipped, total is 66)",
	)
}

func TestMainFromOtherDirectory(t *testing.T) {
	testMain(
		t,
		"../..",
		[]string{"--debug", "--exec-timeout", "1", "github.com/JekaMas/go-mutesting/example"},
		returnOk,
		"The mutation score is 0.564516 (35 passed, 27 failed, 8 duplicated, 0 skipped, total is 62)",
	)
}

func TestMainMatch(t *testing.T) {
	testMain(
		t,
		"../../example",
		[]string{"--debug", "--exec", "../scripts/exec/test-mutated-package.sh", "--exec-timeout", "1", "--match", "baz", "./..."},
		returnOk,
		"The mutation score is 0.500000 (4 passed, 4 failed, 0 duplicated, 0 skipped, total is 8)",
	)
}

func TestMainSkipWithoutTest(t *testing.T) {
	testMain(
		t,
		"../../example",
		[]string{"--debug", "--exec-timeout", "1", "--config", "../testdata/configs/configSkipWithoutTest.yml.test"},
		returnOk,
		"The mutation score is 0.583333 (35 passed, 25 failed, 8 duplicated, 0 skipped, total is 60)",
	)
}

func TestMainJSONReport(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "go-mutesting-main-test-")
	require.NoError(t, err)

	reportFileName := "reportTestMainJSONReport.json"
	jsonFile := tmpDir + "/" + reportFileName
	if _, err := os.Stat(jsonFile); err == nil {
		err = os.Remove(jsonFile)
		require.NoError(t, err)
	}

	models.ReportFileName = jsonFile

	testMain(
		t,
		"../../example",
		[]string{"--debug", "--exec-timeout", "1", "--config", "../testdata/configs/configForJson.yml.test"},
		returnOk,
		"The mutation score is 0.583333 (35 passed, 25 failed, 8 duplicated, 0 skipped, total is 60)",
	)

	info, err := os.Stat(jsonFile)
	require.NoError(t, err)
	require.NotNil(t, info)

	defer func() {
		err = os.Remove(jsonFile)
		if err != nil {
			fmt.Println("Error while deleting temp file")
		}
	}()

	jsonData, err := ioutil.ReadFile(jsonFile)
	require.NoError(t, err)

	var mutationReport models.Report
	err = json.Unmarshal(jsonData, &mutationReport)
	require.NoError(t, err)

	expectedStats := models.Stats{
		TotalMutantsCount:    60,
		KilledCount:          35,
		NotCoveredCount:      0,
		EscapedCount:         25,
		ErrorCount:           0,
		SkippedCount:         0,
		TimeOutCount:         0,
		Msi:                  big.NewFloat(0.5833333333333334),
		MutationCodeCoverage: 0,
		CoveredCodeMsi:       0,
		DuplicatedCount:      0,
	}

	require.Equal(t, expectedStats, mutationReport.Stats)
	require.Equal(t, 25, len(mutationReport.Escaped))
	require.Nil(t, mutationReport.Timeouted)
	require.Equal(t, 35, len(mutationReport.Killed))
	require.Nil(t, mutationReport.Errored)

	for i := 0; i < len(mutationReport.Escaped); i++ {
		require.Contains(t, mutationReport.Escaped[i].ProcessOutput, "FAIL")
	}
	for i := 0; i < len(mutationReport.Killed); i++ {
		require.Contains(t, mutationReport.Killed[i].ProcessOutput, "PASS")
	}
}

func testMain(t *testing.T, root string, exec []string, expectedExitCode int, contains string) {
	t.Helper()

	saveStderr := os.Stderr
	saveStdout := os.Stdout
	saveCwd, err := os.Getwd()
	require.Nil(t, err)

	r, w, err := os.Pipe()
	require.Nil(t, err)

	os.Stderr = w
	os.Stdout = w
	require.Nil(t, os.Chdir(root))

	bufChannel := make(chan string)

	go func() {
		buf := new(bytes.Buffer)
		_, err = io.Copy(buf, r)
		require.Nil(t, err)
		require.Nil(t, r.Close())

		bufChannel <- buf.String()
	}()

	exitCode := mainCmd(exec)

	require.Nil(t, w.Close())

	os.Stderr = saveStderr
	os.Stdout = saveStdout
	require.Nil(t, os.Chdir(saveCwd))

	out := <-bufChannel

	require.Equal(t, expectedExitCode, exitCode)
	require.Contains(t, out, contains)
}
