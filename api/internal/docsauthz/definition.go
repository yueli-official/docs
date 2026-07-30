// Package docsauthz owns the Docs consumer declaration and product-specific
// resource constraints. The Foundation module owns execution and persistence.
package docsauthz

import (
	"context"

	"github.com/yueli-official/foundation/go/authorization"
)

const (
	RootScopeID authorization.ScopeID = "docs"

	ScopeSite       authorization.ScopeType = "site"
	ScopeCollection authorization.ScopeType = "collection"
	ScopeDocument   authorization.ScopeType = "document"

	RoleAdministrator authorization.RoleKey = "administrator"
	RoleAuthor        authorization.RoleKey = "author"

	CapabilityPublicRead                authorization.CapabilityKey = "docs.public.read"
	CapabilityCollectionManage          authorization.CapabilityKey = "docs.collection.manage"
	CapabilityDocumentCreate            authorization.CapabilityKey = "docs.document.create"
	CapabilityDocumentRead              authorization.CapabilityKey = "docs.document.read"
	CapabilityDocumentUpdate            authorization.CapabilityKey = "docs.document.update"
	CapabilityDocumentPublish           authorization.CapabilityKey = "docs.document.publish"
	CapabilityDocumentArchive           authorization.CapabilityKey = "docs.document.archive"
	CapabilityDocumentDeletePermanently authorization.CapabilityKey = "docs.document.delete_permanently"
	CapabilityDocumentReassign          authorization.CapabilityKey = "docs.document.reassign"
	CapabilityVersionManage             authorization.CapabilityKey = "docs.version.manage"
	CapabilityImportManage              authorization.CapabilityKey = "docs.import.manage"
	CapabilitySiteSettingsManage        authorization.CapabilityKey = "docs.site_settings.manage"
	CapabilityAssetSettingsManage       authorization.CapabilityKey = "docs.asset_settings.manage"

	RelationOwner authorization.RelationKind = "owner"

	ConstraintNormalRoleOwnsDocument authorization.ConstraintKey = "docs.normal_role_owns_document"
	PredicateRegistrationAuthor      authorization.PredicateKey  = "docs.registration_auto_author"
	TriggerUserRegistered            authorization.TriggerKey    = "identity.user.registered"
	AutomaticRegistrationAuthorKey                               = "docs.registration_author"
)

func CollectionScopeID(id string) authorization.ScopeID {
	return authorization.ScopeID("collection:" + id)
}

func DocumentScopeID(id string) authorization.ScopeID {
	return authorization.ScopeID("document:" + id)
}

func Definition() authorization.Definition {
	normalSubjects := []authorization.SubjectKind{
		authorization.SubjectUser,
		authorization.SubjectService,
	}
	return authorization.Definition{
		Consumer: "docs",
		Version:  1,
		Capabilities: []authorization.CapabilityDefinition{
			{
				Key: CapabilityPublicRead, Version: 1,
				Binding:       authorization.BindingAccessLayerEligible,
				AllowedScopes: []authorization.ScopeType{ScopeSite, ScopeCollection, ScopeDocument},
			},
			{
				Key: CapabilityCollectionManage, Version: 1, Binding: authorization.BindingNormal,
				AllowedScopes:    []authorization.ScopeType{ScopeSite, ScopeCollection},
				EligibleSubjects: normalSubjects, Delegable: true,
			},
			{
				Key: CapabilityDocumentCreate, Version: 1, Binding: authorization.BindingNormal,
				AllowedScopes:    []authorization.ScopeType{ScopeCollection},
				EligibleSubjects: normalSubjects, Delegable: true,
			},
			{
				Key: CapabilityDocumentRead, Version: 1, Binding: authorization.BindingNormal,
				AllowedScopes:    []authorization.ScopeType{ScopeSite, ScopeCollection, ScopeDocument},
				EligibleSubjects: normalSubjects, QueryableRelation: RelationOwner, Delegable: true,
			},
			{
				Key: CapabilityDocumentUpdate, Version: 1, Binding: authorization.BindingNormal,
				AllowedScopes:    []authorization.ScopeType{ScopeDocument},
				EligibleSubjects: normalSubjects, QueryableRelation: RelationOwner, Delegable: true,
			},
			{
				Key: CapabilityDocumentPublish, Version: 1, Binding: authorization.BindingNormal,
				AllowedScopes:    []authorization.ScopeType{ScopeDocument},
				EligibleSubjects: normalSubjects, QueryableRelation: RelationOwner, Delegable: true,
			},
			{
				Key: CapabilityDocumentArchive, Version: 1, Binding: authorization.BindingNormal,
				AllowedScopes:    []authorization.ScopeType{ScopeDocument},
				EligibleSubjects: normalSubjects, QueryableRelation: RelationOwner, Delegable: true,
			},
			{
				Key: CapabilityDocumentDeletePermanently, Version: 1,
				Binding: authorization.BindingProtectedOnly, Risk: authorization.RiskHigh,
				Audit: authorization.AuditFull, AllowedScopes: []authorization.ScopeType{ScopeDocument},
				EligibleSubjects: normalSubjects,
			},
			{
				Key: CapabilityDocumentReassign, Version: 1,
				Binding: authorization.BindingProtectedOnly, Risk: authorization.RiskHigh,
				Audit: authorization.AuditFull, AllowedScopes: []authorization.ScopeType{ScopeDocument},
				EligibleSubjects: normalSubjects,
			},
			protectedCapability(CapabilityVersionManage),
			protectedCapability(CapabilityImportManage),
			protectedCapability(CapabilitySiteSettingsManage),
			protectedCapability(CapabilityAssetSettingsManage),
		},
		Scopes: authorization.ScopeSchema{Types: []authorization.ScopeTypeDefinition{
			{Key: ScopeSite, Root: true, Children: []authorization.ScopeType{ScopeCollection}},
			{Key: ScopeCollection, Children: []authorization.ScopeType{ScopeDocument}},
			{Key: ScopeDocument},
		}},
		AccessLayers: []authorization.AccessLayerDefinition{
			{
				Key:          authorization.AccessLayerVisitor,
				Capabilities: []authorization.CapabilityKey{CapabilityPublicRead},
			},
			{
				Key: authorization.AccessLayerAuthenticated,
				Capabilities: []authorization.CapabilityKey{
					authorization.CapabilityApplicationCreate,
					authorization.CapabilityApplicationReadOwn,
					authorization.CapabilityApplicationWithdraw,
					authorization.CapabilityInvitationAccept,
				},
			},
		},
		Roles: []authorization.RoleDefinition{
			{
				Key: RoleAdministrator, DisplayName: "管理员", Protected: true,
				Capabilities: []authorization.CapabilityKey{
					authorization.CapabilityManage,
					authorization.CapabilityAuditRead,
					CapabilityCollectionManage,
					CapabilityDocumentCreate,
					CapabilityDocumentRead,
					CapabilityDocumentUpdate,
					CapabilityDocumentPublish,
					CapabilityDocumentArchive,
					CapabilityDocumentDeletePermanently,
					CapabilityDocumentReassign,
					CapabilityVersionManage,
					CapabilityImportManage,
					CapabilitySiteSettingsManage,
					CapabilityAssetSettingsManage,
				},
			},
			{
				Key: RoleAuthor, DisplayName: "作者",
				Capabilities: []authorization.CapabilityKey{
					CapabilityDocumentCreate,
					CapabilityDocumentRead,
					CapabilityDocumentUpdate,
					CapabilityDocumentPublish,
					CapabilityDocumentArchive,
				},
				Assignment: authorization.AssignmentPolicy{
					Sources: []authorization.GrantSource{
						authorization.GrantSourceApplication,
						authorization.GrantSourceInvitation,
						authorization.GrantSourceDirect,
						authorization.GrantSourceAutomatic,
						authorization.GrantSourceGroup,
					},
				},
			},
		},
		Constraints: []authorization.ConstraintDefinition{
			{
				Key: ConstraintNormalRoleOwnsDocument, Version: 1,
				Mode: authorization.ConstraintSource,
				Capabilities: []authorization.CapabilityKey{
					CapabilityDocumentRead,
					CapabilityDocumentUpdate,
					CapabilityDocumentPublish,
					CapabilityDocumentArchive,
				},
				AllNormalRoles: true,
			},
		},
		Automatic: []authorization.AutomaticRuleDefinition{
			{
				Key: AutomaticRegistrationAuthorKey, Trigger: TriggerUserRegistered,
				Predicate: PredicateRegistrationAuthor, Role: RoleAuthor, Enabled: false,
			},
		},
	}
}

func ConstraintEvaluators() map[authorization.ConstraintKey]authorization.ConstraintEvaluator {
	return map[authorization.ConstraintKey]authorization.ConstraintEvaluator{
		ConstraintNormalRoleOwnsDocument: authorization.ConstraintFunc(
			func(_ context.Context, input authorization.ConstraintInput) authorization.ConstraintResult {
				for _, owner := range input.Resource.Relations[RelationOwner] {
					if owner == input.Subject {
						return authorization.ConstraintResult{}
					}
				}
				return authorization.ConstraintResult{Denied: true}
			},
		),
	}
}

func PredicateEvaluators() map[authorization.PredicateKey]authorization.PredicateEvaluator {
	return map[authorization.PredicateKey]authorization.PredicateEvaluator{
		PredicateRegistrationAuthor: authorization.PredicateFunc(
			func(_ context.Context, input authorization.PredicateInput) bool {
				return input.Subject.Kind == authorization.SubjectUser
			},
		),
	}
}

func protectedCapability(key authorization.CapabilityKey) authorization.CapabilityDefinition {
	return authorization.CapabilityDefinition{
		Key: key, Version: 1, Binding: authorization.BindingProtectedOnly,
		Risk: authorization.RiskHigh, Audit: authorization.AuditFull,
		AllowedScopes: []authorization.ScopeType{ScopeSite},
		EligibleSubjects: []authorization.SubjectKind{
			authorization.SubjectUser,
			authorization.SubjectService,
		},
	}
}
