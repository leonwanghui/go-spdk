package spdk

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

var (
	spdkScript  = "spdk.py"
	setupScript = "setup.py"
)

func NewSPDKCaller(address string, port int, isVerbose bool) *SPDKCaller {
	return &SPDKCaller{
		ServerAddress: address,
		Port:          port,
		IsVerbose:     isVerbose,
	}
}

type SPDKCaller struct {
	ServerAddress string
	Port          int
	IsVerbose     bool
}

func (s *SPDKCaller) StartServer(spdkDir, serverName string) error {
	if err := s.initHugePages(spdkDir); err != nil {
		log.Fatalln("Failed to init hugepages:", err)
		return err
	}

	serverDir := path.Join(spdkDir, "app")
	// Judge if the specified app exists.
	if _, err := os.Stat(path.Join(serverDir, serverName)); err != nil {
		log.Fatalln("Failed to find the app:", err)
		return err
	}
	if err := os.Chdir(serverDir); err != nil {
		log.Fatalln("Failed to change to the dir :", err)
		return err
	}
	// Execute the server start process.
	_, err := s.execCmd("./"+serverName, []string{})
	if err != nil {
		return err
	}

	return nil
}

func (s *SPDKCaller) initHugePages(spdkDir string) error {
	hugeDir := path.Join(spdkDir, "scripts")
	// Judge if the init hugepage script exists.
	if _, err := os.Stat(path.Join(hugeDir, setupScript)); err != nil {
		return err
	}
	if err := os.Chdir(hugeDir); err != nil {
		return err
	}
	// Execute the setup process.
	_, err := s.execCmd(setupScript, []string{})
	if err != nil {
		return err
	}

	return nil
}

func (s *SPDKCaller) GetRPCMethods() ([]string, error) {
	out, err := s.execCmd(spdkScript, []string{"get_rpc_methods"})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var methods []string
	if err = json.Unmarshal(out, &methods); err != nil {
		return nil, err
	}

	return methods, nil
}

func (s *SPDKCaller) SetTraceFlag(flag string) error {
	if _, err := s.execCmd(spdkScript, []string{
		"set_trace_flags",
		"flag", flag,
	}); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *SPDKCaller) GetTraceFlags() ([]string, error) {
	out, err := s.execCmd(spdkScript, []string{"get_trace_flags"})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var flags []string
	if err = json.Unmarshal(out, &flags); err != nil {
		return nil, err
	}

	return flags, nil
}

func (s *SPDKCaller) ClearTraceFlag(flag string) error {
	if _, err := s.execCmd(spdkScript, []string{
		"clear_trace_flags",
		"flag", flag,
	}); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *SPDKCaller) execCmd(script string, subCmd []string) ([]byte, error) {
	var cmd = []string{
		"-s", s.ServerAddress,
		"-p", fmt.Sprint(s.Port),
	}
	cmd = append(cmd, subCmd...)

	return exec.Command(script, cmd...).Output()
}
