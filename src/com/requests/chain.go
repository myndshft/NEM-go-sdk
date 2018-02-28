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

package requests

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// Block is the structure that contains the transaction information.
// A block can contain up to 120 transactions.
// Blocks are generated and signed by accounts and are the instrument by
// which information is spread in the network.
type Block struct {
	// TimeStamp is the number of seconds elapsed since the creation of the nemesis block.
	TimeStamp int `json:"timeStamp"`
	// Signature is the signature of the block.
	// The signature was generated by the signer and can be used to validate
	// that the block data was not modified by a node.
	Signature string `json:"signature"`
	// PrevBlockHash is the sha3-256 hash of the last block as hex-string.
	PrevBlockHash prevBlockHashData `json:"prevBlockHash"`
	// Type is the block type.
	// There are currently two block types used:
	// -1: Only the nemesis block has this type.
	// 1: Regular block type.
	Type int `json:"type"`
	// Transactions is the array of transaction structures.
	Transactions []Transaction `json:"transactions"`
	// Version is the block version.
	// The following versions are supported.
	// 0x68 << 24 + 1 (1744830465 as 4 byte integer): the main network version
	// 0x60 << 24 + 1 (1610612737 as 4 byte integer): the mijin network version
	// 0x98 << 24 + 1 (-1744830463 as 4 byte integer): the test network version
	Version int `json:"version"`
	// Signer is the public key of the harvester of the block as hexadecimal number.
	Signer string `json:"signer"`
	// Height is the height of the block.
	// Each block has a unique height.
	// Subsequent blocks differ in height by 1.
	Height int `json:"height"`
}

type prevBlockHashData struct {
	Data string `json:"data"`
}

// LastBlock gets the current last block of the chain
func LastBlock(u url.URL) (Block, error) {
	u.Path = "/chain/last-block"
	options := Options{
		URL:    u,
		Method: http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return Block{}, err
	}
	var data Block
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return Block{}, err
	}
	return data, nil
}

// BlockByHeight gets a block by its height
func BlockByHeight(u url.URL, height int) (Block, error) {
	u.Path = "/block/at/public"
	// TODO finish this
	payload, err := json.Marshal(map[string]string{"height": strconv.Itoa(height)})
	if err != nil {
		return Block{}, err
	}
	options := Options{
		URL:    u,
		Method: http.MethodGet,
		Body:   payload}
	resp, err := Send(options)
	if err != nil {
		return Block{}, err
	}
	var data Block
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return Block{}, err
	}
	return data, nil
}

// BlockHeight describes the position of the block within the block chain.
// The first block of the chain has height one. Each subsequent block has a
// height which is one higher than the previous block.
type BlockHeight struct {
	Height int `json:"height"`
}

// Height gets the current height of the block chain
func Height(u url.URL) (BlockHeight, error) {
	u.Path = "/chain/height"
	options := Options{
		URL:    u,
		Method: http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return BlockHeight{}, err
	}

	var data BlockHeight
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return BlockHeight{}, err
	}

	return data, nil
}

// CommunicationTimeStamps contain information about the network time of a remote NIS.
// NEM uses a time synchronization mechanism to synchronize time across the network.
// Each node maintains a network time which is the time of the computer clock
// plus an offset which compensates for the deviation from the computer clocks of other nodes.
type CommunicationTimeStamps struct {
	// SendTimeStamp si the network time at the moment the reply was sent.
	SendTimeStamp int `json:"sendTimeStamp"`
	// ReceiveTimeStamp is the network time at the moment the request was receivied.
	ReceiveTimeStamp int `json:"receiveTimeStamp"`
}

// Time gets network time (in ms)
func Time(u url.URL) (CommunicationTimeStamps, error) {
	u.Path = "/time-sync/network-time"
	options := Options{
		URL:    u,
		Method: http.MethodGet}
	resp, err := Send(options)
	if err != nil {
		return CommunicationTimeStamps{}, err
	}
	var data CommunicationTimeStamps
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return CommunicationTimeStamps{}, err
	}
	return data, nil
}