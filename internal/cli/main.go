package cli

import (
    "context"
    "github.com/alecthomas/kingpin"
    "gitlab.com/distributed_lab/kit/kv"
    "gitlab.com/distributed_lab/logan/v3"
    "gitlab.com/tokend/subgroup/tokenproject/internal/config"
    "gitlab.com/tokend/subgroup/tokenproject/internal/service"
    token "gitlab.com/tokend/subgroup/tokenproject/internal/service/token-svc/token_pools"
)

func Run(args []string) bool {
    log := logan.New()

    defer func() {
        if rvr := recover(); rvr != nil {
            log.WithRecover(rvr).Error("app panicked")
        }
    }()

    cfg := config.New(kv.MustFromEnv())
    log = cfg.Log()

    app := kingpin.New("tokenproject", "")

    runCmd := app.Command("run", "run command")
    apiCmd := runCmd.Command("api", "run api")
    serviceCmd := runCmd.Command("service", "run service") // you can insert custom help

    //tokenCmd := runCmd.Command("service", "run service token") // you can insert custom help

    migrateCmd := app.Command("migrate", "migrate command")
    migrateUpCmd := migrateCmd.Command("up", "migrate db up")
    migrateDownCmd := migrateCmd.Command("down", "migrate db down")

    ctx := context.Background()
    // custom commands go here...

    cmd, err := app.Parse(args[1:])
    if err != nil {
        log.WithError(err).Error("failed to parse arguments")
        return false
    }

    switch cmd {
    case apiCmd.FullCommand():
        service.Run(cfg)
        return true
    case serviceCmd.FullCommand():
        svc := token.New(cfg)
        svc.Run(cfg, ctx)
        return true
    case migrateUpCmd.FullCommand():
        err = MigrateUp(cfg)
    case migrateDownCmd.FullCommand():
        err = MigrateDown(cfg)
    // handle any custom commands here in the same way
    default:
        log.Errorf("unknown command %s", cmd)
        return false
    }
    if err != nil {
        log.WithError(err).Error("failed to exec cmd")
        return false
    }
    return true
}
