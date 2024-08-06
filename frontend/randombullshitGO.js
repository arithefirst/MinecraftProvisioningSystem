// Initalize the tooltips
document.addEventListener('DOMContentLoaded', function () {
    var tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'))
    var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
        return new bootstrap.Tooltip(tooltipTriggerEl)
    })
})

function sendRequest() {
  // Get the values of all of the checkboxes
  var allowFlight = $("#allowFlight").is(":checked");
  var allowNether = $("#allowNether").is(":checked");
  var broadcastConsoleOps = $("#broadcastConsoleOps").is(":checked");
  var broadcastRconOps = $("#broadcastRconOps").is(":checked");
  var difficulty = $("#difficulty").is(":checked");
  var enableCommandBlock = $("#enableCommandBlock").is(":checked");
  var enableJmxMonitoring = $("#enableJmxMonitoring").is(":checked");
  var enableRcon = $("#enableRcon").is(":checked");
  var enableStatus = $("#enableStatus").is(":checked");
  var enableQuery = $("#enableQuery").is(":checked");
  var enforceSecureProfile = $("#enforceSecureProfile").is(":checked");
  var enforceWhitelist = $("#enforceWhitelist").is(":checked");
  var entityBroadcastRangePercent = $("#entityBroadcastRangePercent").is(":checked");
  var forceGamemode = $("#forceGamemode").is(":checked");
  var funcPermLevel = $("#funcPermLevel").is(":checked");
  var gamemode = $("#gamemode").is(":checked");
  var generateStructures = $("#generateStructures").is(":checked");
  var hardcore = $("#hardcore").is(":checked");
  var hideOnlinePlayers = $("#hideOnlinePlayers").is(":checked");
  var initalDisabledPacks = $("#initalDisabledPacks").is(":checked");
  var initalEnabledPacks = $("#initalEnabledPacks").is(":checked");
  var levelName = $("#levelName").is(":checked");
  var levelSeed = $("#levelSeed").is(":checked");
  var maxChainedNeighborUpdates = $("#maxChainedNeighborUpdates").is(":checked");
  var maxPlayers = $("#maxPlayers").is(":checked");
  var maxTickTime = $("#maxTickTime").is(":checked");
  var maxWorldSize = $("#maxWorldSize").is(":checked");
  var motd = $("#motd").is(":checked");
  var networkCompressionThreshold = $("#networkCompressionThreshold").is(":checked");
  var onlineMode = $("#onlineMode").is(":checked");
  var opPermLevel = $("#opPermLevel").is(":checked");
  var playerIdleTimeout = $("#playerIdleTimeout").is(":checked");
  var preventProxyConnections = $("#preventProxyConnections").is(":checked");
  var previewsChat = $("#previewsChat").is(":checked");
  var pvp = $("#pvp").is(":checked");
  var queryPort = $("#queryPort").is(":checked");
  var rateLimit = $("#rateLimit").is(":checked");
  var rconPassword = $("#rconPassword").is(":checked");
  var rconPort = $("#rconPort").is(":checked");
  var resourcePack = $("#resourcePack").is(":checked");
  var resourcePackPrompt = $("#resourcePackPrompt").is(":checked");
  var resourcePackSha1 = $("#resourcePackSha1").is(":checked");
  var requireResourcePack = $("#requireResourcePack").is(":checked");
  var serverIP = $("#serverIP").is(":checked");
  var serverPort = $("#serverPort").is(":checked");
  var simDistance = $("#simDistance").is(":checked");
  var snooperEnabled = $("#snooperEnabled").is(":checked");
  var spawnAnimals = $("#spawnAnimals").is(":checked");
  var spawnMonsters = $("#spawnMonsters").is(":checked");
  var spawnNPCs = $("#spawnNPCs").is(":checked");
  var spawnProtection = $("#spawnProtection").is(":checked");
  var syncChunkWrites = $("#syncChunkWrites").is(":checked");
  var useNativeTransport = $("#useNativeTransport").is(":checked");
  var viewDistance = $("#viewDistance").is(":checked");
  var whitelist = $("#whitelist").is(":checked");

  console.log(
    `${allowFlight}, ${allowNether}, ${broadcastConsoleOps}, ${broadcastRconOps}, ${difficulty}, ${enableCommandBlock}, ${enableJmxMonitoring}, ${enableRcon}, ${enableStatus}, ${enableQuery}, ${enforceSecureProfile}, ${enforceWhitelist}, ${entityBroadcastRangePercent}, ${forceGamemode}, ${funcPermLevel}, ${gamemode}, ${generateStructures}, ${hardcore}, ${hideOnlinePlayers}, ${initalDisabledPacks}, ${initalEnabledPacks}, ${levelName}, ${levelSeed}, ${maxChainedNeighborUpdates}, ${maxPlayers}, ${maxTickTime}, ${maxWorldSize}, ${motd}, ${networkCompressionThreshold}, ${onlineMode}, ${opPermLevel}, ${playerIdleTimeout}, ${preventProxyConnections}, ${previewsChat}, ${pvp}, ${queryPort}, ${rateLimit}, ${rconPassword}, ${rconPort}, ${resourcePack}, ${resourcePackPrompt}, ${resourcePackSha1}, ${requireResourcePack}, ${serverIP}, ${serverPort}, ${simDistance}, ${snooperEnabled}, ${spawnAnimals}, ${spawnMonsters}, ${spawnNPCs}, ${spawnProtection}, ${syncChunkWrites}, ${useNativeTransport}, ${viewDistance}, ${whitelist},`
  );

  location.reload()
}

