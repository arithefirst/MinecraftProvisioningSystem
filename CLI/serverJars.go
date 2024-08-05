package main

import (
	"strings"
)

func retriveServerJar(platform string, ver string) {
	switch strings.ToLower(platform) {
	case "vanilla":
		// Create maps to store the URLs of each server type jarfile
		vanillaJars := make(map[string]string)

		// Vanilla 1.12
		vanillaJars["1.12"] = "https://piston-data.mojang.com/v1/objects/8494e844e911ea0d63878f64da9dcc21f53a3463/server.jar"
		vanillaJars["1.12.1"] = "https://piston-data.mojang.com/v1/objects/561c7b2d54bae80cc06b05d950633a9ac95da816/server.jar"
		vanillaJars["1.12.2"] = "https://piston-data.mojang.com/v1/objects/886945bfb2b978778c3a0288fd7fab09d315b25f/server.jar"

		// Vanilla 1.13
		vanillaJars["1.13"] = "https://piston-data.mojang.com/v1/objects/d0caafb8438ebd206f99930cfaecfa6c9a13dca0/server.jar"
		vanillaJars["1.13.1"] = "https://piston-data.mojang.com/v1/objects/fe123682e9cb30031eae351764f653500b7396c9/server.jar"
		vanillaJars["1.13.2"] = "https://piston-data.mojang.com/v1/objects/3737db93722a9e39eeada7c27e7aca28b144ffa7/server.jar"

		// Vanilla 1.14
		vanillaJars["1.14"] = "https://piston-data.mojang.com/v1/objects/f1a0073671057f01aa843443fef34330281333ce/server.jar"
		vanillaJars["1.14.1"] = "https://piston-data.mojang.com/v1/objects/ed76d597a44c5266be2a7fcd77a8270f1f0bc118/server.jar"
		vanillaJars["1.14.2"] = "https://piston-data.mojang.com/v1/objects/808be3869e2ca6b62378f9f4b33c946621620019/server.jar"
		vanillaJars["1.14.3"] = "https://piston-data.mojang.com/v1/objects/d0d0fe2b1dc6ab4c65554cb734270872b72dadd6/server.jar"
		vanillaJars["1.14.4"] = "https://piston-data.mojang.com/v1/objects/3dc3d84a581f14691199cf6831b71ed1296a9fdf/server.jar"

		// Vanilla 1.15
		vanillaJars["1.15"] = "https://piston-data.mojang.com/v1/objects/e9f105b3c5c7e85c7b445249a93362a22f62442d/server.jar"
		vanillaJars["1.15.1"] = "https://piston-data.mojang.com/v1/objects/4d1826eebac84847c71a77f9349cc22afd0cf0a1/server.jar"
		vanillaJars["1.15.2"] = "https://piston-data.mojang.com/v1/objects/bb2b6b1aefcd70dfd1892149ac3a215f6c636b07/server.jar"

		// Vanilla 1.16
		vanillaJars["1.16"] = "https://piston-data.mojang.com/v1/objects/a0d03225615ba897619220e256a266cb33a44b6b/server.jar"
		vanillaJars["1.16.1"] = "https://piston-data.mojang.com/v1/objects/a412fd69db1f81db3f511c1463fd304675244077/server.jar"
		vanillaJars["1.16.2"] = "https://piston-data.mojang.com/v1/objects/c5f6fb23c3876461d46ec380421e42b289789530/server.jar"
		vanillaJars["1.16.3"] = "https://piston-data.mojang.com/v1/objects/f02f4473dbf152c23d7d484952121db0b36698cb/server.jar"
		vanillaJars["1.16.4"] = "https://piston-data.mojang.com/v1/objects/35139deedbd5182953cf1caa23835da59ca3d7cd/server.jar"
		vanillaJars["1.16.5"] = "https://piston-data.mojang.com/v1/objects/1b557e7b033b583cd9f66746b7a9ab1ec1673ced/server.jar"

		// Vanilla 1.17
		vanillaJars["1.17"] = "https://piston-data.mojang.com/v1/objects/0a269b5f2c5b93b1712d0f5dc43b6182b9ab254e/server.jar"
		vanillaJars["1.17.1"] = "https://piston-data.mojang.com/v1/objects/a16d67e5807f57fc4e550299cf20226194497dc2/server.jar"

		// Vanilla 1.18
		vanillaJars["1.18"] = "https://piston-data.mojang.com/v1/objects/3cf24a8694aca6267883b17d934efacc5e44440d/server.jar"
		vanillaJars["1.18.1"] = "https://piston-data.mojang.com/v1/objects/125e5adf40c659fd3bce3e66e67a16bb49ecc1b9/server.jar"
		vanillaJars["1.18.2"] = "https://piston-data.mojang.com/v1/objects/c8f83c5655308435b3dcf03c06d9fe8740a77469/server.jar"

		// Vanilla 1.19
		vanillaJars["1.19"] = "https://piston-data.mojang.com/v1/objects/e00c4052dac1d59a1188b2aa9d5a87113aaf1122/server.jar"
		vanillaJars["1.19.1"] = "https://piston-data.mojang.com/v1/objects/8399e1211e95faa421c1507b322dbeae86d604df/server.jar"
		vanillaJars["1.19.2"] = "https://piston-data.mojang.com/v1/objects/f69c284232d7c7580bd89a5a4931c3581eae1378/server.jar"
		vanillaJars["1.19.3"] = "https://piston-data.mojang.com/v1/objects/c9df48efed58511cdd0213c56b9013a7b5c9ac1f/server.jar"
		vanillaJars["1.19.4"] = "https://piston-data.mojang.com/v1/objects/8f3112a1049751cc472ec13e397eade5336ca7ae/server.jar"

		// Vanilla 1.20
		vanillaJars["1.20"] = "https://piston-data.mojang.com/v1/objects/15c777e2cfe0556eef19aab534b186c0c6f277e1/server.jar"
		vanillaJars["1.20.1"] = "https://piston-data.mojang.com/v1/objects/84194a2f286ef7c14ed7ce0090dba59902951553/server.jar"
		vanillaJars["1.20.2"] = "https://piston-data.mojang.com/v1/objects/5b868151bd02b41319f54c8d4061b8cae84e665c/server.jar"
		vanillaJars["1.20.3"] = "https://piston-data.mojang.com/v1/objects/4fb536bfd4a83d61cdbaf684b8d311e66e7d4c49/server.jar"
		vanillaJars["1.20.4"] = "https://piston-data.mojang.com/v1/objects/8dd1a28015f51b1803213892b50b7b4fc76e594d/server.jar"
		vanillaJars["1.20.5"] = "https://piston-data.mojang.com/v1/objects/79493072f65e17243fd36a699c9a96b4381feb91/server.jar"
		vanillaJars["1.20.6"] = "https://piston-data.mojang.com/v1/objects/145ff0858209bcfc164859ba735d4199aafa1eea/server.jar"

		// Vanilla 1.21
		vanillaJars["1.21"] = "https://piston-data.mojang.com/v1/objects/450698d1863ab5180c25d7c804ef0fe6369dd1ba/server.jar"

		// Return the url for the requested version
		err := DownloadFile(vanillaJars[ver], "server.jar")
		if err != nil {
			panic(err)
		}

	case "forge":
	case "fabric":
	case "spigot":
	}
}
