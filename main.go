package goxget

func main() {
	for _, r := range listAllPkgConfig() {
		println("repo.name =", r.pkgName)
	}
}
