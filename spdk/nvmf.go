package spdk

import (
	"encoding/json"
	"fmt"

	pb "github.com/leonwanghui/go-spdk/spdk/proto"
)

func (s *SPDKCaller) ConstructNVMFSubsystem(req *pb.ConstructNVMFSubsystemRequest) error {
	if _, err := s.execCmd(spdkScript, []string{
		"construct_nvmf_subsystem",
		"-c", fmt.Sprint(req.Core),
		"nqn", req.Nqn,
		"listen", req.Listen,
		"hosts", req.Hosts,
		"-a", fmt.Sprint(req.IsAnyHostAllowed),
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
