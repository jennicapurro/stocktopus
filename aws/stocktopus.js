var cp = require("child_process");

exports.handler = function(event, context) {

    // Parse our the request from the body
    var queryStr = unescape(event.body)
    var jsonStr  = '{"' + queryStr.replace(/&/g, '", "').replace(/=/g, '": "') + '"}';
    var tickers = JSON.parse(jsonStr).text.replace(/\+/g, " ")

    // Spawn the go routine to lookup stock quote
    var proc = cp.spawnSync("./colinmc", [tickers], {stdio: 'pipe', encoding: "utf8"});
    var quote = proc.stdout;

    // Parse quote into json for slack
    var resp = '{ "response_type" : "in_channel", "text" : "' + quote + '" }';

    // Return json
    context.succeed(resp);
};