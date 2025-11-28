package logging

import (
	"github.com/Mohamadreza-shad/auth/pkg/logging/keyval"
	"go.uber.org/fx/fxevent"
)

type FxZapLogger struct {
	log Logger
}

func (f *FxZapLogger) LogEvent(e fxevent.Event) {

	switch ev := e.(type) {
	case *fxevent.OnStartExecuting:
		f.log.Info("OnStart hook executing",
			keyval.String("callee", ev.FunctionName),
			keyval.String("caller", ev.CallerName),
		)

	case *fxevent.OnStartExecuted:
		if ev.Err != nil {
			f.log.Error("OnStart hook failed",
				keyval.String("callee", ev.FunctionName),
				keyval.String("caller", ev.CallerName),
				keyval.Error(ev.Err),
			)
		} else {
			f.log.Info("OnStart hook executed",
				keyval.String("callee", ev.FunctionName),
				keyval.String("caller", ev.CallerName),
			)
		}

	case *fxevent.OnStopExecuting:
		f.log.Info("OnStop hook executing",
			keyval.String("callee", ev.FunctionName),
			keyval.String("caller", ev.CallerName),
		)

	case *fxevent.OnStopExecuted:
		if ev.Err != nil {
			f.log.Error("OnStop hook failed",
				keyval.String("callee", ev.FunctionName),
				keyval.String("caller", ev.CallerName),
				keyval.Error(ev.Err),
			)
		} else {
			f.log.Info("OnStop hook executed",
				keyval.String("callee", ev.FunctionName),
				keyval.String("caller", ev.CallerName),
			)
		}

	case *fxevent.Supplied:
		if ev.Err != nil {
			f.log.Error("Failed to supply constructor",
				keyval.String("type", ev.TypeName),
				keyval.Error(ev.Err),
			)
		} else {
			f.log.Info("Supplied",
				keyval.String("type", ev.TypeName),
			)
		}

	case *fxevent.Provided:
		for _, rtype := range ev.OutputTypeNames {
			f.log.Info("Provided",
				keyval.String("constructor", ev.ConstructorName),
				keyval.String("type", rtype),
			)
		}

	case *fxevent.Invoking:
		f.log.Info("Invoking",
			keyval.String("function", ev.FunctionName),
		)

	case *fxevent.Invoked:
		if ev.Err != nil {
			f.log.Error("Invoke failed",
				keyval.String("function", ev.FunctionName),
				keyval.Error(ev.Err),
			)
		}

	case *fxevent.Started:
		if ev.Err != nil {
			f.log.Error("Failed to start application",
				keyval.Error(ev.Err),
			)
		} else {
			f.log.Info("Application started")
		}

	case *fxevent.Stopped:
		if ev.Err != nil {
			f.log.Error("Application stopped with error",
				keyval.Error(ev.Err),
			)
		} else {
			f.log.Info("Application stopped cleanly")
		}

	case *fxevent.RollingBack:
		f.log.Error("Start failed, rolling back",
			keyval.Error(ev.StartErr),
		)

	case *fxevent.RolledBack:
		if ev.Err != nil {
			f.log.Error("Rollback failed",
				keyval.Error(ev.Err),
			)
		} else {
			f.log.Info("Rollback complete")
		}

	case *fxevent.LoggerInitialized:
		f.log.Info("Logger initialized",
			keyval.String("name", ev.ConstructorName),
		)

	default:
		f.log.Info("fx event", keyval.Any("event", e))

	}
}

func NewFxLogger(logger Logger) fxevent.Logger {

	return &FxZapLogger{log: logger}
}
