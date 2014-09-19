//
// Copyright 2014, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cloudstack

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type StartSystemVmParams struct {
	p map[string]interface{}
}

func (p *StartSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *StartSystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new StartSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewStartSystemVmParams(id string) *StartSystemVmParams {
	p := &StartSystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Starts a system virtual machine.
func (s *SystemVMService) StartSystemVm(p *StartSystemVmParams) (*StartSystemVmResponse, error) {
	resp, err := s.cs.newRequest("startSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StartSystemVmResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type StartSystemVmResponse struct {
	JobID                string `json:"jobid,omitempty"`
	Linklocalip          string `json:"linklocalip,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Privatenetmask       string `json:"privatenetmask,omitempty"`
	Networkdomain        string `json:"networkdomain,omitempty"`
	Zonename             string `json:"zonename,omitempty"`
	Linklocalnetmask     string `json:"linklocalnetmask,omitempty"`
	Privateip            string `json:"privateip,omitempty"`
	Id                   string `json:"id,omitempty"`
	Zoneid               string `json:"zoneid,omitempty"`
	Dns2                 string `json:"dns2,omitempty"`
	Jobid                string `json:"jobid,omitempty"`
	Systemvmtype         string `json:"systemvmtype,omitempty"`
	Templateid           string `json:"templateid,omitempty"`
	Publicmacaddress     string `json:"publicmacaddress,omitempty"`
	Podid                string `json:"podid,omitempty"`
	Hostid               string `json:"hostid,omitempty"`
	Jobstatus            int    `json:"jobstatus,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	State                string `json:"state,omitempty"`
	Activeviewersessions int    `json:"activeviewersessions,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
	Privatemacaddress    string `json:"privatemacaddress,omitempty"`
	Publicnetmask        string `json:"publicnetmask,omitempty"`
	Created              string `json:"created,omitempty"`
	Name                 string `json:"name,omitempty"`
	Dns1                 string `json:"dns1,omitempty"`
	Linklocalmacaddress  string `json:"linklocalmacaddress,omitempty"`
}

type RebootSystemVmParams struct {
	p map[string]interface{}
}

func (p *RebootSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *RebootSystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new RebootSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewRebootSystemVmParams(id string) *RebootSystemVmParams {
	p := &RebootSystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Reboots a system VM.
func (s *SystemVMService) RebootSystemVm(p *RebootSystemVmParams) (*RebootSystemVmResponse, error) {
	resp, err := s.cs.newRequest("rebootSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r RebootSystemVmResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type RebootSystemVmResponse struct {
	JobID                string `json:"jobid,omitempty"`
	Linklocalnetmask     string `json:"linklocalnetmask,omitempty"`
	Name                 string `json:"name,omitempty"`
	Activeviewersessions int    `json:"activeviewersessions,omitempty"`
	Dns2                 string `json:"dns2,omitempty"`
	Publicmacaddress     string `json:"publicmacaddress,omitempty"`
	Created              string `json:"created,omitempty"`
	Networkdomain        string `json:"networkdomain,omitempty"`
	Dns1                 string `json:"dns1,omitempty"`
	Jobstatus            int    `json:"jobstatus,omitempty"`
	Id                   string `json:"id,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	Jobid                string `json:"jobid,omitempty"`
	Systemvmtype         string `json:"systemvmtype,omitempty"`
	Podid                string `json:"podid,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Linklocalip          string `json:"linklocalip,omitempty"`
	Publicnetmask        string `json:"publicnetmask,omitempty"`
	Privatemacaddress    string `json:"privatemacaddress,omitempty"`
	Linklocalmacaddress  string `json:"linklocalmacaddress,omitempty"`
	Privatenetmask       string `json:"privatenetmask,omitempty"`
	Zonename             string `json:"zonename,omitempty"`
	Hostid               string `json:"hostid,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
	Privateip            string `json:"privateip,omitempty"`
	State                string `json:"state,omitempty"`
	Templateid           string `json:"templateid,omitempty"`
	Zoneid               string `json:"zoneid,omitempty"`
}

type StopSystemVmParams struct {
	p map[string]interface{}
}

func (p *StopSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["forced"]; found {
		vv := strconv.FormatBool(v.(bool))
		u.Set("forced", vv)
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *StopSystemVmParams) SetForced(v bool) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["forced"] = v
	return
}

func (p *StopSystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new StopSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewStopSystemVmParams(id string) *StopSystemVmParams {
	p := &StopSystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Stops a system VM.
func (s *SystemVMService) StopSystemVm(p *StopSystemVmParams) (*StopSystemVmResponse, error) {
	resp, err := s.cs.newRequest("stopSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r StopSystemVmResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type StopSystemVmResponse struct {
	JobID                string `json:"jobid,omitempty"`
	Podid                string `json:"podid,omitempty"`
	Privatemacaddress    string `json:"privatemacaddress,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
	Dns1                 string `json:"dns1,omitempty"`
	Hostid               string `json:"hostid,omitempty"`
	Linklocalmacaddress  string `json:"linklocalmacaddress,omitempty"`
	Networkdomain        string `json:"networkdomain,omitempty"`
	Created              string `json:"created,omitempty"`
	Id                   string `json:"id,omitempty"`
	Linklocalnetmask     string `json:"linklocalnetmask,omitempty"`
	State                string `json:"state,omitempty"`
	Dns2                 string `json:"dns2,omitempty"`
	Linklocalip          string `json:"linklocalip,omitempty"`
	Jobid                string `json:"jobid,omitempty"`
	Name                 string `json:"name,omitempty"`
	Templateid           string `json:"templateid,omitempty"`
	Publicnetmask        string `json:"publicnetmask,omitempty"`
	Systemvmtype         string `json:"systemvmtype,omitempty"`
	Zoneid               string `json:"zoneid,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	Activeviewersessions int    `json:"activeviewersessions,omitempty"`
	Privatenetmask       string `json:"privatenetmask,omitempty"`
	Jobstatus            int    `json:"jobstatus,omitempty"`
	Privateip            string `json:"privateip,omitempty"`
	Publicmacaddress     string `json:"publicmacaddress,omitempty"`
	Zonename             string `json:"zonename,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
}

type DestroySystemVmParams struct {
	p map[string]interface{}
}

func (p *DestroySystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	return u
}

func (p *DestroySystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

// You should always use this function to get a new DestroySystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewDestroySystemVmParams(id string) *DestroySystemVmParams {
	p := &DestroySystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	return p
}

// Destroyes a system virtual machine.
func (s *SystemVMService) DestroySystemVm(p *DestroySystemVmParams) (*DestroySystemVmResponse, error) {
	resp, err := s.cs.newRequest("destroySystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r DestroySystemVmResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type DestroySystemVmResponse struct {
	JobID                string `json:"jobid,omitempty"`
	Linklocalnetmask     string `json:"linklocalnetmask,omitempty"`
	Publicmacaddress     string `json:"publicmacaddress,omitempty"`
	Networkdomain        string `json:"networkdomain,omitempty"`
	Zoneid               string `json:"zoneid,omitempty"`
	Podid                string `json:"podid,omitempty"`
	Jobid                string `json:"jobid,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Id                   string `json:"id,omitempty"`
	Privateip            string `json:"privateip,omitempty"`
	Dns1                 string `json:"dns1,omitempty"`
	Privatemacaddress    string `json:"privatemacaddress,omitempty"`
	Created              string `json:"created,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
	State                string `json:"state,omitempty"`
	Hostid               string `json:"hostid,omitempty"`
	Linklocalmacaddress  string `json:"linklocalmacaddress,omitempty"`
	Name                 string `json:"name,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	Publicnetmask        string `json:"publicnetmask,omitempty"`
	Systemvmtype         string `json:"systemvmtype,omitempty"`
	Zonename             string `json:"zonename,omitempty"`
	Dns2                 string `json:"dns2,omitempty"`
	Privatenetmask       string `json:"privatenetmask,omitempty"`
	Linklocalip          string `json:"linklocalip,omitempty"`
	Activeviewersessions int    `json:"activeviewersessions,omitempty"`
	Jobstatus            int    `json:"jobstatus,omitempty"`
	Templateid           string `json:"templateid,omitempty"`
}

type ListSystemVmsParams struct {
	p map[string]interface{}
}

func (p *ListSystemVmsParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["keyword"]; found {
		u.Set("keyword", v.(string))
	}
	if v, found := p.p["name"]; found {
		u.Set("name", v.(string))
	}
	if v, found := p.p["page"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("page", vv)
	}
	if v, found := p.p["pagesize"]; found {
		vv := strconv.Itoa(v.(int))
		u.Set("pagesize", vv)
	}
	if v, found := p.p["podid"]; found {
		u.Set("podid", v.(string))
	}
	if v, found := p.p["state"]; found {
		u.Set("state", v.(string))
	}
	if v, found := p.p["storageid"]; found {
		u.Set("storageid", v.(string))
	}
	if v, found := p.p["systemvmtype"]; found {
		u.Set("systemvmtype", v.(string))
	}
	if v, found := p.p["zoneid"]; found {
		u.Set("zoneid", v.(string))
	}
	return u
}

func (p *ListSystemVmsParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
	return
}

func (p *ListSystemVmsParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ListSystemVmsParams) SetKeyword(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["keyword"] = v
	return
}

func (p *ListSystemVmsParams) SetName(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["name"] = v
	return
}

func (p *ListSystemVmsParams) SetPage(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["page"] = v
	return
}

func (p *ListSystemVmsParams) SetPagesize(v int) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["pagesize"] = v
	return
}

func (p *ListSystemVmsParams) SetPodid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["podid"] = v
	return
}

func (p *ListSystemVmsParams) SetState(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["state"] = v
	return
}

func (p *ListSystemVmsParams) SetStorageid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["storageid"] = v
	return
}

func (p *ListSystemVmsParams) SetSystemvmtype(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["systemvmtype"] = v
	return
}

func (p *ListSystemVmsParams) SetZoneid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["zoneid"] = v
	return
}

// You should always use this function to get a new ListSystemVmsParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewListSystemVmsParams() *ListSystemVmsParams {
	p := &ListSystemVmsParams{}
	p.p = make(map[string]interface{})
	return p
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SystemVMService) GetSystemVmID(name string) (string, error) {
	p := &ListSystemVmsParams{}
	p.p = make(map[string]interface{})

	p.p["name"] = name

	l, err := s.ListSystemVms(p)
	if err != nil {
		return "", err
	}

	if l.Count == 0 {
		return "", fmt.Errorf("No match found for %s: %+v", name, l)
	}

	if l.Count == 1 {
		return l.SystemVms[0].Id, nil
	}

	if l.Count > 1 {
		for _, v := range l.SystemVms {
			if v.Name == name {
				return v.Id, nil
			}
		}
	}
	return "", fmt.Errorf("Could not find an exact match for %s: %+v", name, l)
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SystemVMService) GetSystemVmByName(name string) (*SystemVm, int, error) {
	id, err := s.GetSystemVmID(name)
	if err != nil {
		return nil, -1, err
	}

	r, count, err := s.GetSystemVmByID(id)
	if err != nil {
		return nil, count, err
	}
	return r, count, nil
}

// This is a courtesy helper function, which in some cases may not work as expected!
func (s *SystemVMService) GetSystemVmByID(id string) (*SystemVm, int, error) {
	p := &ListSystemVmsParams{}
	p.p = make(map[string]interface{})

	p.p["id"] = id

	l, err := s.ListSystemVms(p)
	if err != nil {
		return nil, -1, err
	}

	if l.Count == 0 {
		return nil, l.Count, fmt.Errorf("No match found for %s: %+v", id, l)
	}

	if l.Count == 1 {
		return l.SystemVms[0], l.Count, nil
	}
	return nil, l.Count, fmt.Errorf("There is more then one result for SystemVm UUID: %s!", id)
}

// List system virtual machines.
func (s *SystemVMService) ListSystemVms(p *ListSystemVmsParams) (*ListSystemVmsResponse, error) {
	resp, err := s.cs.newRequest("listSystemVms", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ListSystemVmsResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ListSystemVmsResponse struct {
	Count     int         `json:"count"`
	SystemVms []*SystemVm `json:"systemvm"`
}

type SystemVm struct {
	Networkdomain        string `json:"networkdomain,omitempty"`
	Privateip            string `json:"privateip,omitempty"`
	Privatemacaddress    string `json:"privatemacaddress,omitempty"`
	Privatenetmask       string `json:"privatenetmask,omitempty"`
	Hostid               string `json:"hostid,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Jobstatus            int    `json:"jobstatus,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	Activeviewersessions int    `json:"activeviewersessions,omitempty"`
	Linklocalip          string `json:"linklocalip,omitempty"`
	Zoneid               string `json:"zoneid,omitempty"`
	Zonename             string `json:"zonename,omitempty"`
	Id                   string `json:"id,omitempty"`
	Publicmacaddress     string `json:"publicmacaddress,omitempty"`
	Linklocalmacaddress  string `json:"linklocalmacaddress,omitempty"`
	Publicnetmask        string `json:"publicnetmask,omitempty"`
	Created              string `json:"created,omitempty"`
	Jobid                string `json:"jobid,omitempty"`
	Podid                string `json:"podid,omitempty"`
	Dns2                 string `json:"dns2,omitempty"`
	Templateid           string `json:"templateid,omitempty"`
	Name                 string `json:"name,omitempty"`
	State                string `json:"state,omitempty"`
	Dns1                 string `json:"dns1,omitempty"`
	Systemvmtype         string `json:"systemvmtype,omitempty"`
	Linklocalnetmask     string `json:"linklocalnetmask,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
}

type MigrateSystemVmParams struct {
	p map[string]interface{}
}

func (p *MigrateSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["hostid"]; found {
		u.Set("hostid", v.(string))
	}
	if v, found := p.p["virtualmachineid"]; found {
		u.Set("virtualmachineid", v.(string))
	}
	return u
}

func (p *MigrateSystemVmParams) SetHostid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["hostid"] = v
	return
}

func (p *MigrateSystemVmParams) SetVirtualmachineid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["virtualmachineid"] = v
	return
}

// You should always use this function to get a new MigrateSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewMigrateSystemVmParams(hostid string, virtualmachineid string) *MigrateSystemVmParams {
	p := &MigrateSystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["hostid"] = hostid
	p.p["virtualmachineid"] = virtualmachineid
	return p
}

// Attempts Migration of a system virtual machine to the host specified.
func (s *SystemVMService) MigrateSystemVm(p *MigrateSystemVmParams) (*MigrateSystemVmResponse, error) {
	resp, err := s.cs.newRequest("migrateSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r MigrateSystemVmResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type MigrateSystemVmResponse struct {
	JobID                string `json:"jobid,omitempty"`
	Linklocalip          string `json:"linklocalip,omitempty"`
	Id                   string `json:"id,omitempty"`
	Jobid                string `json:"jobid,omitempty"`
	Publicnetmask        string `json:"publicnetmask,omitempty"`
	Privatenetmask       string `json:"privatenetmask,omitempty"`
	Zonename             string `json:"zonename,omitempty"`
	Hostid               string `json:"hostid,omitempty"`
	Linklocalmacaddress  string `json:"linklocalmacaddress,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
	Name                 string `json:"name,omitempty"`
	Zoneid               string `json:"zoneid,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	Templateid           string `json:"templateid,omitempty"`
	Activeviewersessions int    `json:"activeviewersessions,omitempty"`
	Privateip            string `json:"privateip,omitempty"`
	Podid                string `json:"podid,omitempty"`
	Systemvmtype         string `json:"systemvmtype,omitempty"`
	State                string `json:"state,omitempty"`
	Linklocalnetmask     string `json:"linklocalnetmask,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Privatemacaddress    string `json:"privatemacaddress,omitempty"`
	Networkdomain        string `json:"networkdomain,omitempty"`
	Dns2                 string `json:"dns2,omitempty"`
	Publicmacaddress     string `json:"publicmacaddress,omitempty"`
	Jobstatus            int    `json:"jobstatus,omitempty"`
	Dns1                 string `json:"dns1,omitempty"`
	Created              string `json:"created,omitempty"`
}

type ChangeServiceForSystemVmParams struct {
	p map[string]interface{}
}

func (p *ChangeServiceForSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["details"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	return u
}

func (p *ChangeServiceForSystemVmParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
	return
}

func (p *ChangeServiceForSystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ChangeServiceForSystemVmParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
	return
}

// You should always use this function to get a new ChangeServiceForSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewChangeServiceForSystemVmParams(id string, serviceofferingid string) *ChangeServiceForSystemVmParams {
	p := &ChangeServiceForSystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["serviceofferingid"] = serviceofferingid
	return p
}

// Changes the service offering for a system vm (console proxy or secondary storage). The system vm must be in a "Stopped" state for this command to take effect.
func (s *SystemVMService) ChangeServiceForSystemVm(p *ChangeServiceForSystemVmParams) (*ChangeServiceForSystemVmResponse, error) {
	resp, err := s.cs.newRequest("changeServiceForSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ChangeServiceForSystemVmResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type ChangeServiceForSystemVmResponse struct {
	Hostid               string `json:"hostid,omitempty"`
	State                string `json:"state,omitempty"`
	Jobid                string `json:"jobid,omitempty"`
	Templateid           string `json:"templateid,omitempty"`
	Privatemacaddress    string `json:"privatemacaddress,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	Privatenetmask       string `json:"privatenetmask,omitempty"`
	Publicnetmask        string `json:"publicnetmask,omitempty"`
	Dns2                 string `json:"dns2,omitempty"`
	Created              string `json:"created,omitempty"`
	Name                 string `json:"name,omitempty"`
	Zonename             string `json:"zonename,omitempty"`
	Linklocalip          string `json:"linklocalip,omitempty"`
	Linklocalmacaddress  string `json:"linklocalmacaddress,omitempty"`
	Networkdomain        string `json:"networkdomain,omitempty"`
	Id                   string `json:"id,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
	Podid                string `json:"podid,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Zoneid               string `json:"zoneid,omitempty"`
	Linklocalnetmask     string `json:"linklocalnetmask,omitempty"`
	Privateip            string `json:"privateip,omitempty"`
	Activeviewersessions int    `json:"activeviewersessions,omitempty"`
	Systemvmtype         string `json:"systemvmtype,omitempty"`
	Jobstatus            int    `json:"jobstatus,omitempty"`
	Publicmacaddress     string `json:"publicmacaddress,omitempty"`
	Dns1                 string `json:"dns1,omitempty"`
}

type ScaleSystemVmParams struct {
	p map[string]interface{}
}

func (p *ScaleSystemVmParams) toURLValues() url.Values {
	u := url.Values{}
	if p.p == nil {
		return u
	}
	if v, found := p.p["details"]; found {
		i := 0
		for k, vv := range v.(map[string]string) {
			u.Set(fmt.Sprintf("details[%d].key", i), k)
			u.Set(fmt.Sprintf("details[%d].value", i), vv)
			i++
		}
	}
	if v, found := p.p["id"]; found {
		u.Set("id", v.(string))
	}
	if v, found := p.p["serviceofferingid"]; found {
		u.Set("serviceofferingid", v.(string))
	}
	return u
}

func (p *ScaleSystemVmParams) SetDetails(v map[string]string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["details"] = v
	return
}

func (p *ScaleSystemVmParams) SetId(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["id"] = v
	return
}

func (p *ScaleSystemVmParams) SetServiceofferingid(v string) {
	if p.p == nil {
		p.p = make(map[string]interface{})
	}
	p.p["serviceofferingid"] = v
	return
}

// You should always use this function to get a new ScaleSystemVmParams instance,
// as then you are sure you have configured all required params
func (s *SystemVMService) NewScaleSystemVmParams(id string, serviceofferingid string) *ScaleSystemVmParams {
	p := &ScaleSystemVmParams{}
	p.p = make(map[string]interface{})
	p.p["id"] = id
	p.p["serviceofferingid"] = serviceofferingid
	return p
}

// Scale the service offering for a system vm (console proxy or secondary storage). The system vm must be in a "Stopped" state for this command to take effect.
func (s *SystemVMService) ScaleSystemVm(p *ScaleSystemVmParams) (*ScaleSystemVmResponse, error) {
	resp, err := s.cs.newRequest("scaleSystemVm", p.toURLValues())
	if err != nil {
		return nil, err
	}

	var r ScaleSystemVmResponse
	if err := json.Unmarshal(resp, &r); err != nil {
		return nil, err
	}

	// If we have a async client, we need to wait for the async result
	if s.cs.async {
		b, warn, err := s.cs.GetAsyncJobResult(r.JobID, s.cs.timeout)
		if err != nil {
			return nil, err
		}
		// If 'warn' has a value it means the job is running longer than the configured
		// timeout, the resonse will contain the jobid of the running async job
		if warn != nil {
			return &r, warn
		}

		b, err = getRawValue(b)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(b, &r); err != nil {
			return nil, err
		}
	}
	return &r, nil
}

type ScaleSystemVmResponse struct {
	JobID                string `json:"jobid,omitempty"`
	Jobid                string `json:"jobid,omitempty"`
	Privatemacaddress    string `json:"privatemacaddress,omitempty"`
	Linklocalip          string `json:"linklocalip,omitempty"`
	Publicip             string `json:"publicip,omitempty"`
	Name                 string `json:"name,omitempty"`
	Gateway              string `json:"gateway,omitempty"`
	Publicmacaddress     string `json:"publicmacaddress,omitempty"`
	Templateid           string `json:"templateid,omitempty"`
	Dns2                 string `json:"dns2,omitempty"`
	Zonename             string `json:"zonename,omitempty"`
	Publicnetmask        string `json:"publicnetmask,omitempty"`
	State                string `json:"state,omitempty"`
	Networkdomain        string `json:"networkdomain,omitempty"`
	Podid                string `json:"podid,omitempty"`
	Linklocalmacaddress  string `json:"linklocalmacaddress,omitempty"`
	Created              string `json:"created,omitempty"`
	Systemvmtype         string `json:"systemvmtype,omitempty"`
	Dns1                 string `json:"dns1,omitempty"`
	Hostid               string `json:"hostid,omitempty"`
	Zoneid               string `json:"zoneid,omitempty"`
	Privateip            string `json:"privateip,omitempty"`
	Id                   string `json:"id,omitempty"`
	Linklocalnetmask     string `json:"linklocalnetmask,omitempty"`
	Jobstatus            int    `json:"jobstatus,omitempty"`
	Privatenetmask       string `json:"privatenetmask,omitempty"`
	Activeviewersessions int    `json:"activeviewersessions,omitempty"`
	Hostname             string `json:"hostname,omitempty"`
}