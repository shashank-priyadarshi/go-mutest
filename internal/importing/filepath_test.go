package importing

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/shashank-priyadarshi/go-mutest/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestFilesOfArgs(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	p := filepath.Clean(filepath.Join(wd, "../..")) + "/"

	cases := []struct {
		name   string
		args   []string
		expect []string
	}{
		{
			"empty",
			[]string{},
			[]string{"filepath.go", "import.go"},
		},
		{
			"files",
			[]string{"./filepathfixtures/first.go"},
			[]string{"./filepathfixtures/first.go"},
		},
		{
			"directory",
			[]string{"./filepathfixtures"},
			[]string{"filepathfixtures/first.go", "filepathfixtures/second.go", "filepathfixtures/third.go"},
		},
		{
			"sub-directory",
			[]string{"../importing/filepathfixtures"},
			[]string{
				"../importing/filepathfixtures/first.go",
				"../importing/filepathfixtures/second.go",
				"../importing/filepathfixtures/third.go",
			},
		},
		{
			"package",
			[]string{"github.com/shashank-priyadarshi/go-mutest/internal/importing/filepathfixtures"},
			[]string{
				p + "internal/importing/filepathfixtures/first.go",
				p + "internal/importing/filepathfixtures/second.go",
				p + "internal/importing/filepathfixtures/third.go",
			},
		},
		{
			"packages",
			[]string{"github.com/shashank-priyadarshi/go-mutest/internal/importing/filepathfixtures/..."},
			[]string{
				p + "internal/importing/filepathfixtures/first.go",
				p + "internal/importing/filepathfixtures/second.go",
				p + "internal/importing/filepathfixtures/third.go",
				p + "internal/importing/filepathfixtures/secondfixturespackage/fourth.go",
			},
		},
	}

	for _, test := range cases {
		test := test

		t.Run(test.name, func(t *testing.T) {
			var opts = &models.Options{}
			got := FilesOfArgs(test.args, opts)

			require.Equal(t, test.expect, got, fmt.Sprintf("With args: %#v", test.args))
		})
	}
}

func TestPackagesWithFilesOfArgs(t *testing.T) {
	p := os.Getenv("GOPATH") + "/src/"

	testCases := []struct {
		name   string
		args   []string
		expect []Package
	}{
		{
			"empty",
			[]string{},
			[]Package{{Name: ".", Files: []string{"filepath.go", "import.go"}}},
		},
		{
			"files",
			[]string{"./filepathfixtures/first.go"},
			[]Package{{Name: "filepathfixtures", Files: []string{"./filepathfixtures/first.go"}}},
		},
		{
			"directories",
			[]string{"./filepathfixtures"},
			[]Package{{Name: "filepathfixtures", Files: []string{
				"filepathfixtures/first.go",
				"filepathfixtures/second.go",
				"filepathfixtures/third.go",
			}}},
		},
		{
			"relative directories",
			[]string{"../importing/filepathfixtures"},
			[]Package{{Name: "../importing/filepathfixtures", Files: []string{
				"../importing/filepathfixtures/first.go",
				"../importing/filepathfixtures/second.go",
				"../importing/filepathfixtures/third.go",
			}}},
		},
		{
			"package",
			[]string{"internal/importing/filepathfixtures"},
			[]Package{{
				Name: p + "internal/importing/filepathfixtures",
				Files: []string{
					p + "internal/importing/filepathfixtures/first.go",
					p + "internal/importing/filepathfixtures/second.go",
					p + "internal/importing/filepathfixtures/third.go",
				},
			}},
		},
		{
			"packages",
			[]string{"internal/importing/filepathfixtures/..."},
			[]Package{
				{
					Name: p + "internal/importing/filepathfixtures",
					Files: []string{
						p + "internal/importing/filepathfixtures/first.go",
						p + "internal/importing/filepathfixtures/second.go",
						p + "internal/importing/filepathfixtures/third.go",
					},
				},
				{
					Name: p + "internal/importing/filepathfixtures/secondfixturespackage",
					Files: []string{
						p + "internal/importing/filepathfixtures/secondfixturespackage/fourth.go",
					},
				},
			},
		},
	}

	for _, test := range testCases {
		test := test

		t.Run(test.name, func(t *testing.T) {
			var opts = &models.Options{}
			got := PackagesWithFilesOfArgs(test.args, opts)

			assert.Equal(t, test.expect, got, fmt.Sprintf("With args: %#v", test.args))
		})
	}
}

func TestFilesWithSkipWithoutTests(t *testing.T) {
	cases := []struct {
		name   string
		args   []string
		expect []string
	}{
		{
			"single file",
			[]string{"./filepathfixtures/first.go"},
			[]string(nil),
		},
		{
			"duplicate files",
			[]string{"./filepathfixtures/second.go"},
			[]string{"./filepathfixtures/second.go"},
		},
		{
			"directory",
			[]string{"./filepathfixtures"},
			[]string{"filepathfixtures/second.go", "filepathfixtures/third.go"},
		},
		{
			"package",
			[]string{"internal/importing/filepathfixtures/..."},
			[]string{
				"internal/importing/filepathfixtures/second.go",
				"internal/importing/filepathfixtures/third.go",
			},
		},
	}

	for _, testCase := range cases {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			var opts = &models.Options{}
			opts.Config.SkipFileWithoutTest = true
			got := FilesOfArgs(testCase.args, opts)

			assert.Equal(t, testCase.expect, got, fmt.Sprintf("With args: %#v", testCase.args))
		})
	}
}

func TestFilesWithSkipWithBuildTagsTests(t *testing.T) {
	p := os.Getenv("GOPATH") + "/src/"

	testCases := []struct {
		name   string
		args   []string
		expect []string
	}{
		{
			"file doesn't exist",
			[]string{"./filepathfixtures/first.go"},
			[]string(nil),
		},
		{
			"file doesn't exist - 2",
			[]string{"./filepathfixtures/third.go"},
			[]string(nil),
		},
		{
			"file",
			[]string{"./filepathfixtures/second.go"},
			[]string{"./filepathfixtures/second.go"},
		},
		{
			"directories",
			[]string{"./filepathfixtures"},
			[]string{"filepathfixtures/second.go"},
		},
		{
			"packages",
			[]string{"internal/importing/filepathfixtures/..."},
			[]string{
				p + "internal/importing/filepathfixtures/second.go",
			},
		},
	}

	for _, test := range testCases {
		test := test

		t.Run(test.name, func(t *testing.T) {

			var opts = &models.Options{}
			opts.Config.SkipFileWithBuildTag = true
			got := FilesOfArgs(test.args, opts)

			assert.Equal(t, test.expect, got, fmt.Sprintf("With args: %#v", test.args))
		})
	}
}

func TestFilesWithExcludedDirs(t *testing.T) {
	p := os.Getenv("GOPATH") + "/src/"

	testCases := []struct {
		name   string
		args   []string
		expect []string
		config []string
	}{
		{
			"files doesn't exist",
			[]string{"./filepathfixtures/first.go"},
			[]string{"./filepathfixtures/first.go"},
			[]string(nil),
		},
		{
			"relative path to file",
			[]string{"./filepathfixtures/second.go"},
			[]string{"./filepathfixtures/second.go"},
			[]string{"filepathfixtures"},
		},
		{
			"incorrect path to file",
			[]string{"filepathfixtures/second.go"},
			[]string(nil),
			[]string{"filepathfixtures"},
		},
		{
			"incorrect path to file - 2",
			[]string{"./filepathfixtures/second.go"},
			[]string(nil),
			[]string{"./filepathfixtures"},
		},
		{
			"directories",
			[]string{"./filepathfixtures/..."},
			[]string{
				"filepathfixtures/first.go",
				"filepathfixtures/second.go",
				"filepathfixtures/third.go",
			},
			[]string{"filepathfixtures/secondfixturespackage"},
		},
		{
			"directories - 1",
			[]string{"./filepathfixtures/..."},
			[]string(nil),
			[]string{"filepathfixtures"},
		},
		{
			"directories - 2",
			[]string{"./filepathfixtures"},
			[]string(nil),
			[]string{"filepathfixtures"},
		},
		{
			"directories - 3",
			[]string{"./filepathfixtures"},
			[]string{
				"filepathfixtures/first.go",
				"filepathfixtures/second.go",
				"filepathfixtures/third.go",
			},
			[]string(nil),
		},

		{
			"packages",
			[]string{"internal/importing/filepathfixtures/..."},
			[]string{
				p + "internal/importing/filepathfixtures/first.go",
				p + "internal/importing/filepathfixtures/second.go",
				p + "internal/importing/filepathfixtures/third.go",
				p + "internal/importing/filepathfixtures/secondfixturespackage/fourth.go",
			},
			[]string{"filepathfixtures"},
		},
		{
			"packages - 2",
			[]string{"internal/importing/filepathfixtures/..."},
			[]string{
				p + "internal/importing/filepathfixtures/first.go",
				p + "internal/importing/filepathfixtures/second.go",
				p + "internal/importing/filepathfixtures/third.go",
			},
			[]string{p + "internal/importing/filepathfixtures/secondfixturespackage/"},
		},
	}

	for _, test := range testCases {
		test := test

		t.Run(test.name, func(t *testing.T) {

			var opts = &models.Options{}
			opts.Config.ExcludeDirs = test.config

			got := FilesOfArgs(test.args, opts)

			assert.Equal(t, test.expect, got, fmt.Sprintf("With args: %#v", test.args))
		})
	}
}
