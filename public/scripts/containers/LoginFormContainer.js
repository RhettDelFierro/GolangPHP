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
            token: ""
        };

    },


    //tie state to inputs.
    handleUpdateUser: function (e) {
        this.setState({
            user: e.target.value
        })
    },
    handleUpdatePassword: function (e) {
        this.setState({
            password: e.target.value
        })
    },

    //navbar form.
    //onSubmit: handling AJAX here. Probably not the best?
    //do a this.props.updateLogin. This will update the state of Main and re-render this.
    //no Ajax here. Handle that in componenntDidMount since it will re-render from Main.
    handleSubmitUser: function (e) {
        e.preventDefault();
        this.getToken();
    },

    //this will run also on logout. Must fix. This happens because of the re-render and the Main container will send in a new state as props.
    //componentDidMount: function () {
    //    this.getToken(this.props.user)
    //},
    //componentWillReceiveProps: function (nextProps) {
    //    this.getUser(nextProps.user)
    //},

    //called on navbar login.
    getToken: function(){
        userFunctions.loginUser({
            user: this.state.user,
            password: this.state.password
        }).then(function (data) {
            var date = new Date();
            date.setMinutes(15);
            //var expires = "expires=" + date;
            //var authorization = "Authorization=" + data.token;
            //document.cookie = authorization + ";" + expires;
            document.cookie = "expires=" + date;
            document.cookie = "token=" + data.token;
            this.props.onUpdateLogin(true, data.user.username);
        }.bind(this));
    },

    //this is called after a user registers:
    getUser: function (user) {
        if (user.length >= 5) {
            userFunctions.loginPassword(user)
                .then(function (data) {
                    var date = new Date();
                    date.setMinutes(15);
                    document.cookie = "Authorization=Bearer " + data.token + ";expires=" + date;
                    this.props.onUpdateLogin(true, data.user.username);
                }.bind(this));
        }
    },
    handleLogout: function () {
        this.setState({
            isLoggedIn: false,
            user: "",
            password: "",
            login: "",
            token: ""
        });
        document.cookie = "Authorization=;expires=Thu, 01 Jan 1970 00:00:01 GMT;";
        document.cookie = "expires=;expires=Thu, 01 Jan 1970 00:00:01 GMT;";
        this.props.onUpdateLogin(false, "");
    },
    render: function () {
        return (
            <LoginForm isLoggedIn={this.props.isLoggedIn}
                       user={this.state.user}
                       onUpdateLogin={this.props.onUpdateLogin}
                       onSubmitUser={this.handleSubmitUser}
                       onUpdateUser={this.handleUpdateUser}
                       onUpdatePassword={this.handleUpdatePassword}
                       password={this.state.password}
                       onLogout={this.handleLogout}/>
        )
    }
});


module.exports = LoginFormContainer;