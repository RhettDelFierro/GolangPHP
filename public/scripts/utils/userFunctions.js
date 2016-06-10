var axios = require("axios");

//axios.defaults.headers.common['Authorization'] = AUTH_TOKEN; for the jwt.
var userFunctions = {
    verifyName: function (user) {
        return axios.post("/username", {data: {username: user}}).then(function (response) {
            return response.data;
        }).catch(function (error) {
            console.log(error);
        })
    },
    registerUser: function (user) {
        return axios.post("/users/register", {data: {username: user.user, email: user.email, password: user.password}})
            .then(function (response) {
                console.log(response);
                return response.data
            })
            .catch(function (error) {
                console.log(error)
            })
    },
    loginUser: function(user){
        console.log(user);
        return axios.post("/users/login", {data: {login: user.user, password: user.password}})
            .then(function (response){
                console.log(response);
                return response.data
            })
            .catch(function (error){
                console.log(error);
            })
    },
    loginPassword: function(user){
        console.log("loginPassword: ", user);
        return axios.post("/users/pw", {data: {login: user}})
            .then(function (response){
                console.log(response);
                return response.data
            })
            .catch(function (error){
                console.log(error);
            })
    }
};

module.exports = userFunctions;