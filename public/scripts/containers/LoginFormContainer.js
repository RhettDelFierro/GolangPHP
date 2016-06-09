var React = require("react");
var LoginForm = require("../components/LoginForm");
var userFunctions = require("../utils/userFunctions");

var LoginFormContainer = React.createClass({
    //don't forget to go to the logged in route.
    contextTypes: {
        router: React.PropTypes.object.isRequired
    },
    getInitialState: function () {
        return {
            isLoggedIn: false,
            user: "",
            password: "",
            login: "",
            token: ""
        }

    },
    handleUpdateUser: function (e) {
        this.setState({
            login: e.target.value
        })
    },
    handleUpdatePassword: function (e) {
        this.setState({
            password: e.target.value
        })
    },
    handleSubmitUser: function () {
        e.preventDefault();
        userFunctions.loginUser({
            user: this.state.user,
            password: this.state.password
        }).then(function (data) {
            this.setState({
                isLoggedIn: true,
                user: data.username,
                token: data.token
            })
        }.bind(this))
    },
    componentWillReceiveProps: function (nextProps) {
        userFunctions.loginPassword(nextProps.user)
            .then(function (data) {
                this.setState({
                    isLoggedIn: true,
                    user: data.username,
                    token: data.token
                })
            }.bind(this))
    },
    render: function () {
        return (
            <LoginForm isLoggedIn={this.state.isLoggedIn}
                       user={this.state.user}
                       onUpdateLogin={this.props.onUpdateLogin}
                       onSubmitUser={this.handleSubmitUser}
                       onUpdateUser={this.handleUpdateUser}
                       onUpdatePassword={this.handleUpdatePassword}
                       password={this.state.password}/>
        )
    }
});


module.exports = LoginFormContainer;