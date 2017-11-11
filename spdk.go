package gospdk

import (
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

func (s *SPDKCaller) getLuns() (string, error) {
	out, err := s.execCmd(spdkScript, []string{"get_luns"})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return out, nil
}

func (s *SPDKCaller) getPortalGroups() (string, error) {
	out, err := s.execCmd(spdkScript, []string{"get_portal_groups"})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return out, nil
}

func (s *SPDKCaller) getInitiatorGroups() (string, error) {
	out, err := s.execCmd(spdkScript, []string{"get_initiator_groups"})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return out, nil
}

func (s *SPDKCaller) getTargetNodes() (string, error) {
	out, err := s.execCmd(spdkScript, []string{"get_target_nodes"})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return out, nil
}

func (s *SPDKCaller) constructTargetNode(p *ConstructTargetParams) (string, error) {
	cmd := []string{
		"construct_target_nodes",
		"name", p.Name,
		"alias_name", p.AliasName,
		"lun_name_id_pairs", p.LunNameIdPairs,
		"pg_ig_mappings", p.PGIGMappings,
		"queue_depth", fmt.Sprint(p.QueueDepth),
		"chap_disabled", fmt.Sprint(p.ChapDisabled),
		"chap_required", fmt.Sprint(p.ChapRequired),
		"chap_mutual", fmt.Sprint(p.ChapMutual),
		"chap_auth_group", fmt.Sprint(p.ChapAuthGroup),
	}

	out, err := s.execCmd(spdkScript, cmd)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return out, nil
}

func (s *SPDKCaller) execCmd(script string, subCmd []string) (string, error) {
	var cmd = []string{
		"-s", s.ServerAddress,
		"-p", fmt.Sprint(s.Port),
	}
	cmd = append(cmd, subCmd...)

	ret, err := exec.Command(script, cmd...).Output()
	if err != nil {
		return "", err
	}
	return string(ret), nil
}
