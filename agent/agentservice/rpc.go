package agentservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cylScripter/openapi/agent"
)

func Chat(ctx context.Context, req *agent.ChatReq, callOptions ...client.Option) (resp *agent.ChatResp, err error) {
	clients := MustNewClient("agent", callOptions...)
	return clients.Chat(ctx, req)
}

func Health(ctx context.Context, req *agent.HealthReq, callOptions ...client.Option) (resp *agent.HealthResp, err error) {
	clients := MustNewClient("agent", callOptions...)
	return clients.Health(ctx, req)
}

func CreateSession(ctx context.Context, req *agent.CreateSessionReq, callOptions ...client.Option) (resp *agent.CreateSessionResp, err error) {
	clients := MustNewClient("agent", callOptions...)
	return clients.CreateSession(ctx, req)
}

func GetSession(ctx context.Context, req *agent.GetSessionReq, callOptions ...client.Option) (resp *agent.GetSessionResp, err error) {
	clients := MustNewClient("agent", callOptions...)
	return clients.GetSession(ctx, req)
}

func ListSession(ctx context.Context, req *agent.ListSessionReq, callOptions ...client.Option) (resp *agent.ListSessionResp, err error) {
	clients := MustNewClient("agent", callOptions...)
	return clients.ListSession(ctx, req)
}

func UpdateSession(ctx context.Context, req *agent.UpdateSessionReq, callOptions ...client.Option) (resp *agent.UpdateSessionResp, err error) {
	clients := MustNewClient("agent", callOptions...)
	return clients.UpdateSession(ctx, req)
}

func DeleteSession(ctx context.Context, req *agent.DeleteSessionReq, callOptions ...client.Option) (resp *agent.DeleteSessionResp, err error) {
	clients := MustNewClient("agent", callOptions...)
	return clients.DeleteSession(ctx, req)
}

func ListSessionMessages(ctx context.Context, req *agent.ListSessionMessagesReq, callOptions ...client.Option) (resp *agent.ListSessionMessagesResp, err error) {
	clients := MustNewClient("agent", callOptions...)
	return clients.ListSessionMessages(ctx, req)
}
