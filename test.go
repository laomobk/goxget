package main

func TestListRepo() {
	repos := listAllPkgConfig()

	for _, r := range repos {
		println(r.pkgName)
	}
}
