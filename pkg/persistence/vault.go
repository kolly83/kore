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

package persistence

// import (
// 	"context"
// 	"errors"

// 	"github.com/appvia/kore/pkg/persistence/model"
// 	"github.com/jinzhu/gorm"
// )

// // Vault defines the vault interface to the store
// type Vault interface {
// 	// Delete removes a key from the store
// 	Delete(context.Context, *model.Vault) (*model.Vault, error)
// 	// Exists check if key exists
// 	Exists(context.Context, string) (bool, error)
// 	// Get returns a key from the store
// 	Get(context.Context, string) (*model.Vault, error)
// 	// List returns a list of keys from the store
// 	List(context.Context, ...ListFunc) ([]*model.Vault, error)
// 	// // Update updates a key in the store
// 	// Update(context.Context, *model.Vault) error
// }

// // vaultImpl handles access to the vault model
// type vaultImpl struct {
// 	Interface

// 	load []string

// 	conn *gorm.DB
// }

// // Delete removes a key from the store
// func (v vaultImpl) Delete(ctx context.Context, vault *model.Vault) (*model.Vault, error) {

// 	if vault.Key == "" {
// 		return nil, errors.New("invalid key for deletion")
// 	}

// 	q := v.conn

// 	if vault.Key != "" {
// 		q = q.Where("username = ?", vault.Key)
// 	}

// 	return vault, q.Delete(&model.Vault{}).Error
// }

// // Exists check if the key exists
// func (v vaultImpl) Exists(ctx context.Context, name string) (bool, error) {
// 	if _, err := v.Get(ctx, name); err != nil {
// 		if !gorm.IsRecordNotFoundError(err) {
// 			return false, err
// 		}

// 		return false, nil
// 	}

// 	return true, nil
// }

// // Get returns a key from the store
// func (v vaultImpl) Get(ctx context.Context, name string) (*model.Vault, error) {
// 	// timed := prometheus.NewTimer(getLatency)
// 	// defer timed.ObserveDuration()

// 	key := &model.Vault{}

// 	err := v.conn.Where("key = ?", name).Find(key).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return key, nil
// }

// func (v vaultImpl) List(ctx context.Context, opts ...ListFunc) ([]*model.Vault, error) {
// 	// timed := prometheus.NewTimer(listLatency)
// 	// defer timed.ObserveDuration()

// 	// terms := ApplyListOptions(opts...)

// 	q := Preload(v.load, v.conn).
// 		Model(&model.Team{}).
// 		Select("v.*").
// 		Table("vault t")

// 	var list []*model.Vault

// 	return list, q.Find(&list).Error
// }
