// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: controller/pxc_cluster_restore_api.proto

package controllerv1beta1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ListPXCClusterRestoresRequest) Validate() error {
	if nil == this.KubeAuth {
		return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", fmt.Errorf("message must exist"))
	}
	if this.KubeAuth != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.KubeAuth); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", err)
		}
	}
	return nil
}
func (this *ListPXCClusterRestoresResponse) Validate() error {
	return nil
}
