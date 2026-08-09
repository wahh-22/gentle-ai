package cli

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gentleman-programming/gentle-ai/v2/internal/model"
	"github.com/gentleman-programming/gentle-ai/v2/internal/reviewtransaction"
	"github.com/gentleman-programming/gentle-ai/v2/internal/state"
)

func TestInstallStateLockPathResolvesTrustedHomeSymlink(t *testing.T) {
	physicalHome := filepath.Join(t.TempDir(), "home")
	if err := os.MkdirAll(physicalHome, 0o755); err != nil {
		t.Fatal(err)
	}
	aliasHome := filepath.Join(t.TempDir(), "home")
	if err := os.Symlink(physicalHome, aliasHome); err != nil {
		t.Skipf("symlinks unavailable: %v", err)
	}

	got := installStateLockPath(aliasHome)
	resolvedHome, err := filepath.EvalSymlinks(physicalHome)
	if err != nil {
		t.Fatal(err)
	}
	want := state.Path(resolvedHome) + ".lock"
	if got != want {
		t.Fatalf("installStateLockPath() = %q, want %q", got, want)
	}
}

func TestPersistSyncManagedAssetStateReReadsLatestStateAfterLockContention(t *testing.T) {
	home := t.TempDir()
	held, err := reviewtransaction.AcquireAuthorityFileLock(installStateLockPath(home))
	if err != nil {
		t.Fatal(err)
	}
	if err := persistSyncManagedAssetState(home, model.Selection{}, "sha256:current-writer"); err == nil || !strings.Contains(err.Error(), "acquire install state lock") {
		t.Fatalf("contended sync state persistence error = %v", err)
	}
	if err := state.Write(home, state.InstallState{Persona: "concurrent"}); err != nil {
		t.Fatal(err)
	}
	if err := held.Release(); err != nil {
		t.Fatal(err)
	}
	if err := persistSyncManagedAssetState(home, model.Selection{CommunityTools: []model.CommunityToolID{model.CommunityToolCodeGraph}}, "sha256:current-writer"); err != nil {
		t.Fatal(err)
	}
	got, err := state.Read(home)
	if err != nil || got.Persona != "concurrent" || got.ManagedAssetDigest != "sha256:current-writer" || !got.CommunityToolsConfigured || len(got.CommunityTools) != 1 || got.CommunityTools[0] != string(model.CommunityToolCodeGraph) {
		t.Fatalf("persisted sync state = %#v, err = %v", got, err)
	}
}

func TestPersistSyncManagedAssetStateRefusesCorruptLatestState(t *testing.T) {
	home := t.TempDir()
	if err := state.Write(home, state.InstallState{}); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(state.Path(home), []byte("{not valid json\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := persistSyncManagedAssetState(home, model.Selection{}, "sha256:current-writer"); err == nil || !strings.Contains(err.Error(), "run `gentle-ai install`") {
		t.Fatalf("corrupt sync state persistence error = %v", err)
	}
}

func TestPersistInstallStateMergesExplicitAgentsFromLatestState(t *testing.T) {
	home := t.TempDir()
	held, err := reviewtransaction.AcquireAuthorityFileLock(installStateLockPath(home))
	if err != nil {
		t.Fatal(err)
	}
	if err := state.Write(home, state.InstallState{InstalledAgents: []string{"opencode"}, Persona: "concurrent"}); err != nil {
		t.Fatal(err)
	}
	if err := held.Release(); err != nil {
		t.Fatal(err)
	}
	if err := persistInstallState(home, state.InstallState{InstalledAgents: []string{"codex"}}, []string{"codex"}, InstallFlags{Agents: []string{"codex"}}, "sha256:current-writer"); err != nil {
		t.Fatal(err)
	}
	got, err := state.Read(home)
	if err != nil || len(got.InstalledAgents) != 2 || got.InstalledAgents[0] != "opencode" || got.InstalledAgents[1] != "codex" || got.Persona != "concurrent" || got.ManagedAssetDigest != "sha256:current-writer" {
		t.Fatalf("persisted install state = %#v, err = %v", got, err)
	}
}
