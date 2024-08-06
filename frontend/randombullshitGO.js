// Initalize the tooltips
document.addEventListener('DOMContentLoaded', function () {
    var tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'))
    var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
        return new bootstrap.Tooltip(tooltipTriggerEl)
    })
})


function sendRequest() {
  // Get the values of all of the switches
  var allowFlight = $("#allowFlight").is(":checked");
  var allowNether = $("#allowNether").is(":checked");
  var broadcastConsoleOps = $("#broadcastConsoleOps").is(":checked");
  var broadcastRconOps = $("#broadcastRconOps").is(":checked");
  var enableCommandBlock = $("#enableCommandBlock").is(":checked");
  var enableJmxMonitoring = $("#enableJmxMonitoring").is(":checked");
  var enableRcon = $("#enableRcon").is(":checked");
  var enableStatus = $("#enableStatus").is(":checked");
  var enableQuery = $("#enableQuery").is(":checked");
  var enforceSecureProfile = $("#enforceSecureProfile").is(":checked");
  var enforceWhitelist = $("#enforceWhitelist").is(":checked");
  var forceGamemode = $("#forceGamemode").is(":checked");
  var generateStructures = $("#generateStructures").is(":checked");
  var hardcore = $("#hardcore").is(":checked");
  var hideOnlinePlayers = $("#hideOnlinePlayers").is(":checked");
  var onlineMode = $("#onlineMode").is(":checked");
  var preventProxyConnections = $("#preventProxyConnections").is(":checked");
  var previewsChat = $("#previewsChat").is(":checked");
  var pvp = $("#pvp").is(":checked");
  var requireResourcePack = $("#requireResourcePack").is(":checked");
  var snooperEnabled = $("#snooperEnabled").is(":checked");
  var spawnAnimals = $("#spawnAnimals").is(":checked");
  var spawnMonsters = $("#spawnMonsters").is(":checked");
  var spawnNPCs = $("#spawnNPCs").is(":checked");
  var syncChunkWrites = $("#syncChunkWrites").is(":checked");
  var useNativeTransport = $("#useNativeTransport").is(":checked");
  var whitelist = $("#whitelist").is(":checked");

  // Get the values of all of the text boxes
  var difficulty = $("#difficulty").val();
  var entityBroadcastRangePercent = $("#entityBroadcastRangePercent").val();
  var funcPermLevel = $("#funcPermLevel").val();
  var gamemode = $("#gamemode").val();
  var funcPermLevel = $("#funcPermLevel").val();
  var maxChainedNeighborUpdates = $("#maxChainedNeighborUpdates").val();
  var maxPlayers = $("#maxPlayers").val();
  var maxTickTime = $("#maxTickTime").val();
  var maxWorldSize = $("#maxWorldSize").val();
  var motd = $("#motd").val();
  var networkCompressionThreshold = $("#networkCompressionThreshold").val();
  var opPermLevel = $("#opPermLevel").val();
  var playerIdleTimeout = $("#playerIdleTimeout").val();
  var queryPort = $("#queryPort").val();
  var rateLimit = $("#rateLimit").val();
  var rconPassword = $("#rconPassword").val();
  var rconPort = $("#rconPort").val();
  var resourcePack = $("#resourcePack").val();
  var resourcePackPrompt = $("#resourcePackPrompt").val();
  var resourcePackSha1 = $("#resourcePackSha1").val();
  var serverIP = $("#serverIP").val();
  var serverPort = $("#serverPort").val();
  var simDistance = $("#simDistance").val();
  var spawnProtection = $("#spawnProtection").val();
  var viewDistance = $("#viewDistance").val();
  var levelName = $("#levelName").val();
  var levelSeed = $("#levelName").val();

  console.log(
    `allowFlight: "${allowFlight}", allowNether: "${allowNether}", broadcastConsoleOps: "${broadcastConsoleOps}", broadcastRconOps: "${broadcastRconOps}", difficulty: "${difficulty}", enableCommandBlock: "${enableCommandBlock}", enableJmxMonitoring: "${enableJmxMonitoring}", enableRcon: "${enableRcon}", enableStatus: "${enableStatus}", enableQuery: "${enableQuery}", enforceSecureProfile: "${enforceSecureProfile}", enforceWhitelist: "${enforceWhitelist}", entityBroadcastRangePercent: "${entityBroadcastRangePercent}", forceGamemode: "${forceGamemode}", funcPermLevel: "${funcPermLevel}", gamemode: "${gamemode}", generateStructures: "${generateStructures}", hardcore: "${hardcore}", hideOnlinePlayers: "${hideOnlinePlayers}", levelName: "${levelName}", levelSeed: "${levelSeed}", maxChainedNeighborUpdates: "${maxChainedNeighborUpdates}", maxPlayers: "${maxPlayers}", maxTickTime: "${maxTickTime}", maxWorldSize: "${maxWorldSize}", motd: "${motd}", networkCompressionThreshold: "${networkCompressionThreshold}", onlineMode: "${onlineMode}", opPermLevel: "${opPermLevel}", playerIdleTimeout: "${playerIdleTimeout}", preventProxyConnections: "${preventProxyConnections}", previewsChat: "${previewsChat}", pvp: "${pvp}", queryPort: "${queryPort}", rateLimit: "${rateLimit}", rconPassword: "${rconPassword}", rconPort: "${rconPort}", resourcePack: "${resourcePack}", resourcePackPrompt: "${resourcePackPrompt}", resourcePackSha1: "${resourcePackSha1}", requireResourcePack: "${requireResourcePack}", serverIP: "${serverIP}", serverPort: "${serverPort}", simDistance: "${simDistance}", snooperEnabled: "${snooperEnabled}", spawnAnimals: "${spawnAnimals}", spawnMonsters: "${spawnMonsters}", spawnNPCs: "${spawnNPCs}", spawnProtection: "${spawnProtection}", syncChunkWrites: "${syncChunkWrites}", useNativeTransport: "${useNativeTransport}", viewDistance: "${viewDistance}", whitelist: "${whitelist}"`
  );

  location.reload()
}

function showToast(message) {
    var toastEl = document.getElementById('toastContainer');
    var bsToast = new bootstrap.Toast(toastEl);
    var toastBody = toastEl.querySelector('.toast-body');
    toastBody.innerHTML = message;
    bsToast.show();
  }  