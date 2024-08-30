// Set the application's base URL
const baseUrl = "http://127.0.0.1:8080";

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

function truncate(inputString, searchString) {
  const position = inputString.indexOf(searchString);
  if (position === -1) {
    return inputString; // Return the original string if searchString is not found
  }
  return inputString.substring(0, position);
}

function sendRequest() {
  // Put all of the values of the toggles in the "bool" map
  const bools = new Map();
  bools.set("AllowFlight", $("#allowFlight").is(":checked"));
  bools.set("AllowNether", $("#allowNether").is(":checked"));
  bools.set("BroadcastConsoleOps", $("#broadcastConsoleOps").is(":checked"));
  bools.set("BroadcastRconOps", $("#broadcastRconOps").is(":checked"));
  bools.set("EnableCommandBlock", $("#enableCommandBlock").is(":checked"));
  bools.set("EnableJmxMonitoring", $("#enableJmxMonitoring").is(":checked"));
  bools.set("EnableRcon", $("#enableRcon").is(":checked"));
  bools.set("EnableStatus", $("#enableStatus").is(":checked"));
  bools.set("EnableQuery", $("#enableQuery").is(":checked"));
  bools.set("EnforceSecureProfile", $("#enforceSecureProfile").is(":checked"));
  bools.set("EnforceWhitelist", $("#enforceWhitelist").is(":checked"));
  bools.set("ForceGamemode", $("#forceGamemode").is(":checked"));
  bools.set("GenerateStructures", $("#generateStructures").is(":checked"));
  bools.set("Hardcore", $("#hardcore").is(":checked"));
  bools.set("HideOnlinePlayers", $("#hideOnlinePlayers").is(":checked"));
  bools.set("OnlineMode", $("#onlineMode").is(":checked"));
  bools.set("PreventProxyConnections", $("#preventProxyConnections").is(":checked"));
  bools.set("PreviewsChat", $("#previewsChat").is(":checked"));
  bools.set("Pvp", $("#pvp").is(":checked"));
  bools.set("RequireResourcePack", $("#requireResourcePack").is(":checked"));
  bools.set("SnooperEnabled", $("#requireResourcePack").is(":checked"));
  bools.set("SpawnAnimals", $("#spawnAnimals").is(":checked"));
  bools.set("SpawnMonsters", $("#spawnMonsters").is(":checked"));
  bools.set("SpawnNPCs", $("#spawnNPCs").is(":checked"));
  bools.set("SyncChunkWrites", $("#syncChunkWrites").is(":checked"));
  bools.set("UseNativeTransport", $("#useNativeTransport").is(":checked"));
  bools.set("Whitelist", $("#whitelist").is(":checked"));

  // Put all of the strings in the "strings" map
  const strings = new Map();
  strings.set("Difficulty", $("#difficulty").val());
  strings.set("Gamemode", $("#gamemode").val());
  strings.set("LevelName", $("#levelName").val());
  strings.set("LevelSeed", $("#levelSeed").val());
  strings.set("Motd", $("#motd").val());
  strings.set("RconPassword", $("#rconPassword").val());
  strings.set("ResourcePack", $("#resourcePack").val());
  strings.set("ResourcePackPrompt", $("#resourcePackPrompt").val());
  strings.set("ResourcePackSha1", $("#resourcePackSha1").val());
  strings.set("ServerIP", $("#serverIP").val());

  // Put all of the ints in the "ints" map
  const ints = new Map();
  ints.set("EntityBroadcastRangePercent", $("#entityBroadcastRangePercent").val());
  ints.set("FuncPermLevel", $("#funcPermLevel").val());
  ints.set("MaxChainedNeighborUpdates", $("#maxChainedNeighborUpdates").val());
  ints.set("MaxPlayers", $("#maxPlayers").val());
  ints.set("MaxTickTime", $("#maxTickTime").val());
  ints.set("MaxWorldSize", $("#maxWorldSize").val());
  ints.set("NetworkCompressionThreshold", $("#networkCompressionThreshold").val());
  ints.set("OpPermLevel", $("#opPermLevel").val());
  ints.set("PlayerIdleTimeout", $("#playerIdleTimeout").val());
  ints.set("QueryPort", $("#queryPort").val());
  ints.set("RateLimit", $("#rateLimit").val());
  ints.set("RconPort", $("#rconPort").val());
  ints.set("ServerPort", $("#serverPort").val());
  ints.set("SimDistance", $("#simulationDistance").val());
  ints.set("SpawnProtection", $("#spawnProtection").val());
  ints.set("ViewDistance", $("#viewDistance").val());

  var jsonString = "{";

  // Iterate through all of the bools
  for (let [key, value] of bools) {
    // "key": "value",
    jsonString = jsonString + '"' + key + '"' + ": " + '"' + value + '", ';
  }

  // Iterate through all of the strings
  for (let [key, value] of strings) {
    // "key": "value",
    jsonString = jsonString + '"' + key + '"' + ": " + '"' + value + '", ';
  }

  // Iterate through all of the ints
  for (let [key, value] of ints) {
    // "key": value,
    jsonString = jsonString + '"' + key + '"' + ": " + value + ", ";
  }

  // Remove the last ", " and close the json with a "}"
  jsonString = jsonString.slice(0, -2);
  jsonString = jsonString + "}";

  console.log(jsonString);
  const currentUrl = window.location.href;
  const baseUrl = currentUrl.split("?")[0];

  window.location.replace(encodeURI(baseUrl + "?jsonString=" + jsonString + "&overlay=true"));
}

function pageSet(httpReturn) {
  const displayReadyHTML = truncate(httpReturn.replace(/(\r\n|\r|\n)/g, "<br>"), ",<br>level-name");
  $("#server-response-display").html(displayReadyHTML + "...");
  localStorage.setItem("responseData", httpReturn);
}

function spawnAlert() {
  if (getUrlVars()["overlay"]) {
    console.log("Enabling overlay....");
    $("#overlay").css("visibility", "visible");
    getData();
  }
}

function getData() {
  const http = new XMLHttpRequest();
  json = encodeURI(getUrlVars()["jsonString"]);
  console.log(json);
  const url = `${baseUrl}/api/v1/server-properties?jsonString=${json}`;
  http.open("GET", url);
  http.send();

  http.onreadystatechange = (e) => pageSet(http.responseText);
}

function downloadData() {
  const data = localStorage.getItem("responseData");
  const blob = new Blob([data], { type: "text/plain" });
  const tempLink = document.createElement("a");
  tempLink.href = URL.createObjectURL(blob);
  tempLink.download = "server.properties";
  document.body.appendChild(tempLink);
  tempLink.click();
  document.body.removeChild(tempLink);
}

function dissmissAlert() {
  $("#overlay").css("visibility", "hidden");
}
