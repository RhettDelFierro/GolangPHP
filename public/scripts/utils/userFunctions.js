var axios = require("axios");

function Duplicate(user){
    axios.post("/username", {username: user}).then(function(response){
        if (response.data === true) {

        }
        console.log(data);
    }).catch(function(error){

    })
}

var userFunctions = {
    verfifyName: function(user){
        return Duplicate(user)
    }
};

module.exports = userFunctions;