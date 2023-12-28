/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-azuread/apis/users/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Member.
func (mg *Member) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.GroupObjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.GroupObjectIDRef,
		Selector:     mg.Spec.ForProvider.GroupObjectIDSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.GroupObjectID")
	}
	mg.Spec.ForProvider.GroupObjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.GroupObjectIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.MemberObjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.MemberObjectIDRef,
		Selector:     mg.Spec.ForProvider.MemberObjectIDSelector,
		To: reference.To{
			List:    &v1beta1.UserList{},
			Managed: &v1beta1.User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.MemberObjectID")
	}
	mg.Spec.ForProvider.MemberObjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.MemberObjectIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.GroupObjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.GroupObjectIDRef,
		Selector:     mg.Spec.InitProvider.GroupObjectIDSelector,
		To: reference.To{
			List:    &GroupList{},
			Managed: &Group{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.GroupObjectID")
	}
	mg.Spec.InitProvider.GroupObjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.GroupObjectIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.MemberObjectID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.InitProvider.MemberObjectIDRef,
		Selector:     mg.Spec.InitProvider.MemberObjectIDSelector,
		To: reference.To{
			List:    &v1beta1.UserList{},
			Managed: &v1beta1.User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.MemberObjectID")
	}
	mg.Spec.InitProvider.MemberObjectID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.MemberObjectIDRef = rsp.ResolvedReference

	return nil
}
