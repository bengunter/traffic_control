// Copyright 2015 Comcast Cable Communications Management, LLC

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file was initially generated by gen_to_start.go (add link), as a start
// of the Traffic Ops golang data model

package api

import (
	"encoding/json"
	_ "github.com/Comcast/traffic_control/traffic_ops/experimental/server/output_format" // needed for swagger
	"github.com/jmoiron/sqlx"
	null "gopkg.in/guregu/null.v3"
	"log"
	"time"
)

type Federations struct {
	Id          int64            `db:"id" json:"id"`
	Cname       string           `db:"cname" json:"cname"`
	Description null.String      `db:"description" json:"description"`
	Ttl         int64            `db:"ttl" json:"ttl"`
	CreatedAt   time.Time        `db:"created_at" json:"createdAt"`
	Links       FederationsLinks `json:"_links" db:-`
}

type FederationsLinks struct {
	Self string `db:"self" json:"_self"`
}

// @Title getFederationsById
// @Description retrieves the federations information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Federations
// @Resource /api/2.0
// @Router /api/2.0/federations/{id} [get]
func getFederation(id int64, db *sqlx.DB) (interface{}, error) {
	ret := []Federations{}
	arg := Federations{}
	arg.Id = id
	queryStr := "select *, concat('" + API_PATH + "federations/', id) as self"
	queryStr += " from federations WHERE id=:id"
	nstmt, err := db.PrepareNamed(queryStr)
	err = nstmt.Select(&ret, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nstmt.Close()
	return ret, nil
}

// @Title getFederationss
// @Description retrieves the federations
// @Accept  application/json
// @Success 200 {array}    Federations
// @Resource /api/2.0
// @Router /api/2.0/federations [get]
func getFederations(db *sqlx.DB) (interface{}, error) {
	ret := []Federations{}
	queryStr := "select *, concat('" + API_PATH + "federations/', id) as self"
	queryStr += " from federations"
	err := db.Select(&ret, queryStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ret, nil
}

// @Title postFederations
// @Description enter a new federations
// @Accept  application/json
// @Param                 Body body     Federations   true "Federations object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/federations [post]
func postFederation(payload []byte, db *sqlx.DB) (interface{}, error) {
	var v Federations
	err := json.Unmarshal(payload, &v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sqlString := "INSERT INTO federations("
	sqlString += "cname"
	sqlString += ",description"
	sqlString += ",ttl"
	sqlString += ",created_at"
	sqlString += ") VALUES ("
	sqlString += ":cname"
	sqlString += ",:description"
	sqlString += ",:ttl"
	sqlString += ",:created_at"
	sqlString += ")"
	result, err := db.NamedExec(sqlString, v)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title putFederations
// @Description modify an existing federationsentry
// @Accept  application/json
// @Param   id              path    int     true        "The row id"
// @Param                 Body body     Federations   true "Federations object that should be added to the table"
// @Success 200 {object}    output_format.ApiWrapper
// @Resource /api/2.0
// @Router /api/2.0/federations/{id}  [put]
func putFederation(id int64, payload []byte, db *sqlx.DB) (interface{}, error) {
	var arg Federations
	err := json.Unmarshal(payload, &arg)
	arg.Id = id
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sqlString := "UPDATE federations SET "
	sqlString += "cname = :cname"
	sqlString += ",description = :description"
	sqlString += ",ttl = :ttl"
	sqlString += ",created_at = :created_at"
	sqlString += " WHERE id=:id"
	result, err := db.NamedExec(sqlString, arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}

// @Title delFederationsById
// @Description deletes federations information for a certain id
// @Accept  application/json
// @Param   id              path    int     false        "The row id"
// @Success 200 {array}    Federations
// @Resource /api/2.0
// @Router /api/2.0/federations/{id} [delete]
func delFederation(id int64, db *sqlx.DB) (interface{}, error) {
	arg := Federations{}
	arg.Id = id
	result, err := db.NamedExec("DELETE FROM federations WHERE id=:id", arg)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, err
}
