package spdk

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	pb "github.com/leonwanghui/go-spdk/spdk/proto"
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

func (s *SPDKCaller) Getbdevs() ([]*pb.BlockDevice, error) {
	out, err := s.execCmd(spdkScript, []string{"get_bdevs"})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var bdevs []*pb.BlockDevice
	if err = json.Unmarshal(out, &bdevs); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return bdevs, nil
}

func (s *SPDKCaller) Deletebdev(name string) error {
	if _, err := s.execCmd(spdkScript, []string{
		"delete_bdev",
		"name", name,
	}); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *SPDKCaller) KillInstance(signame string) error {
	if _, err := s.execCmd(spdkScript, []string{
		"delete_bdev",
		"sig_name", signame,
	}); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *SPDKCaller) ConstructErrorbdev(basename string) error {
	if _, err := s.execCmd(spdkScript, []string{
		"construct_error_bdev",
		"basename", basename,
	}); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *SPDKCaller) ConstructNullbdev(name string, totalSize, blockSize int32) (string, error) {
	out, err := s.execCmd(spdkScript, []string{
		"construct_null_bdev",
		"name", name,
		"total_size", fmt.Sprint(totalSize),
		"block_size", fmt.Sprint(blockSize),
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}

func (s *SPDKCaller) ConstructMallocbdev(totalSize, blockSize int32) error {
	if _, err := s.execCmd(spdkScript, []string{
		"construct_malloc_bdev",
		"total_size", fmt.Sprint(totalSize),
		"block_size", fmt.Sprint(blockSize),
	}); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *SPDKCaller) ConstructAIObdev(req *pb.ConstructAIObdevRequest) (string, error) {
	out, err := s.execCmd(spdkScript, []string{
		"construct_aio_bdev",
		"filename", req.FileName,
		"name", req.Name,
		"block_size", fmt.Sprint(req.BlockSize),
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}

func (s *SPDKCaller) ConstructNVMEbdev(req *pb.ConstructNVMEbdevRequest) (string, error) {
	out, err := s.execCmd(spdkScript, []string{
		"construct_nvme_bdev",
		"-b", req.Name,
		"-t	", req.Trtype,
		"-a", req.Traddr,
		"-f", req.Adrfam,
		"-s", req.Trsvcid,
		"-n", req.Subnqn,
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}

func (s *SPDKCaller) ConstructNVMFSubsystem(req *pb.ConstructNVMFSubsystemRequest) error {
	if _, err := s.execCmd(spdkScript, []string{
		"construct_nvmf_subsystem",
		"nqn", req.Nqn,
		"listen", req.Listen,
		"hosts", req.Hosts,
		"-s", req.SerialNumber,
		"-n", req.Namespaces,
	}); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *SPDKCaller) GetNVMFSubsystems() ([]*pb.NVMFSubsystem, error) {
	out, err := s.execCmd(spdkScript, []string{"get_nvmf_subsystems"})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var nofs []*pb.NVMFSubsystem
	if err := json.Unmarshal(out, &nofs); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return nofs, nil
}

func (s *SPDKCaller) DeleteNVMFSubsystem(nqn string) error {
	if _, err := s.execCmd(spdkScript, []string{
		"delete_nvmf_subsystem",
		"subsystem_nqn", nqn,
	}); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *SPDKCaller) getLuns() (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{"get_luns"})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}

func (s *SPDKCaller) getPortalGroups() (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{"get_portal_groups"})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}

func (s *SPDKCaller) getInitiatorGroups() (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{"get_initiator_groups"})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}

func (s *SPDKCaller) getTargetNodes() (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{"get_target_nodes"})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}

func (s *SPDKCaller) constructTargetNode(req *pb.ConstructTargetRequest) (string, error) {
	// TODO Add the definition of returned value.
	cmd := []string{
		"construct_target_nodes",
		"name", req.Name,
		"alias_name", req.AliasName,
		"lun_name_id_pairs", req.LunNameIdPairs,
		"pg_ig_mappings", req.PgigMappings,
		"queue_depth", fmt.Sprint(req.QueueDepth),
		"chap_disabled", fmt.Sprint(req.ChapDisabled),
		"chap_required", fmt.Sprint(req.ChapRequired),
		"chap_mutual", fmt.Sprint(req.ChapMutual),
		"chap_auth_group", fmt.Sprint(req.ChapAuthGroup),
	}

	out, err := s.execCmd(spdkScript, cmd)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}

func (s *SPDKCaller) execCmd(script string, subCmd []string) ([]byte, error) {
	var cmd = []string{
		"-s", s.ServerAddress,
		"-p", fmt.Sprint(s.Port),
	}
	cmd = append(cmd, subCmd...)

	return exec.Command(script, cmd...).Output()
}
