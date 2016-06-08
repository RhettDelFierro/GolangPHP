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
        return {
            //isLoggedIn maybe will pass into /teachers/:users to re-render "/"
            isLoggedIn: false,
            user: "",
            email: "",
            password: "",
            //maybe have all the user info here: current grades, class average, etc.
            userInfo: [],
            duplicate: false,
            helpBlock: "hidden"
        }
    },
    handleUpdateUser: function (e) {
        this.setState({
            user: e.target.value
        });

        if (this.state.user.length >= 8) {
            this.ajaxValidUserName();
            this.setState({helpBlock: "show"})
        } else {
            this.setState({helpBlock: "hidden"})
        }
    },
    ajaxValidUserName: function () {
        if (this.state.user.length >= 8) {
            userFunctions.verfifyName(this.state.user)
                .then(function (userdata) {
                    this.setState({
                        duplicate: userdata.taken,
                        user: userdata.username
                    })
                }.bind(this));
            this.setState({helpBlock: "show"})
        } else {
            this.setState({helpBlock: "hidden"})
        }
    },
    handleUpdateEmail: function (e) {
        this.setState({
            email: e.target.value
        });
    },
    handleUpdatePassword: function (e) {
        this.setState({
            password: e.target.value
        });
    },
    handleSubmitUser: function (e) {
        e.preventDefault();
        //in case for a backspace, but this should be done after the push to a new route:
        this.setState({
            user: "",
            email: "",
            password: "",
            //maybe have all the user info here: current grades, class average, etc.
            userInfo: []
        });

        //make the ajax call.
        //on success push to /teachers/users.
        //on fail stay on page and display error messages. Also re-set the state to have the info.
        userFunctions.registerUser({
            user: this.state.user,
            email: this.state.email,
            pasword: this.state.password
        })
            .then(function(userdata){
                this.context.router.push({
                    pathname: "/teachers/" + userdata.username,
                    query: {
                        playerOne: this.props.routeParams.playerOne,
                        playerTwo: this.state.username
                    }
                })
            }.bind(this));

        if (this.props.routeParams.playerOne) {
            //going to go to "/teachers/:users" with the user information to re-render.
            this.context.router.push({
                pathname: "/teachers/" + this.state.username,
                state: {
                    playerOne: this.props.routeParams.playerOne,
                    playerTwo: this.state.username
                }
            })
        } else {
            //go to /playerTwo.
            this.context.router.push("/playerTwo/" + this.state.username);
        }
    },
    render: function () {
        return (
            <RegisterForm onupdateUser={this.handleUpdateUser}
                          onupdateEmail={this.handleUpdateEmail}
                          onupdatePassword={this.handleUpdatePassword}
                          onSubmitUser={this.handleSubmitUser}
                          user={this.state.user}
                          duplicate={this.state.duplicate}/>
        )
    }
});

module.exports = RegisterFormContainer;