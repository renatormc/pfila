package processes

import (
	"os/exec"
)

type IpedParams struct {
	Destination string   `json:"destination"`
	Sources     []string `json:"sources"`
	Portable    bool     `json:"portable"`
	Profile     string   `json:"profile"`
}

func (p *IpedParams) ToCmd() *exec.Cmd {
	args := []string{}
	for _, src := range p.Sources {
		args = append(args, "-d")
		args = append(args, src)
	}
	args = append(args, "-o")
	args = append(args, p.Destination)
	if p.Portable {
		args = append(args, "--portable")
	}
	args = append(args, "-profile")
	args = append(args, p.Profile)
	return exec.Command("iped", args...)
}
