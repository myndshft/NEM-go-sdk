// Copyright 2018 Myndshft Technologies, Inc.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nemgo

import (
	"net/url"
	"reflect"
	"testing"
)

var clientMock = Client{
	url:     url.URL{},
	request: sendReqMock}

func TestGetBatchAccountData(t *testing.T) {
	want := []AccountMetadataPair{
		AccountMetadataPair{
			Account: AccountInfo{
				Address:         "TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S",
				Balance:         124446551689680,
				VestedBalance:   1041345514976241,
				Importance:      0.010263666447108395,
				PublicKey:       "a11a1a6c17a24252e674d151713cdf51991ad101751e4af02a20c61b59f1fe1a",
				Label:           "",
				HarvestedBlocks: 645},
			Meta: AccountMetadata{
				Status:        "LOCKED",
				RemoteStatus:  "ACTIVE",
				CosignatoryOf: []AccountInfo{},
				Cosignatories: []AccountInfo{}}},
		AccountMetadataPair{
			Account: AccountInfo{
				Address:         "TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S",
				Balance:         124446551689680,
				VestedBalance:   1041345514976241,
				Importance:      0.010263666447108395,
				PublicKey:       "a11a1a6c17a24252e674d151713cdf51991ad101751e4af02a20c61b59f1fe1a",
				Label:           "",
				HarvestedBlocks: 645},
			Meta: AccountMetadata{
				Status:        "LOCKED",
				RemoteStatus:  "ACTIVE",
				CosignatoryOf: []AccountInfo{},
				Cosignatories: []AccountInfo{}}}}
	got, err := clientMock.GetBatchAccountData([]string{"TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S", "TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S"})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\nWanted: %v\n Got: %v", want, got)
	}
}

func TestAccountInfo(t *testing.T) {
	want := AccountMetadataPair{
		Account: AccountInfo{
			Address:         "TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S",
			Balance:         124446551689680,
			VestedBalance:   1041345514976241,
			Importance:      0.010263666447108395,
			PublicKey:       "a11a1a6c17a24252e674d151713cdf51991ad101751e4af02a20c61b59f1fe1a",
			Label:           "",
			HarvestedBlocks: 645},
		Meta: AccountMetadata{
			Status:        "LOCKED",
			RemoteStatus:  "ACTIVE",
			CosignatoryOf: []AccountInfo{},
			Cosignatories: []AccountInfo{}}}
	got, err := clientMock.AccountData(Address("TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S"))
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\nWanted: %v\n Got: %v", want, got)
	}
}

func TestAccountInfoPrivateKey(t *testing.T) {
	want := AccountMetadataPair{
		Account: AccountInfo{
			Address:         "TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S",
			Balance:         124446551689680,
			VestedBalance:   1041345514976241,
			Importance:      0.010263666447108395,
			PublicKey:       "a11a1a6c17a24252e674d151713cdf51991ad101751e4af02a20c61b59f1fe1a",
			Label:           "",
			HarvestedBlocks: 645},
		Meta: AccountMetadata{
			Status:        "LOCKED",
			RemoteStatus:  "ACTIVE",
			CosignatoryOf: []AccountInfo{},
			Cosignatories: []AccountInfo{}}}
	got, err := clientMock.AccountData(PublicKey("a11a1a6c17a24252e674d151713cdf51991ad101751e4af02a20c61b59f1fe1a"))
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\nWanted: %v\n Got: %v", want, got)
	}
}

func TestGetDelegated(t *testing.T) {
	want := AccountMetadataPair{
		Account: AccountInfo{
			Address:         "TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S",
			Balance:         124446551689680,
			VestedBalance:   1041345514976241,
			Importance:      0.010263666447108395,
			PublicKey:       "a11a1a6c17a24252e674d151713cdf51991ad101751e4af02a20c61b59f1fe1a",
			Label:           "",
			HarvestedBlocks: 645},
		Meta: AccountMetadata{
			Status:        "LOCKED",
			RemoteStatus:  "ACTIVE",
			CosignatoryOf: []AccountInfo{},
			Cosignatories: []AccountInfo{}}}
	got, err := clientMock.GetDelegated("TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("\nWanted: %v\n Got: %v", want, got)
	}
}

func TestAccountStatus(t *testing.T) {
	want := AccountMetadata{
		Status:        "LOCKED",
		RemoteStatus:  "ACTIVE",
		CosignatoryOf: []AccountInfo{},
		Cosignatories: []AccountInfo{}}
	got, err := clientMock.AccountStatus("TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nWanted: %v\n Got: %v", want, got)
	}
}

func TestHarvested(t *testing.T) {
	want := []HarvestInfo{
		HarvestInfo{
			TimeStamp:  8879051,
			Difficulty: 26453656336676,
			TotalFee:   102585065,
			ID:         1262068,
			Height:     37015},
		HarvestInfo{
			TimeStamp:  8879051,
			Difficulty: 26453656336676,
			TotalFee:   102585065,
			ID:         1262068,
			Height:     37015}}
	got, err := clientMock.Harvested("TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S", "HASHGOESHERE")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nWanted: %v\n Got: %v", want, got)
	}
}

func TestMosaicsOwned(t *testing.T) {
	want := []OwnedMosaic{
		OwnedMosaic{
			MosaicID: struct {
				NamespaceID string
				Name        string
			}{NamespaceID: "alice.drinks",
				Name: "orange juice"},
			Quantity: 123},
		OwnedMosaic{
			MosaicID: struct {
				NamespaceID string
				Name        string
			}{NamespaceID: "alice.drinks",
				Name: "orange juice"},
			Quantity: 123}}
	got, err := clientMock.MosaicsOwned("TBCI2A67UQZAKCR6NS4JWAEICEIGEIM72G3MVW5S")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nWanted: %v\n Got: %v", want, got)
	}
}
