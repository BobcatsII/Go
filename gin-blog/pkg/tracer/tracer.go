package tracer

import (
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"time"
)

func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	//config.Configuration 是 jaeger client 的配置项，设置基本信息
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{ //Sampler 固定采样、对所有数据都采样
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{ //是否启用LoggingReporter、刷新缓冲区的频率、上报的Agent地址
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentHostPort,
		},
	}
	tracer, closer, err := cfg.NewTracer() //cfg.NewTracer 初始化 Tracer 对象，此处返回的是 opentracing.Tracer
	if err != nil {
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer) //设置全局tracer对象，根据实际情况设置即可。
	return tracer, closer, nil
}