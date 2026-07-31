package model

import (
	"slices"
	"testing"
)

func TestComponentsForPresetFullGentlemanUsesCompleteVisualInventory(t *testing.T) {
	tests := []struct {
		name    string
		persona PersonaID
	}{
		{name: "gentleman persona", persona: PersonaGentleman},
		{name: "custom persona", persona: PersonaCustom},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ComponentsForPreset(PresetFullGentleman, tt.persona)

			for _, want := range VisualPolishComponents() {
				if !slices.Contains(got, want) {
					t.Errorf("ComponentsForPreset() missing visual component %q: %v", want, got)
				}
			}
		})
	}
}

func TestVisualPolishComponentsReturnsCompleteManagedCleanupInventory(t *testing.T) {
	want := []ComponentID{ComponentTheme, ComponentClaudeTheme, ComponentOpenCodeGentleLogo}
	if got := VisualPolishComponents(); !slices.Equal(got, want) {
		t.Fatalf("VisualPolishComponents() = %v, want complete cleanup inventory %v", got, want)
	}
}
