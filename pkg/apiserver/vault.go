// /**
//  * Copyright 2020 Appvia Ltd <info@appvia.io>
//  *
//  * Licensed under the Apache License, Version 2.0 (the "License");
//  * you may not use this file except in compliance with the License.
//  * You may obtain a copy of the License at
//  *
//  *     http://www.apache.org/licenses/LICENSE-2.0
//  *
//  * Unless required by applicable law or agreed to in writing, software
//  * distributed under the License is distributed on an "AS IS" BASIS,
//  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  * See the License for the specific language governing permissions and
//  * limitations under the License.
//  */

package apiserver

// import (
// 	"net/http"

// 	"github.com/appvia/kore/pkg/apiserver/types"
// 	"github.com/appvia/kore/pkg/kore"
// 	"github.com/appvia/kore/pkg/utils"
// 	restful "github.com/emicklei/go-restful"
// 	log "github.com/sirupsen/logrus"
// )

// func init() {
// 	RegisterHandler(&vaultImpl{})
// }

// type vaultImpl struct {
// 	kore.Interface
// 	// DefaultHandler implements default features
// 	DefaultHandler
// }

// // Register is called by the api server on registration
// func (v *vaultImpl) Register(i kore.Interface, builder utils.PathBuilder) (*restful.WebService, error) {
// 	path := builder.Path("vault")

// 	log.WithFields(log.Fields{
// 		"path": path,
// 	}).Info("registering the user webservice with container")

// 	v.Interface = i

// 	ws := &restful.WebService{}
// 	ws.Consumes(restful.MIME_JSON)
// 	ws.Produces(restful.MIME_JSON)
// 	ws.Path(path)

// 	ws.Route(
// 		withStandardErrors(ws.GET("")).To(v.listVault).
// 			Doc("Returns a list of key").
// 			Operation("list").
// 			Returns(http.StatusOK, "A list of all the keys in the vault", types.WhoAmI{}),
// 	)

// 	return ws, nil
// }

// func (v vaultImpl) listVault(req *restful.Request, resp *restful.Response) {

// 	resp.WriteHeaderAndEntity(http.StatusOK, "test1")
// }

// func (v vaultImpl) Name() string {
// 	return "vault"
// }
