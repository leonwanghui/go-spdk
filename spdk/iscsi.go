package spdk

import (
	"fmt"

	pb "github.com/leonwanghui/go-spdk/spdk/proto"
)

func (s *SPDKCaller) getLuns() (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{"get_luns"})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}

func (s *SPDKCaller) addPortalGroup(tag int32, portals string) (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{
		"add_portal_group",
		"tag", fmt.Sprint(tag),
		"portal_list", portals,
	})
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

func (s *SPDKCaller) deletePortalGroup(tag int32) (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{
		"delete_portal_group",
		"tag", fmt.Sprint(tag),
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}

func (s *SPDKCaller) addInitiatorGroup(tag int32, initiators, netmasks string) (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{
		"add_initiator_group",
		"tag", fmt.Sprint(tag),
		"initiator_list", initiators,
		"netmask_list", netmasks,
	})
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

func (s *SPDKCaller) deleteInitiatorGroup(tag int32) (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{
		"delete_initiator_groups",
		"tag", fmt.Sprint(tag),
	})
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

func (s *SPDKCaller) getTargetNodes() (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{"get_target_nodes"})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}

func (s *SPDKCaller) deleteTargetNode(name string) (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{
		"delete_target_node",
		"target_node_name", name,
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}

func (s *SPDKCaller) getIscsiConnections() (string, error) {
	// TODO Add the definition of returned value.
	out, err := s.execCmd(spdkScript, []string{"get_iscsi_connections"})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(out), nil
}
