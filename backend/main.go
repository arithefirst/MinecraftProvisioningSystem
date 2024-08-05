package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func string2bool(instring string) bool {
	if strings.ToLower(instring) == "true" {
		return true
	} else {
		return false
	}
}

func string2int(instring string) int {
	if instring == "" {
		return 0
	} else {
		out, err := strconv.Atoi(instring)
		if err != nil {
			panic(err)
		}

		return out
	}
}

func serverPropertiesHttp(w http.ResponseWriter, r *http.Request) {

	// Frontend sends the server.properties config VIA headers in a get request
	// Func converts it into a server.properties file and returns it in the get request

	headers := serverProperties{
		allowFlight:                 string2bool(r.Header.Get("allowFlight")),
		allowNether:                 string2bool(r.Header.Get("allowNether")),
		broadcastConsoleOps:         string2bool(r.Header.Get("broadcastConsoleOps")),
		broadcastRconOps:            string2bool(r.Header.Get("broadcastRconOpsHeader")),
		difficulty:                  r.Header.Get("difficulty"),
		enableCommandBlock:          string2bool(r.Header.Get("enableCommandBlock")),
		enableJmxMonitoring:         string2bool(r.Header.Get("enableJmxMonitoring")),
		enableRcon:                  string2bool(r.Header.Get("enableRcon")),
		enableStatus:                string2bool(r.Header.Get("enableStatus")),
		enableQuery:                 string2bool(r.Header.Get("enableQuery")),
		enforceSecureProfile:        string2bool(r.Header.Get("enforceSecureProfile")),
		enforceWhitelist:            string2bool(r.Header.Get("enforceWhitelist")),
		entityBroadcastRangePercent: string2int(r.Header.Get("entityBroadcastRangePercent")),
		forceGamemode:               string2bool(r.Header.Get("forceGamemode")),
		funcPermLevel:               string2int(r.Header.Get("funcPermLevel")),
		gamemode:                    string2int(r.Header.Get("gamemode")),
		generateStructures:          string2bool(r.Header.Get("generateStructures")),
		hardcore:                    string2bool(r.Header.Get("hardcore")),
		hideOnlinePlayers:           string2bool(r.Header.Get("hideOnlinePlayers")),
		initalDisabledPacks:         r.Header.Get("initalDisabledPacks"),
		initalEnabledPacks:          r.Header.Get("initalEnabledPacks"),
		levelName:                   r.Header.Get("levelName"),
		levelSeed:                   r.Header.Get("levelSeed"),
		maxChainedNeighborUpdates:   string2int(r.Header.Get("maxChainedNeighborUpdates")),
		maxPlayers:                  string2int(r.Header.Get("maxPlayers")),
		maxTickTime:                 string2int(r.Header.Get("maxTickTime")),
		maxWorldSize:                string2int(r.Header.Get("maxWorldSize")),
		motd:                        r.Header.Get("motd"),
		networkCompressionThreshold: string2int(r.Header.Get("networkCompressionThreshold")),
		onlineMode:                  string2bool(r.Header.Get("onlineMode")),
		opPermLevel:                 string2int(r.Header.Get("opPermLevel")),
		playerIdleTimeout:           string2int(r.Header.Get("playerIdleTimeout")),
		preventProxyConnections:     string2bool(r.Header.Get("preventProxyConnections")),
		previewsChat:                string2bool(r.Header.Get("previewsChat")),
		pvp:                         string2bool(r.Header.Get("pvp")),
		queryPort:                   string2int(r.Header.Get("")),
		rateLimit:                   string2int(r.Header.Get("")),
		rconPassword:                r.Header.Get("rconPassword"),
		rconPort:                    string2int(r.Header.Get("rconPort")),
		resourcePack:                r.Header.Get("resourcePack"),
		resourcePackPrompt:          r.Header.Get("resourcePackPrompt"),
		resourcePackSha1:            r.Header.Get("resourcePackSha1"),
		requireResourcePack:         string2bool(r.Header.Get("requireResourcePack")),
		serverIP:                    r.Header.Get("serverIP"),
		serverPort:                  string2int(r.Header.Get("serverPort")),
		simDistance:                 string2int(r.Header.Get("simDistance")),
		snooperEnabled:              string2bool(r.Header.Get("snooperEnabled")),
		spawnAnimals:                string2bool(r.Header.Get("spawnAnimals")),
		spawnMonsters:               string2bool(r.Header.Get("spawnMonsters")),
		spawnNPCs:                   string2bool(r.Header.Get("spawnNPCs")),
		spawnProtection:             string2int(r.Header.Get("spawnProtection")),
		syncChunkWrites:             string2bool(r.Header.Get("syncChunkWrites")),
		useNativeTransport:          string2bool(r.Header.Get("useNativeTransport")),
		viewDistance:                string2int(r.Header.Get("viewDistance")),
		whitelist:                   string2bool(r.Header.Get("whitelist")),
	}

	fmt.Println(headers)
	fmt.Fprintf(w, "Minecraft server.properties:\n%v", generateServerProperties(headers))
}

func main() {
	http.HandleFunc("/server-properties", serverPropertiesHttp)
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
