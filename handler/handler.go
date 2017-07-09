package handler

import (
	"github.com/corpix/queues/message"
)

type Handler func(message.Message)
