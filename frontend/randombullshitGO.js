// Initalize the tooltips
document.addEventListener("DOMContentLoaded", function () {
  var tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'));
  var tooltipList = tooltipTriggerList.map(function (tooltipTriggerEl) {
    return new bootstrap.Tooltip(tooltipTriggerEl);
  });
});

function getUrlVars() {
  var vars = [],
    hash;
  var hashes = window.location.href.slice(window.location.href.indexOf("?") + 1).split("&");
  for (var i = 0; i < hashes.length; i++) {
    hash = hashes[i].split("=");
    vars.push(hash[0]);
    vars[hash[0]] = hash[1];
  }
  return vars;
}

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
  var simDistance = $("#simulationDistance").val();
  var spawnProtection = $("#spawnProtection").val();
  var viewDistance = $("#viewDistance").val();
  var levelName = $("#levelName").val();
  var levelSeed = $("#levelName").val();

  var jsonString = `{"AllowFlight": "${allowFlight}", "AllowNether": "${allowNether}", "BroadcastConsoleOps": "${broadcastConsoleOps}", "BroadcastRconOps": "${broadcastRconOps}", "Difficulty": "${difficulty}", "EnableCommandBlock": "${enableCommandBlock}", "EnableJmxMonitoring": "${enableJmxMonitoring}", "EnableRcon": "${enableRcon}", "EnableStatus": "${enableStatus}", "EnableQuery": "${enableQuery}", "EnforceSecureProfile": "${enforceSecureProfile}", "EnforceWhitelist": "${enforceWhitelist}", "EntityBroadcastRangePercent": ${entityBroadcastRangePercent}, "ForceGamemode": "${forceGamemode}", "FuncPermLevel": ${funcPermLevel}, "Gamemode": "${gamemode}", "GenerateStructures": "${generateStructures}", "Hardcore": "${hardcore}", "HideOnlinePlayers": "${hideOnlinePlayers}", "LevelName": "${levelName}", "LevelSeed": "${levelSeed}", "MaxChainedNeighborUpdates": ${maxChainedNeighborUpdates}, "MaxPlayers": ${maxPlayers}, "MaxTickTime": ${maxTickTime}, "MaxWorldSize": ${maxWorldSize}, "Motd": "${motd}", "NetworkCompressionThreshold": ${networkCompressionThreshold}, "OnlineMode": "${onlineMode}", "OpPermLevel": ${opPermLevel}, "PlayerIdleTimeout": ${playerIdleTimeout}, "PreventProxyConnections": "${preventProxyConnections}", "PreviewsChat": "${previewsChat}", "Pvp": "${pvp}", "QueryPort": ${queryPort}, "RateLimit": ${rateLimit}, "RconPassword": "${rconPassword}", "RconPort": ${rconPort}, "ResourcePack": "${resourcePack}", "ResourcePackPrompt": "${resourcePackPrompt}", "ResourcePackSha1": "${resourcePackSha1}", "RequireResourcePack": "${requireResourcePack}", "ServerIP": "${serverIP}", "ServerPort": ${serverPort}, "SimDistance": ${simDistance}, "SnooperEnabled": "${snooperEnabled}", "SpawnAnimals": "${spawnAnimals}", "SpawnMonsters": "${spawnMonsters}", "SpawnNPCs": "${spawnNPCs}", "SpawnProtection": ${spawnProtection}, "SyncChunkWrites": "${syncChunkWrites}", "UseNativeTransport": "${useNativeTransport}", "ViewDistance": ${viewDistance}, "Whitelist": "${whitelist}"}`;

  const baseURL = "http://127.0.0.1:5500/frontend/";
  console.log(jsonString);
  window.location.replace(baseURL + "propertiesReturned.html?jsonString=" + encodeURI(jsonString));
}

function pageSet(httpReturn) {
  const displayReadyHTML = httpReturn.replace(/(\r\n|\r|\n)/g, "<br>");
  $("#pageBody").html(displayReadyHTML);
  console.log(httpReturn);
}

function responsePageLoad() {
  const http = new XMLHttpRequest();
  json = encodeURI(getUrlVars()["jsonString"]);
  const url = "http://127.0.0.1:8080/server-properties?jsonString=" + json;
  http.open("GET", url);
  http.send();

  http.onreadystatechange = (e) => pageSet(http.responseText);
}
