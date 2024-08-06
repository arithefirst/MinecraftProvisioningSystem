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
  var rconPassword = $("#rconPassword").is(":checked");
  var requireResourcePack = $("#requireResourcePack").is(":checked");
  var snooperEnabled = $("#snooperEnabled").is(":checked");
  var spawnAnimals = $("#spawnAnimals").is(":checked");
  var spawnMonsters = $("#spawnMonsters").is(":checked");
  var spawnNPCs = $("#spawnNPCs").is(":checked");
  var syncChunkWrites = $("#syncChunkWrites").is(":checked");
  var useNativeTransport = $("#useNativeTransport").is(":checked");
  var whitelist = $("#whitelist").is(":checked");

  console.log(
    `${allowFlight}, ${allowNether}, ${broadcastConsoleOps}, ${broadcastRconOps}, ${enableCommandBlock}, ${enableJmxMonitoring}, ${enableRcon}, ${enableStatus}, ${enableQuery}, ${enforceSecureProfile}, ${enforceWhitelist}, ${forceGamemode}, ${generateStructures}, ${hardcore}, ${hideOnlinePlayers}, ${onlineMode}, ${preventProxyConnections}, ${previewsChat}, ${pvp}, ${requireResourcePack}, ${snooperEnabled}, ${spawnAnimals}, ${spawnMonsters}, ${spawnNPCs}, ${syncChunkWrites}, ${useNativeTransport}, ${whitelist},`
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