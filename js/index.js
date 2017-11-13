var ethUtil = require('ethereumjs-util')
var sigUtil = require('eth-sig-util')
var $ = require("jquery");
var endpoint = '/api';

var provider
window.addEventListener('load', function() {
  // Checking if Web3 has been injected by the browser (Mist/MetaMask)
  if (typeof web3 !== 'undefined') {
    // Use the browser's ethereum provider
    provider = web3.currentProvider
  } else {
    console.log('No web3? You should consider trying MetaMask!')
  }
})

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

function postMessage(data, signature) {
  console.log(JSON.stringify(signature))
  var serializedJSON = JSON.stringify({ "data": data, "signature": signature});
  $.ajax({
    'type': 'POST',
    'url': endpoint + '/commit',
    'contentType': 'application/json',
    'data': serializedJSON,
    'dataType': 'json',
    'success': function(data){
      console.log(data);
    },
    'fail': function(data){
      console.log(data);
    }
  });
}

personalSignButton.addEventListener('click', function(event) {
  event.preventDefault()
  signMessage('0x544f444f', postMessage)
})
