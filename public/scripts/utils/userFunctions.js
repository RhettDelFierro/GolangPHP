var axios = require("axios");

function cookieFinder(name) {
    var cookie = document.cookie;
    var initialValue = {};

    return cookie.split(';').reduce(function (prev, c) {
        var arr = c.split('=');
        return (arr[0].trim() == name ? arr[1] : prev);
    }, initialValue);
}

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
                return response.data
            })
            .catch(function (error) {
                console.log(error)
            })
    },
    loginUser: function (user) {
        return axios.post("/users/login", {data: {login: user.user, password: user.password}})
            .then(function (response) {
                return response.data
            })
            .catch(function (error) {
                console.log(error);
            })
    },
    loginPassword: function (user) {
        return axios.post("/users/pw", {data: {login: user}})
            .then(function (response) {
                return response.data
            })
            .catch(function (error) {
                console.log(error);
            })
    },
    //addStudent: function (student, token) {
    //    //note the difference between this and user. The users have Data Mapping (and expects "data" field)
    //    return axios.post("/api/add", {
    //            data: {student: student.student, course: student.course, grade: student.grade},
    //            headers: {"AUTHORIZATION": "Bearer " + token}
    //        })
    //        .then(function (response) {
    //            console.log(response);
    //            return response.data
    //        })
    //        .catch(function (error) {
    //            console.log(error);
    //        })
    //}
    addStudent: function (student, user) {
        console.log("user", student);
        axios.defaults.headers.common['Authorization'] = "Bearer " + cookieFinder("token");
        //note the difference between this and user. The users have Data Mapping (and expects "data" field)
        return axios.post("/api/add", {name: student.student, course: student.course, grade: student.grade, user: user})
            .then(function (response) {
                console.log("addStudent: ", response);
                return response.data
            })
            .catch(function (error) {
                console.log(error);
            })
    },
    deleteStudent: function (id) {
        axios.defaults.headers.common['Authorization'] = "Bearer " + cookieFinder("token");
        return axios.get("/api/delete/" + id)
            .then(function (response) {
                return response.data
            })
            .catch(function (error) {
                console.log(error);
            })
    }
};

module.exports = userFunctions;