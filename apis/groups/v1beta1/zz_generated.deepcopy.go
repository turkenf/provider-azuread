//go:build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicMembershipInitParameters) DeepCopyInto(out *DynamicMembershipInitParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicMembershipInitParameters.
func (in *DynamicMembershipInitParameters) DeepCopy() *DynamicMembershipInitParameters {
	if in == nil {
		return nil
	}
	out := new(DynamicMembershipInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicMembershipObservation) DeepCopyInto(out *DynamicMembershipObservation) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicMembershipObservation.
func (in *DynamicMembershipObservation) DeepCopy() *DynamicMembershipObservation {
	if in == nil {
		return nil
	}
	out := new(DynamicMembershipObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DynamicMembershipParameters) DeepCopyInto(out *DynamicMembershipParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DynamicMembershipParameters.
func (in *DynamicMembershipParameters) DeepCopy() *DynamicMembershipParameters {
	if in == nil {
		return nil
	}
	out := new(DynamicMembershipParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Group) DeepCopyInto(out *Group) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Group.
func (in *Group) DeepCopy() *Group {
	if in == nil {
		return nil
	}
	out := new(Group)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Group) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupInitParameters) DeepCopyInto(out *GroupInitParameters) {
	*out = *in
	if in.AdministrativeUnitIds != nil {
		in, out := &in.AdministrativeUnitIds, &out.AdministrativeUnitIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AssignableToRole != nil {
		in, out := &in.AssignableToRole, &out.AssignableToRole
		*out = new(bool)
		**out = **in
	}
	if in.AutoSubscribeNewMembers != nil {
		in, out := &in.AutoSubscribeNewMembers, &out.AutoSubscribeNewMembers
		*out = new(bool)
		**out = **in
	}
	if in.Behaviors != nil {
		in, out := &in.Behaviors, &out.Behaviors
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.DynamicMembership != nil {
		in, out := &in.DynamicMembership, &out.DynamicMembership
		*out = make([]DynamicMembershipInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExternalSendersAllowed != nil {
		in, out := &in.ExternalSendersAllowed, &out.ExternalSendersAllowed
		*out = new(bool)
		**out = **in
	}
	if in.HideFromAddressLists != nil {
		in, out := &in.HideFromAddressLists, &out.HideFromAddressLists
		*out = new(bool)
		**out = **in
	}
	if in.HideFromOutlookClients != nil {
		in, out := &in.HideFromOutlookClients, &out.HideFromOutlookClients
		*out = new(bool)
		**out = **in
	}
	if in.MailEnabled != nil {
		in, out := &in.MailEnabled, &out.MailEnabled
		*out = new(bool)
		**out = **in
	}
	if in.MailNickname != nil {
		in, out := &in.MailNickname, &out.MailNickname
		*out = new(string)
		**out = **in
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OnpremisesGroupType != nil {
		in, out := &in.OnpremisesGroupType, &out.OnpremisesGroupType
		*out = new(string)
		**out = **in
	}
	if in.Owners != nil {
		in, out := &in.Owners, &out.Owners
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PreventDuplicateNames != nil {
		in, out := &in.PreventDuplicateNames, &out.PreventDuplicateNames
		*out = new(bool)
		**out = **in
	}
	if in.ProvisioningOptions != nil {
		in, out := &in.ProvisioningOptions, &out.ProvisioningOptions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecurityEnabled != nil {
		in, out := &in.SecurityEnabled, &out.SecurityEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Theme != nil {
		in, out := &in.Theme, &out.Theme
		*out = new(string)
		**out = **in
	}
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(string)
		**out = **in
	}
	if in.WritebackEnabled != nil {
		in, out := &in.WritebackEnabled, &out.WritebackEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupInitParameters.
func (in *GroupInitParameters) DeepCopy() *GroupInitParameters {
	if in == nil {
		return nil
	}
	out := new(GroupInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupList) DeepCopyInto(out *GroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Group, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupList.
func (in *GroupList) DeepCopy() *GroupList {
	if in == nil {
		return nil
	}
	out := new(GroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupObservation) DeepCopyInto(out *GroupObservation) {
	*out = *in
	if in.AdministrativeUnitIds != nil {
		in, out := &in.AdministrativeUnitIds, &out.AdministrativeUnitIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AssignableToRole != nil {
		in, out := &in.AssignableToRole, &out.AssignableToRole
		*out = new(bool)
		**out = **in
	}
	if in.AutoSubscribeNewMembers != nil {
		in, out := &in.AutoSubscribeNewMembers, &out.AutoSubscribeNewMembers
		*out = new(bool)
		**out = **in
	}
	if in.Behaviors != nil {
		in, out := &in.Behaviors, &out.Behaviors
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.DynamicMembership != nil {
		in, out := &in.DynamicMembership, &out.DynamicMembership
		*out = make([]DynamicMembershipObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExternalSendersAllowed != nil {
		in, out := &in.ExternalSendersAllowed, &out.ExternalSendersAllowed
		*out = new(bool)
		**out = **in
	}
	if in.HideFromAddressLists != nil {
		in, out := &in.HideFromAddressLists, &out.HideFromAddressLists
		*out = new(bool)
		**out = **in
	}
	if in.HideFromOutlookClients != nil {
		in, out := &in.HideFromOutlookClients, &out.HideFromOutlookClients
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Mail != nil {
		in, out := &in.Mail, &out.Mail
		*out = new(string)
		**out = **in
	}
	if in.MailEnabled != nil {
		in, out := &in.MailEnabled, &out.MailEnabled
		*out = new(bool)
		**out = **in
	}
	if in.MailNickname != nil {
		in, out := &in.MailNickname, &out.MailNickname
		*out = new(string)
		**out = **in
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ObjectID != nil {
		in, out := &in.ObjectID, &out.ObjectID
		*out = new(string)
		**out = **in
	}
	if in.OnpremisesDomainName != nil {
		in, out := &in.OnpremisesDomainName, &out.OnpremisesDomainName
		*out = new(string)
		**out = **in
	}
	if in.OnpremisesGroupType != nil {
		in, out := &in.OnpremisesGroupType, &out.OnpremisesGroupType
		*out = new(string)
		**out = **in
	}
	if in.OnpremisesNetbiosName != nil {
		in, out := &in.OnpremisesNetbiosName, &out.OnpremisesNetbiosName
		*out = new(string)
		**out = **in
	}
	if in.OnpremisesSamAccountName != nil {
		in, out := &in.OnpremisesSamAccountName, &out.OnpremisesSamAccountName
		*out = new(string)
		**out = **in
	}
	if in.OnpremisesSecurityIdentifier != nil {
		in, out := &in.OnpremisesSecurityIdentifier, &out.OnpremisesSecurityIdentifier
		*out = new(string)
		**out = **in
	}
	if in.OnpremisesSyncEnabled != nil {
		in, out := &in.OnpremisesSyncEnabled, &out.OnpremisesSyncEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Owners != nil {
		in, out := &in.Owners, &out.Owners
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PreferredLanguage != nil {
		in, out := &in.PreferredLanguage, &out.PreferredLanguage
		*out = new(string)
		**out = **in
	}
	if in.PreventDuplicateNames != nil {
		in, out := &in.PreventDuplicateNames, &out.PreventDuplicateNames
		*out = new(bool)
		**out = **in
	}
	if in.ProvisioningOptions != nil {
		in, out := &in.ProvisioningOptions, &out.ProvisioningOptions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ProxyAddresses != nil {
		in, out := &in.ProxyAddresses, &out.ProxyAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecurityEnabled != nil {
		in, out := &in.SecurityEnabled, &out.SecurityEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Theme != nil {
		in, out := &in.Theme, &out.Theme
		*out = new(string)
		**out = **in
	}
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(string)
		**out = **in
	}
	if in.WritebackEnabled != nil {
		in, out := &in.WritebackEnabled, &out.WritebackEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupObservation.
func (in *GroupObservation) DeepCopy() *GroupObservation {
	if in == nil {
		return nil
	}
	out := new(GroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupParameters) DeepCopyInto(out *GroupParameters) {
	*out = *in
	if in.AdministrativeUnitIds != nil {
		in, out := &in.AdministrativeUnitIds, &out.AdministrativeUnitIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AssignableToRole != nil {
		in, out := &in.AssignableToRole, &out.AssignableToRole
		*out = new(bool)
		**out = **in
	}
	if in.AutoSubscribeNewMembers != nil {
		in, out := &in.AutoSubscribeNewMembers, &out.AutoSubscribeNewMembers
		*out = new(bool)
		**out = **in
	}
	if in.Behaviors != nil {
		in, out := &in.Behaviors, &out.Behaviors
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.DynamicMembership != nil {
		in, out := &in.DynamicMembership, &out.DynamicMembership
		*out = make([]DynamicMembershipParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExternalSendersAllowed != nil {
		in, out := &in.ExternalSendersAllowed, &out.ExternalSendersAllowed
		*out = new(bool)
		**out = **in
	}
	if in.HideFromAddressLists != nil {
		in, out := &in.HideFromAddressLists, &out.HideFromAddressLists
		*out = new(bool)
		**out = **in
	}
	if in.HideFromOutlookClients != nil {
		in, out := &in.HideFromOutlookClients, &out.HideFromOutlookClients
		*out = new(bool)
		**out = **in
	}
	if in.MailEnabled != nil {
		in, out := &in.MailEnabled, &out.MailEnabled
		*out = new(bool)
		**out = **in
	}
	if in.MailNickname != nil {
		in, out := &in.MailNickname, &out.MailNickname
		*out = new(string)
		**out = **in
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.OnpremisesGroupType != nil {
		in, out := &in.OnpremisesGroupType, &out.OnpremisesGroupType
		*out = new(string)
		**out = **in
	}
	if in.Owners != nil {
		in, out := &in.Owners, &out.Owners
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PreventDuplicateNames != nil {
		in, out := &in.PreventDuplicateNames, &out.PreventDuplicateNames
		*out = new(bool)
		**out = **in
	}
	if in.ProvisioningOptions != nil {
		in, out := &in.ProvisioningOptions, &out.ProvisioningOptions
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SecurityEnabled != nil {
		in, out := &in.SecurityEnabled, &out.SecurityEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Theme != nil {
		in, out := &in.Theme, &out.Theme
		*out = new(string)
		**out = **in
	}
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(string)
		**out = **in
	}
	if in.WritebackEnabled != nil {
		in, out := &in.WritebackEnabled, &out.WritebackEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupParameters.
func (in *GroupParameters) DeepCopy() *GroupParameters {
	if in == nil {
		return nil
	}
	out := new(GroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupSpec) DeepCopyInto(out *GroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupSpec.
func (in *GroupSpec) DeepCopy() *GroupSpec {
	if in == nil {
		return nil
	}
	out := new(GroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupStatus) DeepCopyInto(out *GroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupStatus.
func (in *GroupStatus) DeepCopy() *GroupStatus {
	if in == nil {
		return nil
	}
	out := new(GroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Member) DeepCopyInto(out *Member) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Member.
func (in *Member) DeepCopy() *Member {
	if in == nil {
		return nil
	}
	out := new(Member)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Member) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberInitParameters) DeepCopyInto(out *MemberInitParameters) {
	*out = *in
	if in.GroupObjectID != nil {
		in, out := &in.GroupObjectID, &out.GroupObjectID
		*out = new(string)
		**out = **in
	}
	if in.GroupObjectIDRef != nil {
		in, out := &in.GroupObjectIDRef, &out.GroupObjectIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.GroupObjectIDSelector != nil {
		in, out := &in.GroupObjectIDSelector, &out.GroupObjectIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.MemberObjectID != nil {
		in, out := &in.MemberObjectID, &out.MemberObjectID
		*out = new(string)
		**out = **in
	}
	if in.MemberObjectIDRef != nil {
		in, out := &in.MemberObjectIDRef, &out.MemberObjectIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.MemberObjectIDSelector != nil {
		in, out := &in.MemberObjectIDSelector, &out.MemberObjectIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberInitParameters.
func (in *MemberInitParameters) DeepCopy() *MemberInitParameters {
	if in == nil {
		return nil
	}
	out := new(MemberInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberList) DeepCopyInto(out *MemberList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Member, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberList.
func (in *MemberList) DeepCopy() *MemberList {
	if in == nil {
		return nil
	}
	out := new(MemberList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MemberList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberObservation) DeepCopyInto(out *MemberObservation) {
	*out = *in
	if in.GroupObjectID != nil {
		in, out := &in.GroupObjectID, &out.GroupObjectID
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MemberObjectID != nil {
		in, out := &in.MemberObjectID, &out.MemberObjectID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberObservation.
func (in *MemberObservation) DeepCopy() *MemberObservation {
	if in == nil {
		return nil
	}
	out := new(MemberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberParameters) DeepCopyInto(out *MemberParameters) {
	*out = *in
	if in.GroupObjectID != nil {
		in, out := &in.GroupObjectID, &out.GroupObjectID
		*out = new(string)
		**out = **in
	}
	if in.GroupObjectIDRef != nil {
		in, out := &in.GroupObjectIDRef, &out.GroupObjectIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.GroupObjectIDSelector != nil {
		in, out := &in.GroupObjectIDSelector, &out.GroupObjectIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.MemberObjectID != nil {
		in, out := &in.MemberObjectID, &out.MemberObjectID
		*out = new(string)
		**out = **in
	}
	if in.MemberObjectIDRef != nil {
		in, out := &in.MemberObjectIDRef, &out.MemberObjectIDRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.MemberObjectIDSelector != nil {
		in, out := &in.MemberObjectIDSelector, &out.MemberObjectIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberParameters.
func (in *MemberParameters) DeepCopy() *MemberParameters {
	if in == nil {
		return nil
	}
	out := new(MemberParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberSpec) DeepCopyInto(out *MemberSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberSpec.
func (in *MemberSpec) DeepCopy() *MemberSpec {
	if in == nil {
		return nil
	}
	out := new(MemberSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberStatus) DeepCopyInto(out *MemberStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberStatus.
func (in *MemberStatus) DeepCopy() *MemberStatus {
	if in == nil {
		return nil
	}
	out := new(MemberStatus)
	in.DeepCopyInto(out)
	return out
}
