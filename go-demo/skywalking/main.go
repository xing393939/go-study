package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"skywalking.apache.org/repo/goapi/collect/language/agent/v3"
	"time"
)

func meter(cc grpc.ClientConnInterface, md metadata.MD) {
	serviceClient := v3.NewMeterReportServiceClient(cc)
	stream, err := serviceClient.Collect(metadata.NewOutgoingContext(context.Background(), md))
	if err != nil {
		return
	}
	histogram := v3.MeterData_Histogram{
		Histogram: &v3.MeterHistogram{
			Name:   "alloc_MB",
			Labels: []*v3.Label{},
			Values: []*v3.MeterBucketValue{
				{
					Bucket:             0,
					Count:              26,
					IsNegativeInfinity: true,
				},
			},
		},
	}
	err = stream.Send(&v3.MeterData{
		Timestamp:       time.Now().UnixNano() / 1e6,
		Service:         "op_common",
		ServiceInstance: "192.168.2.119",
		Metric:          &histogram,
	})
	println(err)
}

func trace(cc grpc.ClientConnInterface, md metadata.MD) {
	serviceClient := v3.NewTraceSegmentReportServiceClient(cc)
	stream, err := serviceClient.Collect(metadata.NewOutgoingContext(context.Background(), md))
	if err != nil {
		return
	}

	span := v3.SpanObject{
		SpanId:        0,
		ParentSpanId:  -1,
		StartTime:     time.Now().UnixNano() / 1e6,
		EndTime:       time.Now().UnixNano() / 1e6,
		OperationName: "/test2",
		SpanType:      v3.SpanType_Entry,
		SpanLayer:     v3.SpanLayer_Http,
		ComponentId:   5006,
		Refs:          []*v3.SegmentReference{},
	}
	err = stream.Send(&v3.SegmentObject{
		TraceId:         "893fb640c39511eca1a864006a6e3730",
		TraceSegmentId:  "893fb640c39511eca1a864006a6e3730",
		IsSizeLimited:   false,
		Service:         "my_service",
		ServiceInstance: "5dd8cb57a2b111eca17b00163e050534@8.8.8.8",
		Spans: []*v3.SpanObject{
			&span,
		},
	})
	println(err)
	_ = stream.CloseSend()
}

func main() {
	conn, err := grpc.Dial("8.8.8.8:11800", grpc.WithInsecure())
	if err != nil {
		return
	}
	md := metadata.New(map[string]string{})
	meter(conn, md)
	trace(conn, md)
	time.Sleep(time.Second)
}
