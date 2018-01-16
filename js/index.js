const ethUtil = require('ethereumjs-util')
const sigUtil = require('eth-sig-util')
const angular = require('angular')
const $ = require("jquery")
const BN = require('bn.js');
const Eth = require('ethjs');
const endpoint = '/api'
var contract
var channels
var local_state;
var fWaitNextState = true;
var account;
/*
var accountInterval = setInterval(function() {
  if (web3.eth.accounts[0] !== account) {
    account = web3.eth.accounts[0];
    updateInterface();
  }
}, 1000);
*/

window.addEventListener('load', function() {
  // Checking if Web3 has been injected by the browser (Mist/MetaMask)
    if (typeof web3 !== 'undefined') {
      // Use the browser's ethereum provider
      provider = web3.currentProvider
      eth = new Eth(provider)
      account = web3.eth.accounts[0];
      $.ajax({url: endpoint + "/status",})
        .done(function( data ) {
        var abi = JSON.parse(data.abi);
        var contractAddress = data.contract_address;
        channels = data.channels;
        contract = eth.contract(abi, '', {
          from: account,
          gas: 300000, /* TODO gas limit */
        }).at(contractAddress);
        checkAndApply(0);
      });
    } else {
      console.log('No web3? You should consider trying MetaMask!')
    }
})

//https://stackoverflow.com/questions/14603205/how-to-convert-hex-string-into-a-bytes-array-and-a-bytes-array-in-the-hex-strin/34356351#34356351
// Convert a hex string to a byte array
function hexToBytes(hex) {
    for (var bytes = [], c = 0; c < hex.length; c += 2)
    bytes.push(parseInt(hex.substr(c, 2), 16));
    return bytes;
}

// Convert a byte array to a hex string
function bytesToHex(bytes) {
    for (var hex = [], i = 0; i < bytes.length; i++) {
        hex.push((bytes[i] >>> 4).toString(16));
        hex.push((bytes[i] & 0xF).toString(16));
    }
    return hex.join("");
}

function extractSignatureHex(sig) {
  var r = sig.substring(2, 66)
  var s = sig.substring(66, 130)
  var v = new BN(sig.substring(130, 132), 16)
  return {v: v, r: '0x' + r, s: '0x' + s}
}

function checkAndApply(chanID) {
  var state = channels[chanID].state
  var sig = extractSignatureHex(channels[chanID].signatures[0])
  contract.CheckAndApply(state, sig.v, sig.r, sig.s).then(function(d){
      local_state = d['0']
      var scope = angular.element($("#board")).scope();
      scope.$apply(function(){
        scope.board['grid'] = state2grid(local_state)
        fWaitNextState = true
      })
  })
}

function state2grid(state) {
  state = hexToBytes(state.substring(2, state.length))
  var grid = [[' ',' ',' '],[' ',' ',' '],[' ',' ',' ']]
  for (var y = 0; y < 3; y++)
    for (var x = 0; x < 3; x++) {
      // header_length+x_max*y+x
      if (state[4+3*y+x] != 0) {
        grid[y][x] = state[4+3*y+x] == 1 ? 'X' : 'O'
      }
    }
  return grid
}

function commitState(new_state) {
  signMessage(new_state, function(state, signature){
      postMessage(state, [signature])
      local_state = state
      fWaitNextState = true
      console.log(new_state, state)
  })
}

function op_put(x, y, $scope) {
  y = new BN(y);
  x = new BN(x);
  contract.op_put(local_state, x, y).then(function(d){
    commitState(d['0'])
  })
}

function signMessage(data, cb) {
  var from = web3.eth.accounts[0]

  var params = [data, from]
  var method = 'personal_sign'

  provider.sendAsync({
    method,
    params,
    from,
  }, function (err, result) {
    if (err) return console.error(err)
    if (result.error) return console.error(result.error)
    cb(data, result.result)
  })
}

function ecRecoMsg(data, sig, cb) {
  var from = web3.eth.accounts[0]

  var params = [data, sig]
  var method = 'personal_ecRecover'

  provider.sendAsync({
    method,
    params,
    from,
  }, function (err, result) {
    if (err) return console.error(err)
    if (result.error) return console.error(result.error)
    cb(data, result.result)
  })
}

function publishMessage(data, signature) {
  var serializedJSON = JSON.stringify({ "data": data, "signatures": signature})
  $.ajax({
    'type': 'POST',
    'url': endpoint + '/publish',
    'contentType': 'application/json',
    'data': serializedJSON,
    'dataType': 'json',
    'fail': function(data){
      console.log('fail:'+data)
    }
  })
  .fail(function(data, data2) {
      console.log(data)
      console.log(data2)
      console.log('fail:'+data)
  })
  .done(function(data){
      console.log('success:'+data)
      console.log(channels[0].state)
      console.log(data)
  })
  .always(function() {
      console.log('always:')
  });
}

function postMessage(data, signature) {
  var serializedJSON = JSON.stringify({ "data": data, "signatures": signature})
  console.log(serializedJSON)
  $.ajax({
    'type': 'POST',
    'url': endpoint + '/commit',
    'contentType': 'application/json',
    'data': serializedJSON,
    'dataType': 'json'
  })
  .fail(function(data, data2) {
      console.log(data)
      console.log(data2)
      console.log('fail:'+data)
  })
  .done(function(data){
      console.log('success:'+data)
      console.log(channels[0].state)
      console.log(data)
      channels[0].state = data.data
      var scope = angular.element($("#board")).scope();
      scope.$apply(function(){
        scope.board['grid'] = state2grid(data.data)
        fWaitNextState = false
      })
  })
  .always(function() {
      console.log('always:')
  });
}

publish.addEventListener('click', function(event) {
  event.preventDefault()
  var chanID = 0
  signMessage(channels[chanID].state, function(state, signature){
      channels[chanID].signatures.push(signature)
      publishMessage(state, channels[chanID].signatures)
      console.log('publish')
      console.log(state)
      console.log(channels)
  })
})

challenge.addEventListener('click', function(event) {
  event.preventDefault()
  var chanID = 0
  var sig = extractSignatureHex(channels[chanID].signatures[0])
  console.log(channels[chanID].state)
  console.log(sig)
  contract.InfoSigner(channels[chanID].state, sig.v, sig.r, sig.s).then(function(ret) {console.log(ret)})
  contract.ChallengeDebug(channels[chanID].state, ['0x0', '0x0', '0x0'], [sig.v], [sig.r], [sig.s]).then(function(ret) {console.log(ret)})
  contract.Challenge(channels[chanID].state, ['0x0', '0x0', '0x0'], [sig.v], [sig.r], [sig.s]).then(function(tx) {
    console.log('TxHash:' + tx)
  })
  /*
  signMessage(channels[chanID].state, function(state, signature){
      channels[chanID].signatures.push(signature)
      publishMessage(state, channels[chanID].signatures)
      console.log('challenge')
      console.log(state)
      console.log(channels)
  })
  */
})

setPlayer.addEventListener('click', function(event) {
  event.preventDefault()
  other_player = $('#other_player')['0'].value
  contract.SetPlayer(other_player)
})
ttt = angular.module('TictactoeApp', [])
ttt.controller('BoardCtrl', function BoardCtrl($scope, $http) {
    var board = {grid: [['0,0','0,1','0,2'],['1,0','1,1','1,2'],['2,0','2,1','2,2']]}
    $scope.board = board
    $scope.playCell = (x, y) => op_put(x, y, $scope)
});
