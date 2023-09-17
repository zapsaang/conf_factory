package surge

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/zapsaang/conf_factory/conf/conf"
	"github.com/zapsaang/conf_factory/pkg/logs"
	"github.com/zapsaang/conf_factory/pkg/ordered"
	"github.com/zapsaang/conf_factory/repo"
	"github.com/zapsaang/conf_factory/utils"
)

type Config struct {
	Platform string
	Password string
}

type WorkSpace struct {
	Config
	BaseName    string
	surgeConfig *ordered.Map[string, *ordered.Map[string, any]]
	surgeTmp    *ordered.Map[string, *ordered.Map[string, any]]
	// TODO change to surge repo
	repoHandler repo.Repo
}

type ProcessNode func(ctx context.Context, ws *WorkSpace) error

func New(baseName string, config Config) *WorkSpace {
	return &WorkSpace{
		Config:      config,
		repoHandler: repo.New(repo.RoleIntegrator),
		surgeConfig: ordered.NewMap[string, *ordered.Map[string, any]](),
		surgeTmp:    ordered.NewMap[string, *ordered.Map[string, any]](),
		BaseName:    baseName,
	}
}

func Load(ctx context.Context, ws *WorkSpace) error {
	// TODO more log
	logger := logs.WithContext(ctx).
		WithField("base_name", ws.BaseName)
	content, err := ws.repoHandler.Get(ctx, fmt.Sprintf("%s/%s", conf.GetSurgeBaseDir(), ws.BaseName))
	if err != nil {
		logger.WithError(err).Error("get content failed")
		return err
	}
	lines := strings.Split(string(content), "\n")

	var currentSection string
	for _, line := range lines {
		_line := strings.TrimSpace(line)
		if utils.IsComment(_line) {
			continue
		}
		if IsSectionTitle(_line) {
			currentSection = _line
			ws.surgeTmp.LoadOrStore(currentSection, ordered.NewMap[string, any]())
		}
		sectionContent, _ := ws.surgeTmp.Load(currentSection)
		sectionContent.Store(_line, struct{}{})
	}
	// TODO sort surgeTmp
	return nil
}

func Integrate(ctx context.Context, ws *WorkSpace) error {
	ws.surgeTmp.Range(func(key string, value *ordered.Map[string, any]) bool {
		title, parser := GetSectionParser(key)
		ws.surgeConfig.LoadOrStore(title, ordered.NewMap[string, any]())
		var lines = make([]string, 0, value.Len())
		value.Range(func(key string, value any) bool {
			lines = append(lines, key)
			return true
		})
		parser(ws, key, lines)
		return true
	})
	return nil
}

func Filte(ctx context.Context, ws *WorkSpace) error {
	// TODO more log
	// TODO filte by feature which platform excludes
	ws.surgeConfig.Range(func(key string, value *ordered.Map[string, any]) bool {

		return true
	})
	return nil
}

func Store(ctx context.Context, ws *WorkSpace) error {
	// TODO more log
	// TODO optimize by calculating length ahead
	var storeBuf bytes.Buffer
	ws.surgeConfig.Range(func(key string, value *ordered.Map[string, any]) bool {
		storer, ok := GetSectionStorer(key)
		if !ok {
			return true
		}
		storeBuf.WriteString(key)
		storeBuf.WriteRune('\n')
		storeBuf.Write(storer(value))
		return true
	})
	if err := ws.repoHandler.Set(ctx, fmt.Sprintf("%s/%s", conf.GetSurgeResultDir(), ws.BaseName), storeBuf.Bytes()); err != nil {
		logs.WithField("base_name", ws.BaseName).WithError(err).Error("store failed")
		return err
	}
	return nil
}
