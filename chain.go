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
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

// Height gets the current height of the blockchain
func (c Client) Height() (int, error) {
	var data struct{ Height int }
	c.url.Path = "/chain/height"
	req, err := c.buildReq(nil, nil, http.MethodGet)
	if err != nil {
		return data.Height, err
	}
	body, err := c.request(req)
	if err != nil {
		return data.Height, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return data.Height, err
	}
	return data.Height, nil
}

// Score gets the current score of the blockchain.
// The higher the score, the better the chain.
// During synchronization, nodes try to get the best block chain in the network.
func (c Client) Score() (string, error) {
	var data struct{ Score string }
	c.url.Path = "/chain/score"
	req, err := c.buildReq(nil, nil, http.MethodGet)
	if err != nil {
		return data.Score, err
	}
	body, err := c.request(req)
	if err != nil {
		return data.Score, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return data.Score, nil
	}
	return data.Score, nil
}

// Block objects are generated by accounts. If an account generates a block
// and the block gets included in the block chain, the generating account,
// called the harvester, gets all the transaction fees for transactions that
// are included in the block. A harvester will therefore usually include as
// many transactions as possible.
// Transactions reflect all account activities. In order for a client to
// have an up to date balance for every account it is crucial to know about
// every transaction that occurred and therefore the client must have
// knowledge about every single block in the chain (one says: the client must
// be synchronized with the block chain).
// Whenever timestamps are used, the time reflects the network time. NEM has
// a time synchronization mechanism which lets all node agree on how many
// seconds since the nemesis have elapsed. This common time is called network
// time.
type Block struct {
	// BUG(tyler): This should really be an int
	TimeStamp     float64
	Signature     string
	PrevBlockHash string
	// BUG(tyler): This should really be an int
	Type         float64
	Transactions []Transaction
	// BUG(tyler): This should really be an int
	Version float64
	Signer  string
	// BUG(tyler): This should really be an int
	Height float64
}

// UnmarshalJSON implements a custom JSON unmarshaller for Blocks
func (b *Block) UnmarshalJSON(data []byte) error {
	var t interface{}
	var ok bool
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}
	m := t.(map[string]interface{})
	b.TimeStamp, ok = m["timeStamp"].(float64)
	if !ok {
		return errors.New("Unable to assert timestamp to int")
	}
	b.Signature, ok = m["signature"].(string)
	if !ok {
		return errors.New("Unable to assert signature to string")
	}
	b.Type, ok = m["type"].(float64)
	if !ok {
		return errors.New("Unable to assert type to int")
	}
	b.Version, ok = m["version"].(float64)
	if !ok {
		return errors.New("Unable to assert version to int")
	}
	b.Signer, ok = m["signer"].(string)
	if !ok {
		return errors.New("Unable to assert signer to string")
	}
	b.Height, ok = m["height"].(float64)
	if !ok {
		return errors.New("Unable to assert height to int")
	}
	var pbh struct{ Data string }
	tempVal, err := json.Marshal(m["prevBlockHash"])
	if err != nil {
		return err
	}
	if err = json.Unmarshal(tempVal, &pbh); err != nil {
		return err
	}
	b.PrevBlockHash = pbh.Data
	if !ok {
		fmt.Printf("%T", pbh)
		return errors.New("Unable to assert prevBlockHash to string")
	}
	var tx []Transaction
	txs, err := json.Marshal(m["transactions"])
	if err != nil {
		return err
	}
	if err := json.Unmarshal(txs, &tx); err != nil {
		return err
	}
	b.Transactions = tx
	return nil
}

// LastBlock will get the most recent confirmed block on NEM and
// return information about the block
func (c Client) LastBlock() (Block, error) {
	var data Block
	c.url.Path = "/chain/last-block"
	req, err := c.buildReq(nil, nil, http.MethodGet)
	if err != nil {
		return data, err
	}
	body, err := c.request(req)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return data, nil
	}
	return data, nil
}

// BlockInfo will supply individual block info identified by block height.
func (c Client) BlockInfo(height int) (Block, error) {
	var data Block
	c.url.Path = "/block/at/public"
	req, err := c.buildReq(map[string]string{"height": strconv.Itoa(height)}, nil, http.MethodGet)
	if err != nil {
		return data, err
	}
	body, err := c.request(req)
	if err != nil {
		return data, err
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return data, err
	}
	return data, nil
}
