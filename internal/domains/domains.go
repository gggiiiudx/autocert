package domains

import (
	"certimate/internal/utils/app"
	"context"
	"fmt"

	"github.com/pocketbase/pocketbase/models"
)

func create(ctx context.Context, record *models.Record) error {
	if !record.GetBool("enabled") {
		return nil
	}

	if record.GetBool("rightnow") {
		defer func() {
			setRightnow(ctx, record, false)
		}()
		if err := deploy(ctx, record); err != nil {
			return err
		}
	}

	scheduler := app.GetScheduler()

	err := scheduler.Add(record.Id, record.GetString("crontab"), func() {
		deploy(ctx, record)
	})

	if err != nil {
		app.GetApp().Logger().Error("add cron job failed", "err", err)
		return fmt.Errorf("add cron job failed: %w", err)
	}
	app.GetApp().Logger().Error("add cron job failed", "domain", record.GetString("domain"))

	scheduler.Start()
	return nil
}

func update(ctx context.Context, record *models.Record) error {
	scheduler := app.GetScheduler()
	scheduler.Remove(record.Id)
	if !record.GetBool("enabled") {
		return nil
	}

	if record.GetBool("rightnow") {
		defer func() {
			setRightnow(ctx, record, false)
		}()
		if err := deploy(ctx, record); err != nil {
			return err
		}
	}

	err := scheduler.Add(record.Id, record.GetString("crontab"), func() {
		deploy(ctx, record)
	})

	if err != nil {
		app.GetApp().Logger().Error("update cron job failed", "err", err)
		return fmt.Errorf("update cron job failed: %w", err)
	}
	app.GetApp().Logger().Error("update cron job failed", "domain", record.GetString("domain"))

	scheduler.Start()
	return nil
}

func delete(_ context.Context, record *models.Record) error {
	scheduler := app.GetScheduler()

	scheduler.Remove(record.Id)
	scheduler.Start()
	return nil
}

func setRightnow(ctx context.Context, record *models.Record, ok bool) error {
	record.Set("rightnow", ok)
	return app.GetApp().Dao().SaveRecord(record)
}