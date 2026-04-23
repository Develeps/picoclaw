package commands

import (
	"context"

	"github.com/sipeed/picoclaw/pkg/logger"
)

func stopCommand() Definition {
	return Definition{
		Name:        "stop",
		Description: "Stop the current LLM generation",
		Usage:       "/stop",
		Handler: func(_ context.Context, req Request, rt *Runtime) error {
			if rt == nil || rt.CancelTurn == nil {
				logger.WarnCF("commands", "/stop command called but CancelTurn is not available",
					map[string]any{
						"channel":   req.Channel,
						"chat_id":   req.ChatID,
						"sender_id": req.SenderID,
					})
				return req.Reply(unavailableMsg)
			}

			// Build session key from channel and chat ID
			sessionKey := req.Channel + ":" + req.ChatID

			logger.InfoCF("commands", "/stop command invoked",
				map[string]any{
					"session_key": sessionKey,
					"channel":     req.Channel,
					"chat_id":     req.ChatID,
					"sender_id":   req.SenderID,
				})

			// Try to cancel the active turn
			if rt.CancelTurn(sessionKey) {
				logger.InfoCF("commands", "/stop command successfully canceled turn",
					map[string]any{
						"session_key": sessionKey,
					})
				return req.Reply("✓ Генерация остановлена")
			}

			// No active turn found
			logger.InfoCF("commands", "/stop command: no active turn to cancel",
				map[string]any{
					"session_key": sessionKey,
				})
			return req.Reply("Нет активной генерации для остановки")
		},
	}
}
