//  Copyright 2019 Google Inc. All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package packages

import (
	"errors"
	"os"
	"os/exec"
	"reflect"
	"testing"

	utilmocks "github.com/GoogleCloudPlatform/osconfig/util/mocks"
	"github.com/golang/mock/gomock"
)

func TestInstallAptPackages(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCommandRunner := utilmocks.NewMockCommandRunner(mockCtrl)
	runner = mockCommandRunner

	expectedCmd := exec.Command(aptGet, append(aptGetInstallArgs, pkgs...)...)
	expectedCmd.Env = append(os.Environ(),
		"DEBIAN_FRONTEND=noninteractive",
	)

	mockCommandRunner.EXPECT().Run(testCtx, expectedCmd).Return([]byte("stdout"), []byte("stderr"), nil).Times(1)
	if err := InstallAptPackages(testCtx, pkgs); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	first := mockCommandRunner.EXPECT().Run(testCtx, expectedCmd).Return([]byte("stdout"), dpkgErr, errors.New("error")).Times(1)
	repair := mockCommandRunner.EXPECT().Run(testCtx, exec.Command(dpkg, dpkgRepairArgs...)).After(first).Return([]byte("stdout"), []byte("stderr"), nil).Times(1)
	mockCommandRunner.EXPECT().Run(testCtx, expectedCmd).After(repair).Return([]byte("stdout"), []byte("stderr"), errors.New("error")).Times(1)
	if err := InstallAptPackages(testCtx, pkgs); err == nil {
		t.Errorf("did not get expected error")
	}
}

func TestRemoveAptPackages(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCommandRunner := utilmocks.NewMockCommandRunner(mockCtrl)
	runner = mockCommandRunner
	expectedCmd := exec.Command(aptGet, append(aptGetRemoveArgs, pkgs...)...)
	expectedCmd.Env = append(os.Environ(),
		"DEBIAN_FRONTEND=noninteractive",
	)

	mockCommandRunner.EXPECT().Run(testCtx, expectedCmd).Return([]byte("stdout"), []byte("stderr"), nil).Times(1)
	if err := RemoveAptPackages(testCtx, pkgs); err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	first := mockCommandRunner.EXPECT().Run(testCtx, expectedCmd).Return([]byte("stdout"), dpkgErr, errors.New("error")).Times(1)
	repair := mockCommandRunner.EXPECT().Run(testCtx, exec.Command(dpkg, dpkgRepairArgs...)).After(first).Return([]byte("stdout"), []byte("stderr"), nil).Times(1)
	mockCommandRunner.EXPECT().Run(testCtx, expectedCmd).After(repair).Return([]byte("stdout"), []byte("stderr"), errors.New("error")).Times(1)
	if err := RemoveAptPackages(testCtx, pkgs); err == nil {
		t.Errorf("did not get expected error")
	}
}

func TestInstalledDebPackages(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCommandRunner := utilmocks.NewMockCommandRunner(mockCtrl)
	runner = mockCommandRunner
	expectedCmd := exec.Command(dpkgquery, dpkgQueryArgs...)
	data := []byte("foo amd64 1.2.3-4")

	mockCommandRunner.EXPECT().Run(testCtx, expectedCmd).Return(data, []byte("stderr"), nil).Times(1)
	ret, err := InstalledDebPackages(testCtx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	want := []PkgInfo{{"foo", "x86_64", "1.2.3-4"}}
	if !reflect.DeepEqual(ret, want) {
		t.Errorf("InstalledDebPackages() = %v, want %v", ret, want)
	}

	mockCommandRunner.EXPECT().Run(testCtx, expectedCmd).Return(data, []byte("stderr"), errors.New("error")).Times(1)
	if _, err := InstalledDebPackages(testCtx); err == nil {
		t.Errorf("did not get expected error")
	}
}

func TestParseInstalledDebpackages(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want []PkgInfo
	}{
		{"NormalCase", []byte("foo amd64 1.2.3-4\nbar noarch 1.2.3-4"), []PkgInfo{{"foo", "x86_64", "1.2.3-4"}, {"bar", "all", "1.2.3-4"}}},
		{"NoPackages", []byte("nothing here"), nil},
		{"nil", nil, nil},
		{"UnrecognizedPackage", []byte("something we dont understand\n bar noarch 1.2.3-4"), []PkgInfo{{"bar", "all", "1.2.3-4"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInstalledDebpackages(tt.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInstalledDebpackages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseAptUpdates(t *testing.T) {
	normalCase := `
Inst libldap-common [2.4.45+dfsg-1ubuntu1.2] (2.4.45+dfsg-1ubuntu1.3 Ubuntu:18.04/bionic-updates, Ubuntu:18.04/bionic-security [all])
Inst google-cloud-sdk [245.0.0-0] (246.0.0-0 cloud-sdk-stretch:cloud-sdk-stretch [amd64]) []
Inst firmware-linux-free (3.4 Debian:9.9/stable [all])
Conf firmware-linux-free (3.4 Debian:9.9/stable [all])
`

	tests := []struct {
		name    string
		data    []byte
		showNew bool
		want    []PkgInfo
	}{
		{"NormalCase", []byte(normalCase), false, []PkgInfo{{"libldap-common", "all", "2.4.45+dfsg-1ubuntu1.3"}, {"google-cloud-sdk", "x86_64", "246.0.0-0"}}},
		{"NormalCaseShowNew", []byte(normalCase), true, []PkgInfo{{"libldap-common", "all", "2.4.45+dfsg-1ubuntu1.3"}, {"google-cloud-sdk", "x86_64", "246.0.0-0"}, {"firmware-linux-free", "all", "3.4"}}},
		{"NoPackages", []byte("nothing here"), false, nil},
		{"nil", nil, false, nil},
		{"UnrecognizedPackage", []byte("Inst something [we dont understand\n Inst google-cloud-sdk [245.0.0-0] (246.0.0-0 cloud-sdk-stretch:cloud-sdk-stretch [amd64])"), false, []PkgInfo{{"google-cloud-sdk", "x86_64", "246.0.0-0"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseAptUpdates(testCtx, tt.data, tt.showNew); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseAptUpdates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAptUpdates(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockCommandRunner := utilmocks.NewMockCommandRunner(mockCtrl)
	runner = mockCommandRunner
	updateCmd := exec.Command(aptGet, aptGetUpdateArgs...)
	expectedCmd := exec.Command(aptGet, append(aptGetUpgradableArgs, aptGetUpgradeCmd)...)
	data := []byte("Inst google-cloud-sdk [245.0.0-0] (246.0.0-0 cloud-sdk-stretch:cloud-sdk-stretch [amd64])")

	first := mockCommandRunner.EXPECT().Run(testCtx, updateCmd).Return(data, []byte("stderr"), nil).Times(1)
	mockCommandRunner.EXPECT().Run(testCtx, expectedCmd).After(first).Return(data, []byte("stderr"), nil).Times(1)
	ret, err := AptUpdates(testCtx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	want := []PkgInfo{{"google-cloud-sdk", "x86_64", "246.0.0-0"}}
	if !reflect.DeepEqual(ret, want) {
		t.Errorf("AptUpdates() = %v, want %v", ret, want)
	}

	first = mockCommandRunner.EXPECT().Run(testCtx, updateCmd).Return(data, []byte("stderr"), nil).Times(1)
	mockCommandRunner.EXPECT().Run(testCtx, expectedCmd).After(first).Return(data, []byte("stderr"), errors.New("error")).Times(1)
	if _, err := AptUpdates(testCtx); err == nil {
		t.Errorf("did not get expected error")
	}
}
