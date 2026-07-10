package fslib

import "os"

func ReadDir(path string) ([]map[string]any, error) {
	nodes := []map[string]any{}
	dirs, err := os.ReadDir(path)
	if err != nil {
		return nodes, err
	}

	for _, dir := range dirs {
		info, err := dir.Info()
		if err != nil {
			return nodes, err
		}

		nodes = append(nodes, map[string]any{
			"name":             info.Name(),
			"size":             info.Size(),
			"isDir":            info.IsDir,
			"isFile":           func() bool { return !info.IsDir() },
			"modificationUnix": info.ModTime().Unix(),
		})
	}

	return nodes, nil
}
