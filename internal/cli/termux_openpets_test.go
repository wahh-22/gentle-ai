package cli

import (
	"testing"

	"github.com/gentleman-programming/gentle-ai/v2/internal/model"
	"github.com/gentleman-programming/gentle-ai/v2/internal/system"
)

func TestWithTermuxOpenPetsEnvForInstallEnablesAndRestoresWhenEligible(t *testing.T) {
	restoreSetenv := osSetenv
	restoreUnsetenv := osUnsetenv
	restoreLookupEnv := osLookupEnv
	t.Cleanup(func() {
		osSetenv = restoreSetenv
		osUnsetenv = restoreUnsetenv
		osLookupEnv = restoreLookupEnv
	})

	env := map[string]string{"GENTLE_AI_TERMUX_OPENPETS": "custom"}
	osLookupEnv = func(key string) (string, bool) {
		v, ok := env[key]
		return v, ok
	}
	osSetenv = func(key, value string) error {
		env[key] = value
		return nil
	}
	osUnsetenv = func(key string) error {
		delete(env, key)
		return nil
	}

	restore := withTermuxOpenPetsEnvForInstall(
		system.PlatformProfile{OS: "android", LinuxDistro: system.LinuxDistroTermux, PackageManager: "pkg"},
		[]model.AgentID{model.AgentOpenCode},
	)
	if got := env[termuxOpenPetsEnvKey]; got != "1" {
		t.Fatalf("env during install = %q, want 1", got)
	}

	restore()
	if got := env[termuxOpenPetsEnvKey]; got != "custom" {
		t.Fatalf("env after restore = %q, want custom", got)
	}
}

func TestWithTermuxOpenPetsEnvForInstallNoopWhenNotEligible(t *testing.T) {
	restoreSetenv := osSetenv
	restoreLookupEnv := osLookupEnv
	t.Cleanup(func() {
		osSetenv = restoreSetenv
		osLookupEnv = restoreLookupEnv
	})

	setenvCalls := 0
	osSetenv = func(string, string) error {
		setenvCalls++
		return nil
	}
	osLookupEnv = func(string) (string, bool) { return "", false }

	restore := withTermuxOpenPetsEnvForInstall(
		system.PlatformProfile{OS: "linux", LinuxDistro: system.LinuxDistroUbuntu, PackageManager: "apt"},
		[]model.AgentID{model.AgentOpenCode},
	)
	restore()

	if setenvCalls != 0 {
		t.Fatalf("setenvCalls = %d, want 0", setenvCalls)
	}
}

func TestWithTermuxOpenPetsEnvForSyncEnablesAndRestoresWhenEligible(t *testing.T) {
	restoreSetenv := osSetenv
	restoreUnsetenv := osUnsetenv
	restoreLookupEnv := osLookupEnv
	t.Cleanup(func() {
		osSetenv = restoreSetenv
		osUnsetenv = restoreUnsetenv
		osLookupEnv = restoreLookupEnv
	})

	t.Setenv("PREFIX", "/data/data/com.termux/files/usr")

	env := map[string]string{}
	osLookupEnv = func(key string) (string, bool) {
		v, ok := env[key]
		return v, ok
	}
	osSetenv = func(key, value string) error {
		env[key] = value
		return nil
	}
	osUnsetenv = func(key string) error {
		delete(env, key)
		return nil
	}

	restore := withTermuxOpenPetsEnvForSync([]model.AgentID{model.AgentOpenCode})
	if got := env[termuxOpenPetsEnvKey]; got != "1" {
		t.Fatalf("env during sync = %q, want 1", got)
	}

	restore()
	if _, exists := env[termuxOpenPetsEnvKey]; exists {
		t.Fatalf("env should be removed after restore, got map: %#v", env)
	}
}

func TestIsTermuxShellEnvironment(t *testing.T) {
	t.Setenv("PREFIX", "")
	t.Setenv("TERMUX_VERSION", "")
	t.Setenv("TERMUX_APP_PID", "")
	if isTermuxShellEnvironment() {
		t.Fatal("isTermuxShellEnvironment() = true, want false")
	}

	t.Setenv("PREFIX", "/data/data/com.termux/files/usr")
	if !isTermuxShellEnvironment() {
		t.Fatal("isTermuxShellEnvironment() = false, want true with termux prefix")
	}
}
