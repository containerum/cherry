package models

import (
	"fmt"
	"strings"
	"unicode"

	"git.containerum.net/ch/cherry"
	"git.containerum.net/ch/cherry/pkg/noicerrs"
)

func extractDefinedKinds(errs []TOMLerror) (map[cherry.ErrKind]TOMLerror, error) {
	defined := map[cherry.ErrKind]TOMLerror{}
	for _, tomlerr := range errs {
		if tomlerr.Kind != 0 {
			if conflict, alreadyDefined := defined[tomlerr.Kind]; alreadyDefined {
				return nil, fmt.Errorf("kind %d of %q conflict with kind of %q",
					tomlerr.Kind,
					tomlerr.Name,
					conflict.Name)
			}
			defined[tomlerr.Kind] = tomlerr
		}
	}
	return defined, nil
}

func findConfictingKinds(errs []TOMLerror) error {
	defined := map[cherry.ErrKind]TOMLerror{}
	for _, tomlerr := range errs {
		if tomlerr.Kind != 0 {
			if conflict, alreadyDefined := defined[tomlerr.Kind]; alreadyDefined {
				return noicerrs.ErrConflictingKinds().
					AddDetailF("kind %d of %q conflict with kind of %q",
						tomlerr.Kind,
						tomlerr.Name,
						conflict.Name)
			}
			defined[tomlerr.Kind] = tomlerr
		}
	}
	return nil
}

func calculateFreeKinds(errNum uint64, errs map[cherry.ErrKind]TOMLerror) []cherry.ErrKind {
	kinds := []cherry.ErrKind{}
	for i := cherry.ErrKind(1); i <= cherry.ErrKind(errNum); i++ {
		if _, isOccupied := errs[i]; !isOccupied {
			kinds = append(kinds, i)
		}
	}
	return kinds
}

func cleanPackageName(name string) string {
	fields := strings.FieldsFunc(name, func(r rune) bool {
		return unicode.IsSpace(r) || unicode.IsPunct(r)
	})
	cleanedTokens := make([]string, 0, len(fields))
	for i, field := range fields {
		field = strings.TrimSpace(field)
		if field != "" {
			if i != 0 {
				field = strings.Title(field)
			}
			cleanedTokens = append(cleanedTokens, field)
		}
	}
	return strings.Join(cleanedTokens, "")
}
