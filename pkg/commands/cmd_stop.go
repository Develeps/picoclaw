package commands

import (
	"context"
)

func stopCommand() Definition {
	return Definition{
		Name:        "stop",
		Description: "Stop the current LLM generation",
		Usage:       "/stop",
		Handler: func(_ context.Context, req Request, rt *Runtime) error {
			if rt == nil || rt.CancelTurn == nil {
				return req.Reply(unavailableMsg)
			}

			// Build session key from channel and chat ID
			sessionKey := req.Channel + ":" + req.ChatID

			// Try to cancel the active turn
			if rt.CancelTurn(sessionKey) {
				return req.Reply("✓ Генерация остановлена")
			}

			// No active turn found
			return req.Reply("Нет активной генерации для остановки")
		},
	}
}