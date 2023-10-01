package common

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	tracer "go.opentelemetry.io/otel/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func NewExporter(ctx context.Context, client otlptrace.Client) (*otlptrace.Exporter, error) {
	return otlptrace.New(ctx, client)
}

func NewTraceProvider(ctx context.Context) (*trace.TracerProvider, error) {
	// // Create client: Jaeger
	// client := otlptracehttp.NewClient(otlptracehttp.WithEndpoint("localhost:4318"),
	// 									otlptracehttp.WithInsecure())

	// Create exporter to Jaeger
	jgExp, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint("localhost:4318"),
									otlptracehttp.WithInsecure())
	if err != nil {
		panic(err)
	}

	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("bcon-service"),
		),
	)
	if err != nil {
		panic(err)
	}

	return trace.NewTracerProvider(
		trace.WithBatcher(jgExp),
		trace.WithResource(r),
	), nil
}

func NewTracer(ctx context.Context, name string) (tracer.Tracer, error) {
	tp := otel.GetTracerProvider()
	return tp.Tracer(name), nil
}