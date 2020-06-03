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
	"net/http/pprof"

	"github.com/appvia/kore/pkg/kore"
	"github.com/appvia/kore/pkg/utils"

	restful "github.com/emicklei/go-restful"
	log "github.com/sirupsen/logrus"
)

func init() {
	RegisterHandler(&profilingHandler{})
}

type profilingHandler struct {
	kore.Interface
	// provides the default handler
	DefaultHandler
}

// Register is called by the api server on registration
func (p *profilingHandler) Register(i kore.Interface, builder utils.PathBuilder) (*restful.WebService, error) {
	path := builder.Add("debug")

	p.Interface = i

	log.WithFields(log.Fields{
		"path": path,
	}).Info("registering the profiling webservice")

	ws := &restful.WebService{}
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)
	ws.Path(path.Base())

	ws.Route(
		ws.GET("/pprof").To(func(req *restful.Request, resp *restful.Response) {
			pprof.Index(resp.ResponseWriter, req.Request)
		}).Doc("Provides a profiling index"),
	)

	ws.Route(
		ws.GET("/pprof/cmdline").To(func(req *restful.Request, resp *restful.Response) {
			pprof.Cmdline(resp.ResponseWriter, req.Request)
		}).Doc("Provides the profiling cmdline"),
	)

	ws.Route(
		ws.GET("/pprof/profile").To(func(req *restful.Request, resp *restful.Response) {
			pprof.Profile(resp.ResponseWriter, req.Request)
		}).Doc("Provides the profiling enrtypoint"),
	)

	ws.Route(
		ws.GET("/pprof/symbols").To(func(req *restful.Request, resp *restful.Response) {
			pprof.Symbol(resp.ResponseWriter, req.Request)
		}).Doc("Provides the pprof symbols"),
	)

	ws.Route(
		ws.GET("/pprof/trace").To(func(req *restful.Request, resp *restful.Response) {
			pprof.Trace(resp.ResponseWriter, req.Request)
		}).Doc("Provides the pprof trace"),
	)

	return ws, nil
}

// Name returns the name of the handler
func (p *profilingHandler) Name() string {
	return "profiling"
}

// Enabled returns if the handler is enabled
func (p *profilingHandler) Enabled() bool {
	return p.Config().EnableProfiling
}

// EnableAudit defaults to audit everything.
func (p *profilingHandler) EnableAudit() bool {
	return false
}

// EnableLogging defaults to true
func (p *profilingHandler) EnableLogging() bool {
	return false
}

// EnableAdminsOnly requires the user is a member of the admin group
func (p *profilingHandler) EnableAdminsOnly() bool {
	return true
}
