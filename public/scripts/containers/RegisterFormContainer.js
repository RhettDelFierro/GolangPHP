var React = require("react");
var RegisterForm = require("../components/RegisterForm");
var userFunctions = require("../utils/userFunctions");

var RegisterFormContainer = React.createClass({
    contextTypes: {
        router: React.PropTypes.object.isRequired
    },
    //make the ajax calls from here.
    getInitialState: function () {
        //bet we get the values of the input texts and make an object out of it
        return ({
            //isLoggedIn maybe will pass into /teachers/:users to re-render "/"
            isLoggedIn: false,
            user: "",
            email: "",
            password: "",
            userInfo: {},
            duplicate: true,
            helpBlock: "hidden"
        })
    },
    componentWillUnmount: function(){
        this.setState({
            user: "",
            email: "",
            password:"",
            userInfo: [],
            duplicate: true,
            helpBlock: "hidden"
        })
    },
    handleUpdateUser: function (e) {
        this.setState({
            user: e.target.value
        });
        console.log(this.state.user);
        if (this.state.user.length >= 8) {
            this.ajaxValidUserName();
            this.setState({helpBlock: "show"})
        } else {
            this.setState({helpBlock: "hidden"})
        }
    },
    ajaxValidUserName: function () {
        userFunctions.verifyName(this.state.user)
            .then(function (userdata) {
                console.log("from the container: ", userdata);
                this.setState({
                    duplicate: userdata.taken
                })
            }.bind(this));
    },
    handleUpdateEmail: function (e) {
        this.setState({
            email: e.target.value
        });
        console.log(this.state.email)
    },
    handleUpdatePassword: function (e) {
        this.setState({
            password: e.target.value
        });
    },
    handleSubmitUser: function (e) {
        e.preventDefault();
        console.log(this.state);
        //in case for a backspace, but this should be done after the push to a new route:
        this.setState({
            user: "",
            email: "",
            password: "",
            userInfo: {}
        });

        //make the ajax call.
        //on success push to /teachers/users.
        //on fail stay on page and display error messages. Also re-set the state to have the info.
        userFunctions.registerUser({
                user: this.state.user,
                email: this.state.email,
                password: this.state.password
            })
            .then(function (userdata) {
                console.log(userdata);
                this.props.onUpdateLogin(true, userdata.username)
            }.bind(this));
    },
    render: function () {
        return (
            <RegisterForm onUpdateUser={this.handleUpdateUser}
                          onUpdateEmail={this.handleUpdateEmail}
                          onUpdatePassword={this.handleUpdatePassword}
                          onSubmitUser={this.handleSubmitUser}
                          user={this.state.user}
                          duplicate={this.state.duplicate}
                          helpBlock={this.state.helpBlock}/>
        )
    }
});

module.exports = RegisterFormContainer;