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
        userFunctions.loginUser({
            user: this.state.user,
            password: this.state.password
        }).then(function (data) {
            this.setState({
                isLoggedIn: true,
                user: data.user.username,
                token: data.token
            });
            this.props.onUpdateLogin(true, this.state.user);
        }.bind(this))
    },

    //this will run also on logout. Must fix. This happens because of the re-render and the Main container will send in a new state as props.
    componentDidMount: function () {
        this.getUser(this.props.user)
    },
    componentWillReceiveProps: function (nextProps) {
        this.getUser(nextProps.user)
    },
    getUser: function (user) {
        if (user.length >= 5) {
            userFunctions.loginPassword(user)
                .then(function (data) {
                    this.setState({
                        isLoggedIn: true,
                        user: data.user.username,
                        token: data.token
                    });
                }.bind(this));
            //must re-render the Home container to re-render the rest of the page.
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
        this.props.onUpdateLogin(false, "");
    },
    render: function () {
        return (
            <LoginForm isLoggedIn={this.state.isLoggedIn}
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