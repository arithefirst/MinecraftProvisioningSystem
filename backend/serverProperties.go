package main

import "fmt"

type serverProperties struct {
	allowFlight                 bool   `schema:"allowFlight"`
	allowNether                 bool   `schema:"allowNether"`
	broadcastConsoleOps         bool   `schema:"broadcastConsoleOps"`
	broadcastRconOps            bool   `schema:"broadcastRconOps"`
	difficulty                  string `schema:"difficulty"`
	enableCommandBlock          bool   `schema:"enableCommandBlock"`
	enableJmxMonitoring         bool   `schema:"enableJmxMonitoring"`
	enableRcon                  bool   `schema:"enableRcon"`
	enableStatus                bool   `schema:"enableStatus"`
	enableQuery                 bool   `schema:"enableQuery"`
	enforceSecureProfile        bool   `schema:"enforceSecureProfile"`
	enforceWhitelist            bool   `schema:"enforceWhitelist"`
	entityBroadcastRangePercent int    `schema:"entityBroadcastRangePercent"` // Range 10-1000
	forceGamemode               bool   `schema:"forceGamemode"`
	funcPermLevel               int    `schema:"funcPermLevel"` // Range 1-4
	gamemode                    int    `schema:"gamemode"`
	generateStructures          bool   `schema:"generateStructures"`
	hardcore                    bool   `schema:"hardcore"`
	hideOnlinePlayers           bool   `schema:"hideOnlinePlayers"`
	initalDisabledPacks         string `schema:"initalDisabledPacks"`
	initalEnabledPacks          string `schema:"initalEnabledPacks"`
	levelName                   string `schema:"levelName"`
	levelSeed                   string `schema:"levelSeed"`
	maxChainedNeighborUpdates   int    `schema:"maxChainedNeighborUpdates"`
	maxPlayers                  int    `schema:"maxPlayers"`   // Range 0-2147483647
	maxTickTime                 int    `schema:"maxTickTime"`  // -1 to Disable || Range 0–(2^63 - 1)
	maxWorldSize                int    `schema:"maxWorldSize"` // Range 1-29999984
	motd                        string `schema:"motd"`
	networkCompressionThreshold int    `schema:"networkCompressionThreshold"`
	onlineMode                  bool   `schema:"onlineMode"`
	opPermLevel                 int    `schema:"opPermLevel"` // Range 0-4
	playerIdleTimeout           int    `schema:"playerIdleTimeout"`
	preventProxyConnections     bool   `schema:"preventProxyConnections"`
	previewsChat                bool   `schema:"previewsChat"`
	pvp                         bool   `schema:"pvp"`
	queryPort                   int    `schema:"queryPort"` // Valid port numbers only
	rateLimit                   int    `schema:"rateLimit"`
	rconPassword                string `schema:"rconPassword"`
	rconPort                    int    `schema:"rconPort"`     // Valid port numbers only
	resourcePack                string `schema:"resourcePack"` // URI
	resourcePackPrompt          string `schema:"resourcePackPrompt"`
	resourcePackSha1            string `schema:"resourcePackSha1"`
	requireResourcePack         bool   `schema:"requireResourcePack"`
	serverIP                    string `schema:"serverIP"`
	serverPort                  int    `schema:"serverPort"`  // Valid port numbers only
	simDistance                 int    `schema:"simDistance"` // Range 3-32
	snooperEnabled              bool   `schema:"snooperEnabled"`
	spawnAnimals                bool   `schema:"spawnAnimals"`
	spawnMonsters               bool   `schema:"spawnMonsters"`
	spawnNPCs                   bool   `schema:"spawnNPCs"`
	spawnProtection             int    `schema:"spawnProtection"`
	syncChunkWrites             bool   `schema:"syncChunkWrites"`
	useNativeTransport          bool   `schema:"useNativeTransport"`
	viewDistance                int    `schema:"viewDistance"` // Range 3-32
	whitelist                   bool   `schema:"whitelist"`
}

func generateServerProperties(input serverProperties) string {
	var serverProps string

	// Set allow-flight
	if input.allowFlight {
		serverProps = serverProps + "allow-flight=true\n"
	} else {
		serverProps = serverProps + "allow-flight=false\n"
	}

	// Set allow-nether
	if input.allowNether {
		serverProps = serverProps + "allow-nether=true\n"
	} else {
		serverProps = serverProps + "allow-nether=false\n"
	}

	// Set broadcast-console-to-ops
	if input.broadcastConsoleOps {
		serverProps = serverProps + "broadcast-console-to-ops=true\n"
	} else {
		serverProps = serverProps + "broadcast-console-to-ops=false\n"
	}

	// Set broadcast-rcon-to-ops
	if input.broadcastRconOps {
		serverProps = serverProps + "broadcast-rcon-to-ops=true\n"
	} else {
		serverProps = serverProps + "broadcast-rcon-to-ops=false\n"
	}

	// Check if difficulty is valid and set it
	if input.difficulty == "easy" || input.difficulty == "normal" || input.difficulty == "hard" || input.difficulty == "peaceful" {
		serverProps = serverProps + "difficulty=" + input.difficulty + "\n"
	} else {
		fmt.Println("Difficulty option invalid; Defaulting to easy")
		serverProps = serverProps + "difficulty=easy"
	}

	// Set enable-command-block
	if input.enableCommandBlock {
		serverProps = serverProps + "enable-command-block=true\n"
	} else {
		serverProps = serverProps + "enable-command-block=false\n"
	}

	// Set enable-jmx-monitoring
	if input.enableJmxMonitoring {
		serverProps = serverProps + "enable-jmx-monitoring=true\n"
	} else {
		serverProps = serverProps + "enable-jmx-monitoring=false\n"
	}

	// Set enable-rcon
	if input.enableRcon {
		serverProps = serverProps + "enable-rcon=true\n"
	} else {
		serverProps = serverProps + "enable-rcon=false\n"
	}

	// Set enable-status
	if input.enableStatus {
		serverProps = serverProps + "enable-status=true\n"
	} else {
		serverProps = serverProps + "enable-status=false\n"
	}

	// Set enable-query
	if input.enableQuery {
		serverProps = serverProps + "enable-query=true\n"
	} else {
		serverProps = serverProps + "enable-query=false\n"
	}

	// Set enforce-secure-profile
	if input.enforceSecureProfile {
		serverProps = serverProps + "enforce-secure-profile=true\n"
	} else {
		serverProps = serverProps + "enforce-secure-profile=false\n"
	}

	// Set enforce-whitelist
	if input.enforceWhitelist {
		serverProps = serverProps + "enforce-whitelist=true\n"
	} else {
		serverProps = serverProps + "enforce-whitelist=false\n"
	}

	// Set entity-broadcast-range-percentage
	EBPP := input.entityBroadcastRangePercent
	if EBPP > 1000 || EBPP < 10 {
		fmt.Println("Entity Broadcast Range Percent must be below 1000 and above 10; Defaulting to 100")
		serverProps = serverProps + "entity-broadcast-range-percentage=100\n"
	} else {
		serverProps = serverProps + "entity-broadcast-range-percentage=" + fmt.Sprint(EBPP) + "\n"
	}

	// Set force-gamemode
	if input.forceGamemode {
		serverProps = serverProps + "force-gamemode=true\n"
	} else {
		serverProps = serverProps + "force-gamemode=false\n"
	}

	// Set function-permission-level
	if input.funcPermLevel < 1 || input.funcPermLevel > 4 {
		fmt.Println("Function permission level must be => 1 and <= 5; Defaulting to 2")
		serverProps = serverProps + "function-permission-level=2\n"
	} else {
		serverProps = serverProps + "function-permisison-level=" + fmt.Sprint(input.funcPermLevel) + "\n"
	}

	// Set op-permission-level
	if input.opPermLevel < 1 || input.opPermLevel > 4 {
		fmt.Println("Op permission level must be => 1 and <= 5; Defaulting to 2")
		serverProps = serverProps + "op-permission-level=2\n"
	} else {
		serverProps = serverProps + "op-permission-level=" + fmt.Sprint(input.opPermLevel) + "\n"
	}

	// Set gamemode
	switch input.gamemode {
	case 0:
		serverProps = serverProps + "gamemode=survival\n"
	case 1:
		serverProps = serverProps + "gamemode=creative\n"
	case 2:
		serverProps = serverProps + "gamemode=adventure\n"
	case 3:
		serverProps = serverProps + "gamemode=spectator\n"
	default:
		fmt.Println("Gamemode must be either 1, 2, 3, or 4; Defaulting to survival (0)")
		serverProps = serverProps + "gamemode=survival\n"
	}

	// Set generate-structures
	if input.generateStructures {
		serverProps = serverProps + "generate-structures=true\n"
	} else {
		serverProps = serverProps + "generate-structures=false\n"
	}

	// Set hardcore
	if input.hardcore {
		serverProps = serverProps + "hardcore=true\n"
	} else {
		serverProps = serverProps + "hardcore=false\n"
	}

	// Set hide-online-players
	if input.hideOnlinePlayers {
		serverProps = serverProps + "hide-online-players=true\n"
	} else {
		serverProps = serverProps + "hide-online-players=false\n"
	}

	// Set initial-disabled-packs
	serverProps = serverProps + "initial-disabled-packs=" + input.initalDisabledPacks + "\n"

	// Set initial-enabled-packs
	serverProps = serverProps + "initial-enabled-packs=" + input.initalEnabledPacks + "\n"

	// Set level-name
	serverProps = serverProps + "level-name=" + input.levelName + "\n"

	// Set level-seed
	serverProps = serverProps + "level-seed=" + input.levelSeed + "\n"

	// Set max-chained-neighbor-updates
	serverProps = serverProps + "max-chained-neighbor-updates=" + fmt.Sprint(input.maxChainedNeighborUpdates) + "\n"

	// Set max-players
	if input.maxPlayers < (2 ^ 31 - 1) {
		serverProps = serverProps + "max-players=" + fmt.Sprint(input.maxPlayers) + "\n"
	} else {
		fmt.Printf("Max players must be less than %v; Defaulting to 20\n", (2 ^ 31 - 1))
		serverProps = serverProps + "max-players=20\n"
	}

	// Set max-world-size
	if input.maxWorldSize < 1 || input.maxWorldSize > 29999984 {
		fmt.Println("Max world size must be less than 29999984; Defaulting to 29999984")
		serverProps = serverProps + "max-world-size=29999984\n"
	} else {
		serverProps = serverProps + "max-world-size=" + fmt.Sprint(input.maxWorldSize) + "\n"
	}

	// Set motd
	if input.motd == "" {
		serverProps = serverProps + "motd=A Minecraft Server\n"
	} else {
		serverProps = serverProps + "motd=" + input.motd + "\n"
	}

	// Set network-compression-threshold
	serverProps = serverProps + "network-compression-threshold=" + fmt.Sprint(input.networkCompressionThreshold) + "\n"

	//Set online-mode
	if input.onlineMode {
		serverProps = serverProps + "online-mode=true\n"
	} else {
		serverProps = serverProps + "online-mode=false\n"
	}

	// Set player-idle-timeout
	serverProps = serverProps + "player-idle-timeout=" + fmt.Sprint(input.playerIdleTimeout) + "\n"

	//Set prevent-proxy-connections
	if input.preventProxyConnections {
		serverProps = serverProps + "prevent-proxy-connections=true\n"
	} else {
		serverProps = serverProps + "prevent-proxy-connections=false\n"
	}

	//Set previews-chat
	if input.previewsChat {
		serverProps = serverProps + "previews-chat=true\n"
	} else {
		serverProps = serverProps + "previews-chat=false\n"
	}

	//Set pvp
	if input.pvp {
		serverProps = serverProps + "pvp=true\n"
	} else {
		serverProps = serverProps + "pvp=false\n"
	}

	// Set query.port
	if input.queryPort > 65534 {
		fmt.Printf("Query Port must be less than %v; Defaulting to 25565\n", 65534)
		serverProps = serverProps + "query.port=25565\n"
	} else {
		serverProps = serverProps + "query.port=" + fmt.Sprint(input.queryPort) + "\n"
	}

	// Set rcon.port
	if input.rconPort > 65534 {
		fmt.Printf("Rcon Port must be less than %v; Defaulting to 25575\n", 65534)
		serverProps = serverProps + "rcon.port=25575\n"
	} else {
		serverProps = serverProps + "rcon.port=" + fmt.Sprint(input.rconPort) + "\n"
	}

	// Set server-port
	if input.serverPort > 65534 {
		fmt.Printf("Server Port must be less than %v; Defaulting to 25565\n", 65534)
		serverProps = serverProps + "server-port=25565\n"
	} else {
		serverProps = serverProps + "server-port=" + fmt.Sprint(input.serverPort) + "\n"
	}

	// Set rate-limit
	serverProps = serverProps + "rate-limit=" + fmt.Sprint(input.rateLimit) + "\n"

	// Set rcon.password
	serverProps = serverProps + "rcon.password=" + input.rconPassword + "\n"

	// Set resource pack stuff
	serverProps = serverProps + "resource-pack=" + input.resourcePack + "\n"
	serverProps = serverProps + "resource-pack-prompt=" + input.resourcePackPrompt + "\n"
	serverProps = serverProps + "resource-pack-sha1=" + input.resourcePackSha1 + "\n"

	if input.requireResourcePack {
		serverProps = serverProps + "require-resource-pack=true\n"
	} else {
		serverProps = serverProps + "require-resource-pack=false\n"
	}

	// Set server-ip
	serverProps = serverProps + "server-ip=" + input.serverIP + "\n"

	// Set simulation-distance
	if input.simDistance < 3 || input.simDistance > 32 {
		fmt.Println("Simulation distance must be => 3 and => 32; Defaulting to 10")
		serverProps = serverProps + "simulation-distance=10\n"
	} else {
		serverProps = serverProps + "simulation-distance=" + fmt.Sprint(input.simDistance) + "\n"
	}

	// Set snooper-enabled
	if input.snooperEnabled {
		serverProps = serverProps + "snooper-enabled=true\n"
	} else {
		serverProps = serverProps + "snooper-enabled=false\n"
	}

	// Set spawn-animals
	if input.spawnAnimals {
		serverProps = serverProps + "spawn-animals=true\n"
	} else {
		serverProps = serverProps + "spawn-animals=false\n"
	}

	// Set spawn-monsters
	if input.spawnMonsters {
		serverProps = serverProps + "spawn-monsters=true\n"
	} else {
		serverProps = serverProps + "spawn-monsters=false\n"
	}

	// Set spawn-npcs
	if input.spawnNPCs {
		serverProps = serverProps + "spawn-npcs=true\n"
	} else {
		serverProps = serverProps + "spawn-npcs=false\n"
	}

	// Set server-protection
	serverProps = serverProps + "spawn-protection=" + fmt.Sprint(input.spawnProtection) + "\n"

	// Set sync-chunk-writes
	if input.syncChunkWrites {
		serverProps = serverProps + "sync-chunk-writes=true\n"
	} else {
		serverProps = serverProps + "sync-chunk-writes=false\n"
	}

	// Set use-native-transport
	if input.useNativeTransport {
		serverProps = serverProps + "use-native-transport=true\n"
	} else {
		serverProps = serverProps + "use-native-transport=false\n"
	}

	// Set view-distance
	if input.viewDistance < 3 || input.viewDistance > 32 {
		fmt.Println("View distance must be => 3 and => 32; Defaulting to 10")
		serverProps = serverProps + "view-distance=10\n"
	} else {
		serverProps = serverProps + "view-distance=" + fmt.Sprint(input.viewDistance) + "\n"
	}

	// Set whitelist
	if input.whitelist {
		serverProps = serverProps + "whitelist=true"
	} else {
		serverProps = serverProps + "whitelist=false"
	}

	finalServerProperties := "#This Minecraft server.properties file was generated by arithefirst\n" + serverProps
	return finalServerProperties
}
