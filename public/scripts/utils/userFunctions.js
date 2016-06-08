var axios = require("axios");

function Duplicate(user) {
    axios.post("/username", {username: user}).then(function (response) {
        return response.data;
    }).catch(function (error) {
        console.log(error);
    })
}

function Register(user) {
    return axios.post("/users/register", {username: user.username, email: user.email, password: user.password})
        .then(function (response) {
            return response.data
        })
        .catch(function (error) {
            console.log(error)
        })
}

//axios.defaults.headers.common['Authorization'] = AUTH_TOKEN; for the jwt.
var userFunctions = {
    verfifyName: function (user) {
        return Duplicate(user)
    },
    registerUser: function (user) {
        return Register(user)
    }
};

module.exports = userFunctions;