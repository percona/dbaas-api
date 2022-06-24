// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: controller/pxc_cluster_api.proto

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

func (this *ListPXCClustersRequest) Validate() error {
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
func (this *ListPXCClustersResponse) Validate() error {
	for _, item := range this.Clusters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Clusters", err)
			}
		}
	}
	return nil
}
func (this *ListPXCClustersResponse_Cluster) Validate() error {
	if this.Operation != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Operation); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Operation", err)
		}
	}
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}
func (this *GetPXCClusterCredentialsRequest) Validate() error {
	if nil == this.KubeAuth {
		return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", fmt.Errorf("message must exist"))
	}
	if this.KubeAuth != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.KubeAuth); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", err)
		}
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *GetPXCClusterCredentialsResponse) Validate() error {
	if this.Credentials != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Credentials); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Credentials", err)
		}
	}
	return nil
}
func (this *CreatePXCClusterRequest) Validate() error {
	if nil == this.KubeAuth {
		return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", fmt.Errorf("message must exist"))
	}
	if this.KubeAuth != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.KubeAuth); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", err)
		}
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if nil == this.Params {
		return github_com_mwitkow_go_proto_validators.FieldError("Params", fmt.Errorf("message must exist"))
	}
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	if this.Pmm != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pmm); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pmm", err)
		}
	}
	return nil
}
func (this *CreatePXCClusterResponse) Validate() error {
	return nil
}
func (this *UpdatePXCClusterRequest) Validate() error {
	if nil == this.KubeAuth {
		return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", fmt.Errorf("message must exist"))
	}
	if this.KubeAuth != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.KubeAuth); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", err)
		}
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if nil == this.Params {
		return github_com_mwitkow_go_proto_validators.FieldError("Params", fmt.Errorf("message must exist"))
	}
	if this.Params != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Params); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Params", err)
		}
	}
	return nil
}
func (this *UpdatePXCClusterRequest_UpdatePXCClusterParams) Validate() error {
	if this.Pxc != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pxc); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pxc", err)
		}
	}
	if this.Proxysql != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Proxysql); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Proxysql", err)
		}
	}
	if this.Haproxy != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Haproxy); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Haproxy", err)
		}
	}
	return nil
}
func (this *UpdatePXCClusterRequest_UpdatePXCClusterParams_PXC) Validate() error {
	if this.ComputeResources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ComputeResources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ComputeResources", err)
		}
	}
	return nil
}
func (this *UpdatePXCClusterRequest_UpdatePXCClusterParams_ProxySQL) Validate() error {
	if this.ComputeResources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ComputeResources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ComputeResources", err)
		}
	}
	return nil
}
func (this *UpdatePXCClusterRequest_UpdatePXCClusterParams_HAProxy) Validate() error {
	if this.ComputeResources != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ComputeResources); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ComputeResources", err)
		}
	}
	return nil
}
func (this *UpdatePXCClusterResponse) Validate() error {
	return nil
}
func (this *DeletePXCClusterRequest) Validate() error {
	if nil == this.KubeAuth {
		return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", fmt.Errorf("message must exist"))
	}
	if this.KubeAuth != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.KubeAuth); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", err)
		}
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *DeletePXCClusterResponse) Validate() error {
	return nil
}
func (this *RestartPXCClusterRequest) Validate() error {
	if nil == this.KubeAuth {
		return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", fmt.Errorf("message must exist"))
	}
	if this.KubeAuth != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.KubeAuth); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("KubeAuth", err)
		}
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *RestartPXCClusterResponse) Validate() error {
	return nil
}
