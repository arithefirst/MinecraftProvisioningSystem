package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type serverProperties struct {
	AllowFlight                 Bool   `schema:"allowFlight"`
	AllowNether                 Bool   `schema:"allowNether"`
	BroadcastConsoleOps         Bool   `schema:"broadcastConsoleOps"`
	BroadcastRconOps            Bool   `schema:"broadcastRconOps"`
	Difficulty                  string `schema:"difficulty"`
	EnableCommandBlock          Bool   `schema:"enableCommandBlock"`
	EnableJmxMonitoring         Bool   `schema:"enableJmxMonitoring"`
	EnableRcon                  Bool   `schema:"enableRcon"`
	EnableStatus                Bool   `schema:"enableStatus"`
	EnableQuery                 Bool   `schema:"enableQuery"`
	EnforceSecureProfile        Bool   `schema:"enforceSecureProfile"`
	EnforceWhitelist            Bool   `schema:"enforceWhitelist"`
	EntityBroadcastRangePercent int    `schema:"entityBroadcastRangePercent"` // Range 10-1000
	ForceGamemode               Bool   `schema:"forceGamemode"`
	FuncPermLevel               int    `schema:"funcPermLevel"` // Range 1-4
	Gamemode                    string `schema:"gamemode"`
	GenerateStructures          Bool   `schema:"generateStructures"`
	Hardcore                    Bool   `schema:"hardcore"`
	HideOnlinePlayers           Bool   `schema:"hideOnlinePlayers"`
	LevelName                   string `schema:"levelName"`
	LevelSeed                   string `schema:"levelSeed"`
	MaxChainedNeighborUpdates   int    `schema:"maxChainedNeighborUpdates"`
	MaxPlayers                  int    `schema:"maxPlayers"`   // Range 0-2147483647
	MaxTickTime                 int    `schema:"maxTickTime"`  // -1 to Disable || Range 0–(2^63 - 1)
	MaxWorldSize                int    `schema:"maxWorldSize"` // Range 1-29999984
	Motd                        string `schema:"motd"`
	NetworkCompressionThreshold int    `schema:"networkCompressionThreshold"`
	OnlineMode                  Bool   `schema:"onlineMode"`
	OpPermLevel                 int    `schema:"opPermLevel"` // Range 0-4
	PlayerIdleTimeout           int    `schema:"playerIdleTimeout"`
	PreventProxyConnections     Bool   `schema:"preventProxyConnections"`
	PreviewsChat                Bool   `schema:"previewsChat"`
	Pvp                         Bool   `schema:"pvp"`
	QueryPort                   int    `schema:"queryPort"` // Valid port numbers only
	RateLimit                   int    `schema:"rateLimit"`
	RconPassword                string `schema:"rconPassword"`
	RconPort                    int    `schema:"rconPort"`     // Valid port numbers only
	ResourcePack                string `schema:"resourcePack"` // URI
	ResourcePackPrompt          string `schema:"resourcePackPrompt"`
	ResourcePackSha1            string `schema:"resourcePackSha1"`
	RequireResourcePack         Bool   `schema:"requireResourcePack"`
	ServerIP                    string `schema:"serverIP"`
	ServerPort                  int    `schema:"serverPort"`  // Valid port numbers only
	SimDistance                 int    `schema:"simDistance"` // Range 3-32
	SnooperEnabled              Bool   `schema:"snooperEnabled"`
	SpawnAnimals                Bool   `schema:"spawnAnimals"`
	SpawnMonsters               Bool   `schema:"spawnMonsters"`
	SpawnNPCs                   Bool   `schema:"spawnNPCs"`
	SpawnProtection             int    `schema:"spawnProtection"`
	SyncChunkWrites             Bool   `schema:"syncChunkWrites"`
	UseNativeTransport          Bool   `schema:"useNativeTransport"`
	ViewDistance                int    `schema:"viewDistance"` // Range 3-32
	Whitelist                   Bool   `schema:"whitelist"`
}

func serverPropertiesHttp(w http.ResponseWriter, r *http.Request) {

	// Frontend sends the server.properties config VIA querystring in a get request
	// Func converts it into a server.properties file and returns it in the get request

	var jsonProperties serverProperties

	encodedJsonString := r.URL.Query().Get("jsonString")
	jsonString, err := url.QueryUnescape(encodedJsonString)
	if err != nil {
		fmt.Fprintf(w, "Error decoding url: %v", err)
	}
	fmt.Printf("Recived jsonString: %v\n", jsonString)

	err = json.Unmarshal([]byte(jsonString), &jsonProperties)
	if err != nil {
		fmt.Fprintf(w, "Error parsing JSON: %v", err)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "%v", generateServerProperties(jsonProperties))
}

func generateServerProperties(input serverProperties) string {
	var serverProps string

	// Set allow-flight
	if input.AllowFlight {
		serverProps = serverProps + "allow-flight=true\n"
	} else {
		serverProps = serverProps + "allow-flight=false\n"
	}

	// Set allow-nether
	if input.AllowNether {
		serverProps = serverProps + "allow-nether=true\n"
	} else {
		serverProps = serverProps + "allow-nether=false\n"
	}

	// Set broadcast-console-to-ops
	if input.BroadcastConsoleOps {
		serverProps = serverProps + "broadcast-console-to-ops=true\n"
	} else {
		serverProps = serverProps + "broadcast-console-to-ops=false\n"
	}

	// Set broadcast-rcon-to-ops
	if input.BroadcastRconOps {
		serverProps = serverProps + "broadcast-rcon-to-ops=true\n"
	} else {
		serverProps = serverProps + "broadcast-rcon-to-ops=false\n"
	}

	// Check if difficulty is valid and set it
	switch strings.ToLower(input.Difficulty) {
	case "easy":
		serverProps = serverProps + "difficulty=easy\n"
	case "normal":
		serverProps = serverProps + "difficulty=normal\n"
	case "hard":
		serverProps = serverProps + "difficulty=hard\n"
	case "peaceful":
		serverProps = serverProps + "difficulty=peaceful\n"
	default:
		fmt.Println("Difficulty option invalid; Defaulting to easy")
		serverProps = serverProps + "difficulty=easy\n"
	}

	// Set enable-command-block
	if input.EnableCommandBlock {
		serverProps = serverProps + "enable-command-block=true\n"
	} else {
		serverProps = serverProps + "enable-command-block=false\n"
	}

	// Set enable-jmx-monitoring
	if input.EnableJmxMonitoring {
		serverProps = serverProps + "enable-jmx-monitoring=true\n"
	} else {
		serverProps = serverProps + "enable-jmx-monitoring=false\n"
	}

	// Set enable-rcon
	if input.EnableRcon {
		serverProps = serverProps + "enable-rcon=true\n"
	} else {
		serverProps = serverProps + "enable-rcon=false\n"
	}

	// Set enable-status
	if input.EnableStatus {
		serverProps = serverProps + "enable-status=true\n"
	} else {
		serverProps = serverProps + "enable-status=false\n"
	}

	// Set enable-query
	if input.EnableQuery {
		serverProps = serverProps + "enable-query=true\n"
	} else {
		serverProps = serverProps + "enable-query=false\n"
	}

	// Set enforce-secure-profile
	if input.EnforceSecureProfile {
		serverProps = serverProps + "enforce-secure-profile=true\n"
	} else {
		serverProps = serverProps + "enforce-secure-profile=false\n"
	}

	// Set enforce-whitelist
	if input.EnforceWhitelist {
		serverProps = serverProps + "enforce-whitelist=true\n"
	} else {
		serverProps = serverProps + "enforce-whitelist=false\n"
	}

	// Set entity-broadcast-range-percentage
	EBPP := input.EntityBroadcastRangePercent
	if EBPP > 1000 || EBPP < 10 {
		fmt.Println("Entity Broadcast Range Percent must be below 1000 and above 10; Defaulting to 100")
		serverProps = serverProps + "entity-broadcast-range-percentage=100\n"
	} else {
		serverProps = serverProps + "entity-broadcast-range-percentage=" + fmt.Sprint(EBPP) + "\n"
	}

	// Set force-gamemode
	if input.ForceGamemode {
		serverProps = serverProps + "force-gamemode=true\n"
	} else {
		serverProps = serverProps + "force-gamemode=false\n"
	}

	// Set function-permission-level
	if input.FuncPermLevel < 1 || input.FuncPermLevel > 4 {
		fmt.Println("Function permission level must be => 1 and <= 5; Defaulting to 2")
		serverProps = serverProps + "function-permission-level=2\n"
	} else {
		serverProps = serverProps + "function-permisison-level=" + fmt.Sprint(input.FuncPermLevel) + "\n"
	}

	// Set op-permission-level
	if input.OpPermLevel < 1 || input.OpPermLevel > 4 {
		fmt.Println("Op permission level must be => 1 and <= 5; Defaulting to 2")
		serverProps = serverProps + "op-permission-level=2\n"
	} else {
		serverProps = serverProps + "op-permission-level=" + fmt.Sprint(input.OpPermLevel) + "\n"
	}

	// Set gamemode
	switch strings.ToLower(input.Gamemode) {
	case "survival":
		serverProps = serverProps + "gamemode=survival\n"
	case "creative":
		serverProps = serverProps + "gamemode=creative\n"
	case "adventure":
		serverProps = serverProps + "gamemode=adventure\n"
	case "spectator":
		serverProps = serverProps + "gamemode=spectator\n"
	default:
		fmt.Println("Gamemode must be either 1, 2, 3, or 4; Defaulting to survival (0)")
		serverProps = serverProps + "gamemode=survival\n"
	}

	// Set generate-structures
	if input.GenerateStructures {
		serverProps = serverProps + "generate-structures=true\n"
	} else {
		serverProps = serverProps + "generate-structures=false\n"
	}

	// Set hardcore
	if input.Hardcore {
		serverProps = serverProps + "hardcore=true\n"
	} else {
		serverProps = serverProps + "hardcore=false\n"
	}

	// Set hide-online-players
	if input.HideOnlinePlayers {
		serverProps = serverProps + "hide-online-players=true\n"
	} else {
		serverProps = serverProps + "hide-online-players=false\n"
	}

	// Set initial-disabled-packs
	serverProps = serverProps + "initial-disabled-packs=" + "" + "\n"

	// Set initial-enabled-packs
	serverProps = serverProps + "initial-enabled-packs=" + "vanilla," + "\n"

	// Set level-name
	serverProps = serverProps + "level-name=" + input.LevelName + "\n"

	// Set level-seed
	serverProps = serverProps + "level-seed=" + input.LevelSeed + "\n"

	// Set max-chained-neighbor-updates
	serverProps = serverProps + "max-chained-neighbor-updates=" + fmt.Sprint(input.MaxChainedNeighborUpdates) + "\n"

	// Set max-players
	if input.MaxPlayers < (2 ^ 31 - 1) {
		serverProps = serverProps + "max-players=" + fmt.Sprint(input.MaxPlayers) + "\n"
	} else {
		fmt.Printf("Max players must be less than %v; Defaulting to 20\n", (2 ^ 31 - 1))
		serverProps = serverProps + "max-players=20\n"
	}

	// Set max-world-size
	if input.MaxWorldSize < 1 || input.MaxWorldSize > 29999984 {
		fmt.Println("Max world size must be less than 29999984; Defaulting to 29999984")
		serverProps = serverProps + "max-world-size=29999984\n"
	} else {
		serverProps = serverProps + "max-world-size=" + fmt.Sprint(input.MaxWorldSize) + "\n"
	}

	// Set motd
	if input.Motd == "" {
		serverProps = serverProps + "motd=A Minecraft Server\n"
	} else {
		serverProps = serverProps + "motd=" + input.Motd + "\n"
	}

	// Set network-compression-threshold
	serverProps = serverProps + "network-compression-threshold=" + fmt.Sprint(input.NetworkCompressionThreshold) + "\n"

	//Set online-mode
	if input.OnlineMode {
		serverProps = serverProps + "online-mode=true\n"
	} else {
		serverProps = serverProps + "online-mode=false\n"
	}

	// Set player-idle-timeout
	serverProps = serverProps + "player-idle-timeout=" + fmt.Sprint(input.PlayerIdleTimeout) + "\n"

	//Set prevent-proxy-connections
	if input.PreventProxyConnections {
		serverProps = serverProps + "prevent-proxy-connections=true\n"
	} else {
		serverProps = serverProps + "prevent-proxy-connections=false\n"
	}

	//Set previews-chat
	if input.PreviewsChat {
		serverProps = serverProps + "previews-chat=true\n"
	} else {
		serverProps = serverProps + "previews-chat=false\n"
	}

	//Set pvp
	if input.Pvp {
		serverProps = serverProps + "pvp=true\n"
	} else {
		serverProps = serverProps + "pvp=false\n"
	}

	// Set query.port
	if input.QueryPort > 65534 {
		fmt.Printf("Query Port must be less than %v; Defaulting to 25565\n", 65534)
		serverProps = serverProps + "query.port=25565\n"
	} else {
		serverProps = serverProps + "query.port=" + fmt.Sprint(input.QueryPort) + "\n"
	}

	// Set rcon.port
	if input.RconPort > 65534 {
		fmt.Printf("Rcon Port must be less than %v; Defaulting to 25575\n", 65534)
		serverProps = serverProps + "rcon.port=25575\n"
	} else {
		serverProps = serverProps + "rcon.port=" + fmt.Sprint(input.RconPassword) + "\n"
	}

	// Set server-port
	if input.ServerPort > 65534 {
		fmt.Printf("Server Port must be less than %v; Defaulting to 25565\n", 65534)
		serverProps = serverProps + "server-port=25565\n"
	} else {
		serverProps = serverProps + "server-port=" + fmt.Sprint(input.ServerPort) + "\n"
	}

	// Set rate-limit
	serverProps = serverProps + "rate-limit=" + fmt.Sprint(input.RateLimit) + "\n"

	// Set rcon.password
	serverProps = serverProps + "rcon.password=" + input.RconPassword + "\n"

	// Set resource pack stuff
	serverProps = serverProps + "resource-pack=" + input.ResourcePack + "\n"
	serverProps = serverProps + "resource-pack-prompt=" + input.ResourcePackPrompt + "\n"
	serverProps = serverProps + "resource-pack-sha1=" + input.ResourcePackSha1 + "\n"

	if input.RequireResourcePack {
		serverProps = serverProps + "require-resource-pack=true\n"
	} else {
		serverProps = serverProps + "require-resource-pack=false\n"
	}

	// Set server-ip
	serverProps = serverProps + "server-ip=" + input.ServerIP + "\n"
	// Set simulation-distance
	if input.SimDistance < 3 || input.SimDistance > 32 {
		fmt.Println("Simulation distance must be => 3 and => 32; Defaulting to 10")
		serverProps = serverProps + "simulation-distance=10\n"
	} else {
		serverProps = serverProps + "simulation-distance=" + fmt.Sprint(input.SimDistance) + "\n"
	}

	// Set snooper-enabled
	if input.SnooperEnabled {
		serverProps = serverProps + "snooper-enabled=true\n"
	} else {
		serverProps = serverProps + "snooper-enabled=false\n"
	}

	// Set spawn-animals
	if input.SpawnAnimals {
		serverProps = serverProps + "spawn-animals=true\n"
	} else {
		serverProps = serverProps + "spawn-animals=false\n"
	}

	// Set spawn-monsters
	if input.SpawnMonsters {
		serverProps = serverProps + "spawn-monsters=true\n"
	} else {
		serverProps = serverProps + "spawn-monsters=false\n"
	}

	// Set spawn-npcs
	if input.SpawnNPCs {
		serverProps = serverProps + "spawn-npcs=true\n"
	} else {
		serverProps = serverProps + "spawn-npcs=false\n"
	}

	// Set server-protection
	serverProps = serverProps + "spawn-protection=" + fmt.Sprint(input.SpawnProtection) + "\n"

	// Set sync-chunk-writes
	if input.SyncChunkWrites {
		serverProps = serverProps + "sync-chunk-writes=true\n"
	} else {
		serverProps = serverProps + "sync-chunk-writes=false\n"
	}

	// Set use-native-transport
	if input.UseNativeTransport {
		serverProps = serverProps + "use-native-transport=true\n"
	} else {
		serverProps = serverProps + "use-native-transport=false\n"
	}

	// Set view-distance
	if input.ViewDistance < 3 || input.ViewDistance > 32 {
		fmt.Println("View distance must be => 3 and => 32; Defaulting to 10")
		serverProps = serverProps + "view-distance=10\n"
	} else {
		serverProps = serverProps + "view-distance=" + fmt.Sprint(input.ViewDistance) + "\n"
	}

	// Set whitelist
	if input.Whitelist {
		serverProps = serverProps + "whitelist=true"
	} else {
		serverProps = serverProps + "whitelist=false"
	}

	watermarkLN1 := "#This Minecraft server.properties file was generated by arithefirst using MCPS\n"
	watermarkLN2 := "#https://github.com/arithefirst/MinecraftProvisioningSystem\n"

	finalServerProperties := watermarkLN1 + watermarkLN2 + serverProps
	return finalServerProperties
}
