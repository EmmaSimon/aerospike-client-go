// Copyright 2013-2020 Aerospike, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aerospike

type batchIndexCommandGet struct {
	batchCommandGet
}

func newBatchIndexCommandGet(
	batch *batchNode,
	policy *BatchPolicy,
	records []*BatchRead,
) *batchIndexCommandGet {
	var node *Node
	if batch != nil {
		node = batch.Node
	}

	return &batchIndexCommandGet{
		batchCommandGet{
			batchCommand: batchCommand{
				baseMultiCommand: *newMultiCommand(node, nil),
				policy:           policy,
				batch:            batch,
			},
			records:      nil,
			indexRecords: records,
		},
	}
}

func (cmd *batchIndexCommandGet) cloneBatchCommand(batch *batchNode) batcher {
	res := *cmd
	res.batch = batch
	res.node = batch.Node

	return &res
}

func (cmd *batchIndexCommandGet) writeBuffer(ifc command) error {
	return cmd.setBatchIndexRead(cmd.policy, cmd.indexRecords, cmd.batch)
}

func (cmd *batchIndexCommandGet) Execute() error {
	return cmd.execute(cmd, true)
}

func (cmd *batchIndexCommandGet) generateBatchNodes(cluster *Cluster) ([]*batchNode, error) {
	return newBatchNodeListRecords(cluster, cmd.policy, cmd.indexRecords, cmd.sequenceAP, cmd.sequenceSC, cmd.batch)
}
