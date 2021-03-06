// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/m3db/m3/src/metrics/generated/proto/metricpb/metadata.proto

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package metricpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import aggregationpb "github.com/m3db/m3/src/metrics/generated/proto/aggregationpb"
import policypb "github.com/m3db/m3/src/metrics/generated/proto/policypb"
import pipelinepb "github.com/m3db/m3/src/metrics/generated/proto/pipelinepb"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PipelineMetadata struct {
	AggregationId   aggregationpb.AggregationID `protobuf:"bytes,1,opt,name=aggregation_id,json=aggregationId" json:"aggregation_id"`
	StoragePolicies []policypb.StoragePolicy    `protobuf:"bytes,2,rep,name=storage_policies,json=storagePolicies" json:"storage_policies"`
	Pipeline        pipelinepb.AppliedPipeline  `protobuf:"bytes,3,opt,name=pipeline" json:"pipeline"`
	DropPolicy      policypb.DropPolicy         `protobuf:"varint,4,opt,name=drop_policy,json=dropPolicy,proto3,enum=policypb.DropPolicy" json:"drop_policy,omitempty"`
}

func (m *PipelineMetadata) Reset()                    { *m = PipelineMetadata{} }
func (m *PipelineMetadata) String() string            { return proto.CompactTextString(m) }
func (*PipelineMetadata) ProtoMessage()               {}
func (*PipelineMetadata) Descriptor() ([]byte, []int) { return fileDescriptorMetadata, []int{0} }

func (m *PipelineMetadata) GetAggregationId() aggregationpb.AggregationID {
	if m != nil {
		return m.AggregationId
	}
	return aggregationpb.AggregationID{}
}

func (m *PipelineMetadata) GetStoragePolicies() []policypb.StoragePolicy {
	if m != nil {
		return m.StoragePolicies
	}
	return nil
}

func (m *PipelineMetadata) GetPipeline() pipelinepb.AppliedPipeline {
	if m != nil {
		return m.Pipeline
	}
	return pipelinepb.AppliedPipeline{}
}

func (m *PipelineMetadata) GetDropPolicy() policypb.DropPolicy {
	if m != nil {
		return m.DropPolicy
	}
	return policypb.DropPolicy_NONE
}

type Metadata struct {
	Pipelines []PipelineMetadata `protobuf:"bytes,1,rep,name=pipelines" json:"pipelines"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptorMetadata, []int{1} }

func (m *Metadata) GetPipelines() []PipelineMetadata {
	if m != nil {
		return m.Pipelines
	}
	return nil
}

type StagedMetadata struct {
	CutoverNanos int64    `protobuf:"varint,1,opt,name=cutover_nanos,json=cutoverNanos,proto3" json:"cutover_nanos,omitempty"`
	Tombstoned   bool     `protobuf:"varint,2,opt,name=tombstoned,proto3" json:"tombstoned,omitempty"`
	Metadata     Metadata `protobuf:"bytes,3,opt,name=metadata" json:"metadata"`
}

func (m *StagedMetadata) Reset()                    { *m = StagedMetadata{} }
func (m *StagedMetadata) String() string            { return proto.CompactTextString(m) }
func (*StagedMetadata) ProtoMessage()               {}
func (*StagedMetadata) Descriptor() ([]byte, []int) { return fileDescriptorMetadata, []int{2} }

func (m *StagedMetadata) GetCutoverNanos() int64 {
	if m != nil {
		return m.CutoverNanos
	}
	return 0
}

func (m *StagedMetadata) GetTombstoned() bool {
	if m != nil {
		return m.Tombstoned
	}
	return false
}

func (m *StagedMetadata) GetMetadata() Metadata {
	if m != nil {
		return m.Metadata
	}
	return Metadata{}
}

type StagedMetadatas struct {
	Metadatas []StagedMetadata `protobuf:"bytes,1,rep,name=metadatas" json:"metadatas"`
}

func (m *StagedMetadatas) Reset()                    { *m = StagedMetadatas{} }
func (m *StagedMetadatas) String() string            { return proto.CompactTextString(m) }
func (*StagedMetadatas) ProtoMessage()               {}
func (*StagedMetadatas) Descriptor() ([]byte, []int) { return fileDescriptorMetadata, []int{3} }

func (m *StagedMetadatas) GetMetadatas() []StagedMetadata {
	if m != nil {
		return m.Metadatas
	}
	return nil
}

type ForwardMetadata struct {
	AggregationId     aggregationpb.AggregationID `protobuf:"bytes,1,opt,name=aggregation_id,json=aggregationId" json:"aggregation_id"`
	StoragePolicy     policypb.StoragePolicy      `protobuf:"bytes,2,opt,name=storage_policy,json=storagePolicy" json:"storage_policy"`
	Pipeline          pipelinepb.AppliedPipeline  `protobuf:"bytes,3,opt,name=pipeline" json:"pipeline"`
	SourceId          uint32                      `protobuf:"varint,4,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	NumForwardedTimes int32                       `protobuf:"varint,5,opt,name=num_forwarded_times,json=numForwardedTimes,proto3" json:"num_forwarded_times,omitempty"`
}

func (m *ForwardMetadata) Reset()                    { *m = ForwardMetadata{} }
func (m *ForwardMetadata) String() string            { return proto.CompactTextString(m) }
func (*ForwardMetadata) ProtoMessage()               {}
func (*ForwardMetadata) Descriptor() ([]byte, []int) { return fileDescriptorMetadata, []int{4} }

func (m *ForwardMetadata) GetAggregationId() aggregationpb.AggregationID {
	if m != nil {
		return m.AggregationId
	}
	return aggregationpb.AggregationID{}
}

func (m *ForwardMetadata) GetStoragePolicy() policypb.StoragePolicy {
	if m != nil {
		return m.StoragePolicy
	}
	return policypb.StoragePolicy{}
}

func (m *ForwardMetadata) GetPipeline() pipelinepb.AppliedPipeline {
	if m != nil {
		return m.Pipeline
	}
	return pipelinepb.AppliedPipeline{}
}

func (m *ForwardMetadata) GetSourceId() uint32 {
	if m != nil {
		return m.SourceId
	}
	return 0
}

func (m *ForwardMetadata) GetNumForwardedTimes() int32 {
	if m != nil {
		return m.NumForwardedTimes
	}
	return 0
}

type TimedMetadata struct {
	AggregationId aggregationpb.AggregationID `protobuf:"bytes,1,opt,name=aggregation_id,json=aggregationId" json:"aggregation_id"`
	StoragePolicy policypb.StoragePolicy      `protobuf:"bytes,2,opt,name=storage_policy,json=storagePolicy" json:"storage_policy"`
}

func (m *TimedMetadata) Reset()                    { *m = TimedMetadata{} }
func (m *TimedMetadata) String() string            { return proto.CompactTextString(m) }
func (*TimedMetadata) ProtoMessage()               {}
func (*TimedMetadata) Descriptor() ([]byte, []int) { return fileDescriptorMetadata, []int{5} }

func (m *TimedMetadata) GetAggregationId() aggregationpb.AggregationID {
	if m != nil {
		return m.AggregationId
	}
	return aggregationpb.AggregationID{}
}

func (m *TimedMetadata) GetStoragePolicy() policypb.StoragePolicy {
	if m != nil {
		return m.StoragePolicy
	}
	return policypb.StoragePolicy{}
}

func init() {
	proto.RegisterType((*PipelineMetadata)(nil), "metricpb.PipelineMetadata")
	proto.RegisterType((*Metadata)(nil), "metricpb.Metadata")
	proto.RegisterType((*StagedMetadata)(nil), "metricpb.StagedMetadata")
	proto.RegisterType((*StagedMetadatas)(nil), "metricpb.StagedMetadatas")
	proto.RegisterType((*ForwardMetadata)(nil), "metricpb.ForwardMetadata")
	proto.RegisterType((*TimedMetadata)(nil), "metricpb.TimedMetadata")
}
func (m *PipelineMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PipelineMetadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintMetadata(dAtA, i, uint64(m.AggregationId.Size()))
	n1, err := m.AggregationId.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.StoragePolicies) > 0 {
		for _, msg := range m.StoragePolicies {
			dAtA[i] = 0x12
			i++
			i = encodeVarintMetadata(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintMetadata(dAtA, i, uint64(m.Pipeline.Size()))
	n2, err := m.Pipeline.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.DropPolicy != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(m.DropPolicy))
	}
	return i, nil
}

func (m *Metadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Metadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Pipelines) > 0 {
		for _, msg := range m.Pipelines {
			dAtA[i] = 0xa
			i++
			i = encodeVarintMetadata(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *StagedMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StagedMetadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CutoverNanos != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(m.CutoverNanos))
	}
	if m.Tombstoned {
		dAtA[i] = 0x10
		i++
		if m.Tombstoned {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintMetadata(dAtA, i, uint64(m.Metadata.Size()))
	n3, err := m.Metadata.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *StagedMetadatas) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StagedMetadatas) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Metadatas) > 0 {
		for _, msg := range m.Metadatas {
			dAtA[i] = 0xa
			i++
			i = encodeVarintMetadata(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ForwardMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ForwardMetadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintMetadata(dAtA, i, uint64(m.AggregationId.Size()))
	n4, err := m.AggregationId.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x12
	i++
	i = encodeVarintMetadata(dAtA, i, uint64(m.StoragePolicy.Size()))
	n5, err := m.StoragePolicy.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	dAtA[i] = 0x1a
	i++
	i = encodeVarintMetadata(dAtA, i, uint64(m.Pipeline.Size()))
	n6, err := m.Pipeline.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if m.SourceId != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(m.SourceId))
	}
	if m.NumForwardedTimes != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintMetadata(dAtA, i, uint64(m.NumForwardedTimes))
	}
	return i, nil
}

func (m *TimedMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimedMetadata) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintMetadata(dAtA, i, uint64(m.AggregationId.Size()))
	n7, err := m.AggregationId.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	dAtA[i] = 0x12
	i++
	i = encodeVarintMetadata(dAtA, i, uint64(m.StoragePolicy.Size()))
	n8, err := m.StoragePolicy.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	return i, nil
}

func encodeVarintMetadata(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PipelineMetadata) Size() (n int) {
	var l int
	_ = l
	l = m.AggregationId.Size()
	n += 1 + l + sovMetadata(uint64(l))
	if len(m.StoragePolicies) > 0 {
		for _, e := range m.StoragePolicies {
			l = e.Size()
			n += 1 + l + sovMetadata(uint64(l))
		}
	}
	l = m.Pipeline.Size()
	n += 1 + l + sovMetadata(uint64(l))
	if m.DropPolicy != 0 {
		n += 1 + sovMetadata(uint64(m.DropPolicy))
	}
	return n
}

func (m *Metadata) Size() (n int) {
	var l int
	_ = l
	if len(m.Pipelines) > 0 {
		for _, e := range m.Pipelines {
			l = e.Size()
			n += 1 + l + sovMetadata(uint64(l))
		}
	}
	return n
}

func (m *StagedMetadata) Size() (n int) {
	var l int
	_ = l
	if m.CutoverNanos != 0 {
		n += 1 + sovMetadata(uint64(m.CutoverNanos))
	}
	if m.Tombstoned {
		n += 2
	}
	l = m.Metadata.Size()
	n += 1 + l + sovMetadata(uint64(l))
	return n
}

func (m *StagedMetadatas) Size() (n int) {
	var l int
	_ = l
	if len(m.Metadatas) > 0 {
		for _, e := range m.Metadatas {
			l = e.Size()
			n += 1 + l + sovMetadata(uint64(l))
		}
	}
	return n
}

func (m *ForwardMetadata) Size() (n int) {
	var l int
	_ = l
	l = m.AggregationId.Size()
	n += 1 + l + sovMetadata(uint64(l))
	l = m.StoragePolicy.Size()
	n += 1 + l + sovMetadata(uint64(l))
	l = m.Pipeline.Size()
	n += 1 + l + sovMetadata(uint64(l))
	if m.SourceId != 0 {
		n += 1 + sovMetadata(uint64(m.SourceId))
	}
	if m.NumForwardedTimes != 0 {
		n += 1 + sovMetadata(uint64(m.NumForwardedTimes))
	}
	return n
}

func (m *TimedMetadata) Size() (n int) {
	var l int
	_ = l
	l = m.AggregationId.Size()
	n += 1 + l + sovMetadata(uint64(l))
	l = m.StoragePolicy.Size()
	n += 1 + l + sovMetadata(uint64(l))
	return n
}

func sovMetadata(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMetadata(x uint64) (n int) {
	return sovMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PipelineMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PipelineMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PipelineMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregationId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AggregationId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoragePolicies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoragePolicies = append(m.StoragePolicies, policypb.StoragePolicy{})
			if err := m.StoragePolicies[len(m.StoragePolicies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pipeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pipeline.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DropPolicy", wireType)
			}
			m.DropPolicy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DropPolicy |= (policypb.DropPolicy(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Metadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Metadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Metadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pipelines", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pipelines = append(m.Pipelines, PipelineMetadata{})
			if err := m.Pipelines[len(m.Pipelines)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StagedMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StagedMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StagedMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CutoverNanos", wireType)
			}
			m.CutoverNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CutoverNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tombstoned", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Tombstoned = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *StagedMetadatas) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: StagedMetadatas: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StagedMetadatas: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadatas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadatas = append(m.Metadatas, StagedMetadata{})
			if err := m.Metadatas[len(m.Metadatas)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ForwardMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ForwardMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ForwardMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregationId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AggregationId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoragePolicy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StoragePolicy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pipeline", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pipeline.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceId", wireType)
			}
			m.SourceId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SourceId |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumForwardedTimes", wireType)
			}
			m.NumForwardedTimes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumForwardedTimes |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TimedMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetadata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TimedMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimedMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AggregationId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AggregationId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoragePolicy", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMetadata
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.StoragePolicy.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetadata
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetadata
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMetadata
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMetadata
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMetadata(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMetadata = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetadata   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/m3db/m3/src/metrics/generated/proto/metricpb/metadata.proto", fileDescriptorMetadata)
}

var fileDescriptorMetadata = []byte{
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x54, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0xdd, 0x69, 0x77, 0x25, 0x9d, 0xda, 0x76, 0x8d, 0x82, 0xa1, 0x2b, 0xb5, 0xc4, 0x97, 0xbe,
	0x38, 0x81, 0x56, 0xf1, 0x45, 0x85, 0x5d, 0x4a, 0xd9, 0x0a, 0xae, 0x4b, 0xd6, 0x27, 0x5f, 0x42,
	0x92, 0x99, 0x8d, 0x03, 0x4d, 0x26, 0xcc, 0x4c, 0x94, 0x7e, 0x83, 0x2f, 0xfe, 0x80, 0xe0, 0x17,
	0xf8, 0x1d, 0xfb, 0xe8, 0x17, 0x88, 0xd4, 0x1f, 0x91, 0x24, 0x33, 0x49, 0xba, 0x2f, 0x52, 0x45,
	0xd8, 0xb7, 0x7b, 0xef, 0xdc, 0x7b, 0x38, 0xe7, 0x70, 0x18, 0xb8, 0x88, 0xa8, 0x7c, 0x9f, 0x05,
	0x28, 0x64, 0xb1, 0x13, 0xcf, 0x70, 0xe0, 0xc4, 0x33, 0x47, 0xf0, 0xd0, 0x89, 0x89, 0xe4, 0x34,
	0x14, 0x4e, 0x44, 0x12, 0xc2, 0x7d, 0x49, 0xb0, 0x93, 0x72, 0x26, 0x99, 0x9a, 0xa7, 0x41, 0x5e,
	0xf8, 0xd8, 0x97, 0x3e, 0x2a, 0xe6, 0xa6, 0xa1, 0x1f, 0x86, 0x8f, 0x1b, 0x88, 0x11, 0x8b, 0x58,
	0x79, 0x18, 0x64, 0x97, 0x45, 0x57, 0xa2, 0xe4, 0x55, 0x79, 0x38, 0x3c, 0xdb, 0x91, 0x80, 0x1f,
	0x45, 0x9c, 0x44, 0xbe, 0xa4, 0x2c, 0x49, 0x83, 0x66, 0xa7, 0xf0, 0xe6, 0x3b, 0xe2, 0xa5, 0x6c,
	0x45, 0xc3, 0x75, 0x1a, 0xa8, 0x42, 0xa1, 0x9c, 0xee, 0x8a, 0x42, 0x53, 0xb2, 0xa2, 0x09, 0xc9,
	0x71, 0x54, 0x59, 0x22, 0xd9, 0x5f, 0x5a, 0xf0, 0xf0, 0x5c, 0x8d, 0x5e, 0x2b, 0xcf, 0xcc, 0x25,
	0xec, 0x37, 0x98, 0x7b, 0x14, 0x5b, 0x60, 0x0c, 0x26, 0xdd, 0xe9, 0x03, 0xb4, 0x25, 0x0f, 0x1d,
	0xd7, 0xdd, 0x72, 0x7e, 0xb2, 0x7f, 0xf5, 0xe3, 0xe1, 0x9e, 0xdb, 0x6b, 0xac, 0x2c, 0xb1, 0x79,
	0x0a, 0x0f, 0x85, 0x64, 0xdc, 0x8f, 0x88, 0x57, 0x28, 0xa0, 0x44, 0x58, 0xad, 0x71, 0x7b, 0xd2,
	0x9d, 0xde, 0x47, 0x5a, 0x1b, 0xba, 0x28, 0x37, 0xce, 0x8b, 0x5e, 0xe1, 0x0c, 0x44, 0x63, 0x48,
	0x89, 0x30, 0x5f, 0x40, 0x43, 0x73, 0xb7, 0xda, 0x05, 0x9d, 0x23, 0x54, 0xeb, 0x42, 0xc7, 0x69,
	0xba, 0xa2, 0x04, 0x6b, 0x2d, 0x0a, 0xa5, 0x3a, 0x31, 0x9f, 0xc2, 0x2e, 0xe6, 0x2c, 0x2d, 0x59,
	0xac, 0xad, 0xfd, 0x31, 0x98, 0xf4, 0xa7, 0xf7, 0x6a, 0x0e, 0x73, 0xce, 0xd2, 0x92, 0x80, 0x0b,
	0x71, 0x55, 0xdb, 0xaf, 0xa0, 0x51, 0xd9, 0xf2, 0x12, 0x76, 0x34, 0x9c, 0xb0, 0x40, 0x21, 0x62,
	0x88, 0x74, 0xb0, 0xd0, 0x75, 0x17, 0x15, 0x83, 0xfa, 0xc4, 0xfe, 0x04, 0x60, 0xff, 0x42, 0xfa,
	0x11, 0xc1, 0x15, 0xe4, 0x23, 0xd8, 0x0b, 0x33, 0xc9, 0x3e, 0x10, 0xee, 0x25, 0x7e, 0xc2, 0x44,
	0x61, 0x74, 0xdb, 0xbd, 0xad, 0x86, 0x67, 0xf9, 0xcc, 0x1c, 0x41, 0x28, 0x59, 0x1c, 0x08, 0xc9,
	0x12, 0x82, 0xad, 0xd6, 0x18, 0x4c, 0x0c, 0xb7, 0x31, 0x31, 0x9f, 0x40, 0x43, 0xc7, 0x5d, 0x39,
	0x63, 0xd6, 0xb4, 0xae, 0xd1, 0xa9, 0x36, 0xed, 0x37, 0x70, 0xb0, 0x4d, 0x46, 0x98, 0xcf, 0x61,
	0x47, 0x3f, 0x6b, 0x81, 0x56, 0x8d, 0xb4, 0xbd, 0xad, 0xe5, 0x55, 0x07, 0xf6, 0xb7, 0x16, 0x1c,
	0x2c, 0x18, 0xff, 0xe8, 0x73, 0xfc, 0x3f, 0x92, 0x34, 0x87, 0xfd, 0xad, 0x24, 0xad, 0x0b, 0x27,
	0xfe, 0x98, 0xa3, 0x5e, 0x33, 0x47, 0xeb, 0x7f, 0x4d, 0xd1, 0x11, 0xec, 0x08, 0x96, 0xf1, 0x90,
	0xe4, 0x52, 0xf2, 0x0c, 0xf5, 0x5c, 0xa3, 0x1c, 0x2c, 0xb1, 0x89, 0xe0, 0xdd, 0x24, 0x8b, 0xbd,
	0xcb, 0xd2, 0x03, 0x82, 0x3d, 0x49, 0x63, 0x22, 0xac, 0x83, 0x31, 0x98, 0x1c, 0xb8, 0x77, 0x92,
	0x2c, 0x5e, 0xe8, 0x97, 0xb7, 0xf9, 0x83, 0xfd, 0x15, 0xc0, 0x5e, 0x5e, 0xdd, 0x5c, 0xbb, 0x4e,
	0x96, 0x57, 0x9b, 0x11, 0xf8, 0xbe, 0x19, 0x81, 0x9f, 0x9b, 0x11, 0xf8, 0xfc, 0x6b, 0xb4, 0xf7,
	0xee, 0xd9, 0x5f, 0xfe, 0xc8, 0xc1, 0xad, 0xa2, 0x9f, 0xfd, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x29,
	0x6f, 0x13, 0xd0, 0xd3, 0x05, 0x00, 0x00,
}
