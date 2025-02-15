package client

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/otiai10/copy"
	"gotest.tools/v3/assert"
	"kcl-lang.io/kpm/pkg/downloader"
	"kcl-lang.io/kpm/pkg/features"
	"kcl-lang.io/kpm/pkg/utils"
)

func testRunWithModSpecVersion(t *testing.T, kpmcli *KpmClient) {
	pkgPath := getTestDir("test_run_with_modspec_version")
	modbkPath := filepath.Join(pkgPath, "kcl.mod.bk")
	modPath := filepath.Join(pkgPath, "kcl.mod")
	modExpect := filepath.Join(pkgPath, "kcl.mod.expect")
	lockbkPath := filepath.Join(pkgPath, "kcl.mod.lock.bk")
	lockPath := filepath.Join(pkgPath, "kcl.mod.lock")
	lockExpect := filepath.Join(pkgPath, "kcl.mod.lock.expect")
	err := copy.Copy(modbkPath, modPath)
	if err != nil {
		t.Fatal(err)
	}

	err = copy.Copy(lockbkPath, lockPath)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		// remove the copied files
		err := os.RemoveAll(modPath)
		if err != nil {
			t.Fatal(err)
		}
		err = os.RemoveAll(lockPath)
		if err != nil {
			t.Fatal(err)
		}
	}()

	res, err := kpmcli.Run(
		WithRunSource(
			&downloader.Source{
				Local: &downloader.Local{
					Path: pkgPath,
				},
			},
		),
	)

	if err != nil {
		t.Errorf("Failed to run package: %v", err)
	}

	assert.Equal(t, res.GetRawYamlResult(), "res: Hello World!")
	expectedMod, err := os.ReadFile(modExpect)
	if err != nil {
		t.Fatal(err)
	}
	gotMod, err := os.ReadFile(modPath)
	if err != nil {
		t.Fatal(err)
	}

	expectedLock, err := os.ReadFile(lockExpect)
	if err != nil {
		t.Fatal(err)
	}

	gotLock, err := os.ReadFile(lockPath)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, utils.RmNewline(string(expectedMod)), utils.RmNewline(string(gotMod)))
	assert.Equal(t, utils.RmNewline(string(expectedLock)), utils.RmNewline(string(gotLock)))
}

func TestRun(t *testing.T) {
	features.Enable(features.SupportNewStorage)
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunWithOciDownloader", TestFunc: testRunWithOciDownloader}})
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunDefaultRegistryDep", TestFunc: testRunDefaultRegistryDep}})
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunInVendor", TestFunc: testRunInVendor}})
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunRemoteWithArgsInvalid", TestFunc: testRunRemoteWithArgsInvalid}})
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunRemoteWithArgs", TestFunc: testRunRemoteWithArgs}})
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunWithNoSumCheck", TestFunc: testRunWithNoSumCheck}})
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunWithGitPackage", TestFunc: testRunWithGitPackage}})
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunGit", TestFunc: testRunGit}})
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunOciWithSettingsFile", TestFunc: testRunOciWithSettingsFile}})
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunWithModSpecVersion", TestFunc: testRunWithModSpecVersion}})

	features.Disable(features.SupportNewStorage)
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunWithOciDownloader", TestFunc: testRunWithOciDownloader}})
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunDefaultRegistryDep", TestFunc: testRunDefaultRegistryDep}})
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunInVendor", TestFunc: testRunInVendor}})
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunRemoteWithArgsInvalid", TestFunc: testRunRemoteWithArgsInvalid}})
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunRemoteWithArgs", TestFunc: testRunRemoteWithArgs}})
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunWithNoSumCheck", TestFunc: testRunWithNoSumCheck}})
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunWithGitPackage", TestFunc: testRunWithGitPackage}})
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunGit", TestFunc: testRunGit}})
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunOciWithSettingsFile", TestFunc: testRunOciWithSettingsFile}})
	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "TestRunWithModSpecVersion", TestFunc: testRunWithModSpecVersion}})
}
func TestRunWithHyphenEntries(t *testing.T) {
	testFunc := func(t *testing.T, kpmcli *KpmClient) {
		pkgPath := getTestDir("test_run_hyphen_entries")

		res, err := kpmcli.Run(
			WithRunSource(
				&downloader.Source{
					Local: &downloader.Local{
						Path: pkgPath,
					},
				},
			),
		)

		if err != nil {
			t.Fatal(err)
		}

		expect, err := os.ReadFile(filepath.Join(pkgPath, "stdout"))
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, utils.RmNewline(res.GetRawYamlResult()), utils.RmNewline(string(expect)))
	}

	RunTestWithGlobalLockAndKpmCli(t, []TestSuite{{Name: "testRunWithHyphenEntries", TestFunc: testFunc}})
}
