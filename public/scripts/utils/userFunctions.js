var axios = require("axios");

//axios.defaults.headers.common['Authorization'] = AUTH_TOKEN; for the jwt.
var userFunctions = {
    verfifyName: function (user) {
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
        return axios.post("/users/login", {data: {}})
    }
};

module.exports = userFunctions;