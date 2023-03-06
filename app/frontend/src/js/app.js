App = {
    web3Provider: null,
    contracts: {},
    account: '0x0',
  
    init: function() {
        console.log("init function");
      return App.initWeb3();
    },
  
    initWeb3: function() {
      // TODO: refactor conditional
      if (typeof web3 !== 'undefined') {
        // If a web3 instance is already provided by MetaMask
        App.web3Provider = web3.currentProvider;
        web3 = new Web3(web3.currentProvider);
        console.log("web3 instance from MetaMask:", web3);
      } else {
        // Specify default instance if no web3 instance provided
        App.web3Provider = new Web3.providers.HttpProvider('http://localhost:7545');
        web3 = new Web3(App.web3Provider);
        console.log("web3 instance from HttpProvider:", web3);
      }
      return App.initContract();
    },
  
    initContract: function() {
        
      $.getJSON("../Ballot.json", function(ballot){
        console.log("Ballot.json:", ballot);
  
        // Connect provider to interact with contract
        const contractAbi = ballot.contracts.Ballot.abi;
        console.log("Contract ABI:", contractAbi);
  
        const networkId = web3.eth.net.getId();
        console.log("Network ID:", networkId);
        console.log("Networks:", ballot.networks);
        const contractAddress = ballot.networks[5777]['address'];
        console.log("Contract Address:", contractAddress);
  
        if (!contractAddress) {
          throw new Error(`Contract not deployed to network with ID ${networkId}`);
        }
        App.contracts.Ballot = new web3.eth.Contract(
          contractAbi,
          contractAddress
        );
        console.log("Contract instance:", App.contracts.Ballot);
        console.log("web3 provider:", App.web3Provider);
        console.log("Contract methods:", App.contracts.Ballot.methods);
  
        App.contracts.Ballot.setProvider(App.web3Provider); 
  
        return App.render();
      });
    },
  
    render: function() {
      console.log("in render function");
      var ballotInstance;
      var loader = $("#loader");
      var content = $("#content");
  
      loader.show();
      content.hide();
  
      // Load account data
      web3.eth.getCoinbase(function(err, account) {
        if (err === null) {
          App.account = account;
          $("#accountAddress").html("Your Account: " + account);
        }
      });
  
      // Load contract data
        App.contracts.Ballot.methods.candidatesCount().call()
        .then(function(candidatesCount){
        var candidatesResults = $("#candidatesResults");
        candidatesResults.empty();

        var promises = [];

        for (var i = 1; i <= candidatesCount; i++) {
            promises.push(App.contracts.Ballot.methods.candidates(i).call());
        }

        return Promise.all(promises);
        }).then(function(candidates){
        var candidatesResults = $("#candidatesResults");
        candidatesResults.empty();

        var candidatesSelect = $('#candidatesSelect');
        candidatesSelect.empty();

        for (var i = 0; i < candidates.length; i++) {
            const id = candidates[i].id;
            const name = candidates[i].name;
            const voteCount = candidates[i].voteCount;
            console.log("Candidate ID:", id);
            console.log("Candidate Name:", name);
            console.log("Candidate Vote Count:", voteCount);

            // Render candidate Result
            const candidateTemplate = "<tr><th>" + id + "</th><td>" + name + "</td><td>" + voteCount + "</td></tr>"
            candidatesResults.append(candidateTemplate);

            // Render candidate ballot option
            var candidateOption = "<option value='" + id +"' >" + name + "</ option >"
            candidatesSelect.append(candidateOption);
        }
        return App.contracts.Ballot.methods.voters(App.account).call();
      }).then(function(hasVoted){
        //Do not allow a user to vote
        if(hasVoted) {
          $('form').hide();
        }

        loader.hide();
        content.show();
        }).catch(function(error){
        console.warn(error);
        });
    },

    castVote: function() {
      var candidateId = $('#candidatesSelect').val();
      App.contracts.Ballot.methods.vote(candidateId).send({from: App.account})
      .then(function(result){
        // Wait for votes to update
        $("#content").hide();
        $("#loader").show();
      }).catch(function(err){
        console.error(err);
      });
    }
  };
  $(() => {
    $(window).load(() => {
      App.init();
    })
  })