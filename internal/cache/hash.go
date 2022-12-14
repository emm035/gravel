package cache

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/emm035/gravel/internal/gravel"
	"github.com/emm035/gravel/internal/resolve"
	"github.com/emm035/gravel/internal/semver"
	"github.com/emm035/gravel/internal/types"
)

var emptyCache = resolve.CacheFile{
	Packages: nil,
	Versions: nil,
}

func NewHashes(graph types.Graph[resolve.Pkg], paths gravel.Paths, ignoreOld bool) (resolve.Hashes, error) {
	oldHashes, err := loadHashes(paths, ignoreOld)
	if err != nil {
		return resolve.Hashes{}, err
	}

	newHashes, err := computeHashes(graph, paths)
	if err != nil {
		return resolve.Hashes{}, err
	}

	return resolve.Hashes{
		Old: oldHashes,
		New: newHashes,
	}, nil
}

func loadHashes(paths gravel.Paths, fakeLoad bool) (*resolve.CacheFile, error) {
	if fakeLoad {
		return &emptyCache, nil
	}

	data, err := os.ReadFile(paths.HashesFile)
	if os.IsNotExist(err) {
		return &emptyCache, nil
	} else if err != nil {
		return nil, err
	}

	cache := new(resolve.CacheFile)
	if err := json.Unmarshal(data, &cache); err != nil {
		return nil, err
	}

	return cache, nil
}

var (
	ErrPkgDirNotFound = errors.New("package directory not found")
)

func computeHashes(graph types.Graph[resolve.Pkg], paths gravel.Paths) (*resolve.CacheFile, error) {
	cacheFile := &resolve.CacheFile{
		Packages: make(map[string]string),
		Versions: make(map[string]semver.Version),
	}

	for pkg := range graph.Nodes() {
		hash, err := pkg.Hash()
		if err != nil {
			return nil, err
		}
		cacheFile.Packages[pkg.PkgPath] = hash
		cacheFile.Versions[pkg.PkgPath] = resolve.Version(pkg).Version
	}
	return cacheFile, nil
}
