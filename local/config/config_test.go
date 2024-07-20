package config

import (
	"strings"
	"testing"
)

func TestParseConfig(t *testing.T) {
	yamlData := `
targets:
  - name: home-assistant
    host_addr: http://192.168.1.20
    path: /api/webhook/someurl
    host_port: 8123
  - name: hubitat
    host_addr: http://192.168.1.21
    path: /data
    host_port: 39501
server:
  port: 8123
  path: /api/webhook/someurl
`

	reader := strings.NewReader(yamlData)
	cfg, err := parseConfig(reader)
	if err != nil {
		t.Fatalf("Failed to parse config: %v", err)
	}

	if len(cfg.Targets) != 2 {
		t.Fatalf("Expected 2 targets, got %d", len(cfg.Targets))
	}

	if cfg.Server.Port != 8123 {
		t.Errorf("Expected server port 8123, got %d", cfg.Server.Port)
	}

	if cfg.Server.Path != "/api/webhook/someurl" {
		t.Errorf("Expected server path '/api/webhook/someurl', got '%s'", cfg.Server.Path)
	}

	if cfg.Targets[0].Name != "home-assistant" {
		t.Errorf("Expected first target name 'home-assistant', got '%s'", cfg.Targets[0].Name)
	}

	if cfg.Targets[1].HostAddr != "http://192.168.1.21" {
		t.Errorf("Expected second target host_addr 'http://192.168.1.21', got '%s'", cfg.Targets[1].HostAddr)
	}
}
