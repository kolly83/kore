/**
 * Copyright 2020 Appvia Ltd <info@appvia.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package apiserver

import (
	"net/http"
	"time"

	"github.com/appvia/kore/pkg/apiserver/filters"
	"github.com/appvia/kore/pkg/apiserver/types"

	clustersv1 "github.com/appvia/kore/pkg/apis/clusters/v1"
	configv1 "github.com/appvia/kore/pkg/apis/config/v1"
	gke "github.com/appvia/kore/pkg/apis/gke/v1alpha1"
	orgv1 "github.com/appvia/kore/pkg/apis/org/v1"
	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/kore/validation"
	"github.com/appvia/kore/pkg/utils"

	restful "github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"
)

func init() {
	RegisterHandler(&teamHandler{})
}

type teamHandler struct {
	kore.Interface
	// DefaultHandlder implements default features
	DefaultHandler
}

// Register is called by the api server to register the service
func (u *teamHandler) Register(i kore.Interface, builder utils.PathBuilder) (*restful.WebService, error) {
	u.Interface = i
	path := builder.Path("teams")

	log.WithFields(log.Fields{
		"path": path,
	}).Info("registering the teams webservice with container")

	ws := &restful.WebService{}
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	ws.Path(path)

	ws.Route(
		ws.PUT("/invitation/{token}").To(u.invitationSubmit).
			Doc("Used to verify and handle the team invitation generated links").
			Filter(filters.NewAuditingFilter(i.Audit, path, "InvitationSubmit")).
			Param(ws.PathParameter("token", "The generated base64 invitation token which was provided from the team")).
			Returns(http.StatusOK, "Indicates the generated link is valid and the user has been granted access", types.TeamInvitationResponse{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("").To(u.listTeams).
			Doc("Returns all the teams in the kore").
			Filter(filters.NewAuditingFilter(i.Audit, path, "ListTeams")).
			Returns(http.StatusOK, "A list of all the teams in the kore", orgv1.TeamList{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}").To(u.findTeam).
			Doc("Return information related to the specific team in the kore").
			Filter(filters.NewAuditingFilter(i.Audit, path, "GetTeam")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "Contains the team definintion from the kore", orgv1.Team{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.PUT("/{team}").To(u.updateTeam).
			Doc("Used to create or update a team in the kore").
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Filter(filters.NewAuditingFilter(i.Audit, path, "UpdateTeam")).
			Reads(orgv1.Team{}, "Contains the definition for a team in the kore").
			Returns(http.StatusOK, "Contains the team definintion from the kore", orgv1.Team{}).
			Returns(http.StatusNotModified, "Indicates the request was processed but no changes applied", orgv1.Team{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.DELETE("/{team}").To(u.deleteTeam).
			Doc("Used to delete a team from the kore").
			Filter(filters.NewAuditingFilter(i.Audit, path, "RemoveTeam")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", orgv1.Team{}).
			Returns(http.StatusNotAcceptable, "Indicates you cannot delete the team for one or more reasons", Error{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	// Team Audit Events

	ws.Route(
		ws.GET("/{team}/audit").To(u.findTeamAudit).
			Doc("Used to return a collection of events against the team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "GetTeamAudit")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.QueryParameter("since", "The duration to retrieve from the audit log").DefaultValue("60m")).
			Returns(http.StatusOK, "A collection of audit events against the team", orgv1.AuditEventList{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	// Team Members

	ws.Route(
		ws.GET("/{team}/members").To(u.findTeamMembers).
			Doc("Returns a list of user memberships in the team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "GetTeamMembers")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "Contains a collection of team memberships for this team", orgv1.TeamMemberList{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.PUT("/{team}/members/{user}").To(u.addTeamMember).
			Doc("Used to add a user to the team via membership").
			Filter(filters.NewAuditingFilter(i.Audit, path, "AddTeamMember")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("user", "Is the user you are adding to the team")).
			Reads(orgv1.TeamMember{}, "The definition for the user in the team").
			Returns(http.StatusOK, "The user has been successfully added to the team", orgv1.TeamMember{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.DELETE("/{team}/members/{user}").To(u.removeTeamMember).
			Doc("Used to remove team membership from the team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "RemoveTeamMember")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("user", "Is the user you are removing from the team")).
			Returns(http.StatusOK, "The user has been successfully removed from the team", orgv1.TeamMember{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	// Team Invitations

	ws.Route(
		ws.GET("/{team}/invites/user").To(u.listInvites).
			Doc("Used to return a list of all the users whom have pending invitations").
			Filter(filters.NewAuditingFilter(i.Audit, path, "ListInvites")).
			Param(ws.PathParameter("team", "The name of the team you are pulling the invitations for")).
			Returns(http.StatusOK, "A list of users whom have an invitation for the team", orgv1.TeamInvitationList{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.PUT("/{team}/invites/user/{user}").To(u.inviteUser).
			Doc("Used to create an invitation for the team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "InviteUser")).
			Param(ws.PathParameter("team", "The name of the team you are creating an invition")).
			Param(ws.PathParameter("user", "The name of the username of the user the invitation is for")).
			Param(ws.QueryParameter("expire", "The expiration of the generated link").DefaultValue("1h")).
			Returns(http.StatusOK, "Indicates the team invitation for the user has been successful", nil).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.DELETE("/{team}/invites/user/{user}").To(u.removeInvite).
			Doc("Used to remove a user invitation for the team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "RemoveInvite")).
			Param(ws.PathParameter("team", "The name of the team you are deleting the invitation")).
			Param(ws.PathParameter("user", "The username of the user whos invitation you are removing")).
			Returns(http.StatusOK, "Indicates the team invitation has been successful removed", nil).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	// Invitation Links

	ws.Route(
		ws.GET("/{team}/invites/generate").To(u.inviteLink).
			Doc("Used to generate a link which provides automatic membership of the team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "GenerateInviteLink")).
			Param(ws.PathParameter("team", "The name of the team you are creating an invition link")).
			Param(ws.QueryParameter("expire", "The expiration of the generated link").DefaultValue("1h")).
			Returns(http.StatusOK, "A generated URI which can be used to join a team", "").
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}/invites/generate/{user}").To(u.inviteLinkByUser).
			Doc("Used to generate for a specific user to join a team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "GenerateInviteLinkForUser")).
			Param(ws.PathParameter("team", "The name of the team you are creating an invition link")).
			Param(ws.PathParameter("user", "The username of the user the link should be limited for")).
			Returns(http.StatusOK, "A generated URI which users can use to join the team", "").
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	// Team Allocations

	ws.Route(
		ws.GET("/{team}/allocations").To(u.findAllocations).
			Doc("Used to return a list of all the allocations in the team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "ListAllocations")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.QueryParameter("assigned", "Retrieves all allocations which have been assigned to you")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", configv1.AllocationList{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)
	ws.Route(
		ws.GET("/{team}/allocations/{name}").To(u.findAllocation).
			Doc("Used to return an allocation within the team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "GetAllocation")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is the name of the allocation you wish to return")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", configv1.Allocation{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)
	ws.Route(
		ws.PUT("/{team}/allocations/{name}").To(u.updateAllocation).
			Doc("Used to create/update an allocation within the team.").
			Filter(filters.NewAuditingFilter(i.Audit, path, "UpdateAllocation")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is the name of the allocation you wish to update")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", configv1.Allocation{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)
	ws.Route(
		ws.DELETE("/{team}/allocations/{name}").To(u.deleteAllocation).
			Doc("Remove an allocation from a team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "RemoveAllocation")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is the name of the allocation you wish to delete")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", configv1.Allocation{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	// Team Namespaces

	ws.Route(
		ws.GET("/{team}/namespaceclaims").To(u.findNamespaces).
			Doc("Used to return all namespaces for the team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "ListNamespaces")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "Contains the former definition from the kore", clustersv1.NamespaceClaimList{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}/namespaceclaims/{name}").To(u.findNamespace).
			Doc("Used to return the details of a namespace within a team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "GetNamespace")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the namespace claim you are acting upon")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", clustersv1.NamespaceClaim{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.PUT("/{team}/namespaceclaims/{name}").To(u.updateNamespace).
			Doc("Used to create or update the details of a namespace within a team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "UpdateNamespace")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the namespace claim you are acting upon")).
			Reads(clustersv1.NamespaceClaim{}, "The definition for namespace claim").
			Returns(http.StatusOK, "Contains the definition from the kore", clustersv1.NamespaceClaim{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.DELETE("/{team}/namespaceclaims/{name}").To(u.deleteNamespace).
			Doc("Used to remove a namespace from a team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "RemoveNamespace")).
			Param(ws.PathParameter("name", "Is name the of the namespace claim you are acting upon")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "Contains the former definition from the kore", clustersv1.NamespaceClaim{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	// Kubernetes Credentials

	ws.Route(
		ws.GET("/{team}/kubernetescredentials").To(u.findKubernetesCredentials).
			Doc("Used to retrieve all kubernetes credentials for a team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "ListKubernetesCredentials")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "Contains the former definition from the kore", clustersv1.KubernetesCredentialsList{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}/kubernetescredentials/{name}").To(u.findKubernetesCredential).
			Doc("Used to retrieve specific kubernetes credentials within a team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "GetKubernetesCredential")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the kubernetes credentials you are acting upon")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", clustersv1.KubernetesCredentials{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.PUT("/{team}/kubernetescredentials/{name}").To(u.updateKubernetesCredential).
			Doc("Used to create/update specific kubernetes credentials within a team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "UpdateKubernetesCredential")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the kubernetes credentials you are acting upon")).
			Reads(clustersv1.KubernetesCredentials{}, "The definition for kubernetes credentials").
			Returns(http.StatusOK, "Contains the definition from the kore", clustersv1.KubernetesCredentials{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.DELETE("/{team}/kubernetescredentials/{name}").To(u.deleteKubernetesCredential).
			Doc("Used to remove specific kubernetes credentials from a team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "RemoveKubernetesCredential")).
			Param(ws.PathParameter("name", "Is name the of the kubernetes credentials you are acting upon")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Doc("Used to return the cluster definition from the kore").
			Returns(http.StatusOK, "Contains the former definition from the kore", clustersv1.KubernetesCredentials{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	// Team Clusters

	ws.Route(
		ws.GET("/{team}/clusters").To(u.findClusters).
			Doc("Lists all clusters available for a team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "ListClusters")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", clustersv1.KubernetesList{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}/clusters/{name}").To(u.findCluster).
			Doc("Used to return the cluster definition from the kore").
			Filter(filters.NewAuditingFilter(i.Audit, path, "GetCluster")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the kubernetes cluster you are acting upon")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", clustersv1.Kubernetes{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.PUT("/{team}/clusters/{name}").To(u.updateCluster).
			Doc("Used to create/update a cluster definition for a team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "UpdateCluster")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the kubernetes cluster you are acting upon")).
			Reads(clustersv1.Kubernetes{}, "The definition for kubernetes cluster").
			Returns(http.StatusOK, "Contains the former team definition from the kore", clustersv1.Kubernetes{}).
			Returns(http.StatusBadRequest, "Validation error of the provided details", validation.ErrValidation{}). // @TODO: Change this to be a class in the orgv1 space
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.DELETE("/{team}/clusters/{name}").To(u.deleteCluster).
			Doc("Used to remove a cluster from a team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "RemoveCluster")).
			Param(ws.PathParameter("name", "Is name the of the GKE cluster you are acting upon")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", clustersv1.Kubernetes{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	// Team Cloud Providers

	ws.Route(
		ws.GET("/{team}/gkes").To(u.findGKEs).
			Doc("Returns a list of Google Container Engine clusters which the team has access").
			Filter(filters.NewAuditingFilter(i.Audit, path, "ListGKEs")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", gke.GKEList{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}/gkes/{name}").To(u.findGKE).
			Doc("Returns a specific Google Container Engine cluster to which the team has access").
			Filter(filters.NewAuditingFilter(i.Audit, path, "GetGKE")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the GKE cluster you are acting upon")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", gke.GKE{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.PUT("/{team}/gkes/{name}").To(u.updateGKE).
			Doc("Is used to provision or update a GKE cluster in the kore").
			Filter(filters.NewAuditingFilter(i.Audit, path, "UpdateGKE")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the GKE cluster you are acting upon")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", gke.GKE{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.DELETE("/{team}/gkes/{name}").To(u.deleteGKE).
			Doc("Is used to delete a managed GKE cluster from the kore").
			Filter(filters.NewAuditingFilter(i.Audit, path, "RemoveGKE")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the GKE cluster you are acting upon")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", gke.GKE{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	// GKE Credentials - @TODO these all need to be autogenerated

	ws.Route(
		ws.GET("/{team}/gkecredentials").To(u.findGKECredientalss).
			Doc("Returns a list of GKE Credentials to which the team has access").
			Filter(filters.NewAuditingFilter(i.Audit, path, "ListGKECredentials")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", gke.GKECredentialsList{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.GET("/{team}/gkecredentials/{name}").To(u.findGKECredientals).
			Doc("Returns a specific GKE Credential to which the team has access").
			Filter(filters.NewAuditingFilter(i.Audit, path, "GetGKECredential")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the GKE cluster you are acting upon")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", gke.GKECredentials{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.PUT("/{team}/gkecredentials/{name}").To(u.updateGKECredientals).
			Doc("Creates or updates a specific GKE Credential to which the team has access").
			Filter(filters.NewAuditingFilter(i.Audit, path, "UpdateGKECredential")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the GKE cluster you are acting upon")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", gke.GKECredentials{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	ws.Route(
		ws.DELETE("/{team}/gkecredentials/{name}").To(u.deleteGKECredientals).
			Doc("Deletes a specific GKE Credential from the team").
			Filter(filters.NewAuditingFilter(i.Audit, path, "RemoveGKECredential")).
			Param(ws.PathParameter("team", "Is the name of the team you are acting within")).
			Param(ws.PathParameter("name", "Is name the of the GKE cluster you are acting upon")).
			Returns(http.StatusOK, "Contains the former team definition from the kore", gke.GKECredentials{}).
			DefaultReturns("A generic API error containing the cause of the error", Error{}),
	)

	return ws, nil
}

// Name returns the name of the handler
func (u teamHandler) Name() string {
	return "teams"
}

// findTeamAudit returns the audit log for a team
func (u teamHandler) findTeamAudit(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		team := req.PathParameter("team")

		since := req.QueryParameter("since")
		if since == "" {
			since = "60m"
		}
		tm, err := time.ParseDuration(since)
		if err != nil {
			return err
		}

		list, err := u.Teams().Team(team).AuditEvents(req.Request.Context(), tm)
		if err != nil {
			return err
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, list)
	})
}

// Teams Management

// deleteTeam is responsible for deleting a team from the kore
func (u teamHandler) deleteTeam(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		err := u.Teams().Delete(req.Request.Context(), req.PathParameter("team"))
		if err != nil {
			return err
		}
		resp.WriteHeader(http.StatusOK)

		return nil
	})
}

// findTeam returns a specific team
func (u teamHandler) findTeam(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		team, err := u.Teams().Get(req.Request.Context(), req.PathParameter("team"))
		if err != nil {
			return err
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, team)
	})
}

// listTeams returns all the teams in the kore
func (u teamHandler) listTeams(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		teams, err := u.Teams().List(req.Request.Context())
		if err != nil {
			return err
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, teams)
	})
}

// updateTeam is responsible for updating for creating a team in the kore
func (u teamHandler) updateTeam(req *restful.Request, resp *restful.Response) {
	handleErrors(req, resp, func() error {
		team := &orgv1.Team{}
		if err := req.ReadEntity(team); err != nil {
			return err
		}
		team, err := u.Teams().Update(req.Request.Context(), team)
		if err != nil {
			return err
		}

		return resp.WriteHeaderAndEntity(http.StatusOK, team)
	})
}
