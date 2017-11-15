package spdk

import (
	"encoding/json"
	"fmt"

	pb "github.com/leonwanghui/go-spdk/spdk/proto"
)

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

func (s *SPDKCaller) constructVhostScsiController(ctrlr, cpumask string) (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{
		"construct_vhost_scsi_controller",
		"ctrlr", ctrlr,
		"--cpumask", cpumask,
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}

func (s *SPDKCaller) removeVhostScsiController(ctrlr string) (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{
		"remove_vhost_scsi_controller",
		"ctrlr", ctrlr,
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}

func (s *SPDKCaller) addVhostScsiLun(ctrlr, name string, devNum int32) (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{
		"add_vhost_scsi_lun",
		"ctrlr", ctrlr,
		"scsi_dev_num", fmt.Sprint(devNum),
		"lun_name", name,
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}

func (s *SPDKCaller) removeVhostScsiDev(ctrlr string, devNum int32) (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{
		"remove_vhost_scsi_dev",
		"ctrlr", ctrlr,
		"scsi_dev_num", fmt.Sprint(devNum),
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}

func (s *SPDKCaller) constructVhostBlkController(ctrlr, devName, cpumask string) (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{
		"construct_vhost_blk_controller",
		"ctrlr", ctrlr,
		"dev_name", devName,
		"--cpumask", cpumask,
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}

func (s *SPDKCaller) removeVhostBlkController(ctrlr string) (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{
		"remove_vhost_blk_controller",
		"ctrlr", ctrlr,
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}

func (s *SPDKCaller) getVhostControllers() (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{"get_vhost_controllers"})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}
